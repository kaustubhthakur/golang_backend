package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kaustubhthakur/rest_api/controllers"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)

}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongo-url")
	if err != nil {
		panic(err)
	}

	return s
}
