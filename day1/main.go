package main

import (
	"fmt"
	"strings"

	"github.com/gurkslask/AoC2024/common"
)

func main() {
	f := common.ReadFileSlice("day1", false)
	left := []int{}
	right := []int{}
	result := 0
	for _, v := range f {
		t := strings.Split(v, "  ")
		left = append(left, common.StrToInt(t[0]))
		//fmt.Println(t[1])
		right = append(right, common.StrToInt(t[1]))
	}
	fmt.Println(left)
	fmt.Println(right)
	/*for len(left) >= 1 {
		minleft, leftkey := common.MinInts(left)
		minright, rightkey := common.MinInts(right)
		result += common.IntAbs(minleft - minright)
		left = common.RemoveFromSliceInt(left, leftkey)
		right = common.RemoveFromSliceInt(right, rightkey)
	}*/
	for _, v := range left {
		num := 0
		for _, r := range right {
			if r == v {
				num += 1
			}
		}
		result += v * num
	}
	fmt.Println(result)

}
