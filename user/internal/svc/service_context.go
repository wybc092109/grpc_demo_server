package svc

import (
	"encoding/json"
	"grpc_demo_server/user/internal/config"
	"os"

	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 读取 etc 配置
	ehost := os.Getenv("ETCD_HOST")
	if ehost != "" {
		ehosts := []string{}
		err := json.Unmarshal([]byte(ehost), &ehosts)
		if err != nil {
			logx.Error("etcd hosts unmarshal failed", err)
		} else {
			c.Etcd.Hosts = ehosts
		}
	}
	eKey := os.Getenv("ETCD_KEY")
	if eKey != "" {
		c.Etcd.Key = eKey
	}
	eUser := os.Getenv("ETCD_USER")
	if eUser != "" {
		c.Etcd.User = eUser
	}
	ePass := os.Getenv("ETCD_PASS")
	if ePass != "" {
		c.Etcd.Pass = ePass
	}
	return &ServiceContext{
		Config: c,
	}
}
