package database

import (
	"errors"
	"time"

	l "github.com/ermanimer/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DatabaseFile = "./database/records.db"
)

var db *sqlx.DB

func Open() error {
	var err error
	db, err = sqlx.Open("sqlite3", DatabaseFile)
	if err != nil {
		l.Debug(err.Error())
		return errors.New("Opening database failed!")
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS records (id INTEGER PRIMARY KEY, cpu_usage REAL, memory_usage REAL, disk_usage REAL, timestamp INTEGER);")
	if err != nil {
		l.Debug(err.Error())
		return errors.New("Opening database failed!")
	}
	return nil
}

func Close() {
	if db == nil {
		return
	}
	db.Close()
}

func InsertRecord(cpuUsage float64, memoryUsage float64, diskUsage float64) error {
	timestamp := time.Now().Unix()
	_, err := db.Exec("INSERT INTO records(cpu_usage, memory_usage, disk_usage, timestamp) VALUES(?, ?, ?, ?);", cpuUsage, memoryUsage, diskUsage, timestamp)
	if err != nil {
		l.Debug(err.Error())
		return errors.New("Inserting record failed!")
	}
	return nil
}

func SelectRecords(start int64, end int64) ([]Record, error) {
	var rs []Record
	err := db.Select(&rs, "SELECT * FROM records WHERE timestamp>=? AND timestamp<?;", start, end)
	if err != nil {
		l.Debug(err.Error())
		return nil, errors.New("Selecting records failed!")
	}
	return rs, nil
}
