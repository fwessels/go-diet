package main

import (
	"sort"
)

func rankPackages(pgkMap PackageMap) PackageList {
	pl := make(PackageList, len(pgkMap))
	i := 0
	for pkg, v := range pgkMap {
		total, subs, autogend := uint64(0), 0, 0
		for _, v2 := range v {
			subs +=  len(v2)
			for _, s := range v2 {
				total += s.size
				if s.autogend {
					autogend++
				}
			}
		}
		pl[i] = PackageItem{Name: pkg, Subroutines: subs, Size: total, Autogend: autogend}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type PackageItem struct {
	Name        string
	Size        uint64
	Subroutines int
	Autogend    int
}

type PackageList []PackageItem

func (p PackageList) Len() int           { return len(p) }
func (p PackageList) Less(i, j int) bool { return p[i].Size < p[j].Size }
func (p PackageList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }


func rankReceivers(rcvMap map[string][]Subroutine) ReceiverList {
	pl := make(ReceiverList, len(rcvMap))
	i := 0
	for k, v := range rcvMap {
		total, autogend := uint64(0), 0
		for _, s := range v {
			total += s.size
			if s.autogend {
				autogend++
			}
		}
		pl[i] = ReceiverItem{Name: k, Subroutines: len(v), Size: total, Autogend: autogend}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type ReceiverItem struct {
	Name        string
	Size        uint64
	Subroutines int
	Autogend    int
}

type ReceiverList []ReceiverItem

func (p ReceiverList) Len() int           { return len(p) }
func (p ReceiverList) Less(i, j int) bool { return p[i].Size < p[j].Size }
func (p ReceiverList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
