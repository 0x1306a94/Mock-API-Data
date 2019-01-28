package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	MockAddr      string `yaml:"MockAddr"`
	MockPort      int64  `yaml:"MockPort"`
	DashboardAddr string `yaml:"DashboardAddr"`
	DashboardPort int64  `yaml:"DashboardPort"`
	DBPath        string `yaml:"DBPath"`
}

func Load(path string) (*Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}
	ext := guessFileType(path)
	if ext == "" {
		return nil, errors.New("cant not load" + path + " config")
	}

	yamlS, readErr := ioutil.ReadFile(path)
	if readErr != nil {
		return nil, readErr
	}
	var c Config
	// yaml解析的时候c.data如果没有被初始化，会自动为你做初始化
	err := yaml.Unmarshal(yamlS, &c)
	if err != nil {
		return nil, errors.New("can not parse " + path + " config")
	}
	return &c, nil
}

//判断配置文件名是否为yaml格式
func guessFileType(path string) string {
	s := strings.Split(path, ".")
	ext := s[len(s)-1]
	switch ext {
	case "yaml", "yml":
		return "yaml"
	}
	return ""
}
