package main

// In this method i will create a 9x9 matrix of the rows.
// i will then create a 9x9 matrix of the columns
// and then i will create 9 3x3 matrix for all the squares.
// I will then loop through the rows 'i' and the columns 'j'
// checking if the cell value is '.' (empty) and handle this.
// i will then conver the cell value which is a byte to an integer by
// subtracting the cell value with the character '0'. I then subtract 1 by
// making the the cell value correspond to a zero index in an array. this way
// cell value 1 == 0 and cell value 9 == 8.
// This way an mark the indexes in the rows by true if that number was found in that specific
// row, column and square.
// This makes it easy in by just checking the index corresponding to the value to see if it already
// exists.
func isValidSudoku(board [][]byte) bool {
	rows := [9][9]bool{}
	columns := [9][9]bool{}
	squares := [3][3][9]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cell := board[i][j]

			// handling empty cells which wil be marked as a '.'
			if cell == '.' {
				continue
			}

			// the reason for this is that i can convert the cell value to an integer
			// and the subtract 1. The reason for the subtraction is that the numbers
			// are from 1-9, but if can make them from 0-8 then i can rely on the
			// array indexing, and just place a true or false if already exists.
			digit := int(cell-'0') - 1 // since it will be a "char / byte" i need to convert to int()

			// here i am just checking if the digit already exists in the row
			// in the column, or in the current square
			if rows[i][digit] || columns[j][digit] || squares[i/3][j/3][digit] {
				return false
			}

			// if the digit is currently not found, then i mark on the index in the current row,
			// column and square that this digit(index) has been found.
			rows[i][digit] = true
			columns[j][digit] = true
			squares[i/3][j/3][digit] = true
		}
	}

	return true
}
