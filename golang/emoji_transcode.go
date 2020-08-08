package main

import (
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// 将表情转换为字符串 ...
func UnicodeEmojiCode(s string) string {
	ret := ""
	rs := []rune(s)
	for i := 0; i < len(rs); i++ {
		if len(string(rs[i])) == 4 {
			u := `[\u` + strconv.FormatInt(int64(rs[i]), 16) + `]`
			ret += u
		} else {
			ret += string(rs[i])
		}
	}
	return ret
}

// 将字符串转换为表情 ...
func UnicodeEmojiDecode(s string) string {
	//emoji表情的数据表达式
	re := regexp.MustCompile("\\[[\\\\u0-9a-zA-Z]+\\]")

	//提取emoji数据表达式
	reg := regexp.MustCompile("\\[\\\\u|]")
	src := re.FindAllString(s, -1)
	for i := 0; i < len(src); i++ {
		e := reg.ReplaceAllString(src[i], "")
		p, err := strconv.ParseInt(e, 16, 32)
		if err == nil {
			s = strings.Replace(s, src[i], string(rune(p)), -1)
		}
	}
	return s
}

func main() {
	// 方案1
	emoji := "😀😨😨"
	text := UnicodeEmojiCode(emoji)

	emojiOut := UnicodeEmojiDecode(text)

	log.Printf("text:%+s", text)
	log.Printf("emoji:%+s", emojiOut)

	// 方案2
	inputEmoji := "😀😨😨"
	resultText := url.QueryEscape(inputEmoji)

	emojiResult, _ := url.QueryUnescape(resultText)
	log.Println("QueryEscape: ", resultText)
	log.Println("QueryUnescape: ", emojiResult)
}

// 参考
// - [golang emoji表情处理](https://blog.csdn.net/zyx4843/article/details/77773286)
