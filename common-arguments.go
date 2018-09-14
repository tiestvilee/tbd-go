package main

import (
	"github.com/ogier/pflag"
)

var (
	description string
)

func init() {
	pflag.StringVarP(&description, "description", "d", "", "Describes the change")
}
