package main

import "testing"

func Test_RestoreIpAddresses(t *testing.T) {
	data := []struct {
		str string
		res []string
	}{
		{str: "25525511135", res: []string{"255.255.11.135", "255.255.111.35"}},
		{str: "1111", res: []string{"1.1.1.1"}},
		{str: "111", res: []string{}},
		{str: "010010", res: []string{"0.10.0.10", "0.100.1.0"}},
	}

	for _, d := range data {
		r := restoreIpAddresses(d.str)
		if len(r) != len(d.res) {
			t.Errorf("test tail , %v has %d res ,but return %d\n", d, len(d.res), len(r))
		}
	}
}
