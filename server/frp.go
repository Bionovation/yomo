package server

import (
	"fmt"
	"math/rand"
	"time"

	conf "yomo/server/config"

	_ "github.com/fatedier/frp/assets/frpc/statik"
	"github.com/fatedier/frp/client"
	"github.com/fatedier/frp/models/config"
	"github.com/fatedier/golib/crypto"
)

type GetTcPortResp struct {
	Port int `json:"port"`
}

func frpLogin() error {
	crypto.DefaultSalt = "frp"
	rand.Seed(time.Now().UnixNano())

	cf := conf.Get()
	cfg := config.GetDefaultClientConf()
	cfg.ServerAddr = cf.Proxy.ServerAddr
	cfg.ServerPort = cf.Proxy.ServerPort

	httpCfg := config.HttpProxyConf{}
	httpCfg.BaseProxyConf.ProxyName = cf.Name
	httpCfg.BaseProxyConf.ProxyType = "http"
	httpCfg.LocalIp = "127.0.0.1"
	httpCfg.LocalPort = cf.Port

	httpCfg.CustomDomains = []string{fmt.Sprintf("%v.%v", cf.Name, cf.Proxy.ServerAddr)}
	httpCfg.Locations = []string{""}

	visitorCfgs := make(map[string]config.VisitorConf)
	pxyCfgs := make(map[string]config.ProxyConf)

	pxyCfgs[cf.Name] = &httpCfg


	svr, err := client.NewService(cfg, pxyCfgs, visitorCfgs, "")
	if err != nil {
		return err
	}

	return svr.Run()
}
