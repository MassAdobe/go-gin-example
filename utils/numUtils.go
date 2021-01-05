/**
 * @Time : 2020/12/18 10:19 上午
 * @Author : MassAdobe
 * @Description: utils
**/
package utils

import (
	"fmt"
	"strconv"
)

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/18 10:19 上午
 * @Description: 去除多重元素
**/
func RemoveRepeatedElement(arr []int64) (newArr []int64) {
	newArr = make([]int64, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/18 10:19 上午
 * @Description: 计算浮点值保留两位小数
**/
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
