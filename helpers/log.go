package helpers

/*
 * File: log.go
 * File Created: Monday, 11th May 2020
 * Author: Sainesh Mamgain (saineshmamgain@gmail.com)
 */

import "log"

// LogError ...
func LogError(msg string, err error) {
	if err != nil {
		log.Fatal(msg+" ", err)
	}
}
