package helpers

import "log"

// LogError ...
func LogError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
