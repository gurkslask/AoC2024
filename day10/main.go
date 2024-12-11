package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gurkslask/AoC2024/common"
)

func main() {
	f := common.ReadFileSlice("day10", false)
	//simple(f)
	adv(f)

}

func simple(data []string) {
	g := common.MakeGrid(len(data), len(data[0]))
	for row, v := range data {
		for col := range v {
			s := string(data[row][col])
			g.AddMarker(common.MakeMarker(row, col, s))
		}
	}
	g.InitGrid()
	g.Print()
	res := 0

	for row, v := range data {
		for col := range v {
			if g.CheckLoc(row, col) == "0" {
				//fmt.Printf("Found 0, row: %v, col: %v\n", row, col)
				tres := 0
				tres, _ = check(g, row, col, []common.Pos{})
				res += tres
				//fmt.Println(res)
			}
		}
	}
	fmt.Println(res)
}
func PrintEnd(g common.Grid, row int, col int) {
	//fmt.Printf("Found end row %v, col %v, value: %v\n", row, col, g.CheckLoc(row, col))
}
func check(g common.Grid, row int, col int, taken []common.Pos) (int, []common.Pos) {
	res := 0
	curloc := StringsToIntwithO(g.CheckLoc(row, col))
	if curloc == 9 && !CheckIfPosTaken(taken, row, col) {
		PrintEnd(g, row, col)
		res += 1
		taken = append(taken, common.Pos{row, col})
		return res, taken
	}
	//Left
	if StringsToIntwithO(g.CheckLeft(row, col)) == curloc+1 {
		tres := 0
		tres, taken = check(g, row, col-1, taken)
		res += tres
	}
	//Right
	if StringsToIntwithO(g.CheckRight(row, col)) == curloc+1 {
		tres := 0
		tres, taken = check(g, row, col+1, taken)
		res += tres
	}
	//Up
	if StringsToIntwithO(g.CheckUp(row, col)) == curloc+1 {
		tres := 0
		tres, taken = check(g, row-1, col, taken)
		res += tres
	}
	//Down
	if StringsToIntwithO(g.CheckDown(row, col)) == curloc+1 {
		tres := 0
		tres, taken = check(g, row+1, col, taken)
		res += tres
	}
	//fmt.Printf("Found %v, res\n", res)
	return res, taken
}
func adv(data []string) {
	g := common.MakeGrid(len(data), len(data[0]))
	for row, v := range data {
		for col := range v {
			s := string(data[row][col])
			g.AddMarker(common.MakeMarker(row, col, s))
		}
	}
	g.InitGrid()
	g.Print()
	res := 0

	meh := make(map[common.Pos]int)
	for row, v := range data {
		for col := range v {
			if g.CheckLoc(row, col) == "0" {
				//fmt.Printf("Found 0, row: %v, col: %v\n", row, col)
				_, meh = checkAdv(g, row, col, meh)

				//fmt.Println(res)
			}
		}
	}
	fmt.Println(meh)
	for _, v := range meh {
		res += v
	}
	fmt.Println(res)

}
func checkAdv(g common.Grid, row int, col int, taken map[common.Pos]int) (int, map[common.Pos]int) {
	res := 0
	curloc := StringsToIntwithO(g.CheckLoc(row, col))
	if curloc == 9 {
		PrintEnd(g, row, col)
		res += 1
		taken[common.Pos{row, col}] += 1
		return res, taken
	}
	//Left
	if StringsToIntwithO(g.CheckLeft(row, col)) == curloc+1 {
		tres := 0
		tres, taken = checkAdv(g, row, col-1, taken)
		res += tres
	}
	//Right
	if StringsToIntwithO(g.CheckRight(row, col)) == curloc+1 {
		tres := 0
		tres, taken = checkAdv(g, row, col+1, taken)
		res += tres
	}
	//Up
	if StringsToIntwithO(g.CheckUp(row, col)) == curloc+1 {
		tres := 0
		tres, taken = checkAdv(g, row-1, col, taken)
		res += tres
	}
	//Down
	if StringsToIntwithO(g.CheckDown(row, col)) == curloc+1 {
		tres := 0
		tres, taken = checkAdv(g, row+1, col, taken)
		res += tres
	}
	//fmt.Printf("Found %v, res\n", res)
	return res, taken
}
func CheckIfPosTaken(t []common.Pos, row int, col int) bool {
	res := false
	for _, p := range t {
		if p.Row == row && p.Col == col {
			res = true
			//fmt.Println("Taken")
		}
	}
	return res
}

// Converts slice of string to slice of int
// If only white space, ignore
func StringsToIntwithO(s string) int {
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "ö", "", -1)
	i, err := strconv.Atoi(s)
	//common.CheckErr(err)
	if err != nil {
	}
	return i
}
func CheckIfRowTaken(t []common.Pos, row int) bool {
	res := false
	for _, p := range t {
		if p.Row == row {
			res = true
		}
	}
	return res
}
func CheckIfColTaken(t []common.Pos, col int) bool {
	res := false
	for _, p := range t {
		if p.Col == col {
			res = true
		}
	}
	return res
}
