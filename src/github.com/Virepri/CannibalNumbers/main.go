package main

import (
	"fmt"
	"sort"
)

type NumSet struct {
	Numbers []int
	Queries []int
	Consumed []bool
}

func main(){
	NS := NumSet{}
	var NSN int
	var NSQ int
	fmt.Scan(&NSN,&NSQ)
	NS.Numbers = make([]int, NSN)
	NS.Consumed = make([]bool, NSN)
	for k := range NS.Numbers {
		fmt.Scan(&NS.Numbers[k])
		NS.Consumed[k] = false
	}
	NS.Queries = make([]int, NSQ)
	for k := range NS.Queries {
		fmt.Scan(&NS.Queries[k])
	}

	sort.Ints(NS.Numbers)
	sort.Ints(NS.Queries)

	for _,v := range NS.Queries {
		o := 0

		NS.Consumed = make([]bool, NSN)
		for k := range NS.Consumed {
			NS.Consumed[k] = false
		}

		for _,nv := range NS.Numbers {
			if nv >= v {
				o++
				continue
			}
			if Consume(nv, v, &NS) >= v {
				o++
				continue
			}
		}

		fmt.Print(o)
		fmt.Print(" ")
	}
}

func Consume(I int, S int, NS *NumSet) int {
	o := I

	for k,v := range NS.Numbers {
		if o >= S {
			break
		}
		if I > v && !NS.Consumed[k] {
			o++
			NS.Consumed[k] = true
		}
	}

	return o
}