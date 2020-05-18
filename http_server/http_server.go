package http_server

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	l "github.com/ermanimer/logger"
	c "github.com/ermanimer/server-monitor/configuration"
	db "github.com/ermanimer/server-monitor/database"
	s "github.com/ermanimer/server-monitor/statistics"
	"github.com/gin-gonic/gin"
)

func Start(co *c.HttpServer) error {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Static("/assets", "./http_server/assets")
	r.LoadHTMLGlob("./http_server/templates/*")
	r.GET("/", home)
	r.GET("/get_statistics", getStatistics)
	r.GET("/select_records/:start/:end", selectRecords)
	hs := http.Server{
		Addr:           fmt.Sprintf("%v:%v", co.IpAddress, co.Port),
		Handler:        r,
		ReadTimeout:    time.Duration(co.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(co.WriteTimeout) * time.Second,
		MaxHeaderBytes: co.MaxHeaderBytes,
	}
	l.Info("Http server started.")
	err := hs.ListenAndServe()
	l.Debug(err.Error())
	return errors.New("Http server stopped!")
}

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

func getStatistics(c *gin.Context) {
	c.JSON(http.StatusOK, s.GetStatistics())
}

func selectRecords(c *gin.Context) {
	start, _ := strconv.ParseInt(c.Param("start"), 10, 64)
	end, _ := strconv.ParseInt(c.Param("end"), 10, 64)
	rs, err := db.SelectRecords(start, end)
	if err != nil {
		l.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, rs)
}
