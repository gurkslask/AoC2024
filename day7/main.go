package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gurkslask/AoC2024/common"
)

func main() {
	f := common.ReadFileSlice("day7", true)
	simple(f)
	//adv(f)

}

func fixData(s []string) map[int][]int {
	mm := make(map[int][]int)
	for _, row := range s {
		ss := strings.Split(row, ":")
		sss := strings.Split(ss[1], " ")
		for _, num := range sss {
			if num == "" {
				continue
			}
			mm[common.StrToInt(ss[0])] = append(mm[common.StrToInt(ss[0])], common.StrToInt(num))
		}
	}
	return mm
}
func subtract(nums []int) int {
	res := 0
	for key, num := range nums {
		if len(nums) == key+1 {
			continue
		} else {
			res += num + nums[key+1]
		}
	}
	return res
}
func c(num int) []string {
	res := []string{}
	for i := 0; i < num; i++ {
		b := strconv.FormatInt(2^int64(i), 2)
		res = append(res, b)
	}
	return res
}
func multi(nums []int) int {
	res := 0
	for key, num := range nums {
		if len(nums) == key+1 {
			continue
		} else {
			res += num * nums[key+1]
		}
	}
	return res
}
func checkIfLegal(sum int, nums []int) int {
	res := 0
	fmt.Println(c(len(nums)))
	if sum == multi(nums) {
		res += multi(nums)
	} else if sum == subtract(nums) {
		res += subtract(nums)
	}
	return res
}

func simple(data []string) {
	res := 0
	d := fixData(data)
	for k := range d {
		fmt.Println(k)
		res += checkIfLegal(k, d[k])
	}
	fmt.Println(res)
}
func adv(data []string) {
	fmt.Printf("%b", data)
}
