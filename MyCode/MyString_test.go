package main

import (
	"strings"
	"testing"
)

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

func Test_SummaryRanges(t *testing.T) {
	data := []struct {
		arr []int
		res []string
	}{
		{arr: []int{0, 1, 2, 4, 5, 7}, res: []string{"0->2", "4->5", "7"}},
		{arr: []int{0, 2, 3, 4, 6, 8, 9}, res: []string{"0", "2->4", "6", "8->9"}},
	}

	for _, d := range data {
		r := summaryRanges(d.arr)
		if len(r) != len(d.res) {
			t.Errorf("test tail , %v has %d res ,but return %d\n", d, len(d.res), len(r))
			continue
		}
		for i := 0; i < len(r); i++ {
			if strings.Compare(r[i], d.res[i]) != 0 {
				t.Errorf("test tail , %v %d should %v ,but return %v\n", d, i, d.res[i], r[i])
			}
		}
	}
}
