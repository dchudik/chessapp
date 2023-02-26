package chesslib

// ChessLib Chess library implementation
type ChessLib struct{}

// ChessKingUnderDangerCalculator Calculator for danger to king from figures
type ChessKingUnderDangerCalculator interface {
	IsKingUnderAttackRook(figureCoordinates, kingCoordinates string) bool
	IsKingUnderAttackBishop(figureCoordinates, kingCoordinates string) bool
}
