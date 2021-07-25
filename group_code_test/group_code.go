package main

import (
	"encoding/json"
	"log"

	"github.com/golang-demo/utils"
)

// ========================================================
// 联动处理
// ========================================================

// FuncArrayMap ...
var FuncArrayMap = make(map[string]utils.FuncArray)

// TopGroupNodeMap 存储联动顺序
var TopGroupNodeMap = make(map[string]CommonGroupNode)

// CommonGroupNode ...
type CommonGroupNode struct {
	GroupCode             string
	GroupCodeRelationPath string
	EntityIDs             []string
	ChildGroupNodes       []*CommonGroupNode
}

// GroupInfo ...
type GroupInfo struct {
	Entities              []string `json:"entities,omitempty"`
	GroupCode             string   `json:"group_code,omitempty"`
	GroupCodeRelationPath string   `json:"group_code_relation_path,omitempty"`
	Count                 int      `json:"count,omitempty"`
}

// TODO: “EntityIDs: []string{}”， 改为 EntityIDs: []CommonEntity{},
// InitMap ...
func InitMap() {
	TopGroupNodeMap["fof"] = CommonGroupNode{
		GroupCode:             "fof",
		GroupCodeRelationPath: "fof",
		EntityIDs:             []string{},
		ChildGroupNodes: []*CommonGroupNode{
			{
				GroupCode:             "invested_fund",
				GroupCodeRelationPath: "fof.invested_fund",
				EntityIDs:             []string{},
				ChildGroupNodes: []*CommonGroupNode{
					{
						GroupCode:             "invested_fund_investments",
						GroupCodeRelationPath: "fof.invested_fund.invested_fund_investments",
						EntityIDs:             []string{},
					},
				},
			},
			{
				GroupCode: "fof.invested_company",
				EntityIDs: []string{},
			},
		},
	}

	TopGroupNodeMap["fund"] = CommonGroupNode{
		GroupCode: "fund",
		EntityIDs: []string{},
		ChildGroupNodes: []*CommonGroupNode{
			{
				GroupCode:             "invested_company",
				GroupCodeRelationPath: "fund.invested_company",
				EntityIDs:             []string{},
			},
			{
				GroupCode:             "lp",
				GroupCodeRelationPath: "fund.lp",
				EntityIDs:             []string{},
			},
		},
	}

	TopGroupNodeMap["invested_fund"] = CommonGroupNode{
		GroupCode: "invested_fund",
		EntityIDs: []string{},
		ChildGroupNodes: []*CommonGroupNode{
			{
				GroupCode:             "invested_fund_investments",
				GroupCodeRelationPath: "invested_fund.invested_fund_investments",
				EntityIDs:             []string{},
			},
		},
	}

	FuncArrayMap["fof"] = fillFuncFOFArray()
	FuncArrayMap["fund"] = fillFuncFundArray()
	FuncArrayMap["invested_fund"] = fillFuncInvestedFundArray()
}

// ================================================  联动查找逻辑 ================================================
func main() {
	InitMap()

	groupInfo := GroupInfo{
		GroupCode:             "fof",
		GroupCodeRelationPath: "fof",
		Count:                 2,
		Entities:              []string{"fof_100", "fof_200"},
	}

	res := GetRevelantGroupDatas(groupInfo, "teamID_100")
	jsonRes, _ := json.Marshal(res)
	log.Printf("res: %+v", string(jsonRes))
}

// GetRevelantGroupDatas ...
func GetRevelantGroupDatas(groupInfo GroupInfo, teamID string) []GroupInfo {
	// 获取顶层标签对应的联动群，并直接传入文件导入的 entityID
	topCommonGroupNode := TopGroupNodeMap[groupInfo.GroupCode]
	topCommonGroupNode.EntityIDs = groupInfo.Entities

	resGroupInfos := make([]GroupInfo, 0)
	resGroupInfos = append(resGroupInfos, groupInfo)

	// 递归调用
	err := fillLevelGroupDatas(groupInfo, &topCommonGroupNode, teamID, &resGroupInfos)
	if err != nil {
		// utils.Logger.Errorf("GetRevelantGroupDatas.GetLevelGroupData err: %+v", err)
		log.Printf("GetRevelantGroupDatas.GetLevelGroupData err: %+v", err)
		return nil
	}

	// jsonRes, _ := json.Marshal(topCommonGroupNode)
	// log.Printf("topCommonGroupNodef: %+v", string(jsonRes))

	return resGroupInfos
}

