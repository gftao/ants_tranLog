package handle

import (
	"net/http"
	"github.com/panjf2000/ants"
	"tranLog/digest"
	"fmt"
	"golib/modules/idManager"
	"golib/modules/config"
	"errors"
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
		t := i.(ITranLogTask)
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

	t := &digest.TranLogTask{W: w, R: r}
	t.NodeName = "DoLogs"
	t.Id = IdGenerator.GetUint32()
	t.Wg.Add(1)
	p.Serve(t)
	t.Wg.Wait()
	t.Infof("running goroutines: %d", p.Running())
}
