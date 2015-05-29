# goproc
Process management API for Go.

## Project Status
This is a stub. Work in progress...

## Example
```go
package main

import "github.com/gsscoder/goproc"

func main() {
  processName := process.NameOf(1) // result: "init" (on Linux)
  processId := process.PidOf("launchd") // result: 1 (on OS X)
  count := process.Count() // result: int count of running processes
  pids := process.ListPids() // result: []int array with running pids
}
```

## Tests
Depending on function (``process.NameOf()`` for example) and platform type you may need run as root.
```sh
sudo go test
```
