package main

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/gurkslask/AoC2024/common"
)

func main() {
	f := common.ReadFileSlice("day3", true)
	//fmt.Println(simple(f))
	adv(f)

}

func simple(data []string) int {
	//regexPattern := `mul\(\d{1,3},\d{1,3}\)`
	result := 0
	//tt := res.FindAllString(data[0], -1)
	//rm := res.FindAllStringSubmatch(data[0], -1)
	//fmt.Println(tt)
	//fmt.Println(rm)
	for _, v := range data {
		result += sumSum(v)
	}

	return result
}
func adv(data []string) {
	res := 0
	for _, row := range data {
		dontPattern := `don't\(\)`
		dontReg := regexp.MustCompile(dontPattern)
		dontInd := dontReg.FindAllStringIndex(row, -1)
		doPattern := `do\(\)`
		doReg := regexp.MustCompile(doPattern)
		doInd := doReg.FindAllStringIndex(row, -1)
		// 1 if do, 0 if dont
		d := map[int]int{}
		for _, v := range dontInd {
			d[v[0]] = 0
		}
		for _, v := range doInd {
			d[v[0]] = 1
		}
		start := 0
		s := row
		first := false
		dodon := false
		sss := ""
		sortedKeys := []int{}
		for k := range d {
			sortedKeys = append(sortedKeys, k)
		}
		sort.Ints(sortedKeys)
		for k, v := range sortedKeys {
			//fmt.Printf("Key: %v, Value %v \n", d[v], v)
			if !first && d[v] == 0 {
				sss += s[:v]
				first = true
			} else if first && dodon && d[v] == 1 && k+1 == len(sortedKeys) {
				sss += s[start+4:]
			} else if first && !dodon && d[v] == 1 && start == 0 {
				start = v
				//fmt.Printf("key %v, len: %v\n", k, len(sortedKeys))
				dodon = true
				if k+1 == len(sortedKeys) {
					sss += s[start+4:]
				}
			} else if first && dodon && d[v] == 0 {
				dodon = false
				sss += s[start+4 : v]
				//fmt.Printf("Start %v, End: %v\n", start, v)
				start = 0
			}

		}
		fmt.Println(sss)
		res += sumSum(sss)
	}
	fmt.Println(res)

}
func removeSubstringFromString(s string, start int, end int) string {
	return s[:start] + s[end+4:]

}

func sumSum(v string) int {
	result := 0
	regexPattern := `mul\((\d{1,3}),(\d{1,3})\)`
	res := regexp.MustCompile(regexPattern)
	rm := res.FindAllStringSubmatch(v, -1)
	for _, matches := range rm {
		//fmt.Println(matches)
		result += common.StrToInt(matches[1]) * common.StrToInt(matches[2])
	}
	return result
}
