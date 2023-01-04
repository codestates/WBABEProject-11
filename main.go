package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"WBABEProject-11/config"
	ctl "WBABEProject-11/controller"
	"WBABEProject-11/logger"
	"WBABEProject-11/model"
	rt "WBABEProject-11/router"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {
	var configFlag = flag.String("config", "./confing/config.toml", "toml file to use for configuration")
	flag.Parse()
	cf := config.NewConfig(*configFlag)

	if err := logger.InitLogger(cf); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	logger.Debug("ready server....")

	if mod, err := model.NewModel(); err != nil {
		panic(err)
	} else if controller, err := ctl.NewCTL(mod); err != nil {
		panic(fmt.Errorf("controller.New > %v", err))
	} else if rt, err := rt.NewRouter(controller); err != nil {
		panic(fmt.Errorf("router.NewRouter > %v", err))
	} else {
		mapi := &http.Server {
			Addr: ":8080",
			Handler: rt.Index(),
			ReadTimeout: 5 * time.Second,
			WriteTimeout: 10 * time.Second,			
			MaxHeaderBytes: 1 << 20,
		}
	
		g.Go(func() error {
			return mapi.ListenAndServe()
		})

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		logger.Warn("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := mapi.Shutdown(ctx); err != nil {
			logger.Error("Server Shutdown:", err)
		}

		select {
		case <-ctx.Done():
			logger.Info("timeout of 5 seconds.")
		}

		logger.Info("Server exiting")
	}

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}