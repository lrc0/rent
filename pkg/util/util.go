package util

import (
	"os"
)

//Home .
func Home() string {
	var home string
	if home == "" {
		home = os.Getenv("GOPATH")
	}
	return home
}
