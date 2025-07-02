package config

import (
	"net"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Conf struct {
	ApplicationConf `yaml:"application_conf"`
	DatabaseConf    `yaml:"database_conf"`
}

type DatabaseConf struct {
	DiverType   string `yaml:"driver_type"`
	Destination string `yaml:"destination"`
}

type ApplicationConf struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var conf Conf
var once sync.Once

func Setup(confPath string) {
	once.Do(func() {
		data, err := os.ReadFile(confPath)
		if err != nil {
			logrus.Fatal(err)
		}
		conf = Conf{}
		err = yaml.Unmarshal(data, &conf)
		if err != nil {
			logrus.Fatal(err)
		}
		printConf()
	})
}

func GetConf() *Conf {
	return &conf
}

func printConf() {
	addr := net.JoinHostPort(conf.Host, conf.Port)
	logrus.Infof("server listening %s, database configuration: %s, %s", addr, conf.DiverType, conf.Destination)
}
