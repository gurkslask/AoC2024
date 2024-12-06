package main

import (
	"fmt"

	"github.com/gurkslask/AoC2024/common"
)

func main() {
	f := common.ReadFileSlice("day6", true)
	simple(f)
	//adv(f)

}

type Moves struct {
	Movess []string
	Key    int
	Marker *common.Marker
	Gridd  *common.Grid
}

func (m *Moves) Move() bool {
	// Moves and returns value if edge is located
	p := m.Marker.GetPos()
	mposrow, mposcol := p.Row, p.Col
	//m.Gridd.AddMarker(common.MakeMarker(mposrow, mposrow, "A"))
	if m.Movess[m.Key] == "U" {
		//fmt.Println(mposrow, mposcol)
		next := m.Gridd.CheckUp(mposrow, mposcol)
		if next == "ö" {
			return true
		} else if next == "#" {
			m.Rotate()
		} else {
			err := m.Marker.MoveMarkerUp(*m.Gridd)
			if err != nil {
				fmt.Println(err)
			}
		}
	} else if m.Movess[m.Key] == "R" {
		//fmt.Println("Want right")
		//m.Gridd.Print()
		next := m.Gridd.CheckRight(mposrow, mposcol)
		if next == "ö" {
			return true
		} else if next == "#" {
			m.Rotate()
		} else {
			err := m.Marker.MoveMarkerRight(*m.Gridd)
			if err != nil {
				fmt.Println(err)
			}
		}
	} else if m.Movess[m.Key] == "D" {
		//fmt.Println(mposrow, mposcol)
		//m.Gridd.Print()
		next := m.Gridd.CheckDown(mposrow, mposcol)
		if next == "ö" {
			return true
		} else if next == "#" {
			m.Rotate()
		} else {
			err := m.Marker.MoveMarkerDown(*m.Gridd)
			if err != nil {
				fmt.Println(err)
			}
		}
	} else if m.Movess[m.Key] == "L" {
		//fmt.Println(mposrow, mposcol)
		//m.Gridd.Print()
		next := m.Gridd.CheckLeft(mposrow, mposcol)
		if next == "ö" {
			return true
		} else if next == "#" {
			m.Rotate()
		} else {
			err := m.Marker.MoveMarkerLeft(*m.Gridd)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return false

}
func (m *Moves) Rotate() {
	m.Key += 1
	if m.Key > 3 {
		m.Key = 0
	}
}

func simple(data []string) int {
	g := common.MakeGrid(len(data), len(data[0]))
	var startMarker common.Marker
	for row, v := range data {
		for col := range v {
			s := string(data[row][col])
			if s == "#" {
				g.AddMarker(common.MakeMarker(row, col, s))
			} else if s == "^" {
				startMarker = common.MakeMarker(row, col, s)
				g.AddMarker(startMarker)
			}
		}

	}
	g.InitGrid()
	mm := Moves{[]string{"U", "R", "D", "L"}, 0, &startMarker, &g}
	var done bool = false
	visited := make(map[common.Pos]int)
	//limit := 60
	for !done {
		//for i := 0; i < limit; i++ {
		//fmt.Printf("%v key %v\n", startMarker, mm.Key)
		//g.Print()
		visited[mm.Marker.GetPos()] += 1
		//fmt.Println(i)
		done = mm.Move()
	}
	fmt.Println(done)
	g.InitGrid()

	//fmt.Printf("%v key %v\n", startMarker, mm.Key)
	fmt.Printf("%v\n", len(visited))
	g.Print()
	return 0
}
func adv(data []string) {
	fmt.Printf("%b", data)
}
