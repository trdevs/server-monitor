package recorder

import (
	"time"

	l "github.com/ermanimer/logger"
	c "github.com/ermanimer/server-monitor/configuration"
	db "github.com/ermanimer/server-monitor/database"
	s "github.com/ermanimer/server-monitor/statistics"
)

func Start(co *c.Recorder) {
	for {
		time.Sleep(time.Duration(co.Interval) * time.Millisecond)
		st := s.GetStatistics()
		err := db.InsertRecord(st.CpuUsage, st.MemoryUsage, st.DiskUsage)
		if err != nil {
			l.Debug(err.Error)
			continue
		}
	}
}
