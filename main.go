package main


//func main() {
//	fmt.Println("This works.")
//	r := router.InitRouter()
//	r.Run()
//}

import (
	"fmt"
	"github.com/nearbyren/gomodule/pkg/setting"
	"github.com/nearbyren/gomodule/routers"
	"net/http"

)

func main() {

	router := routers.InitRouter();

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}