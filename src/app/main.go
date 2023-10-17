package main

import (
	"app/config"
	"app/log"
	"app/routes"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"
	"net/http"
	"os"
	"runtime"
)

var (
	cfg     = flag.String("cfg", "config.json", "cfg")
	logPath = flag.String("log", "/tmp/", "log path")
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	//log.SetLogger(config.GetConfig().LogDir)

	config.Init(*cfg)
	fmt.Println(os.Getenv("ENV"))
	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

}
func main() {
	router := routes.Init()
	addr := config.GetConfig().Server.Host + ":" + config.GetConfig().Server.Port

	if runtime.GOOS == "windows" {
		server := &http.Server{
			Addr:    addr,
			Handler: router,
		}
		// 启动HTTP服务器
		log.Logger.Info("server run ", log.String("addr", addr))
		server.ListenAndServe()
	} else {
		//	err := endless.ListenAndServe(addr, router)
		//	log.Logger.Info("server run ", log.String("addr", addr))
		//	if err != nil {
		//		log.Logger.Error("start failed", log.Error(err))
		//		os.Exit(0)
		//	}
	}
}
