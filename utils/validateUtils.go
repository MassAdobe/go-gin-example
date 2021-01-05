/**
 * @Time : 2020/12/18 10:17 上午
 * @Author : MassAdobe
 * @Description: utils
**/
package utils

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

/**
 * @Author: MassAdobe
 * @TIME: 2020-04-26 21:13
 * @Description: 校验邮箱
**/
func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

/**
 * @Author: MassAdobe
 * @TIME: 2020-04-26 21:13
 * @Description: 校验手机号码
**/
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

/**
 * @Author: MassAdobe
 * @TIME: 2020-06-03 11:12
 * @Description: 返回当前年份
**/
func RtnCurYear() string {
	return strconv.Itoa(time.Now().Year())
}

/**
 * @Author: MassAdobe
 * @TIME: 2020-06-11 09:46
 * @Description: 替换数据库返回的时间为时间
**/
func ReplaceTm(time string) string {
	if len(time) < 19 {
		return ""
	}
	return strings.Replace(time[:19], "T", " ", -1)
}

/**
 * @Author: MassAdobe
 * @TIME: 2020-07-07 15:10
 * @Description: 替换数据库返回的时间为日期
**/
func ReplaceDt(time string) string {
	if len(time) < 19 {
		return ""
	}
	return time[:10]
}
