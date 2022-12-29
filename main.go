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
	/* [코드리뷰]
	 * 시스템의 첫  시작인 main 함수를 간결하게 잘 짜주셨네요
	 * 해당 시스템의 config를 함께 작성해주시면, 더 견고한 코드로 발전할 수 있을 것으로 보여집니다.
	 * command 라인에서 config를 별도로 지정할 수 있는 flag 방식을 사용하여 보세요.
	 * ex, 
	 * var configFlag = flag.String("config", default string value, "description")
	 */
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
	/* [코드리뷰]
	 * panic 함수를 사용하여, 각 상황에 대한 에러를 메세지와 함께 잘 정리해주셨습니다.
	 * 그러나 panic 함수 이외의 구문은 더 이상 코드가 동작하지 않게 되는데요,
	 * 실제로 운영되는 API 서비스라면, 서비스가 특정 상황으로 인해 종료되는 일은 발생하지 말아야 합니다.
	 * 예외처리를 통해 적절한 메세지를 개발자에게 전달하는 방식으로 수정해보시는 방법 또한 추천드립니다.
	 */
}