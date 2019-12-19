/***************************************************
 ** @Desc : This file for ...
 ** @Time : 2019/10/26 11:17
 ** @Author : yuebin
 ** @File : sort_go
 ** @Last Modified by : yuebin
 ** @Last Modified time: 2019/10/26 11:17
 ** @Software: GoLand
****************************************************/
package utils

import (
	"sort"
)

/*
* 对map的key值进行排序
 */
func SortMap(m map[string]string) []string {
	var arr []string
	for k := range m {
		arr = append(arr, k)
	}
	sort.Strings(arr)
	return arr
}
