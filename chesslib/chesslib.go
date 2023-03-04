package chesslib

// Types of figures
const (
	FigureTypeKing   = "KING"
	FigureTypeBishop = "BISHOP"
	FigureTypeRook   = "ROOK"
)

// Chess library implementation
type ChessLib struct {
	Debug  bool
	Logger Logger
}

// Calculator for danger to king from figures
type ChessKingUnderDangerCalculator interface {
	IsKingUnderAttackRook(figureCoordinates, kingCoordinates string) (bool, error)
	IsKingUnderAttackBishop(figureCoordinates, kingCoordinates string) (bool, error)
}
