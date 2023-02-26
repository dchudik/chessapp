package chesslib

import "errors"

// Errors coordinates by diffrents figures.
// Errors returned by king danger calculator
var (
	ErrCoordinatesNotValid = errors.New("figure coordinates not valid") // "figure coordinates not valid"

	ErrKingCoordinatesNotValid   = errors.New("king coordinates not valid")   // "king coordinates not valid"
	ErrRookCoordinatesNotValid   = errors.New("rook coordinates not valid")   // "rook coordinates not valid"
	ErrBishopCoordinatesNotValid = errors.New("bishop coordinates not valid") // "bishop coordinates not valid"
)
