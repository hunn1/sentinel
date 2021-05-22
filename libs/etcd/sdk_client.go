package etcd

import (
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	dialTimeout          = 5 * time.Second
	RequestTimeout       = 2 * time.Second
	AutoSyncInterVal     = time.Second * 3
	DialKeepAliveTime    = time.Second * 3
	DialKeepAliveTimeout = time.Second * 3
	MessageSize          = 2 * 1024 * 1024
	endpoints            = []string{"127.0.0.1:2379"}
	Etcd                 clientv3.Client
)

func NewEtcdClient() (*clientv3.Client, error) {
	config := clientv3.Config{
		Endpoints:            endpoints,            //  Etcd 实例列表 (分布式集群)
		AutoSyncInterval:     AutoSyncInterVal,     //  AutoSyncInterval 自动同步定时器
		DialTimeout:          dialTimeout,          //  建立连接的超时 时间
		DialKeepAliveTime:    DialKeepAliveTime,    //  心跳检测时间 客户端会按照配置的时间对etcd服务进行ping 轮询
		DialKeepAliveTimeout: DialKeepAliveTimeout, // 心跳检测超时时间
		MaxCallSendMsgSize:   MessageSize,          //  最大发送消息的限制
		MaxCallRecvMsgSize:   MessageSize,          // 最大接受响应限制
		Username:             "",                   // 用户名
		Password:             "",                   // 密码
		RejectOldCluster:     false,                // 时候拒绝过时的集群
	}
	Etcd, err := clientv3.New(config)
	return Etcd, err
}
