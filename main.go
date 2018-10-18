package main

import (
	"flag"
	"fmt"
	"os"
	"golib/modules/config"
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
	"tranLog/serve"
	"golib/modules/gormdb"
	"golib/modules/logr"
	"tranLog/handle"
	"runtime"
	"github.com/hashicorp/go-version"
)

var conf = flag.String("conf", "./etc/tran.ini", "config file")
var vs = flag.Bool("v", false, "version")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	//初始化版本
	v, err := version.NewVersion("1.0.0")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if *vs {
		fmt.Printf("当前版本: %s", v)
		return
	}

	//初始化配置文件
	err = config.InitModuleByParams(*conf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//初始化日志
	err = logr.InitModules()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logr.Info("初始化日志成功,版本: ", v)

	//初始化db
	err = gormdb.InitModule()
	if err != nil {
		fmt.Println("初始化数据库失败", err)
		return
	}

	//初始化池
	handle.DoHandleInit()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	r := mux.NewRouter()
	r.HandleFunc("/tranLogs", handle.DoHandles)
	n := negroni.New()
	n.UseHandler(r)
	sockSvr := serve.HttpSvr{}
	err = sockSvr.InitConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	sockSvr.RunSvr(n)
}
