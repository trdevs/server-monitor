package main

import (
	l "github.com/ermanimer/logger"
	c "github.com/ermanimer/server-monitor/configuration"
	db "github.com/ermanimer/server-monitor/database"
	hs "github.com/ermanimer/server-monitor/http_server"
	r "github.com/ermanimer/server-monitor/recorder"
	s "github.com/ermanimer/server-monitor/statistics"
)

func main() {
	//initialize logger
	l.Initialize("default.log", l.DebugTraceLevel)
	//load configuration
	c, err := c.Load()
	if err != nil {
		l.Fatal(err.Error())
	}
	//open database
	err = db.Open()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer db.Close()
	//update statistics
	go s.UpdateStatistics(c.Statistics)
	//start recorder
	go r.Start(c.Recorder)
	//start http server
	l.Fatal(hs.Start(c.HttpServer))
}
