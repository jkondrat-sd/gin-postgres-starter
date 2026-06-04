// Why: logger setup gives the app one shared structured logger configured by environment.
// What to do: initialize logging once at startup, then inject or import Log where needed.
package logger

import "go.uber.org/zap"

var Log *zap.Logger

func Init(env string) {
	var err error
	if env == "production" {
		Log, err = zap.NewProduction()
	} else {
		Log, err = zap.NewDevelopment()
	}
	if err != nil {
		panic(err)
	}
}
