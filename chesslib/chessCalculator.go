package chesslib

func NewChessKingUnderDangerCalculator() *ChessLib {
	return &ChessLib{}
}

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

func (ChessLib *ChessLib) validateCoordinates(figureCoordinates string) error {
	return nil
}
