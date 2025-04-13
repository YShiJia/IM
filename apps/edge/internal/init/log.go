package init

import (
	log "github.com/sirupsen/logrus"
)

func InitLog() error {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetReportCaller(true)
	log.SetLevel(log.DebugLevel)
	return nil
}
