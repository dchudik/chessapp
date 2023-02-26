package chesslib

// Chess library implementation
type ChessLib struct{}

// Calculator for danger to king from figures
type ChessKingUnderDangerCalculator interface {
	IsKingUnderAttackRook(figureCoordinates, kingCoordinates string) bool
	IsKingUnderAttackBishop(figureCoordinates, kingCoordinates string) bool
}
