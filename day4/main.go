package main

import (
	"fmt"

	"github.com/gurkslask/AoC2024/common"
)

func main() {
	f := common.ReadFileSlice("day4", false)
	//fmt.Println(simple(f))
	fmt.Println(adv(f))

}

func simple(data []string) int {
	rows := len(data)
	cols := len(data[0])
	m := make([][]string, rows)
	row := 0
	for rows > row {
		col := 0
		for cols > col {
			m[row] = append(m[row], string(data[row][col]))
			col += 1
		}
		row += 1
	}
	fmt.Println(m)
	p := pos{0, 0, ""}
	res := 0

	for _, v := range data {
		fmt.Println(v)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if m[row][col] == "X" {
				p.x = row
				p.y = col
				if searchForwardHorizon(m, p) {
					res += 1
					Debugprint(m, p)
				}
				if searchBackwardHorizon(m, p) {
					res += 1
					Debugprint(m, p)
				}
				if searchUpVertical(m, p) {
					res += 1
					Debugprint(m, p)
				}
				if searchDownVertical(m, p) {
					res += 1
					Debugprint(m, p)
				}
				if searchForwardDownDiag(m, p) {
					res += 1
					Debugprint(m, p)
				}
				if searchForwardUpDiag(m, p) {
					res += 1
					Debugprint(m, p)
				}
				if searchBackwardDownDiag(m, p) {
					res += 1
					Debugprint(m, p)
				}
				if searchBackwardUpDiag(m, p) {
					res += 1
					Debugprint(m, p)
				}
			}
		}
	}
	fmt.Println(res)

	return 0
}
func Debugprint(board [][]string, p pos) {
	//fmt.Printf("pos: %v \n letter: %v\n", p, board[p.x][p.y])
}
func adv(data []string) int {
	rows := len(data)
	cols := len(data[0])
	m := make([][]string, rows)
	row := 0
	for rows > row {
		col := 0
		for cols > col {
			m[row] = append(m[row], string(data[row][col]))
			col += 1
		}
		row += 1
	}
	fmt.Println(m)
	p := pos{0, 0, ""}
	res := 0

	for _, v := range data {
		fmt.Println(v)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if m[row][col] == "A" {
				p.x = row
				p.y = col
				ludr := make(map[string]int)
				rudl := make(map[string]int)
				if AcheckIfForward(m, p) || AcheckIfBackward(m, p) || AcheckIfDown(m, p) || AcheckIfUp(m, p) {
				} else {
					ludr[m[p.x+1][p.y+1]] += 1 //dr
					rudl[m[p.x-1][p.y+1]] += 1 //ru
					rudl[m[p.x+1][p.y-1]] += 1 //dl
					ludr[m[p.x-1][p.y-1]] += 1 //lu
				}
				if rudl["S"] == 1 && rudl["M"] == 1 && ludr["S"] == 1 && ludr["M"] == 1 {
					res += 1
				}
				//fmt.Println(mm)

			}
		}
	}
	return res
}
func AcheckIfForward(board [][]string, p pos) bool {
	if p.y >= len(board[0])-1 {
		return true
	} else {
		return false
	}
}
func AcheckIfBackward(board [][]string, p pos) bool {
	if p.y < 1 {
		return true
	} else {
		return false
	}
}
func AcheckIfDown(board [][]string, p pos) bool {
	if p.x >= len(board)-1 {
		return true
	} else {
		return false
	}
}
func AcheckIfUp(board [][]string, p pos) bool {
	if p.x < 1 {
		return true
	} else {
		return false
	}
}

type pos struct {
	x       int
	y       int
	comment string
}

func searchForwardHorizon(board [][]string, p pos) bool {
	if checkIfForward(board, p) {
		return false
	} else {
		if board[p.x][p.y+1] == "M" && board[p.x][p.y+2] == "A" && board[p.x][p.y+3] == "S" {
			return true
		} else {
			return false
		}
	}
}
func searchBackwardHorizon(board [][]string, p pos) bool {
	if checkIfBackward(board, p) {
		return false
	}
	if board[p.x][p.y-1] == "M" && board[p.x][p.y-2] == "A" && board[p.x][p.y-3] == "S" {
		return true
	} else {
		return false
	}
}
func searchUpVertical(board [][]string, p pos) bool {
	if checkIfUp(board, p) {
		return false
	}
	if board[p.x-1][p.y] == "M" && board[p.x-2][p.y] == "A" && board[p.x-3][p.y] == "S" {
		return true
	} else {
		return false
	}
}
func searchDownVertical(board [][]string, p pos) bool {
	if checkIfDown(board, p) {
		return false
	}
	if board[p.x+1][p.y] == "M" && board[p.x+2][p.y] == "A" && board[p.x+3][p.y] == "S" {
		return true
	} else {
		return false
	}
}
func searchForwardDownDiag(board [][]string, p pos) bool {
	if checkIfDown(board, p) || checkIfForward(board, p) {
		return false
	}
	if board[p.x+1][p.y+1] == "M" && board[p.x+2][p.y+2] == "A" && board[p.x+3][p.y+3] == "S" {
		return true
	} else {
		return false
	}
}
func searchForwardUpDiag(board [][]string, p pos) bool {
	if checkIfUp(board, p) || checkIfForward(board, p) {
		return false
	}
	if board[p.x-1][p.y+1] == "M" && board[p.x-2][p.y+2] == "A" && board[p.x-3][p.y+3] == "S" {
		return true
	} else {
		return false
	}
}
func searchBackwardDownDiag(board [][]string, p pos) bool {
	if checkIfDown(board, p) || checkIfBackward(board, p) {
		return false
	}
	if board[p.x+1][p.y-1] == "M" && board[p.x+2][p.y-2] == "A" && board[p.x+3][p.y-3] == "S" {
		return true
	} else {
		return false
	}
}
func searchBackwardUpDiag(board [][]string, p pos) bool {
	p.comment = "here"
	Debugprint(board, p)
	if checkIfUp(board, p) || checkIfBackward(board, p) {
		return false
	}
	if board[p.x-1][p.y-1] == "M" && board[p.x-2][p.y-2] == "A" && board[p.x-3][p.y-3] == "S" {
		return true
	} else {
		return false
	}
}

func checkIfForward(board [][]string, p pos) bool {
	if p.y >= len(board[0])-3 {
		return true
	} else {
		return false
	}
}
func checkIfBackward(board [][]string, p pos) bool {
	if p.y <= 2 {
		return true
	} else {
		return false
	}
}
func checkIfDown(board [][]string, p pos) bool {
	if p.x >= len(board)-3 {
		return true
	} else {
		return false
	}
}
func checkIfUp(board [][]string, p pos) bool {
	if p.x <= 2 {
		return true
	} else {
		return false
	}
}
