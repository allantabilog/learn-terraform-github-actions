package main

import "fmt"

type Offset struct {
	dr int
	dc int
}

const (
	rows = 8
	cols = 8 
	unvisited = -1
)

var board [][]int 

var allMoves = []Offset {
	{-2, -1},
	{-1, -2},
	{+2, -1},
	{+1, -2},
	{-2, +1},
	{-1, +2},
	{+2, +1},
	{+1, +2},
}

func makeBoard(numRows int, numCols int) [][]int {
	var board [][]int = make([][]int, numRows)

	for i := range(board) {
		board[i] = make([]int, numCols)
	}
	
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			board[i][j] = unvisited
		}
	}

	return board
}

func main() {
	var board = makeBoard(8, 8)
	fmt.Println("initial board:")
	trace(board)
	tour(board, 0, 0, 0)
	fmt.Println("final board:")
	trace(board)
}

func tour(board [][]int, currRow int, currCol int, stepCount int) {
	// mark the current board position with the step number
	board[currRow][currCol] = stepCount

	stepCount++
	if stepCount > 60 {
		return
	} 
	// find the next legal move in the space of all possible moves
	var nextRow, nextCol = findNextMove(currRow, currCol, board) 
	
	if nextRow != -1 && nextCol != -1 {
		tour(board, nextRow, nextCol, stepCount)
	} else {
		// moves exhausted
		fmt.Println("No more moves possible, Steps: ", stepCount)
	}
}
// there can be several findNextMove() functions,
// depending on the strategy
func findNextMove(currRow int, currCol int, board [][]int) (int, int) {
	var nextRow, nextCol = -1, -1
	for _, move := range allMoves {
		var testRow, testCol = currRow + move.dr, currCol + move.dc
		if validateMove(testRow, testCol, board) {
			nextRow, nextCol = testRow, testCol
			break
		} 
	}
	return nextRow, nextCol
}
func validateMove(newRow int, newCol int, board [][]int) bool {
	if newRow < 0 || newCol < 0 || newRow > rows - 1 || newCol > cols - 1 {
		return false
	}
	if board[newRow][newCol] != unvisited {
		return false
	}
	return true
}

func trace(board [][]int) {
	for i:=0; i<rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf(" %2d", board[i][j])
		}
		fmt.Println()
	}
}