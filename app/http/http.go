package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"go_rtb/internal/configuration"
	"go_rtb/internal/router"
)

func main() {
	systemConfig, err := configuration.LoadConfig("config/setting.json")
	if err != nil {
		panic(err)
	}

	r := router.Init(systemConfig)

	fmt.Println("Start listening on port: ", systemConfig.App.AppPort)
	if err := endless.ListenAndServe(fmt.Sprintf(":%v", systemConfig.App.AppPort), r); err != nil {
		fmt.Println(err)
	}
}
