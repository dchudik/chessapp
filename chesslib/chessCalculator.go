package chesslib

import "strings"

// Create ChessKingUnderDangerCalculator instance
func NewChessKingUnderDangerCalculator() *ChessLib {
	return &ChessLib{}
}

// Calculate danger to King from Rook
func (ChessLib *ChessLib) IsKingUnderAttackRook(figureCoordinates, kingCoordinates string) (bool, error) {
	figureCoordinates = strings.ToUpper(figureCoordinates)
	kingCoordinates = strings.ToUpper(kingCoordinates)
	err := ChessLib.validateCoordinatesForTwoFigures(figureCoordinates, kingCoordinates)
	if err != nil {
		return false, ErrKingCoordinatesNotValid
	}
	return false, nil
}

// Calculate danger to King from Bishop
func (ChessLib *ChessLib) IsKingUnderAttackBishop(figureCoordinates, kingCoordinates string) (bool, error) {
	figureCoordinates = strings.ToUpper(figureCoordinates)
	kingCoordinates = strings.ToUpper(kingCoordinates)
	err := ChessLib.validateCoordinatesForTwoFigures(figureCoordinates, kingCoordinates)
	if err != nil {
		return false, ErrKingCoordinatesNotValid
	}
	return false, nil
}

// Validate coordinates for two figures
func (ChessLib *ChessLib) validateCoordinatesForTwoFigures(figureCoordinates, kingCoordinates string) error {
	if figureCoordinates == kingCoordinates {
		return ErrCoordinatesNotValid
	}
	err := ChessLib.validateCoordinates(figureCoordinates)
	if err != nil {
		return ErrCoordinatesNotValid
	}
	err = ChessLib.validateCoordinates(kingCoordinates)
	if err != nil {
		return ErrKingCoordinatesNotValid
	}

	return nil
}

// Validate coordinates for one figure
func (ChessLib *ChessLib) validateCoordinates(figureCoordinates string) error {
	if len(figureCoordinates) != 2 {
		return ErrCoordinatesNotValid
	}
	if !isFirstCoordinateValid(figureCoordinates[0]) {
		return ErrCoordinatesNotValid
	}
	if !isSecondCoordinateValid(figureCoordinates[1]) {
		return ErrCoordinatesNotValid
	}
	return nil
}

// Validate first coordinate
func isFirstCoordinateValid(coordinate byte) bool {
	// First letter in range A,B,C,D,E,F,G,H
	if coordinate >= 41 && coordinate <= 48 {
		return true
	}
	return false
}

// Validate second coordinate
func isSecondCoordinateValid(coordinate byte) bool {
	// Second letter in range 1,2,3,4,5,6,7,8
	if coordinate >= 31 && coordinate <= 38 {
		return true
	}
	return false
}
