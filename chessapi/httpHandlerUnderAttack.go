package chessapi

import (
	"encoding/json"
	"net/http"

	"github.com/dchudik/chessapp/chesslib"
)

type Handlers struct {
	ChessCalculator chesslib.ChessKingUnderDangerCalculator
}

type RespIsKingUnderAttack struct {
	IsUnderAttack bool
	Error         string
}

func (Handlers *Handlers) IsKingUnderAttackHandler(w http.ResponseWriter, r *http.Request) {
	figureType := r.URL.Query().Get("figure_type")
	figureCoordinates := r.URL.Query().Get("figure_coordinates")
	kingCoordinates := r.URL.Query().Get("king_coordinates")
	isUnderAttack := false
	var err error
	switch figureType {
	case chesslib.FigureTypeBishop:
		isUnderAttack, err = Handlers.ChessCalculator.IsKingUnderAttackBishop(figureCoordinates, kingCoordinates)
	case chesslib.FigureTypeRook:
		isUnderAttack, err = Handlers.ChessCalculator.IsKingUnderAttackRook(figureCoordinates, kingCoordinates)
	default:
		err = ErrFigureTypeNotFound
	}
	if err != nil {
		resp, err := json.Marshal(RespIsKingUnderAttack{
			IsUnderAttack: isUnderAttack,
			Error:         err.Error(),
		})
		if err != nil {
		}
		w.Write(resp)
		return
	}
	resp, err := json.Marshal(RespIsKingUnderAttack{
		IsUnderAttack: isUnderAttack,
		Error:         "",
	})
	if err != nil {
	}
	w.Write(resp)
}
