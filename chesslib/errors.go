package chesslib

import "errors"

var ErrCoordinatesNotValid = errors.New("coordinates not valid")

var ErrKingCoordinatesNotValid = errors.New("king coordinates not valid")
var ErrRookCoordinatesNotValid = errors.New("rook coordinates not valid")
var ErrBishopCoordinatesNotValid = errors.New("bishop coordinates not valid")
