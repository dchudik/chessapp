package chesslib

import "log"

type Logger interface {
	DebugLog(figureType, figureCoordinates, kingCoordinates string,
		isUnderAttack bool, err error)
}

type LoggerImpl struct {
	isEnableDebug bool
}

func NewLogger(isDebug bool) *LoggerImpl {
	return &LoggerImpl{
		isEnableDebug: isDebug,
	}
}

func (LoggerImpl *LoggerImpl) DebugLog(figureType, figureCoordinates, kingCoordinates string,
	isUnderAttack bool, err error) {
	if LoggerImpl.isEnableDebug {
		log.Printf("Figure type: %s coordinates: %s. King coordinates: %s. Is under attack: %v. Error: %s",
			figureType, figureCoordinates, kingCoordinates, isUnderAttack, err)
	}
}
