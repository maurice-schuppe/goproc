// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

package process

/*
#include <stdlib.h>
#include "libproc.h"
*/
import "C"
import "unsafe"
import "strings"

func nameOf(pid int) string {
	name := C.CString(strings.Repeat("\x00", 1024))
	namePtr := unsafe.Pointer(name)
	defer C.free(namePtr)
	nameLen := C.proc_name(C.int(pid), namePtr, C.uint32_t(1024))
	var result string

	if nameLen > 0 {
		result = C.GoString(name);
	} else {
		result = ""
	}

	return result
}

func count() int {
	procs := int(C.proc_listallpids(nil, C.int(0)))
	if procs <= 0 {
		return 0
	} else {
		return procs
	}
}

func listPids() []int {
	cnt := count()
	if cnt <= 0 {
		return make([]int, 0)
	}
	pids := C.malloc(C.size_t(cnt * int(unsafe.Sizeof(C.int(0)))))
	if pids == nil {
		return make([]int, 0)
	}
	defer C.free(unsafe.Pointer(pids))
	pidsCnt := C.proc_listallpids(pids, C.int(cnt * int(unsafe.Sizeof(C.int(0)))))
	if (pidsCnt <= 0) {
		return make([]int, 0)
	}
	casted := (*[1<<20]C.int)(unsafe.Pointer(pids))
	pidsCopy := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		pidsCopy[i] = int(casted[i])
	}
	return trimPidArray(pidsCopy)
}

func trimPidArray(pids []int) []int {
	index := len(pids)
	for i,e := range pids {
		if e <= 0 {
			index = i
			break
		}
	}
	return pids[0:index]
}