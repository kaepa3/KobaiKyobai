package config

import (
	"time"

	"github.com/BurntSushi/toml"
	"github.com/kaepa3/myproj/record"
)

// Config this app config
type Config struct {
	LastDate time.Time
}

type DynamicConfig struct {
	BeforeNewestRecord record.Record
}

func ReadConfig(path string) (*interface{}, bool) {
	var config interface{}
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		return nil, false
	}
	return &config, true
}

func ReadAllConfig(staticPath, dynamicPath string) (Config, DynamicConfig, string) {
	// readconfig
	var conf Config
	var dynamicConf DynamicConfig
	itf, _ := ReadConfig("config.toml")
	static := *itf
	conf, _ = static.(Config)
	itf, _ = ReadConfig("dynamic.toml")
	dynamic := *itf
	dynamicConf, _ = dynamic.(DynamicConfig)
	return conf, dynamicConf, ""
}
