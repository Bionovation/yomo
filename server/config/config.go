package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"yomo/server/models"
)

//const IMGPATH  = "F:\\Workdatas\\红细胞标注\\5\\5\\"

type Config struct {
	Name string
	Port int
	Proxy  ConfigProxy
	Workspaces []models.Workspace
}

// frp代理参数
type ConfigProxy struct {
	ServerAddr string
	ServerPort int
}

var defaultConfig = Config{
	Port: 8080,
	Name: "marktool",
	Workspaces: nil,
}

func Get() *Config {
	return &defaultConfig
}

func Load(location string) (*Config, error) {
	Ensure(location)
	conf := &defaultConfig

	if _, err := toml.DecodeFile(location, conf); err != nil {
		return nil, fmt.Errorf("error opening local config file (%v): %v ", location, err)
	}

	// 读取classnames
	for i, _ := range conf.Workspaces{
		if err := conf.Workspaces[i].LoadClassNames(); err != nil{
			log.Println("load classnames failed, err: " + err.Error())
		}
	}

	return conf, nil
}

// 保证配置文件一定存在
func Ensure(location string) {
	_, err := os.Stat(location)
	if os.IsNotExist(err) == false {
		return
	}

	defaultConfig := `
[common]
name = "marktool"
port = 8080

[proxy]
ServerAddr = "slide.bionovationimc.com"
ServerPort = 7000

[[workspaces]]
name = "workspace01"
folder = "./img"
classdef = "./obj.names"
`

	ioutil.WriteFile(location, []byte(defaultConfig), os.ModePerm)
}

// 根据workspace名称，返回Workspace对象
func FindWsByName(name string) *models.Workspace {
	for _, ws := range defaultConfig.Workspaces{
		if ws.Name == strings.TrimSpace(name){
			return &ws
		}
	}
	return nil
}