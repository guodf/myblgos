package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func init() {
	// 允许viper从os.Environ()中读取
	viper.AutomaticEnv()
	mode := "debug"
	if appModel := viper.GetString("APP_MODE"); appModel != "" {
		mode = appModel
	}
	log.Println(mode)

	viper.SetConfigFile(fmt.Sprintf("app_data/env/%s.json", mode))
	// 初始配置文件
	e := viper.ReadInConfig()
	if e != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", e))
	}
	log.Println(viper.GetString("uploads"))
	log.Println(viper.GetString("db_file_path"))

}
