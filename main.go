package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"go-project/common"
	"go-project/internal/handler/http/api"
	"net/http"
	"sync"
)

var (
	counter = 0
	lock    sync.Mutex
)

type Worker struct {
	id int
}

func main() {
	common.InitEnv()
	common.InitFileStorage()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	api.Routes(r)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		common.SendSuccessJsonResponse(w, "Hello")
	})

	r.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		common.SendErrorResponse(w, "Something went wrong")
	})

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}

func incr() {
	lock.Lock()
	defer lock.Unlock()
	counter++
	fmt.Println(counter)
}

func (w Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("обработчик %d получил %d\n", w.id, data)
	}
}
