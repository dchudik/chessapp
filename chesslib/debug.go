package chesslib

import "os"

func getDebugMode() bool {
	isDebug := os.Getenv("DEBUG")
	return isDebug == "1"
}
