package test

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogging(t *testing.T) {
	logger := logrus.New()
	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warning")
	logger.Error("Error")
}

func TestLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warning")
	logger.Error("Error")
}

func TestOutput(t *testing.T) {
	logger := logrus.New()
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Info("Info")
	logger.Warn("Warning")
	logger.Error("Error")
}

func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Info")
	logger.Warn("Warning")
	logger.Error("Error")
}

func TestField(t *testing.T) {
	logger := logrus.New()

	logger.WithField("username", "faridlan").Info("Hello Logging")
	logger.WithField("username", "faridlan").
		WithField("name", "NulHakim").
		Info("Hello Logging")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "faridlan",
		"name":     "NulHakim",
	}).Info("Hello Logging")
}

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)
	entry.WithField("username", "faridlan")
	entry.Info("Hello Logging")
}
