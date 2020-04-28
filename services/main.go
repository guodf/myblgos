package main

import (
	"fmt"
	"github.com/spf13/viper"
	_ "myblogs/conf"
	_ "myblogs/data"
	"myblogs/router"
)

func main() {
	r := *router.InitRouter()
	r.Run(fmt.Sprintf(":%s", viper.GetString("port")))

	//fresh
}
