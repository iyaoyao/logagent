package g

import (
	"fmt"
	etcd_client "github.com/coreos/etcd/clientv3"
	"github.com/iyaoyao/logagent/tailf"
	logs "github.com/Sirupsen/logrus"
	"time"
	"context"
	"encoding/json"
	"strings"
)

type EtcdClient struct {
	client *etcd_client.Client
	keys   []string
}

var (
	etcdClient *EtcdClient
)


func InitEtcd(addr, key string)(collectConf []tailf.CollectConf, err error){
	cli, err:= etcd_client.New(etcd_client.Config{
		Endpoints: []string{addr},
		DialTimeout: 5 *time.Second,
	})

	if err!=nil{
		logs.Error("etcd conn faild")
		return
	}
	etcdClient = &EtcdClient{
		client: cli,
	}
	logs.Info("etcd conn success")
	if strings.HasSuffix(key, "/") == false {
		key = key + "/"
	}

	for _, ip := range localIPArray {
		etcdKey := fmt.Sprintf("%s%s", key, ip)
		fmt.Println(etcdKey)
		etcdClient.keys = append(etcdClient.keys, etcdKey)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := cli.Get(ctx, etcdKey)
		if err != nil {
			logs.Error("client get from etcd failed, err:%v", err)
			continue
		}
		cancel()
		logs.Info("resp from etcd:%v", resp.Kvs)
		for _, v := range resp.Kvs {
			if string(v.Key) == etcdKey {
				err = json.Unmarshal(v.Value, &collectConf)
				if err != nil {
					logs.Error("unmarshal failed, err:%v", err)
					continue
				}

				logs.Info("log config is %v", collectConf)
			}
		}
	}
	return
}