// fillLevelGroupDatas ...
func fillLevelGroupDatas(groupInfo GroupInfo, groupNode *CommonGroupNode, teamID string, resGroupInfos *[]GroupInfo) error {

	// 获取每个标签对应的子级函数（仅限定一级，例如，fof 对应函数： InvestedFund / InvestedCompany）
	funcArray, _ := FuncArrayMap[groupInfo.GroupCode]
	// log.Printf("input: groupInfo.GroupCode [%+v]funcArray: %+v", groupInfo.GroupCode, funcArray)

	// 子级节点
	childGroupNode := new(CommonGroupNode)
	var ok bool

	// 抽取公共方法：将函数获取的结果，填入到标签联动群，后续可以直接利用递归
	for i := 0; i < len(funcArray); i++ {
		childGroupInfo := GroupInfo{}
		if groupInfoValues, err := funcArray.Call(i, groupNode.EntityIDs, teamID, groupNode.GroupCodeRelationPath+"."); err == nil {
			if childGroupInfo, ok = groupInfoValues[0].Interface().(GroupInfo); ok {

				// 每个方法返回的结果就是最终需要的信息
				*resGroupInfos = append(*resGroupInfos, childGroupInfo)

				// 补充标签联动群中的参数ID，仅针对子一级，所以不会重复
				childGroupNode = fillChildGroupNode(groupNode, childGroupInfo)
			}
		} else {
			// utils.Logger.Errorf("failed to call func: index[%+v] err: %+v", i, err)
			log.Printf("failed to call func: index[%+v] err: %+v", i, err)
			return err
		}

		// 利用子一级计算子二级
		log.Printf("=========下一级调用开始：fillLevelGroupDatas...========")
		err := fillLevelGroupDatas(childGroupInfo, childGroupNode, teamID, resGroupInfos)
		if err != nil {
			// utils.Logger.Errorf("failed to fillLevelGroupDatas: index[%+v] err: %+v", i, err)
			log.Printf("failed to fillLevelGroupDatas: index[%+v] err: %+v", i, err)
			return err
		}
		log.Printf("=========下一级调用完成：fillLevelGroupDatas！！！========")
	}

	return nil
}

// 将函数中获取的数据，填入到导入标签对应的联动群。返回当前填入数据的联动标签
func fillChildGroupNode(groupNode *CommonGroupNode, childGroupInfo GroupInfo) *CommonGroupNode {
	resCommonGroupNode := new(CommonGroupNode)

	log.Printf("groupNode.code: %+v, childGroupInfo.code: %+v", groupNode.GroupCode, childGroupInfo.GroupCode)

	if groupNode != nil && groupNode.ChildGroupNodes != nil {
		for _, childGroupNode := range groupNode.ChildGroupNodes {
			if childGroupNode.GroupCode == childGroupInfo.GroupCode {
				childGroupNode.EntityIDs = childGroupInfo.Entities

				return childGroupNode
			}
		}
	}

	return resCommonGroupNode
}

// ================================================  填充联动函数 ================================================

func fillFuncFOFArray() utils.FuncArray {
	funcArray := make(utils.FuncArray, 0)
	funcArray.Add(getInvestedFund)
	funcArray.Add(getInvestedCompanies)
	return funcArray
}

func fillFuncFundArray() utils.FuncArray {
	funcArray := make(utils.FuncArray, 0)
	funcArray.Add(getInvestedCompanies)
	funcArray.Add(getLP)
	return funcArray
}

func fillFuncInvestedFundArray() utils.FuncArray {
	funcArray := make(utils.FuncArray, 0)
	funcArray.Add(getInvestedFundInvestments)
	return funcArray
}

// ================================================  联动sql ================================================

func getInvestedFund(entityIDs []string, teamID string, parentGroupCodeRelationPath string) (GroupInfo, error) {
	groupInfo := GroupInfo{}

	groupInfo.Entities = []string{"invested_fund_100", "invested_fund_200"}
	groupInfo.GroupCode = "invested_fund"
	groupInfo.GroupCodeRelationPath = parentGroupCodeRelationPath + groupInfo.GroupCode
	groupInfo.Count = len(groupInfo.Entities)

	return groupInfo, nil
}

func getInvestedCompanies(entityIDs []string, teamID string, parentGroupCodeRelationPath string) (GroupInfo, error) {
	groupInfo := GroupInfo{}

	groupInfo.Entities = []string{"invested_company_100", "invested_company_200"}
	groupInfo.GroupCode = "invested_company"
	groupInfo.GroupCodeRelationPath = parentGroupCodeRelationPath + groupInfo.GroupCode
	groupInfo.Count = len(groupInfo.Entities)
	return groupInfo, nil
}

func getLP(entityIDs []string, teamID string, parentGroupCodeRelationPath string) (GroupInfo, error) {
	groupInfo := GroupInfo{}

	groupInfo.Entities = []string{"lp_100", "lp_200"}
	groupInfo.GroupCode = "lp"
	groupInfo.GroupCodeRelationPath = parentGroupCodeRelationPath + groupInfo.GroupCode
	groupInfo.Count = len(groupInfo.Entities)

	return groupInfo, nil
}

func getInvestedFundInvestments(entityIDs []string, teamID string, parentGroupCodeRelationPath string) (GroupInfo, error) {
	groupInfo := GroupInfo{}

	groupInfo.Entities = []string{"invested_fund_investments_100", "invested_fund_investments_200"}
	groupInfo.GroupCode = "invested_fund_investments"
	groupInfo.GroupCodeRelationPath = parentGroupCodeRelationPath + groupInfo.GroupCode
	groupInfo.Count = len(groupInfo.Entities)
	return groupInfo, nil
}
