package chessapi

import (
	"net/http"

	"github.com/dchudik/chessapp/chesslib"
)

func runHttpServer() error {
	handlers := &Handlers{
		ChessCalculator: chesslib.NewChessKingUnderDangerCalculator(),
	}
	http.HandleFunc("/api/v1/king_attck", handlers.IsKingUnderAttackHandler)
	err := http.ListenAndServe(":8080", nil)
	return err
}
