package main

import (
	"fmt"

	"github.com/gurkslask/AoC2024/common"
)

func main() {
	f := common.ReadFileSlice("day9", false)
	ff := []int{}
	for _, s := range f[0] {
		ff = append(ff, common.StrToInt(string(s)))
	}
	//simple(ff)
	adv(f)

}

func simple(data []int) {
	id := 0
	disk := []int{}
	for numkey := 0; numkey < len(data); numkey += 2 {
		num := data[numkey]

		for i := 0; i < num; i++ {
			disk = append(disk, id)
		}
		if numkey < len(data)-1 {
			freespace := data[numkey+1]
			for i := 0; i < freespace; i++ {
				disk = append(disk, -1)
			}
		}
		id += 1
	}

	for rightmost := len(disk) - 1; rightmost > 1; rightmost-- {
		if disk[rightmost] != -1 {
			keytoswitch := checkIfFree(disk)
			disk[keytoswitch] = disk[rightmost]
			disk[rightmost] = -1
			if checkDone(disk) {
				break
			}
		}
	}
	fmt.Println(disk)

	fmt.Println(calcChecksum(disk))
}

func checkIfFree(inslic []int) int {
	for key, val := range inslic {
		if val == -1 {
			return key
		}
	}
	return 0
}
func checkDone(inslic []int) bool {
	done := false
	for _, val := range inslic {
		if done && val != -1 {
			done = false
			break
		} else if val == -1 {
			done = true
		}
	}
	return done
}
func calcChecksum(inslic []int) int {
	res := 0
	for key, val := range inslic {
		if val != -1 {
			res += key * val
		}
	}
	return res
}
func adv(data []string) {
	fmt.Printf("%b", data)
}
