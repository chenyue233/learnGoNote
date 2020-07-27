package seheduler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video_sever/videoServer/sehedulerer/seheduler/taskrunner"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video-delete-record/:vid-id",vidDelRecHandler)
	return router
}

func main()  {
	c := make(chan int)
	go taskrunner.Start()
	r := RegisterHandlers()
	<- c
	http.ListenAndServe(":9001",r)
}
