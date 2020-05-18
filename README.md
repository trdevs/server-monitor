# server-monitor
Server monitor for Linux

![Go](https://github.com/ermanimer/server-monitor/workflows/Go/badge.svg)

## Features
server-monitor is a server monitoring tool for Linux. server-monitor records cpu usage, memory usage and disk usage data between configured intervals. server-monitor has a http server for user interface.

## Installation
```bash
go get -v github.com/ermanimer/server-monitor
```

## Configuration
./configuration/configuration.yaml
```yaml
statistics:
  interval: 1000
  cpu_usage_interval: 1000
  disk_usage_path: "/"
recorder:
  interval: 15000
http_server:
  ip_address: "127.0.0.1"
  port: 8000
  read_timeout: 10
  write_timeout: 10
  max_header_bytes: 1048576
```

## Usage
Create systemd service /lib/systemd/system/server-monitor.service
```bash
[Unit]
Description=Server Monitor
After=multi-user.target
Requires=multi-user.target

[Service]
Type=simple
Restart=always
RestartSec=5s
WorkingDirectory=/home/ermanimer/server-monitor/
ExecStart=/home/ermanimer/server-monitor/server-monitor

[Install]
WantedBy=multi-user.target
```
Enable and start service
```bash
systemctl enable server-monitor
systemctl start server-monitor
```
## User Interface
![Terminal Output](/images/user_interface.png)