package log

import (
	"github.com/professionsforall/hexagonal-template/pkg/log/logger"
	"github.com/professionsforall/hexagonal-template/pkg/log/zap"
)

var Logger logger.AppLogger

func Apply() {
	Logger = zap.NewZapLogger()
}
