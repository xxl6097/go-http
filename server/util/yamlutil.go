package util

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func ParseYaml(conf interface{}) {
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
	err = decoder.Decode(conf)
	if err != nil {
		panic(fmt.Sprintf("Error decoding YAML:%v", err))
	}
}
