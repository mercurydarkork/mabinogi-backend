package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/op/go-logging"
	"github.com/rs/cors"
)

var router = httprouter.New()
var laddr string
var resource string
var logfile, level string
var logFormat = "%{color}%{time:2006-01-02T15:04:05} %{shortfile} %{longfunc} %{callpath} %{level} %{id}%{color:reset} %{message}"
var log = logging.MustGetLogger("common")

func init() {
	flag.StringVar(&laddr, "laddr", ":3216", "服务器监听地址")
	flag.StringVar(&resource, "resource", "./resource", "资源目录地址")
	flag.StringVar(&configFile, "c", "./training_grounds.conf", "配置文件")
	flag.StringVar(&env, "env", "dev", "配置文件中环境配置")
	flag.StringVar(&logfile, "logfile", "./training_grounds.log", "日志文件")
	flag.StringVar(&level, "level", "INFO", "日志级别")
	flag.Parse()
	router.PanicHandler = panicHandler
	router.GET("/training_grounds/team", getTeam)
	router.POST("/training_grounds/team", updateTeam)
	router.GET("/training_grounds/game", getGame)
	router.POST("/training_grounds/game", updateGame)
	router.GET("/veryDifficult_apostle/team", getApostleTeam)
	router.POST("/veryDifficult_apostle/team", updateApostleTeam)
}

func panicHandler(w http.ResponseWriter, _ *http.Request, err interface{}) {
	log.Error(err)
	renderError(w, "Internal Server Error", 500)
}
func renderJSON(w http.ResponseWriter, ret interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	b, _ := json.Marshal(ret)
	w.Write(b)
}
func renderError(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", msg)))
}
func renderRedirect(w http.ResponseWriter, location string, code int) {
	w.Header().Set("Location", location)
	w.WriteHeader(code)
}

func main() {
	sysf, err := rotatelogs.New(
		logfile,
		rotatelogs.WithMaxAge(24*7*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		log.Error(err)
		return
	}
	lev, _ := logging.LogLevel(level)
	filelog := logging.AddModuleLevel(logging.NewBackendFormatter(logging.NewLogBackend(sysf, "", 0), logging.MustStringFormatter(logFormat)))
	filelog.SetLevel(lev, "common")
	logging.SetBackend(filelog)
	initConfig()
	h := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler(router)
	log.Info(http.ListenAndServe(laddr, h))
}
