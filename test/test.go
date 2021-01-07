/**
 * @Time : 2020/12/18 1:38 下午
 * @Author : MassAdobe
 * @Description: test
**/
package main

import (
	"fmt"
)

func main() {
	chanx := make(chan bool)
	close(chanx)

	if _, ok := <-chanx; ok {
		fmt.Println("已经关闭")
	} else {
		fmt.Println("还没关闭")
	}

}
