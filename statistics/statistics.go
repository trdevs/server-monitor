package statistics

import (
	"time"

	l "github.com/ermanimer/logger"
	c "github.com/ermanimer/server-monitor/configuration"
	s "github.com/ermanimer/statistics"
)

var st Statistics

func UpdateStatistics(c *c.Statistics) {
	for {
		time.Sleep(time.Duration(c.Interval) * time.Millisecond)
		cpuUsage, err := s.CpuUsage(c.CpuUsageInterval)
		if err != nil {
			l.Debug(err.Error())
			continue
		}
		st.CpuUsage = cpuUsage
		memoryUsage, err := s.MemoryUsage()
		if err != nil {
			l.Debug(err.Error())
			continue
		}
		st.MemoryUsage = memoryUsage
		diskUsage, err := s.DiskUsage(c.DiskUsagePath)
		if err != nil {
			l.Debug(err.Error())
			continue
		}
		st.DiskUsage = diskUsage
	}
}

func GetStatistics() *Statistics {
	return &st
}
