package database

type Record struct {
	Id          int     `db:"id" json:"id"`
	CpuUsage    float64 `db:"cpu_usage" json:"cpu_usage"`
	MemoryUsage float64 `db:"memory_usage" json:"memory_usage"`
	DiskUsage   float64 `db:"disk_usage" json:"disk_usage"`
	Timestamp   int64   `db:"timestamp" json:"timestamp"`
}
