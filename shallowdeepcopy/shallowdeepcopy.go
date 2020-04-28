package shallowdeepcopy

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

// SearchTemplate 方法传入结构体（第一层）
type SearchTemplate struct {
	ID         string
	Conditions []*SearchCondition
	ColumnIds  []string
}

// SearchCondition 方法传入结构体（第二层）
type SearchCondition struct {
	ID     string
	Values []*SearchConditionValue
}

// SearchConditionValue 方法传入结构体（第三层）
type SearchConditionValue struct {
	Value        string
	DisplayValue string
}

// SearchRequest 中间参数结构体
type SearchRequest struct {
	ID         string
	Conditions []*SearchCondition
}

// ExecuteTemplate ...
func ExecuteTemplate() {
	searchTemplate := new(SearchTemplate)
	searchTemplate.ID = "10"

	var SearchConditionTmp = SearchCondition{
		ID: "100",
	}
	var SearchConditionValueTmp = SearchConditionValue{
		Value:        "nihao",
		DisplayValue: "你好",
	}
	SearchConditionTmp.Values = append(SearchConditionTmp.Values, &SearchConditionValueTmp)
	searchTemplate.Conditions = append(searchTemplate.Conditions, &SearchConditionTmp)

	var searchReq = SearchRequest{
		ID: "1",
	}

	if searchTemplate != nil {
		for j, condition := range searchTemplate.Conditions {
			var conditionTemp SearchCondition
			conditionTemp.Values = make([]*SearchConditionValue, 2)

			// conditionTemp = *condition // 加上此句，相当于执行浅拷贝，无论之后怎么创建临时变量，searchTemplate 和 searchReq 中的数组类型变量都是指向同一个地址。
			for i, value := range condition.Values {
				var valueTemp SearchConditionValue
				valueTemp = *value
				conditionTemp.Values[i] = &valueTemp

				fmt.Printf("&valueTemp: %+v; &value: %+v", unsafe.Pointer(&valueTemp), unsafe.Pointer(value))
				fmt.Println("")
			}
			searchReq.Conditions = append(searchReq.Conditions, &conditionTemp)

			fmt.Printf("condition: %+v; searchReq.Conditions[j]: %+v", unsafe.Pointer(condition), unsafe.Pointer(searchReq.Conditions[j]))
			fmt.Println("")
		}
	}

	resultBefore, err := json.Marshal(searchTemplate)
	if err != nil {
		return
	}
	fmt.Println("searchTemplate, before: ", string(resultBefore))

	searchReq.Conditions[0].Values[0].Value += " helloworld"

	resultAfter, err := json.Marshal(searchTemplate)
	if err != nil {
		return
	}
	fmt.Println("searchTemplate, after: ", string(resultAfter))

	return
}
