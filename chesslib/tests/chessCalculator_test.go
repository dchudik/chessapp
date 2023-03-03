package test

import (
	"fmt"
	"testing"

	"github.com/dchudik/chessapp/chesslib"
	"github.com/stretchr/testify/assert"
)

type testCaseCoordinatesError struct {
	FigureCoordinates string `db:"figure_coordinates"`
	KingCoordinates   string `db:"king_coordinates"`
	IsError           bool   `db:"error"`
}

// type testCaseCoordinatesUnderAttack struct {
// 	FigureCoordinates string `db:"figure_coordinates"`
// 	KingCoordinates   string `db:"king_coordinates"`
// 	FigureType        string `db:"figure"`
// 	IsUnderAttack     bool   `db:"under_attack"`
// }

func TestChessCalculatorCoordinatesError(t *testing.T) {
	connectPostgresOnce()
	chessLib := chesslib.NewChessKingUnderDangerCalculator()
	// Arrange
	// Receive test cases from DB
	testCases, err := psql.Queryx(`
		SELECT figure_coordinates,king_coordinates,error FROM figures_coordinates_errors
	`)
	assert.Nil(t, err, "Get test cases err: ", err)
	// Execute each test from db
	for testCases.Next() {
		testCase := testCaseCoordinatesError{}
		err := testCases.Scan(&testCase.FigureCoordinates, &testCase.KingCoordinates, &testCase.IsError)
		assert.Nil(t, err, "Scan err: ", err)
		// Act
		_, err = chessLib.IsKingUnderAttackBishop(testCase.FigureCoordinates, testCase.KingCoordinates)
		// t.Log(testCase.FigureCoordinates, byte(testCase.FigureCoordinates[0]), testCase.KingCoordinates, testCase.IsError, err)
		// Assert
		if testCase.IsError {
			assert.NotNil(t, err, fmt.Sprintf("Must error when coordinates Figure: %s and King: %s",
				testCase.FigureCoordinates, testCase.KingCoordinates))
		} else {
			assert.Nil(t, err, fmt.Sprintf("Must not error when coordinates Figure: %s and King: %s",
				testCase.FigureCoordinates, testCase.KingCoordinates))
		}

	}
}

// func TestChessCalculatorCoordinatesUnderAttack(t *testing.T) {
// 	connectPostgresOnce()
// 	chessLib := chesslib.NewChessKingUnderDangerCalculator()
// 	// Arrange
// 	testCases, err := psql.Queryx(`
// 		SELECT * FROM figures_coordinates_under_attack
// 	`)
// 	assert.Nil(t, err, "Get test cases err: ", err)
// 	// Execute each test from db
// 	for testCases.Next() {
// 		var testCase testCaseCoordinatesUnderAttack
// 		err := testCases.Scan(&testCase)
// 		assert.Nil(t, err, "Scan err: ", err)
// 		// Act
// 		underAttack, err := chessLib.IsKingUnderAttackBishop(testCase.FigureCoordinates, testCase.KingCoordinates)
// 		assert.NotNil(t, err, fmt.Sprintf("Must error when coordinates Figure: %s and King: %s", testCase.FigureCoordinates, testCase.KingCoordinates))
// 		// Assert
// 		if testCase.IsUnderAttack {
// 			assert.True(t, underAttack, fmt.Sprintf("Under attack must true when figure Type: %s coordinates Figure: %s and King: %s",
// 				testCase.FigureType, testCase.FigureCoordinates, testCase.KingCoordinates))
// 		} else {
// 			assert.False(t, underAttack, fmt.Sprintf("Under attack must true when figure Type: %s coordinates Figure: %s and King: %s",
// 				testCase.FigureType, testCase.FigureCoordinates, testCase.KingCoordinates))
// 		}

// 	}

// }
