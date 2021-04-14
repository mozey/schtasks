# schtasks

Wrapper around schtasks.exe

[Schtasks Docs](https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/schtasks)


## Usage

Run a command every x minutes

```
taskName := "My Command"
interval := 5
exe := "my_command.exe"
_, err := RunEveryMinutes(taskName, interval, exe)
```

Run a command once x minutes into the future
(plus or minus 30 seconds)

```
taskName := "My Command"
at := 1 // run in one minute
exe := "my_command.exe"
_, err := RunAtMinutes(taskName, at, exe)
```


## Schtasks Examples

To run a task with system permissions.
Requires an Admin Command Prompt

Run the task every minute

    schtasks /create /sc minute /mo 1 /f /tn "Schtasks Test" /tr dir /ru System

Run the task once at HH:MM

    schtasks /create /sc once /mt 14:00 /f /tn "Schtasks Test" /tr dir /ru System
    
Query task by name

    schtasks /query /v /tn "Schtasks Test" /fo list

Delete task by name

    schtasks /delete /f /tn "Schtasks Test"

 
## Dev

Run tests

    git clone https://github.com/mozey/schtasks

    cd schtasks

    gotest -v ./...
    
