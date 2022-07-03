package setting

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"log"
)

type Configuration struct {
	MysqlDatabase string `yaml:"mysql_database"`
	Port          uint   `yaml:"port"`
	Host          string `yaml:"host"`
	SignupEnabled bool   `yaml:"signup_enabled"`
	PageSize      uint   `yaml:"page_size"`
	MysqlUsername string `yaml:"mysql_username"`
	MysqlPassword string `yaml:"mysql_password"`
	MysqlHost     string `yaml:"mysql_host"`
	MysqlPort     uint   `yaml:"mysql_port"`
	JwtKey        []byte `yaml:"jwt_key"`
}

var Config *Configuration

func loadConfiguration() error {
	date, err := ioutil.ReadFile("conf/conf.yaml")
	if err != nil {
		return err
	}
	var config Configuration
	err = yaml.Unmarshal(date, &config)
	if err != nil {
		return err
	}
	Config = &config
	return nil
}

func init() {
	err := loadConfiguration()
	if err != nil {
		log.Fatal(err)
	}
}
