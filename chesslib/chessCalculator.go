package chesslib

import "strings"

// Create ChessKingUnderDangerCalculator instance
func NewChessKingUnderDangerCalculator() *ChessLib {
	debug := getDebugMode()
	return &ChessLib{
		Debug:  debug,
		Logger: NewLogger(debug),
	}
}

// Calculate danger to King from Rook
func (ChessLib *ChessLib) IsKingUnderAttackRook(figureCoordinates, kingCoordinates string) (bool, error) {
	figureCoordinates = strings.ToUpper(figureCoordinates)
	kingCoordinates = strings.ToUpper(kingCoordinates)
	err := ChessLib.validateCoordinatesForTwoFigures(figureCoordinates, kingCoordinates)
	if err != nil {
		return false, err
	}
	isUnderAttack := false
	ChessLib.Logger.DebugLog(FigureTypeRook, figureCoordinates, kingCoordinates, isUnderAttack, err)
	return isUnderAttack, nil
}

// Calculate danger to King from Bishop
func (ChessLib *ChessLib) IsKingUnderAttackBishop(figureCoordinates, kingCoordinates string) (bool, error) {
	figureCoordinates = strings.ToUpper(figureCoordinates)
	kingCoordinates = strings.ToUpper(kingCoordinates)
	err := ChessLib.validateCoordinatesForTwoFigures(figureCoordinates, kingCoordinates)
	if err != nil {
		return false, err
	}
	isUnderAttack := false
	ChessLib.Logger.DebugLog(FigureTypeBishop, figureCoordinates, kingCoordinates, isUnderAttack, err)
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
		// return err
	}
	err = ChessLib.validateCoordinates(kingCoordinates)
	if err != nil {
		return ErrKingCoordinatesNotValid
		// return err
	}

	return nil
}

// Validate coordinates for one figure
func (ChessLib *ChessLib) validateCoordinates(figureCoordinates string) error {
	if len(figureCoordinates) != 2 {
		return ErrCoordinatesNotValid
	}
	if !isFirstCoordinateValid(figureCoordinates[0]) {
		return ErrVerticalCoordinateNotValid
	}
	if !isSecondCoordinateValid(figureCoordinates[1]) {
		return ErrHorizontalCoordinateNotValid
	}
	return nil
}

// Validate first coordinate
func isFirstCoordinateValid(coordinate byte) bool {
	// First letter in range A,B,C,D,E,F,G,H
	if coordinate >= 65 && coordinate <= 72 {
		return true
	}
	return false
}

// Validate second coordinate
func isSecondCoordinateValid(coordinate byte) bool {
	// Second letter in range 1,2,3,4,5,6,7,8
	if coordinate >= 49 && coordinate <= 56 {
		return true
	}
	return false
}
