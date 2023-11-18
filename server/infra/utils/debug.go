package utils

import (
	"encoding/json"
	"log"

	"vsoutlook.com/vsoutlook/infra/config"
)

func PrintJSON(notes string, obj any) {
	data, _ := json.Marshal(obj)
	if len(notes) > 0 {
		notes = notes + ": "
	}
	log.Printf("%s%s\n", notes, data)
}

func DevLogf(format string, v ...any) {
	if config.Dev {
		log.Printf(format, v...)
	}
}

func DevLogln(v ...any) {
	if config.Dev {
		log.Println(v...)
	}
}
