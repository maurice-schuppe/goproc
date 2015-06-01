// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

// +build !darwin linux

package process

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const procFsRoot = "/proc"
const procFsPidPath = "/proc/%d/%s"
const procFsPath = "/proc/%s"

func procFsOpenPid(pid int, name string) (*os.File, error) {
	return os.Open(fmt.Sprintf(procFsPidPath, pid, name))
}

func procFsOpen(name string) (*os.File, error) {
	return os.Open(fmt.Sprintf(procFsPath, name))
}

func procFsParseStatusItems(status *os.File, keys []string) []string {
	values := make([]string, len(keys))
	i := 0

	scanner := bufio.NewScanner(status)
	for scanner.Scan() {
		if parts := strings.Split(scanner.Text(), ":"); len(parts) == 2 {
			if currkey := strings.TrimSpace(parts[0]); strElemIndexOf(currkey, keys) >= 0 {
				values[i] = strings.TrimSpace(parts[1])
				i++
			}
		}
	}
	//if err := scanner.Err(); err != nil {
	//}

	return values
}

func procFsListPids() []int {
	items, err := ioutil.ReadDir(procFsRoot)
	if err != nil {
		return []int{}
	}

	pids := make([]int, len(items))
	pids[0] = -1 // mark value
	i := 0
	for _, item := range items {
		pid := procFsTryNameToPid(item.Name())
		if pid > 0 {
			pids[i] = pid
			i++
		}
	}

	if pids[0] > 0 { // some pid added
		return pids[0:i]
	}
	return []int{}
}

func procFsCpuTimeTotal() int {
	stat, err := procFsOpen("stat")
	if err != nil {
		return -1
	}
	defer stat.Close()
	scanner := bufio.NewScanner(stat)
	if !scanner.Scan() {
		return -1
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return -1
	}
        if strings.TrimSpace(parts[0]) != "cpu" {
		return -1
	}
	total := 0
	for _,cpuTime := range parts[1:] {
		partial, err := strconv.Atoi(cpuTime)
		if err != nil {
			return -1
		}
		total += partial
	}
	return total
}

func procFsTryNameToPid(name string) int {
	pid, err := strconv.Atoi(name)
	if err != nil || pid <= 0 {
		return -1
	}

	return pid
}
