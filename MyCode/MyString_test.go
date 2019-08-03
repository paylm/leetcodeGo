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

func Test_IptoInt(t *testing.T) {
	data := []struct {
		ip    string
		legal bool
		intIp int
	}{
		{ip: "127.0.0.1", legal: true, intIp: 2130706433},
		{ip: "0.0.0.1", legal: true, intIp: 1},
		{ip: "255.255.255.255", legal: true, intIp: 4294967295},
		{ip: "198.255.255", legal: false, intIp: 0},
		{ip: "192.168.1.1", legal: true, intIp: 3232235777},
	}
	for _, d := range data {
		res := ipToInt(d.ip)
		if d.legal {
			if res < 0 {
				t.Errorf("test fail at parse %v\n", d)
			}
			if res != d.intIp {
				t.Errorf("test fail at parse %v ,return %d\n", d, res)
			}
		} else {
			if res > 0 {
				t.Errorf("test fail at parse %v\n", d)
			}
		}
	}
}

func Test_InttoIP(t *testing.T) {
	data := []struct {
		ip    string
		legal bool
		intIp int
	}{
		{ip: "127.0.0.1", legal: true, intIp: 2130706433},
		{ip: "0.0.0.1", legal: true, intIp: 1},
		{ip: "255.255.255.255", legal: true, intIp: 4294967295},
		{ip: "198.255.255", legal: false, intIp: 0},
		{ip: "192.168.1.1", legal: true, intIp: 3232235777},
	}
	for _, d := range data {
		res := intToIP(d.intIp)
		if d.legal && strings.Compare(res, d.ip) != 0 {
			t.Errorf("test fail , %v be trans to %s\n", d, res)
		} else {
			t.Logf("%v parse to %v\n", d, res)
		}
	}
}

func Test_ValidUtf8(t *testing.T) {
	data := []struct {
		arr   []int
		legal bool
		size  int
	}{
		{arr: []int{197, 130, 1}, legal: true, size: 2},
		{arr: []int{235, 140, 4}, legal: false, size: 0},
		{arr: []int{255}, legal: false, size: 0},
		{arr: []int{145}, legal: false, size: 1},
		{arr: []int{240, 162, 138, 147}, legal: true, size: 1},
		{arr: []int{248, 130, 130, 130}, legal: false, size: 0},
	}
	for _, d := range data {
		res := validUtf8(d.arr)
		if res != d.legal {
			t.Errorf("test fail , %v should be %v , but return %v\n", d, d.legal, res)
		}
	}
}
