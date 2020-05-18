package configuration

type Configuration struct {
	Statistics *Statistics `yaml:"statistics"`
	Recorder   *Recorder   `yaml:"recorder"`
	HttpServer *HttpServer `yaml:"http_server"`
}

type Statistics struct {
	Interval         int    `yaml:"interval"`
	CpuUsageInterval int    `yaml:"cpu_usage_interval"`
	DiskUsagePath    string `yaml:"disk_usage_path"`
}

type Recorder struct {
	Interval int `yaml:"interval"`
}

type HttpServer struct {
	IpAddress      string `yaml:"ip_address"`
	Port           int    `yaml:"port"`
	ReadTimeout    int    `yaml:"read_timeout"`
	WriteTimeout   int    `yaml:"write_timeout"`
	MaxHeaderBytes int    `yaml:"max_header_bytes"`
}
