/**
 * @Time : 2020/12/21 4:57 下午
 * @Author : MassAdobe
 * @Description: go_framework
**/
package main

import (
	"fmt"
	"github.com/MassAdobe/go-gin-example/router"
	"github.com/MassAdobe/go-gin/logs"
	"github.com/MassAdobe/go-gin/nacos"
	"github.com/MassAdobe/go-gin/start"
	"net/http"
	"os"
	"strconv"
	"time"
)

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/17 1:51 下午
 * @Description: 启动项
**/
func main() {
	rtr := router.Routers() // 配置gin启动
	server := &http.Server{ // 创建服务
		Addr:           ":" + strconv.Itoa(int(nacos.InitConfiguration.Serve.Port)),
		Handler:        rtr,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logs.Lg.Info("启动", logs.Desc(fmt.Sprintf("启动端口: %d", nacos.InitConfiguration.Serve.Port)))
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed { // 监听并启动服务
			logs.Lg.Error("启动失败", err)
			os.Exit(1)
		}
	}()
	start.GracefulShutdown(server) // 优雅停服
}
