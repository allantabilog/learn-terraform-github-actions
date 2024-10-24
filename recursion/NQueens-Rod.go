// N queens - Rod Stephen's solution

// place_queens_1
//  4:   0.000416 seconds
//  5:   0.000030 seconds
//  6: 101.059530 seconds
//  7:   (Got bored)

// place_queens_
//  4:   0.000105 seconds
//  5:   0.001577 seconds
//  6:   0.066245 seconds
//  7:   1.770676 seconds
//  8: 134.652694 seconds
//  9:   (Got bored)

// place_queens_3
//  4:   0.000058 seconds
//  5:   0.000073 seconds
//  6:   0.000463 seconds
//  7:   0.001186 seconds
//  8:   0.016526 seconds
//  9:   0.063856 seconds
// 10:   0.673511 seconds
// 11:   6.131033 seconds
// 12:  41.505714 seconds
// 13: 431.869695 seconds
// 14:   (Got bored)

package main

import (
	"fmt"
	"time"
)

func main() {
    const num_rows = 5
    board := make_board(num_rows)

    start := time.Now()
    success := place_queens_1(board, num_rows, 0, 0)
    //success := place_queens_2(board, num_rows, 0, 0, 0)
    //success := place_queens_3(board, num_rows, 0, 0, 0)

    elapsed := time.Since(start)
    if success {
        fmt.Println("Success!")
        dump_board(board)
    } else {
        fmt.Println("No solution")
    }
    fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())
}

// Make a board filled with periods.
func make_board(num_rows int) [][]string {
    num_cols := num_rows
    board := make([][]string, num_rows)
    for r := range board {
        board[r] = make([]string, num_cols)
        for c := 0; c < num_cols; c++ {
            board[r][c] = "."
        }
    }
    return board
}

// Display the board.
func dump_board(board [][]string) {
    for r := 0; r < len(board); r++ {
        for c := 0; c < len(board[r]); c++ {
            fmt.Printf("%s ", board[r][c])
        }
        fmt.Println()
    }
}

// Try placing a queen at position [r][c].
// Return true if we find a legal board.
func place_queens_1(board [][]string, num_rows, r, c int) bool {
    // See if we have examined the whole board.
    if r >= num_rows {
        // We have examined all of the squares.
        // See if this is a solution.
        return board_is_a_solution(board, num_rows)
    }

    // Find the next square.
    next_r := r
    next_c := c + 1
    if next_c >= num_rows {
        next_r += 1
        next_c = 0
    }

    // Leave no queen in this square and
    // recursively assign the next square.
    if place_queens_1(board, num_rows, next_r, next_c) {
        return true
    }

    // Try placing a queen here and
    // recursively assigning the next square.
    board[r][c] = "Q"
    if place_queens_1(board, num_rows, next_r, next_c) {
        return true
    }

    // That didn't work so remove this queen.
    board[r][c] = "."

    // If we get here, then there is no solution from
    // the board position before this function call.
    // Return false to backtrack and try again farther up the call stack.
    return false
}

// Try placing a queen at position [r][c].
// Keep track of the number of queens placed.
// Return true if we find a legal board.
func place_queens_2(board [][]string, num_rows, num_placed, r, c int) bool {
    // See if we have placed all of the queens.
    if num_placed == num_rows {
        // See if this is a solution.
        return board_is_a_solution(board, num_rows)
    }

    // See if we have examined the whole board.
    if r >= num_rows {
        // We have examined all of the squares but this is not a solution.
        return false
    }

    // Find the next square.
    next_r := r
    next_c := c + 1
    if next_c >= num_rows {
        next_r += 1
        next_c = 0
    }

    // Leave no queen in this square and
    // recursively assign the next square.
    if place_queens_2(board, num_rows, num_placed, next_r, next_c) {
        return true
    }

    // Try placing a queen here and
    // recursively assigning the next square.
    board[r][c] = "Q"
    num_placed += 1
    if place_queens_2(board, num_rows, num_placed, next_r, next_c) {
        return true
    }

    // That didn't work so remove this queen.
    board[r][c] = "."
    num_placed -= 1

    // If we get here, then there is no solution from
    // the board position before this function call.
    // Return false to backtrack and try again farther up the call stack.
    return false
}

// Set up and call place_queens_3.
func place_queens_3(board [][]string, num_rows, num_placed, r, c int) bool {
    // Make the num_attacking array.
    // The value num_attacking[r][c] is the number
    // of queens that can attack square (r, c).
    num_cols := num_rows
    num_attacking := make([][]int, num_rows)
    for r := range board {
        num_attacking[r] = make([]int, num_cols)
    }

    // Call do_place_queens_3.
    return do_place_queens_3(board, num_rows, num_placed, 0, 0, num_attacking)
}

