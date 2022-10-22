package file

import (
	_ "embed"
	"encoding/json"
	"github.com/huyoufu/ddns-go-client/config"
	"github.com/huyoufu/ddns-go-client/logger"
	"github.com/huyoufu/ddns-go-client/util"
	"os"
)

//go:embed template.json
var ConfigTemplate string

const ConfigurationFilePath = "DDNS_CONFIG_FILE_PATH"

// GetConfigFilePath 获得配置文件路径
// 优先选择环境变量配置的位置
// 第二选择当前用户目录
// 第三选择进程的当前目录
func GetConfigFilePath() string {
	configFilePath := os.Getenv(ConfigurationFilePath)
	if configFilePath != "" {
		return configFilePath
	}
	return GetConfigFilePathDefault()
}

// GetConfigFilePathDefault 获得默认的配置文件路径
func GetConfigFilePathDefault() string {
	return util.GetCurrentDirectory() + string(os.PathSeparator) + ".ddns_go_config.json"
}
func ConfigFileExists() bool {
	_, err := os.Stat(GetConfigFilePath())
	return err == nil
}

func ConfigFromJsonFile() (c *config.DDNSConfig) {
	configFile := GetConfigFilePath()
	bytes, e := os.ReadFile(configFile)
	if e != nil {
		logger.Log.Fatal("打开文件配置文件失败!", e)
	}
	c = &config.DDNSConfig{}
	err := json.Unmarshal(bytes, c)
	if err != nil {
		logger.Log.Fatal(err)
	}
	return
}
