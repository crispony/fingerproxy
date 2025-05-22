package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/gospider007/proxy"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Proxy string `yaml:"proxy"`
	Spec  string `yaml:"spec"`
}

// goreleaser build --snapshot --clean
func main() {
	path := flag.String("config", "", "yaml配置文件路径")
	flag.Parse() // 必须调用它，解析传参
	yamlPath := *path
	var config Config
	log.Print("yamlPath: ", yamlPath)
	if yamlPath != "" { // 读取 YAML 配置文件
		data, err := os.ReadFile(yamlPath)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		// 解析 YAML 数据
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	}
	cli, err := proxy.NewClient(context.TODO(), proxy.ClientOption{
		Addr:  "127.0.0.1:8080",
		Proxy: config.Proxy,
		Spec:  config.Spec,
		// Debug: true,
	})
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Print("listening on: ", cli.Addr())
	log.Print(cli.Run())
}
