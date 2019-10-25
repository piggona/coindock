package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// BinanceConf binance configuration
type BinanceConf struct {
	BaseEndPoint string            `json:"baseEndPoint"`
	SecretKey    string            `json:"secretKey"`
	APIKey       string            `json:"apiKey"`
	CallWorker   BinanceWorkerConf `json:"callWorker"`
}

// BinanceWorkerConf binance worker amount configuration
type BinanceWorkerConf struct {
	None int `json:"None"`
	Half int `json:"Half"`
	Full int `json:"Full"`
}

var (
	// Binance configuration struct
	Binance *BinanceConf
)

func init() {
	viper.SetDefault("baseEndPoint", "https://api.binance.com")
	// viper.SetDefault("callWorker", BinanceWorkerConf{"None": 5, "Half": 5, "Full": 5})

	Binance = getBinance()
}

func getBinance() *BinanceConf {
	viper.SetConfigName("binance")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/Users/haohao/Documents/go/src/coindock/config/")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Errorf("Error occurs: viper.ConfigFileNotFoundError")
		} else {
			// Config file was found but another error was produced
			fmt.Errorf("Error occurs in viper.ReadConfig: %v\n", err)
		}
	}
	conf := &BinanceConf{}
	viper.Unmarshal(conf)
	return conf
}
