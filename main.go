package main

import (
	"github.com/iyaoyao/logagent/g"
	"github.com/iyaoyao/logagent/tailf"
	"github.com/iyaoyao/logagent/kafka"
	logs "github.com/Sirupsen/logrus"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)


func init(){
	//初始化日志模块
	err := g.InitLogger()
	if err!=nil{
		logs.Error("init logger error:",err)
		return
	}
	logs.Info("init logger success")
	//初始化配置文件
	err = g.InitConfig("cfg.json")
	if err != nil{
		logs.Error("init config error:",err)
		return
	}
	logs.Info("init config  success")
	//初始化etcd
	collectConf, err := g.InitEtcd(g.Config().EtcdAddr, g.Config().EtcdKey)
	if err!=nil{
		logs.Error("init etcd error:",err)
		return
	}
	fmt.Println(collectConf)
	//初始化tailf
	err = tailf.InitTail(collectConf, g.Config().ChanSize)
	if err != nil {
		logs.Error("init tail error:", err)
		return
	}
	logs.Info("init tail success")
	//初始化kafka
	err = kafka.InitKafka(g.Config().KafkaAddr)
	if err != nil {
		logs.Error("init kafka error:", err)
		return
	}
	logs.Info("init kafka success")
	//接受信号重新载入配置文件及退出等
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGUSR1,syscall.SIGINT)
	go func() {
		for {
			switch <-sigs {
			case syscall.SIGUSR1:
				logs.Info("Reloaded config", g.InitConfig("cfg.json"))
				logs.Info(g.Config().EtcdAddr)
			case syscall.SIGINT:
				logs.Info("ctrl+c")
				os.Exit(0)
			}
		}
	}()
}

func main(){
	logs.Println("starting")
	err := Run()
	if err!=nil{
		logs.Info("run error")
	}
}
