package handle

import (
	"net/http"
	"github.com/panjf2000/ants"
	"fmt"
	"golib/modules/idManager"
	"golib/modules/config"
	"errors"
	"tranLog/digest"
	"context"
)

var (
	p           *ants.PoolWithFunc
	IdGenerator idManager.IdGenerator
	//iw          *goSnowFlake.IdWorker
)

type ITranLogTask interface{ DoTask() }

func DoHandleInit() error {
	var err error

	if !config.HasModuleInit() {
		return errors.New("配置文件未初始化，请先初始化")
	}
	config.SetSection("pool")

	sz := config.IntDefault("poolSize", 10)
	clusterId := config.IntDefault("clusterId", 0)

	p, err = ants.NewPoolWithFunc(sz, func(i interface{}) error {
		t, ok := i.(ITranLogTask)
		if !ok {
			return errors.New("parameter is not ITranLogTask")
		}
		t.DoTask()
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	//iw, err = goSnowFlake.NewIdWorker(1)
	//if err != nil {
	//	fmt.Println(err)
	//	return err
	//}
	IdGenerator = idManager.NewUIdGenerator(clusterId)

	return nil
}
func DoHandleClose() {
	p.Release()
}

func DoHandles(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	t := &digest.TranLogTask{W: w, R: r, Cancel: cancel}
	t.NodeName = "DoLogs"
	t.Id = IdGenerator.GetUint32()
	err := p.Serve(t)

	if err != nil {
		t.Errorf("Serve failed:%s", err)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	t.Infof("running goroutines: %d", p.Running())
	select {
	case <-ctx.Done():
		t.Infof("request cancelled\n")
	}
}
