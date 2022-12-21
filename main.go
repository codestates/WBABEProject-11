// request -> router -> controller -> model -> controller -> response
package main

import (
	"fmt"
	"net/http"
	"time"

	rt "WBABEProject-11/router"

	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {
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

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}