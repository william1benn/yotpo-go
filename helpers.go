package yotpo

import "log"

func CheckErrorFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
