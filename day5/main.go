package main

import (
	"fmt"
	"strings"

	"github.com/gurkslask/AoC2024/common"
)

func main() {
	f := common.ReadFileSlice("day5", false)
	//simple(f)
	adv(f)

}

func adv(data []string) int {
	orders, nm := getData(data)
	res := 0
	for _, order := range orders {
		fmt.Println(order)
		res += makeLegal(nm, order)
	}
	//fmt.Println(nm)
	//fmt.Println(orders)
	fmt.Println(res)
	return res
	//fmt.Printf("%b", data)
}
func checkIfLegalReturnIllegalKey(rules map[int][]int, keyof int, order []int) (bool, int) {
	checkafterslice := order[keyof+1:]
	numtocheck := order[keyof]
	//fmt.Printf("Before %v, After %v, num %v\n", checkbeforeslice, checkafterslice, numtocheck)
	for key, num := range checkafterslice {
		if common.ContainsInt(rules[numtocheck], num) {
			//OK
		} else {
			return false, keyof + key + 1
		}
	}
	return true, 0
}
func makeLegal(rules map[int][]int, order []int) int {
	legal := false
	wholerowisillegal := false
	wasillegal := false
	illegalkey := 0
	for key, _ := range order {
		for !wholerowisillegal {
			legal, illegalkey = checkIfLegalReturnIllegalKey(rules, key, order)
			if !legal {
				wasillegal = true
				//fmt.Printf("ILLEGAL Switch key %v, key that were checked %v, will switch \n", illegalkey, key)
				//fmt.Printf("Before switch %v\n", order)
				order = switchPlaces(key, illegalkey, order)
				//fmt.Printf("After switch %v\n", order)
			} else {
				wholerowisillegal = true
			}
		}
		wholerowisillegal = false
	}
	if wasillegal {
		return getMiddle(order)
	} else {
		return 0
	}
}
func switchPlaces(firstkey int, secondkey int, slice []int) []int {
	firstval := slice[firstkey]
	secondval := slice[secondkey]
	slice[firstkey] = secondval
	slice[secondkey] = firstval
	return slice
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
func simple(data []string) int {
	orders, nm := getData(data)
	res := 0
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
func getMiddle(slice []int) int {
	return slice[((len(slice) - 1) / 2)]
}
func getData(data []string) ([][]int, map[int][]int) {
	nm := make(map[int][]int)
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
	return orders, nm
}
