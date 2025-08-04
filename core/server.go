package core

import (
	"boost/data/server/global"
	"boost/data/server/initialize"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))
	buddha := `
                  _ooOoo_
                o8888888o
                  88" . "88
              	    (| -_- |)
                  O\  =  /O
               ____/'---'\____
             .'  \\|     |//  '.
            /  \\|||  :  |||//  \
           /  _||||| -:- |||||_  \
           |   | \\\  -  /'| |   |
           | \_|  '\'---'//  |_/ |
           \  .-\__ '-. -' __/-.  /
         ___'. .'  /--.--\  '. .'___
      ."" '<  '.___\_<|>_/___.' _> \"".
     | | :  '- \'. ;'. _/; .'/ /  .' ; |
     \  \ '-.   \_\_'. _.'_/_/  -' _.' /
   ====='-.____'.___ \_____/___.-'____.-'=====
                   '=---='
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
            佛祖保佑        永无BUG
   `
	global.GVA_LOG.Info(buddha)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Minute,
		WriteTimeout:   10 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
}
