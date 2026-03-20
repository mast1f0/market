package logger

import log "github.com/sirupsen/logrus"

func Init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.DebugLevel)
}
