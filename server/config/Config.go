package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Yml struct {
	Database   DatabaseConfig `yaml:"database"`
	HttpConfig HttpConfig     `yaml:"http"`
	LogConfig  LogConfig      `yaml:"log"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type HttpServerConfig struct {
	Port    int    `yaml:"port"`
	Host    string `yaml:"host"`
	ApiPath string `yaml:"apipath"`
}

type HttpConfig struct {
	Server HttpServerConfig `yaml:"server"`
}
type LogConfig struct {
	Port     int    `yaml:"port"`
	UserName string `yaml:"username"`
	PassWord string `yaml:"password"`
}

var yml Yml

func init() {
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
	err = decoder.Decode(&yml)
	if err != nil {
		panic(fmt.Sprintf("Error decoding YAML:%v", err))
	}
}

func GetYaml() Yml {
	return yml
}
