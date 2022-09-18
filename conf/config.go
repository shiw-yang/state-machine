package conf

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Configure ServiceConfig
var logConf *logrus.Logger

type ServiceConfig struct {
	Env     string `yaml:"env"`
	App     string `yaml:"app"`
	Version string `yaml:"version"`
	Port    string `yaml:"port"`
}

func Init(configFile string) {
	// Read Config file
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	Configure.GetConfigure()

	// Logger Init
	level, err := logrus.ParseLevel(Configure.Env)
	if err != nil {
		panic(err)
	}
	logrus.SetLevel(level)
	logConf = logrus.New()
	logConf.Formatter.(*logrus.TextFormatter).DisableTimestamp = true
	logrus.SetFormatter(logConf.Formatter)
	logrus.Infof("%v %v Logger Start", Configure.App, Configure.Version)
}

func (s *ServiceConfig) GetConfigure() {
	err := viper.Unmarshal(s)
	if err != nil {
		panic(err)
	}
	s.Port = ":" + s.Port
}
