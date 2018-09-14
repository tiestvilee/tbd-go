package main

import (
	"github.com/ogier/pflag"
)

func main() {
	pflag.Parse()
	switch pflag.Arg(0) {
	case "snapshot":
		snapshot()
	case "integrate":
		integrate()
	case "drop":
		drop()
	}
}

func init() {
}

func drop() {

}
