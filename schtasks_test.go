package schtasks

import (
	"testing"
	"strings"
)

var taskName = "Schtasks Test"

func TestRunEveryMinutes(t *testing.T) {
	_, err := RunEveryMinutes(taskName, 5, "dir")
	if err != nil {
		t.Error(err)
	}
}

func TestGet(t *testing.T) {
	st, err := Get(taskName)
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(st.TaskName, taskName) {
		t.Error("Error parsing task")
	}
}

func TestForceDelete(t *testing.T) {
	_, err := ForceDelete(taskName)
	if err != nil {
		t.Error(err)
	}
}


