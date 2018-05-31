package conf

import (
	"errors"
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

var (
	Conf              config
	defaultConfigFile = "./conf/conf.toml"
)

type config struct {
	Zhenai Zhenai `toml:"zhenai"`
}

type Zhenai struct {
	ElasticUrl   string `toml:"elastic_url"`
	ElasticIndex string `toml:"elastic_index"`
	ElasticType  string `toml:"elastic_type"`
}

func InitConfig() (err error) {
	configBytes, err := ioutil.ReadFile(defaultConfigFile)
	if err != nil {
		return errors.New("config load err:" + err.Error())
	}
	_, err = toml.Decode(string(configBytes), &Conf)
	if err != nil {
		return errors.New("config decode err:" + err.Error())
	}
	return nil
}
