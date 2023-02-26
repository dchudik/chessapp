package chesslib

type ChessLib struct{}

type ChessKingUnderDangerCalculator interface {
	IsKingUnderAttackRook(figureCoordinates, kingCoordinates string) bool
	IsKingUnderAttackBishop(figureCoordinates, kingCoordinates string) bool
}
