package pkg

import (
	"log"
)

func ErrorsHandler(err error, fatal bool) {
	if err != nil {
		if fatal {
			log.Fatal(err.Error())
		} else {
			log.Print(err.Error())
		}
	}
}
