package chesslib

func NewChessKingUnderDangerCalculator() *ChessLib {
	return &ChessLib{}
}

func (ChessLib *ChessLib) IsKingUnderAttackRook(figureCoordinates, kingCoordinates string) bool {
	return false
}

func (ChessLib *ChessLib) IsKingUnderAttackBishop(figureCoordinates, kingCoordinates string) bool {
	return false
}

func (ChessLib *ChessLib) isCoordinatesValid(figureCoordinates string) bool {
	return true
}
