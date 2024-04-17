package main

import (
	"github.com/sirupsen/logrus"
)

type MyHook struct {
}

func (mHook *MyHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.InfoLevel}
}

func (mHook *MyHook) Fire(entry *logrus.Entry) error {
	entry.Data["name"] = "app"
	return nil
}

func main() {
	logrus.AddHook(&MyHook{})
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.Errorln("logrus => Error")
	logrus.Warningln("logrus => Warn")
	logrus.Infoln("logrus => Info")
	logrus.Debugln("logrus => Debug")
}
