package plain_di_logger

import (
	seelog "github.com/cihub/seelog"
)

var log seelog.LoggerInterface = seelog.Disabled

func DisableLog() {
	log = seelog.Disabled
}

func UseLogger(logger seelog.LoggerInterface) {
	log = logger
}

func FlushLog() {
	log.Flush()
}
