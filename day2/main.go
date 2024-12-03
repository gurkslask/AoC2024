package main

import (
	"fmt"
	"strings"

	"github.com/gurkslask/AoC2024/common"
)

func main() {
	f := common.ReadFileSlice("day2", false)
	fmt.Println(simple(f))
	//fmt.Println(adv(f))

}

func simple(data []string) int {
	result := 0
	for _, v := range data {
		UporDown := ""
		iv := common.StringsToInt(strings.Split(v, " "))
		//fmt.Println(iv)
		if iv[0] > iv[1] {
			UporDown = "down"
		} else if iv[0] < iv[1] {
			UporDown = "up"
		}
		fmt.Println(UporDown)

		valid := is_safe(iv, UporDown)
		//fmt.Println(v)
		//fmt.Println(valid)
		if valid {
			result += 1
		}
	}
	return result
}
func adv(data []string) int {
	result := 0
	for _, v := range data {
		UporDown := ""
		iv := common.StringsToInt(strings.Split(v, " "))
		fmt.Println(iv)
		if iv[0] > iv[1] {
			UporDown = "down"
		} else if iv[0] < iv[1] {
			UporDown = "up"
		}
		//fmt.Println(UporDown)
		valid := is_safe(iv, UporDown)

		if !valid {
			for l := 0; l < len(iv); l++ {
				t := common.RemoveFromSliceOrder(iv, l)
				//fmt.Printf("iv: %v", iv)
				//fmt.Printf("t: %v l: %v\n", t, l)

				valid = is_safe(t, UporDown)

				fmt.Println(falsevalid)
				if falsevalid == len(iv)-2 {
					valid = true
				}
			}
		}
		//fmt.Println(v)
		//fmt.Println(valid)
		if valid {
			result += 1
		}
	}
	return result
}

func is_safe(iv []int, UporDown string) bool {
	valid := true
	for i := 0; i < len(iv)-1; i++ {
		if UporDown == "down" && (iv[i] == iv[i+1]+1 || iv[i] == iv[i+1]+2 || iv[i] == iv[i+1]+3) {

		} else if UporDown == "up" && (iv[i] == iv[i+1]-1 || iv[i] == iv[i+1]-2 || iv[i] == iv[i+1]-3) {
			//fmt.Printf("%v:%v\n", iv[i], iv[i+1])
		} else {
			valid = false
		}

	}
	return valid
}
