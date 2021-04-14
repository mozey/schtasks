// Package schtasks is a wrapper around schtasks.exe
package schtasks

import (
	"bytes"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/pkg/errors"
	"os/exec"
	"strconv"
	"time"
)

// Path to schtasks.exe, caller can override according to %windir% if required
var Path = "c:\\Windows\\System32\\schtasks.exe"

// PaddingSeconds for RunAtMinutes
var PaddingSeconds = 5

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
		Path, "/delete", "/f",
		"/tn", taskName).Output()
	return out, errors.WithStack(err)
}

// Get the specified task name
func Get(taskName string) (ScheduledTask, error) {
	out, err := exec.Command(
		Path, "/query",
		"/v", "/tn", taskName,
		"/fo", "csv").Output()
	if err != nil {
		return ScheduledTask{}, errors.WithStack(
			errors.Wrap(err, fmt.Sprintf("task %s not found", taskName)))
	}

	var tasks []ScheduledTask
	err = gocsv.Unmarshal(bytes.NewReader(out), &tasks)
	if err != nil {
		return ScheduledTask{}, errors.WithStack(err)
	}

	if len(tasks) == 0 {
		return ScheduledTask{}, errors.WithStack(
			fmt.Errorf("task %s not found", taskName))
	}

	return tasks[0], nil
}

// RunEveryMinutes schedules a task to run every m minutes
func RunEveryMinutes(taskName string, m int, command string) ([]byte, error) {
	// Create or update a scheduled task
	out, err := exec.Command(
		Path, "/create",
		"/sc", "minute",
		"/mo", strconv.Itoa(m),
		"/f",
		"/tn", taskName,
		"/tr", command,
		"/ru", "SYSTEM").Output()
	if err != nil {
		return out, errors.WithStack(err)
	}

	return out, nil
}

// TimeAtMinutes adds m minutes to current time,
// and returns the time string in format HH:MM.
// Adds PaddingSeconds to make sure the
// scheduled task before time at.
// WARNING The correct value to pass in is probably time.Now().Local()
func TimeAtMinutes(now time.Time, m int) (t string) {
	timeAt := now.Add(time.Minute*time.Duration(m) +
		time.Second*time.Duration(PaddingSeconds))
	return timeAt.Format("15:04")
}

// RunAtMinutes schedules a task to run once in m minutes (plus or minus 30s)
func RunAtMinutes(taskName string, m int, command string) ([]byte, error) {
	// Create or update a scheduled task
	out, err := exec.Command(
		Path, "/create",
		"/sc", "once",
		"/mt", TimeAtMinutes(time.Now().Local(), m),
		"/f",
		"/tn", taskName,
		"/tr", command,
		"/ru", "SYSTEM").Output()
	if err != nil {
		return out, errors.WithStack(err)
	}

	return out, nil
}
