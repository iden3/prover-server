package configs

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

// Config structure represent yaml config for prover serever
type Config struct {
	Server struct {
		Port int    `mapstructure:"port"`
		Host string `mapstructure:"host"`
	} `mapstructure:"server"`
	Prover ProverConfig `mapstructure:"prover"`
	Log    struct {
		Level string `json:"level"`
	}
}

// ProverConfig contains only base path to circuits folder
type ProverConfig struct {
	CircuitsBasePath string `mapstructure:"circuitsBasePath"`
	UseSnarkjs       bool   `mapstructure:"useSnarkjs"`
}

// ReadConfigFromFile parse config file
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
