// Package schtasks is a wrapper around schtasks.exe
package schtasks

import (
	"os/exec"
	"strconv"
	"github.com/gocarina/gocsv"
	"bytes"
	"errors"
	"fmt"
)

// TODO Use %windir% to construct path?
// For now just export the variable to allow override
var SchtasksPath = "c:\\Windows\\System32\\schtasks.exe"

type ScheduledTask struct {
	HostName                     string `csv:"HostName"`
	TaskName                     string `csv:"TaskName"`
	NextRunTime                  string `csv:"Next Run Time"`
	Status                       string `csv:"Status"`
	LogonMode                    string `csv:"Logon Mode"`
	LastRunTime                  string `csv:"Last Run Time"`
	LastResult                   string `csv:"Last Result"`
	Author                       string `csv:"Author"`
	TaskToRun                    string `csv:"Task To Run"`
	StartIn                      string `csv:"Start In"`
	Comment                      string `csv:"Comment"`
	ScheduledTaskState           string `csv:"Scheduled Task State"`
	IdleTime                     string `csv:"Idle Time"`
	PowerManagement              string `csv:"Power Management"`
	RunAsUser                    string `csv:"Run As User"`
	DeleteTaskIdNotRescheduled   string `csv:"Delete Task If Not Rescheduled"`
	StopTaskIfRunsXHoursAndXMins string `csv:"Stop Task If Runs X Hours and X Mins"`
	Schedule                     string `csv:"Schedule"`
	ScheduleType                 string `csv:"Schedule Type"`
	StartTime                    string `csv:"Start Time"`
	StartDate                    string `csv:"Start Date"`
	EndDate                      string `csv:"End Date"`
	Days                         string `csv:"Days"`
	Months                       string `csv:"Months"`
	RepeatEvery                  string `csv:"Repeat: Every"`
	RepeatUntilTime              string `csv:"Repeat: Until: Time"`
	RepeatUntilDuration          string `csv:"Repeat: Until: Duration"`
	RepeatStopIfStillRunning     string `csv:"Repeat: Stop If Still Running"`
}

// ForceDelete deletes a task even if it is currently running
func ForceDelete(taskName string) ([]byte, error) {
	out, err := exec.Command(
		SchtasksPath, "/delete", "/f",
		"/tn", taskName).Output()
	return out, err
}

// Get the specified task name
func Get(taskName string) (ScheduledTask, error) {
	out, err := exec.Command(
		SchtasksPath, "/query",
		"/v", "/tn", taskName,
		"/fo", "csv").Output()
	if err != nil {
		return ScheduledTask{}, errors.New(
			fmt.Sprintf("Task '%s' not found", taskName))
	}

	tasks := []ScheduledTask{}
 	err = gocsv.Unmarshal(bytes.NewReader(out), &tasks)
	if err != nil {
        return ScheduledTask{}, err
    }

	return tasks[0], nil
}

// RunEveryMinutes schedules a task to run every m minutes
func RunEveryMinutes(taskName string, m int, command string) ([]byte, error) {
	// Create or update a scheduled task
	out, err := exec.Command(
		SchtasksPath, "/create",
		"/sc", "minute",
		"/mo", strconv.Itoa(m),
		"/f",
		"/tn", taskName,
		"/tr", command).Output()
	return out, err
}
