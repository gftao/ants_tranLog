package serve

import (
	"fmt"
	"errors"
	"time"
	"os"
	"os/signal"
	"net/http"
	"log"
	"context"
	"golib/modules/config"
	"golib/modules/logr"
	"tranLog/handle"
)

type httpSvrConf struct {
	ListenIp     string
	ListenPort   int
	RecvTimeOut  int
	WriteTimeOut int
	MaxAccNum    int
}

type HttpSvr struct {
	conf *httpSvrConf
}

func (t *HttpSvr) InitConfig() error {

	if !config.HasModuleInit() {
		return errors.New("配置文件未初始化，请先初始化")
	}

	config.SetSection("server")

	cf := &httpSvrConf{}
	cf.ListenIp = config.StringDefault("host", "")
	cf.ListenPort = config.IntDefault("port", 9090)
	cf.RecvTimeOut = config.IntDefault("readTimeout", 30)
	cf.WriteTimeOut = config.IntDefault("writeTimeout", 30)
	t.conf = cf
	fmt.Println("HttpSvr加载成功")

	return nil
}
func (t *HttpSvr) RunSvr(h http.Handler) {
	srv := &http.Server{
		Addr:           t.conf.ListenIp + fmt.Sprintf(":%d", t.conf.ListenPort),
		Handler:        h,
		ReadTimeout:    time.Duration(t.conf.RecvTimeOut) * time.Second,
		WriteTimeout:   time.Duration(t.conf.WriteTimeOut) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		defer logr.Info("----HttpSvr关闭----")
		logr.Info("----HttpSvr启动----")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	select {
	case <-c:
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		handle.DoHandleClose()
		srv.Shutdown(ctx)
	}

	log.Println("shutting down")
}
