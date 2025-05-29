package config

import "github.com/spf13/viper"

type Config struct {
	Typecho     *Typecho     `yaml:"typecho"`
	Halo        *Halo        `yaml:"halo"`
	FileManager *FileManager `yaml:"fileManager"`
}

type Typecho struct {
	BaseUrl string `yaml:"baseUrl"`
	Type    string `yaml:"type"`
	Dsn     string `yaml:"dsn"`
}

type Halo struct {
	Host       string `yaml:"host"`
	Schema     string `yaml:"schema"`
	Token      string `yaml:"token"`
	Debug      bool   `yaml:"debug"`
	PolicyName string `yaml:"policyName"`
	GroupName  string `yaml:"groupName"`
}

type FileManager struct {
	Dir string `yaml:"dir"`
}

func New() (*Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./")
	v.AddConfigPath("../../_tmp/config/")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	c := new(Config)
	err := v.Unmarshal(c)

	return c, err
}
