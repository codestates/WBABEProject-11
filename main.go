package main

import (
	"fmt"
	"net/http"
	"time"

	ctl "WBABEProject-11/controller"
	"WBABEProject-11/model"
	rt "WBABEProject-11/router"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {

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
	}

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}