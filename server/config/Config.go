package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type HttpServerConfig struct {
	Port    int    `yaml:"port"`
	Host    string `yaml:"host"`
	ApiPath string `yaml:"apipath"`
}

type HttpConfig struct {
	Server HttpServerConfig `yaml:"server"`
}

func init() {
	//instance = conf
	var path = flag.String("c", "", "yaml文件路径")
	flag.Parse()
	var file *os.File
	var err error
	if path == nil || len(*path) <= 0 {
		file, err = os.Open("app.yaml")
	} else {
		file, err = os.Open(*path)
	}
	if err != nil {
		panic(fmt.Sprintf("Error opening file:%v", err))
	}
	defer file.Close()

	// 创建解析器
	decoder := yaml.NewDecoder(file)
	// 解析 YAML 数据
	if instance == nil {
		instance = &Config{}
	}
	err = decoder.Decode(Get())
	if err != nil {
		panic(fmt.Sprintf("Error decoding YAML:%v", err))
	}
}

type IConfig interface {
	GetConfig() Config
}

type Config struct {
	HttpConfig HttpConfig `yaml:"http"`
}

var instance IConfig

func Get() IConfig {
	return instance
}

func Set(conf IConfig) {
	instance = conf
}

func (this *Config) GetConfig() Config {
	return *this
}
