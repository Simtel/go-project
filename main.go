package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"go-project/common"
	"go-project/domains"
	"math/rand"
	"net/http"
	"sync"
	"time"
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

	domains.Routes(r)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		common.SendSuccessJsonResponse(w, "Hello")
	})

	r.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		common.SendErrorResponse(w, "Something went wrong")
	})

	r.Get("/lock", func(w http.ResponseWriter, r *http.Request) {

		for i := 0; i < 2; i++ {
			go incr()
		}
		time.Sleep(time.Millisecond * 10)

	})

	r.Get("/channel", func(w http.ResponseWriter, r *http.Request) {
		c := make(chan int)
		for i := 0; i < 4; i++ {
			worker := Worker{id: i}
			go worker.process(c)
		}
		for {
			c <- rand.Int()
			time.Sleep(time.Millisecond * 50)
		}
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
