package utils

import "log"

func CriticalErrorCheck(errorCheck error) {

	if errorCheck != nil {
		log.Fatal(errorCheck)
	}

}
func ErrorCheck(errorCheck error) bool {

	if errorCheck != nil {
		log.Printf("ERROR: %s", errorCheck)
		return true
	}

	return false
}
