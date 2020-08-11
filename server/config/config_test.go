package config

import (
	"log"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	path := "E:\\go\\src\\yomo\\yomo.toml"
	conf,err := Load(path)
	if err != nil{
		t.Error(err)
		return
	}

	log.Println(conf)
}