package main

import (
	"gopkg.in/src-d/go-git.v4"
)

type GitStatus struct {
	isClean bool
}

func git_status(path string) (*GitStatus, error) {

	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	w, err := r.Worktree()
	if err != nil {
		return nil, err
	}

	s, err := w.Status()
	if err != nil {
		return nil, err
	}

	return &GitStatus{isClean: s.IsClean()}, nil //fmt.Errorf("Failed to get status : not implemeted yet")
}
