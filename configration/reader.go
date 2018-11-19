package configration

import (
	"encoding/json"
	"os"
	"fmt"
	"io/ioutil"
)

type JSONFileStream interface {
	getJson() *struct{}
}

func readConfigFile(configFilePath string, conf JSONFileStream) JSONFileStream {
	jsonFile, err := os.Open(configFilePath)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)

	/* Parse the config.json to the `conf` variable */
	json.Unmarshal(byteValue, &conf)

	return conf
}