// Try placing a queen at position [r][c].
// Keep track of the number of queens placed.
// Keep running totals of the number of queens attacking a square.
// Return true if we find a legal board.
func do_place_queens_3(board [][]string, num_rows, num_placed, r, c int, num_attacking [][]int) bool {
    // See if we have placed all of the queens.
    if num_placed == num_rows {
        // See if this is a solution.
        return board_is_a_solution(board, num_rows)
    }

    // See if we have examined the whole board.
    if r >= num_rows {
        // We have examined all of the squares but this is not a solution.
        return false
    }

    // Find the next square.
    next_r := r
    next_c := c + 1
    if next_c >= num_rows {
        next_r += 1
        next_c = 0
    }

    // Leave no queen in this square and
    // recursively assign the next square.
    if do_place_queens_3(board, num_rows, num_placed, next_r, next_c, num_attacking) {
        return true
    }

    // See if we can place a queen at (r, c).
    if num_attacking[r][c] == 0 {
        // Try placing a queen here and
        // recursively assigning the next square.
        board[r][c] = "Q"
        num_placed += 1

        // Increment the attack counts for this queen.
        adjust_attack_counts(num_attacking, num_rows, r, c, +1)

        if do_place_queens_3(board, num_rows, num_placed, next_r, next_c, num_attacking) {
            return true
        }
    
        // That didn't work so remove this queen.
        board[r][c] = "."
        num_placed -= 1
        adjust_attack_counts(num_attacking, num_rows, r, c, -1)
    }

    // If we get here, then there is no solution from
    // the board position before this function call.
    // Return false to backtrack and try again farther up the call stack.
    return false
}

// Add amount to the attack counts for this square.
func adjust_attack_counts(num_attacking [][]int, num_rows, r, c, amount int) {
    // Attacks in the same row.
    for i := 0; i < num_rows; i++ {
        num_attacking[r][i] += amount
    }

    // Attacks in the same column.
    for i := 0; i < num_rows; i++ {
        num_attacking[i][c] += amount
    }

    // Attacks in the upper left to lower right diagonal.
    for i := -num_rows; i < num_rows; i++ {
        test_r := r + i
        test_c := c + i
        if test_r >= 0 && test_r < num_rows &&
           test_c >= 0 && test_c < num_rows {
                num_attacking[test_r][test_c] += amount
           }
    }

    // Attacks in the upper right to lower left diagonal.
    for i := -num_rows; i < num_rows; i++ {
        test_r := r + i
        test_c := c - i
        if test_r >= 0 && test_r < num_rows &&
           test_c >= 0 && test_c < num_rows {
                num_attacking[test_r][test_c] += amount
           }
    }
}

// Return true if the board is legal and a solution.
func board_is_a_solution(board [][]string, num_rows int) bool {
    // See if it is legal.
    if !board_is_legal(board, num_rows) { return false }

    // See if the board contains exactly num_rows queens.
    num_queens := 0
    for r := 0; r < num_rows; r++ {
        for c := 0; c < num_rows; c++ {
            if board[r][c] == "Q" { num_queens += 1 }
        }
    }
    return num_queens == num_rows
}

// Return true if the board is legal.
func board_is_legal(board [][]string, num_rows int) bool {
    // See if each row is legal.
    for r := 0; r < num_rows; r++ {
        if !series_is_legal(board, num_rows, r, 0, 0, 1) { return false }
    }

    // See if each column is legal.
    for c := 0; c < num_rows; c++ {
        if !series_is_legal(board, num_rows, 0, c, 1, 0) { return false }
    }

    // See if diagonals down to the right are legal.
    for r := 0; r < num_rows; r++ {
        if !series_is_legal(board, num_rows, r, 0, 1, 1) { return false }
    }
    for c := 0; c < num_rows; c++ {
        if !series_is_legal(board, num_rows, 0, c, 1, 1) { return false }
    }

    // See if diagonals down to teh left are legal.
    for r := 0; r < num_rows; r++ {
        if !series_is_legal(board, num_rows, r, num_rows - 1, 1, -1) { return false }
    }
    for c := 0; c < num_rows; c++ {
        if !series_is_legal(board, num_rows, 0, c, 1, -1) { return false }
    }

    // If we survived this long, then the board is legal.
    return true
}

// Return true if this series of squares contains at most one queen.
func series_is_legal(board [][]string, num_rows, r0, c0, dr, dc int) bool {
    num_cols := num_rows
    has_queen := false

    r := r0
    c := c0
    for {
        if board[r][c] == "Q" {
            // If we already have a queen on this row,
            // then this board is not legal.
            if has_queen { return false }

            // Remember that we have a queen on this row.
            has_queen = true
        }

        // Move to the next square in the series.
        r += dr
        c += dc

        // If we fall off the board, then the series is legal.
        if  r >= num_rows ||
            c >= num_cols ||
            r < 0 ||
            c < 0 {
                return true
        }
    }
}
