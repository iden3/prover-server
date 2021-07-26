package configs

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Server struct {
		Port int    `mapstructure:"port"`
		Host string `mapstructure:"host"`
	} `mapstructure:"server"`
	Prover ProverConfig `mapstructure:"prover"`
}

type ProverConfig struct {
	CircuitsBasePath string `mapstructure:"circuitsBasePath"`
}

func ReadConfigFromFile(path string) (*Config, error) {

	viper.AddConfigPath("./configs")
	viper.SetConfigName(path)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "Error reading config file")
	}

	config := &Config{}

	err = viper.Unmarshal(config)
	if err != nil {
		return nil, errors.Wrap(err, "Error parsing config file")
	}

	return config, nil
}
