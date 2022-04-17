package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/wallissonmarinho/go-gin-boilerplate/internal/client"
	"github.com/wallissonmarinho/go-gin-boilerplate/internal/endpoint"
	"github.com/wallissonmarinho/go-gin-boilerplate/internal/service"
	trans "github.com/wallissonmarinho/go-gin-boilerplate/internal/transport/http"
)

func init() {

	viper.SetConfigName("boilerplate.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".")
	errViper := viper.ReadInConfig()
	if errViper != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", errViper))
	}

}

const (
	banner = `
	____   ____ _____ _      ______ _____  _____  _            _______ ______ 
	|  _ \ / __ \_   _| |    |  ____|  __ \|  __ \| |        /\|__   __|  ____|
	| |_) | |  | || | | |    | |__  | |__) | |__) | |       /  \  | |  | |__   
	|  _ <| |  | || | | |    |  __| |  _  /|  ___/| |      / /\ \ | |  |  __|  
	| |_) | |__| || |_| |____| |____| | \ \| |    | |____ / ____ \| |  | |____ 
	|____/ \____/_____|______|______|_|  \_\_|    |______/_/    \_\_|  |______|
																			   
	`
)

func main() {

	logrus.Info(banner)

	// initialize our OpenCensus configuration and defer a clean-up
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	var db *sqlx.DB
	{
		var err error
		db, err = sqlx.Open("postgres", "host=localhost port=5432 user=zintech dbname=zintech password=zintech00 sslmode=disable TimeZone=America/Fortaleza")
		if err != nil {
			os.Exit(-1)
		}
	}

	level.Info(logger)
	defer level.Info(logger)

	var (
		context    context.Context
		clients    = client.NewClients()
		services   = service.NewServiceFactory(db, logger, clients)
		endpoint   = endpoint.MakeEndpoints(services, logger)
		serverHTTP = trans.NewService(context, db, &endpoint, &logger)
		httpAddr   = flag.String("http.addr", ":1707", "HTTP listen address")
		err        = make(chan error)
	)

	go func() {
		server := &http.Server{
			Addr:         *httpAddr,
			Handler:      serverHTTP,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
		}
		err <- server.ListenAndServe()
	}()

	fatal := level.Error(logger).Log("exit", <-err)
	if fatal != nil {
		logrus.Error(fatal)
	}

}
