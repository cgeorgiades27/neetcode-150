package arrays

import (
	"testing"
)

// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
// Note:

// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules.

func ValidateSudoku(board [][]byte) bool {
	if !(checkVertical(board) && checkHorizontal(board)) {
		return false
	}

	for rowsStart, rowsStop := 0, 3; rowsStop <= 9; rowsStart, rowsStop = rowsStart+2, rowsStop+2 {
		for colsStart, colsStop := 0, 3; colsStop <= 9; colsStart, colsStop = colsStart+2, colsStop+2 {
			var slicedBoard [][]byte
			for i := 0; i < 3; i++ {
				row := board[rowsStart:rowsStop][i]
				slicedBoard = append(slicedBoard, row[colsStart:colsStop])
			}
			if !(checkVertical(slicedBoard) && checkHorizontal(slicedBoard)) {
				return false
			}
		}
	}
	return true
}

func checkVertical(board [][]byte) bool {
	for _, row := range board {
		mHor := make(map[byte]struct{})
		for _, cell := range row {
			if _, ok := mHor[cell]; ok && cell != '.' {
				return false
			} else {
				mHor[cell] = struct{}{}
			}
		}
	}
	return true
}

func checkHorizontal(board [][]byte) bool {
	numCols := len(board[0]) // always 9
	for i := 0; i < numCols; i++ {
		mVer := make(map[byte]struct{})
		for j := 0; j < len(board); j++ {
			if _, ok := mVer[board[j][i]]; ok && board[j][i] != '.' {
				return false
			} else {
				mVer[board[j][i]] = struct{}{}
			}
		}
	}
	return true
}

func TestValidateSudoku(t *testing.T) {
	tests := []struct {
		board    [][]byte
		expected bool
	}{
		{
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			expected: true,
		},
		{
			board: [][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			expected: false,
		},
	}

	for index, test := range tests {
		actual := ValidateSudoku(test.board)
		if actual != test.expected {
			t.Fatalf("%d - expected: %t, got: %t\n", index, test.expected, actual)
		}
	}
}
