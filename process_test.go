// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

package process

import (
	"fmt"
	"testing"
)

func TestNameOf(t *testing.T) {

	cases := testCasesNameOf()

	for _, c := range cases {
		got := NameOf(c.in)
		if got != c.want {
			t.Errorf("NameOf(%d) == %s, want %s", c.in, got, c.want)
		}
	}
}

func TestCount(t *testing.T) {
	want := true
	got := Count()
	if got > 0 != want {
		t.Errorf("(Count() == %d) > 0 == %t, want %t", got, got > 0, want)
	} else {
		fmt.Printf("Count() == %d\n", got)
	}
}

func TestListPids(t *testing.T) {
	want := true
	got := ListPids()
	if len(got) > 0 != want {
		t.Errorf("(len(ListPids()) == %d) > 0 == %t, want %t", len(got), len(got) > 0, want)
	} else {
		for _, pid := range got {
			fmt.Printf("%d ", pid)
		}
		fmt.Printf("\n")
	}
}

func TestPidOf(t *testing.T) {

	cases := testCasesPidOf()

	for _, c := range cases {
		got := PidOf(c.in)
		if got != c.want {
			t.Errorf("PidOf(%s) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestPropertiesOf(t *testing.T) {
	got := propertiesOf(38, []int{VMUsage}) //syslogd on Darwin
	if len(got) > 0 {
		fmt.Printf("VMUsage: %d\n", got[VMUsage])
	} else {
		t.Errorf("private [darwin only] propertiesOf() failed")
	}
}
