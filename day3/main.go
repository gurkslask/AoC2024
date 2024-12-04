package main

import (
	"fmt"
	"regexp"

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
	dontPattern := `don't\(\)`
	dontReg := regexp.MustCompile(dontPattern)
	dontInd := dontReg.FindAllStringIndex(data[0], -1)
	doPattern := `do\(\)`
	doReg := regexp.MustCompile(doPattern)
	doInd := doReg.FindAllStringIndex(data[0], -1)
	// 1 if do, 0 if dont
	d := map[int]int{}
	for _, v := range dontInd {
		d[v[0]] = 0
	}
	for _, v := range doInd {
		d[v[0]] = 1
	}
	fmt.Println(d)
	for _, s := range data[0] {
		fmt.Printf("%s", s)
	}

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
