package utils

import (
	"encoding/json"
	"log"
)

func PrettyPrint(data interface{}) {
	var p []byte
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("++++++++++++++++++++++++++++++++++++++++++++++")
	log.Printf("%s \n", p)
	log.Println("++++++++++++++++++++++++++++++++++++++++++++++")
}
