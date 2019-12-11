package config

import (
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/kaepa3/KobaiKyobai/record"
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
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, false
	}
	return &config, true
}
func ReadDynamicConfig(path string) (*DynamicConfig, bool) {
	var config DynamicConfig
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, false
	}
	return &config, true
}

func ReadAllConfig(staticPath, dynamicPath string) (Config, DynamicConfig, string) {
	// readconfig
	var conf Config
	if itf, ok := ReadConfig(staticPath); ok {
		static := *itf
		conf, _ = static.(Config)
	} else {
		conf = Config{}
	}
	var dynamicConf DynamicConfig
	if itf, ok := ReadDynamicConfig(dynamicPath); ok {
		dynamicConf = *itf
	} else {
		dynamicConf = DynamicConfig{}
	}
	return conf, dynamicConf, ""
}

func WriteConfig(path string, buffer interface{}) bool {
	file, err := os.Create(path)
	if err == nil {
		if err := toml.NewEncoder(file).Encode(buffer); err != nil {
			log.Fatal(err)
			return true
		}
	}
	return false
}
