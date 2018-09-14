package main

import (
	"github.com/ogier/pflag"
	"log"
	"os"
)

func main() {
	pflag.Parse()
	var err error

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	switch pflag.Arg(0) {
	case "snapshot":
		snapshot()
	case "integrate":
		err = integrate(cwd)
	case "drop":
		drop()
	}

	if err != nil {
		log.Fatal(err)
	}
}

func init() {
}

func drop() {

}
