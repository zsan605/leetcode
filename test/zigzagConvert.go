package test

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

func convert(s string, numRows int) string {

	l := len(s)
	col := l / numRows
	//ret := l % numRows
	span := 2 * (numRows - 1)

	if numRows >= l  || numRows == 1{
		return s
	}

	var strArray []string
	for i:=1; i <= col; i++ {
		leftIndex := (i-1)*span
		rightIndex := i*span
		if rightIndex >l {
			rightIndex = l
		}
		fmt.Println(s[leftIndex:rightIndex])
		strArray = append(strArray, s[leftIndex:rightIndex])
	}

	var retStr string
	for i:=1; i< numRows; i++ {
		for _, arr := range strArray {
			if i == 1 {
				retStr += arr[0:1]
			} else if i < numRows {
				retStr += arr[i-1:i]
				//l-
			}
		}

	}

	//
	//var retStr string
	//for i := 1; i <= numRows; i++ {
	//	for j := 1; j <= col; j++ {
	//		var index int
	//		if i == 1{
	//			index := (j -1) * span
	//			retStr +=s[index:index+1]
	//		} else if i < numRows {
	//			if j == 1 {
	//				index = i-1
	//				retStr +=s[index:index+1]
	//			} else {
	//				index = (j -1) * span -i +1
	//				retStr +=s[index:index+1]
	//
	//				if j == col && ret <=0 {
	//					continue
	//				}
	//				index = (j -1) * span +i -1
	//				retStr +=s[index:index+1]
	//			}
	//		} else {
	//			index := (j -1) * span -i +1
	//			if index >= 0{
	//				retStr +=s[index:index+1]
	//			}
	//		}
	//
	//		if ret >0 && col==j{
	//			ret--
	//		}
	//
	//	}
	//}

	return ""
}


func Create(params map[string]string) string {
	// 生成签名：mb_strtoupper(md5(原字符串&key=md(密钥)))
	delete(params, "sign")
	count := len(params)
	if len(params) == 0 {
		return ""
	}


	// sort by key
	keys := make([]string, 0, count)
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)


	queryStr := ""
	for i := 0; i < count; i++ {
		val := params[keys[i]]
		if len(val) > 0 && len(keys[i]) > 0 && val != "0" {
			queryStr += fmt.Sprintf("%s=%s&", keys[i], val)
		}
	}
	fmt.Println(queryStr)
	queryStr += fmt.Sprintf("key=%x", md5.Sum([]byte("密钥[找券rd提供]")))
	fmt.Println(queryStr)
	md5Str := fmt.Sprintf("%x", md5.Sum([]byte(queryStr)))
	return strings.ToUpper(md5Str)
}
