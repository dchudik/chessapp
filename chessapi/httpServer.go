package chessapi

import (
	"net/http"
	"os"

	"github.com/dchudik/chessapp/chesslib"
)

func RunHttpServer() error {
	handlers := &Handlers{
		ChessCalculator: chesslib.NewChessKingUnderDangerCalculator(),
	}
	http.HandleFunc("/api/v1/king_attck", handlers.IsKingUnderAttackHandler)
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(":"+port, nil)
	return err
}
