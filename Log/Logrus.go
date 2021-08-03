package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func main() {
	log.WithFields(
		log.Fields{
			"animal": "walrus",
			"size":   10,
		}).Info("A group of walrus emerges form the ocean")

	log.WithFields(
		log.Fields{
			"omg":    true,
			"number": 122,
		}).Warn("The group`s number increased tremendously!")

	log.WithFields(
		log.Fields{
			"omg":    true,
			"number": 122,
		}).Fatal("The ice breaks!")

	contextLogger := log.WithFields(
		log.Fields{
			"common": "This is a common field",
			"other":  "I also should be logged always",
		})

	contextLogger.Info("I`ll be logged with common and other fields")
	contextLogger.Info("Me too")
}
