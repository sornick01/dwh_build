package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const BatchSize = 1000

const (
	RAW = "RAW"
	ODS = "ODS"
	DDS = "DDS"
	DM  = "DM"
)

func GetStage() string {
	var res string
	switch os.Getenv("ETL_STAGE") {
	case RAW:
		res = RAW
	case ODS:
		res = ODS
	case DDS:
		res = DDS
	case DM:
		res = DM
	default:
		res = "unknown stage"
	}
	return res
}

type Config struct {
	ConfigDBURL string `mapstructure:"CONFIG_DB_URL"`
	SrcDBURL    string `mapstructure:"SRC_DB_URL"`
	DstDBURL    string `mapstructure:"DEST_DB_URL"`
	PathLog     string `mapstructure:"ETL_LOGPATH"`
	PortGRPC    string `mapstructure:"ETL_PORT_GRPC"`
}

func New(path string) *Config {
	configFileName := os.Getenv("CONFIG_NAME")
	viper.SetConfigName(configFileName)
	viper.SetConfigType("env")

	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read config: %s\n", err.Error())
		os.Exit(1)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to unmarshal config: %s\n", err.Error())
		os.Exit(1)
	}

	return &cfg
}
