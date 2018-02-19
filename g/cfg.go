package g

import (
	"encoding/json"
	"github.com/toolkits/file"
	"sync"
	logs "github.com/Sirupsen/logrus"

)

var (
	ConfigFile string
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
)

type GlobalConfig struct {
	LogLevel 	bool 	`json:"log_level"`
	LogPath		string	`json:"log_path"`
	ChanSize	int		`json:"chan_size"`
	KafkaAddr 	string 	`json:"kafka_addr"`
	EtcdAddr  	string	`json:"etcd_addr"`
	EtcdKey		string	`json:"etcd_key"`
}

func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func InitConfig(cfg string) (err error){
	if cfg == "" {
		logs.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		logs.Fatalln("config file:", cfg, "is not existent")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		logs.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		logs.Fatalln("parse config file:", cfg, "fail:", err)
	}

	configLock.Lock()
	defer configLock.Unlock()
	config = &c
	logs.Println("read config file:", cfg, "successfully")
	return err
}