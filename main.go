package main

import (
	"os"
	"path/filepath"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var (
	bin_path     string
	config_path  string
	logfile_path string
	db_path      string
)

type Config struct {
	ListenOn string
}

var config *Config

func init() {
	var err error

	bin_path, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		Warning.Println(err)
	}

	config = &Config{
		ListenOn: "127.0.0.1:6661",
	}

	logfile_path = "/var/log/awesome_server"
	config_path = "/etc/fdb/config.ini"
	db_path = "/etc/awesome_server/dbs"

	err = MakeAllDirs(config_path)
	if err != nil {
		if 1 == 1 {
			Warning.Println(err)
		}

	}

	err = MakeAllDirs(db_path)
	if err != nil {
		Warning.Println(err)
	}

	SetupLogging()
}

func InitRouteMethods(router *fasthttprouter.Router) {
	GetRoutes(router)
	PostRoutes(router)
	PutRoutes(router)
	DeleteRoutes(router)
}

func main() {
	var router *fasthttprouter.Router
	router = fasthttprouter.New()

	InitRouteMethods(router)

	fasthttp.ListenAndServe(config.ListenOn, router.Handler)
}
