package main

//
//import (
//	"flag"
//	"fmt"
//	"github.com/spf13/viper"
//	"github.com/romapopov1212/currency-service/gateway/internal/config"
//	"github.com/romapopov1212/currency-service/gateway/internal/migrations"
//	"log"
//)
//
//func main() {
//	if err := run(); err != nil {
//		log.Fatal(err.Error())
//	}
//}
//
//func run() error {
//	configPath := flag.String("config", "./config", "path to the config file")
//	flag.Parse()
//
//	//cfg, err := loadConfig(*configPath)
//	//if err != nil {return fmt.Errorf("load config: %w", err)}
//	//
//}
//
//type appConfig struct {
//	Database config.DatabaseConfig `mapstructure:"database"`
//}
//
//func loadConfig(path string) (appConfig, error) {
//	var config appConfig
//
//	viper.SetConfigFile(path)
//
//	if err := viper.ReadInConfig(); err != nil {
//		return
//	}
//}
