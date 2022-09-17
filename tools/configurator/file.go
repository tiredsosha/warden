package configurator

import (
	"errors"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/tiredsosha/warden/tools/logger"
	"gopkg.in/yaml.v3"
)

type conf struct {
	Broker   string `yaml:"broker"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func getConf(file string, cnf interface{}) error {
	yamlFile, err := os.ReadFile(file)
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
			err = errors.New("config: " + strings.ToLower(field) + " field is emtpy/nonexist")
			break
		}
		err = nil
	}
	return err
}

func confFile() *conf {
	logger.Warn.Println("making a default config")
	confDef := conf{
		Broker:   "127.0.0.1",
		Username: "admin",
		Password: "password",
	}

	yamlData, _ := yaml.Marshal(confDef)
	err := ioutil.WriteFile("config.yaml", yamlData, 0644)
	if err != nil {
		logger.Error.Fatal("can't to write default conf into the file")
	}
	return &confDef
}

func ConfInit() *conf {
	cfg := &conf{}
	if err := getConf("config.yaml", cfg); err != nil {
		logger.Error.Println(err)
		cfg = confFile()
	}
	if err := validateConf(cfg); err != nil {
		logger.Error.Println(err)
		logger.Error.Fatal("EXITING")
	}
	return cfg
}
