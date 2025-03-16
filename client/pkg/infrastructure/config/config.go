package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	SConfig   ServerConfig   `yaml:"server" env-required:"true"`
	DBConfig  DataBaseConfig `yaml:"db" env-required:"true"`
	KConfig   KafkaConfig    `yaml:"kafka" env-required:"true"`
	JWTSecret string         `yaml:"jwt_secret" env-required:"true"`
	Address   string         `yaml:"address" env-required:"true"`
}

type KafkaConfig struct {
	Port        string `yaml:"port" env-default:"9092"`
	Host        string `yaml:"host" env-default:"kafka"`
	UsersTopic  string `yaml:"users_topic" env-default:"users_topic"`
	GroupsTopic string `yaml:"groups_topic" env-default:"groups_topic"`
	GroupsId    string `yaml:"group_id" env-default:"users_group"`
}

type ServerConfig struct {
	Port string `yaml:"port" env-default:"8080"`
	Host string `yaml:"host" env-default:"localhost"`
}

type DataBaseConfig struct {
	Port     string `yaml:"port" env-default:"5435"`
	Host     string `yaml:"host" env-default:"localhost"`
	Username string `yaml:"username" env-required:"true"`
	DBName   string `yaml:"dbname" env-required:"true"`
	SSLMode  string `yaml:"ssl_mode" env-default:"disable"`
	Password string `yaml:"password" env-required:"true"`
}

func MustLoad() *Config {
	path := fetchConfigFlags()
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist" + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config" + err.Error())
	}

	return &cfg
}

// fetchConfigFlags получает путь до конфига либо из флага командной строки либо через переменную окружения
func fetchConfigFlags() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
