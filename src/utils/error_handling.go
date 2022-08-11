package utils

import "log"

// CriticalErrorCheck it's function to handle Critical errors
func CriticalErrorCheck(errorCheck error) {

	if errorCheck != nil {
		log.Fatal(errorCheck)
	}

}

// ErrorCheck it's function to handle errors
func ErrorCheck(errorCheck error) bool {

	if errorCheck != nil {
		log.Printf("ERROR: %s", errorCheck)
		return true
	}

	return false
}
