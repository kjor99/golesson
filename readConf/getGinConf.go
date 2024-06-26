package readConf

import "encoding/json"

type AutoGenerated struct {
	Gin Gin `json:"Gin"`
}
type Gin struct {
	Port string `json:"port"`
}

func GetGinConf(url string) string {
	file := GetJsonConf(url)
	conf := AutoGenerated{}
	json.NewDecoder(file).Decode(&conf)
	port := conf.Gin.Port
	return port
}
