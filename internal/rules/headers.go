package rules

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func HeaderRules() map[string]int {

	viper.SetConfigType("yaml")
	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Erro on read config file: %v", err)
	}

	headers := viper.GetStringMap("rules.header")

	headersMap := make(map[string]int)

	for headerName := range headers {
		var headerThreshold RulesThreshold

		if err := viper.UnmarshalKey(fmt.Sprintf("rules.header.%s", headerName), &headerThreshold); err != nil {
			log.Fatalf("Error on unmarshal header %s: %v", headerName, err)
		}

		headersMap[headerName] = headerThreshold.Threshold
	}

	return headersMap
}
