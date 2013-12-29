package main

import (
	"github.com/codegangsta/martini"

	. "github.com/rastasheep/todo/controllers"
)

func main() {
	m := martini.Classic()

	m.Get("/tasks", GetTasks)
	m.Post("/task", CreateTask)
	m.Get("/task/:id", FindTask, GetTask)
	m.Put("/task/:id", FindTask, UpdateTask)

	m.Run()
}
