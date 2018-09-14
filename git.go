package main

import "fmt"

type GitStatus struct {
	uncommittedFiles []string
}

func git_status() (*GitStatus, error) {
	return nil, fmt.Errorf("Failed to get status : not implemeted yet")
}
