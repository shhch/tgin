package tool

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var Log *logrus.Logger

func init()  {
	Log = logrus.New()
	logFile, _ := os.OpenFile("/var/log/server/server.log",  os.O_CREATE|os.O_WRONLY, 0666)
	Log.Out = logFile
	Log.Formatter = &logrus.JSONFormatter{DisableTimestamp: true}
	Log.Level = logrus.DebugLevel
	io.MultiWriter()

}
