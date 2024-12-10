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
	adv(ff)

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
func adv(data []int) {
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
	//fmt.Println("Njdsndnsakdn")
	//fmt.Println(disk)

	for rightmost := len(disk) - 1; rightmost > 1; rightmost-- {
		if disk[rightmost] != -1 {
			sizeofGrp := checkLenGrp(disk, rightmost)
			key := checkIfFits(disk, sizeofGrp, rightmost)
			fmt.Printf("GS: %v, key: %v, right %v, value %v\n", sizeofGrp, key, rightmost, disk[rightmost])
			if key != 0 {
				for i := 0; i < sizeofGrp; i++ {
					//fmt.Println("h")
					disk[key+i] = disk[rightmost-i]
					disk[rightmost-i] = -1
				}
			}
			rightmost = rightmost - sizeofGrp + 1

			//fmt.Println(disk)
		}
	}
	fmt.Println(calcChecksum(disk))
}
func checkLenGrp(inslic []int, key int) int {
	res := 0
	startnum := inslic[key]
	for i := key; inslic[i] == startnum || i == 0; i-- {
		res += 1
		if i <= 0 {
			break
		}
	}
	return res
}
func checkIfFits(inslic []int, size int, startkey int) int {
	cohesive := false
	newsize := 0
	savedkey := 0
	found := false
	for key, val := range inslic {
		if cohesive && val == -1 {
			newsize += 1
		} else if val == -1 {
			cohesive = true
			savedkey = key
			newsize += 1
		} else if val != -1 && key < startkey {
			//fmt.Printf("Key %v, Val %v, Size %v\n", key, val, newsize)
			if newsize >= size {
				found = true
				break
			} else {
				cohesive = false
				newsize = 0
				savedkey = 0
			}
		}
	}
	if found {
		return savedkey
	} else {
		return 0
	}
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
	fmt.Println(disk)

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

	fmt.Println(calcChecksum(disk))
}
