/**
 * @Time : 2020/12/18 10:15 上午
 * @Author : MassAdobe
 * @Description: utils
**/
package utils

import (
	"github.com/MassAdobe/go-gin/systemUtils"
	"time"
)

/**
 * @Author: MassAdobe
 * @TIME: 2020-04-26 19:56
 * @Description: 获取当前时间戳
**/
func RtnTimestamp() int64 {
	t := time.Now()
	return t.Unix()
}

/**
 * @Author: MassAdobe
 * @TIME: 2020-04-26 21:12
 * @Description: 获取当前时间戳(毫秒)
**/
func RtnTimestampMs() int64 {
	return time.Now().UnixNano() / 1e6
}

/**
 * @Author: MassAdobe
 * @TIME: 2020-04-26 21:13
 * @Description: 返回日期字符串
**/
func RtnDtString() (timsStr string) {
	timsStr = time.Now().Format(systemUtils.TIME_FORMAT_MONTH)
	return
}
