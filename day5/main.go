package main

import (
	"fmt"
	"strings"

	"github.com/gurkslask/AoC2024/common"
)

func main() {
	f := common.ReadFileSlice("day5", false)
	simple(f)
	//adv(f)

}

type Page struct {
	number    int
	followers []int
}

func simple(data []string) int {
	nm := make(map[int][]int)
	res := 0
	second_split_enable := false
	orders := [][]int{}
	rules := [][]string{}
	for k, v := range data {
		if !second_split_enable {
			rules = append(rules, strings.Split(v, "|"))
			if len(rules[k]) == 1 {
				second_split_enable = true
				continue
			}
			key := common.StrToInt(rules[k][0])
			value := common.StrToInt(rules[k][1])
			nm[key] = append(nm[key], value)
		} else if second_split_enable {
			orders = append(orders, common.StringsToInt(strings.Split(v, ",")))
		}
	}
	for _, order := range orders {
		legal := true
		for key, _ := range order {
			//fmt.Printf("order: %v, page: %v, legal? %v\n", order, page, checkIfLegal(nm, key, order))
			if checkIfLegal(nm, key, order) {
			} else {
				legal = false
			}
		}
		if legal {
			res += order[((len(order) - 1) / 2)]
			fmt.Println(res)
		}
	}
	//fmt.Println(nm)
	//fmt.Println(orders)
	fmt.Println(res)
	return res

}

func checkIfLegal(rules map[int][]int, keyof int, order []int) bool {
	checkafterslice := order[keyof+1:]
	numtocheck := order[keyof]
	//fmt.Printf("Before %v, After %v, num %v\n", checkbeforeslice, checkafterslice, numtocheck)
	for _, num := range checkafterslice {
		if common.ContainsInt(rules[numtocheck], num) {
			//OK
		} else {
			return false
		}
	}
	return true
}
func adv(data []string) {
	fmt.Printf("%b", data)
}
