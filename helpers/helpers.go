package helpers

import "log"

func CheckNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}

}
