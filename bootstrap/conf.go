package bootstrap

import (
	"encoding/json"
	"github.com/Xhofe/alist/conf"
	"github.com/Xhofe/alist/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

// InitConf init config
func InitConf() {
	log.Infof("reading config file: %s", conf.ConfigFile)
	if !utils.Exists(conf.ConfigFile) {
		log.Infof("config file not exists, creating default config file")
		conf.Conf = conf.DefaultConfig()
		if !utils.WriteToJson(conf.ConfigFile, conf.Conf) {
			log.Fatalf("failed to create default config file")
		}
		return
	}
	config, err := ioutil.ReadFile(conf.ConfigFile)
	if err != nil {
		log.Fatalf("reading config file error:%s", err.Error())
	}
	conf.Conf = new(conf.Config)
	err = json.Unmarshal(config, conf.Conf)
	if err != nil {
		log.Fatalf("load config error: %s", err.Error())
	}
	log.Debugf("config:%+v", conf.Conf)
}
