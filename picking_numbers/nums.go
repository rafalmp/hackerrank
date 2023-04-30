package pickingnumbers

import (
	"sort"
)

type counter struct {
	values map[int32]int32
}

type int32slice []int32

func (s int32slice) Len() int {
	return len(s)
}

func (s int32slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s int32slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (c *counter) init(vals []int32) {
	c.values = make(map[int32]int32)
	for _, v := range vals {
		c.values[v]++
	}
}

func (c *counter) keys() (result int32slice) {
	for k := range c.values {
		result = append(result, k)
	}
	sort.Sort(result)
	return
}

func (c *counter) num(key int32) int32 {
	if v, found := c.values[key]; found {
		return v
	}
	return 0
}

func pickingNumbers(a []int32) int32 {
	ctr := counter{}
	ctr.init(a)
	keys := ctr.keys()
	var result int32
	for idx := range keys[1:] {
		k1, k2 := keys[idx], keys[idx+1]
		if k2-k1 != 1 {
			continue
		}
		if sum := ctr.num(k1) + ctr.num(k2); sum > result {
			result = sum
		}
	}
	// handle edge cases (see test case 4 and 5)
	for _, k := range keys {
		if v := ctr.num(k); v > result {
			result = v
		}
	}
	return result
}
