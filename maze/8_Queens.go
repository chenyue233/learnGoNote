package main

import "fmt"

func queen(col int) {
	boards := make([]int, col)
	Put(boards, 0)
}

func Put(boards []int, col int) {
	boardsSize := len(boards)
	// fmt.Println(boardsSize)

	if col == boardsSize {
		printBoards(boards, col)
		return
	}
	for que := 0; que < boardsSize; que++ {
		boards[col] = que
		if safe(boards, col) {
			Put(boards, col+1)
		}
	}
}
func printBoards(boards []int, col int) []string {
	var result = make([]string, col)
	for _, board := range boards {
		str := ""
		for i := 0; i < board; i++ {
			if board == col {
				str += "Q"
			} else {
				str += "."
			}
		}
		result = append(result, str)
	}
	fmt.Println(result)
	return result
}
func safe(boards []int, col int) bool {
	for c := 0; c < col; c++ {
		if isAttack(boards, c, col) {
			return false
		}
	}
	return true
}
func isAttack(boards []int, c, col int) bool {
	switch {
	case c == col:
		return true
	case boards[c] == boards[col]:
		return true
	case boards[col]-boards[c] == col-c:
		return true
	case boards[col]-boards[c] == c-col:
		return true
	}
	return false
}
func main() {
	// queen(8)
	stringss := []int{1, 2, 3, 4, 5, 6, 7, 8}
	printBoards(stringss, 6)
}
