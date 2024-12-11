package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gurkslask/AoC2024/common"
)

func main() {
	f := common.ReadFileSlice("day11", false)
	f = strings.Split(f[0], " ")
	simple(f)
	//adv(f)
}

func simple(data []string) {
	/*newdata := make([]string, len(data))
	_ = copy(newdata, data)
	fmt.Println(newdata)*/
	for i := 0; i < 25; i++ {
		//fmt.Println(data)
		key := 0
		for key < len(data) {
			//for key, v := range data {
			v := data[key]
			if v == "0" {
				zerorule(data, key)
			} else if len(v)%2 == 0 {
				data = split(data, key)
				key += 1
			} else {
				rule2024(data, key)
			}
			key += 1
		}
	}
	fmt.Println(len(data))
}
func adv(data []string) {
	fmt.Printf("%b", data)
}
func rule2024(is []string, key int) {
	v := common.StrToInt(is[key])
	is[key] = strconv.Itoa(v * 2024)
}
func zerorule(is []string, key int) {
	is[key] = "1"
}
func split(is []string, key int) []string {
	value := is[key]
	firsthalf := value[:len(value)/2]
	secondhalf := strconv.Itoa(common.StrToInt(value[len(value)/2:]))
	//fmt.Printf("First:%v, Second:%v, val:%v.\n", firsthalf, secondhalf, value)
	ret := make([]string, 0)
	ret = append(ret, is[:key]...)
	ret = append(ret, firsthalf)
	ret = append(ret, secondhalf)
	ret = append(ret, is[key+1:]...)
	return ret

}
