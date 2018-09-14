package main

import (
	"github.com/ogier/pflag"
)

func main() {
	pflag.Parse()

	var err error
	switch pflag.Arg(0) {
	case "snapshot":
		snapshot()
	case "integrate":
		err = integrate()
	case "drop":
		drop()
	}

	if err != nil {
		print(err.Error())
	}
}

func init() {
}

func drop() {

}
