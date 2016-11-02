# schtasks

Wrapper around schtasks.exe

[TechNet Schtasks](https://technet.microsoft.com/en-us/library/cc725744\(v=ws.11\).aspx)

## To run a task with system permissions

Requires an Admin Command Prompt

    schtasks /create /tn "Schtasks Test" /tr dir /sc monthly /d 15 /ru System
 
## Testing

TODO Use golang 1.7 to [specify test order?](http://stackoverflow.com/a/39734200/639133)

    go test -run RunEveryMinutes
    
    go test -run Get
    
    go test -run ForceDelete
    