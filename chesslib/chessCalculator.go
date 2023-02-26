package chesslib

// NewChessKingUnderDangerCalculator Create ChessKingUnderDangerCalculator instance
func NewChessKingUnderDangerCalculator() *ChessLib {
	return &ChessLib{}
}

// IsKingUnderAttackRook Calculate danger to King from Rook
func (ChessLib *ChessLib) IsKingUnderAttackRook(figureCoordinates, kingCoordinates string) (bool, error) {
	err := ChessLib.validateCoordinates(kingCoordinates)
	if err != nil {
		return false, ErrKingCoordinatesNotValid
	}
	err = ChessLib.validateCoordinates(figureCoordinates)
	if err != nil {
		return false, ErrRookCoordinatesNotValid
	}
	return false, nil
}

// IsKingUnderAttackBishop Calculate danger to King from Bishop
func (ChessLib *ChessLib) IsKingUnderAttackBishop(figureCoordinates, kingCoordinates string) (bool, error) {
	err := ChessLib.validateCoordinates(kingCoordinates)
	if err != nil {
		return false, ErrKingCoordinatesNotValid
	}
	err = ChessLib.validateCoordinates(figureCoordinates)
	if err != nil {
		return false, ErrBishopCoordinatesNotValid
	}
	return false, nil
}

// validateCoordinates Validate coordinates for figure
func (ChessLib *ChessLib) validateCoordinates(figureCoordinates string) error {
	return nil
}
