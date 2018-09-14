package main

import (
	"fmt"
	"github.com/ogier/pflag"
)

var (
	build string
)

func init() {
	pflag.StringVarP(&build, "build", "b", "", "command to run the build before integrating")
}

/*
takes snapshot of current code, does pull -r, runs build, pushes to central repo

if there are uncommited files
  if there is no message
    prompt for message or quit
  commit files with message
rebase
run build
push

*/
func integrate(path string) error {
	var status, err = git_status(path)
	if err != nil {
		return err
	}

	if !status.isClean {
		return fmt.Errorf("uncommitted files. Take a snapshot before integrating")
	}

	fmt.Print("integrated!!!")

	return nil
}
