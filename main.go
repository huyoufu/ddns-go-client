package main

import (
	"flag"
	"fmt"
	"github.com/huyoufu/ddns-go-client/config"
	"github.com/huyoufu/ddns-go-client/config/file"
	"github.com/huyoufu/ddns-go-client/dns/aliyun"
	"github.com/huyoufu/ddns-go-client/logger"
	"github.com/huyoufu/ddns-go-client/util"
	"github.com/kardianos/service"
	"log"
	"os"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	// Do work here
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

var serviceAction = flag.String("s", "run", "安装为系统服务\n可选值 init install uninstall ")

func main() {
	flag.Parse()
	if util.DetectServiceAction(*serviceAction) == util.Init {
		//初始化配置文件
		if !file.ConfigFileExists() {
			//配置文件不存在
			os.WriteFile(file.GetConfigFilePath(), []byte(file.ConfigTemplate), 0644)
		}
		return
	}
	if util.DetectServiceAction(*serviceAction) == util.Run {
		if !file.ConfigFileExists() {
			//配置文件不存在
			log.Fatal("没有找到服务配置文件,请先 ddns-go-client -s init 初始化配置文件")
		}
		//真正开始运行程序
		logger.Log.Info("开启ddns自动更新客户端")
		runService()
		return
	} else {
		//处理服务安装相关的内容
		handleServiceAction()
	}
}

func runService() {
	//真正的服务

	//读取配置文件
	ddnsConfig := file.ConfigFromJsonFile()

	//获取激活的信息
	if ddnsConfig.DetectDomainNameServiceProvider() == config.DnsAli {
		//如果是ali的
		aliyun.DNSTask(ddnsConfig)
	} else {
		//其他的暂不支持
		log.Fatal("非阿里系暂不支持~~~")
	}
	util.LoopNoOps()
}
func handleServiceAction() {
	p := &program{}
	svcConfig := &service.Config{
		Name:        "ddns-go-client",
		DisplayName: "ddns-go-client-service",
		Description: "a simple program to report ip to ddns server",
		Option:      map[string]interface{}{"RunAtLoad": true},
	}
	s, err := service.New(p, svcConfig)
	if err != nil {
		fmt.Println(err.Error())
	}
	if util.DetectServiceAction(*serviceAction) == util.Install {
		logger.Log.Info("installing ddns-go-client service~~~")
		err := s.Install()
		if err != nil {
			fmt.Println(err)
			return
		}
		logger.Log.Info("install ddns-go-client service success!")
		return
	}
	if util.DetectServiceAction(*serviceAction) == util.Uninstall {
		logger.Log.Info("uninstalling ddns-go-client service~~~")
		err := s.Uninstall()
		if err != nil {
			fmt.Println(err)
			return
		}
		logger.Log.Info("uninstall ddns-go-client service success!")
		return
	}
}
