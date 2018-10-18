# ants_tranLog
`ants`是一个高性能的协程池，本例用 [ants](https://github.com/panjf2000/ants) 协成池实现简单的http服务端。

## 实用性:
- 实现了自动调度并发的goroutine，复用goroutine
- 定时清理过期的goroutine，进一步节省资源
- 可作为小微服务框架
## 使用
写 go 并发程序的时候如果程序会启动大量的 goroutine ，势必会消耗大量的系统资源（内存，CPU），通过使用 `ants`，可以实例化一个协程池，复用 goroutine ，节省资源，提升性能：

``` go
//定义任务 task 接口
type TranLogWorker interface{ DoWork() }

func DoHandleInit() error {
	var err error

	if !config.HasModuleInit() {
		return errors.New("配置文件未初始化，请先初始化")
	}
	config.SetSection("pool")

	sz := config.IntDefault("poolSize", 10)
	clusterId := config.IntDefault("clusterId", 0)

	p, err = ants.NewPoolWithFunc(sz, func(i interface{}) error {
		//pool池接受task 并字处理。
		t := i.(TranLogWorker)
		t.DoWork()
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return err
	} 
	IdGenerator = idManager.NewUIdGenerator(clusterId)

	return nil
}

```
## 定义task 实例
``` go
func DoHandles(w http.ResponseWriter, r *http.Request) {
	//初始化task 实例并serve到pool池
	t := &digest.TranLogWorker{W: w, R: r}
	t.NodeName = "DoLogs"
	t.Id = IdGenerator.GetUint32()
	t.Wg.Add(1)
	p.Serve(t)
	t.Wg.Wait()
	t.Infof("running goroutines: %d", p.Running())
}
```

## 该方法架构还不错
