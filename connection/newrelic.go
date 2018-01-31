package connection

import (
	"github.com/newrelic/go-agent"
)

type NewRelicConfig struct {
	AppName string
	License string
}

func CreateNewRelicApp(conf NewRelicConfig) newrelic.Application {
	newRelicConf := newrelic.NewConfig(
		conf.AppName,
		conf.License,
	)
	newRelicApp, err := newrelic.NewApplication(newRelicConf)
	if err != nil {
		panic(err)
	}
	return newRelicApp
}
