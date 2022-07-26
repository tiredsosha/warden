package configurator

import (
	"errors"
	"io/ioutil"
	"log"
	"reflect"

	"gopkg.in/yaml.v3"
)

type conf struct {
	Broker   string `yaml:"broker"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func getConf(file string, cnf interface{}) error {
	yamlFile, err := ioutil.ReadFile(file)
	if err == nil {
		err = yaml.Unmarshal(yamlFile, cnf)
	}
	return err
}

func validateConf(cfg *conf) error {
	var err error
	v := reflect.ValueOf(*cfg)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := typeOfS.Field(i).Name
		value := v.Field(i).Interface()
		if value == "" {
			err = errors.New("config: " + field + " field is emtpy/nonexist")
			break
		}
		err = nil
	}
	return err
}

func ConfInit() *conf {
	cfg := &conf{}
	if err := getConf("config.yaml", cfg); err != nil {
		log.Fatal(err)
	}
	if err := validateConf(cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
