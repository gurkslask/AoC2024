package main

import (
	"fmt"

	"github.com/gurkslask/AoC2024/common"
)

func main() {
	f := common.ReadFileSlice("templ", false)
	simple(f)
	adv(f)

}

func simple(data []string) {
	for _, v := range data {
		fmt.Println(v)
	}
}
func adv(data []string) {
	fmt.Printf("%b", data)
}
