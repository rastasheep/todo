package controllers

import (
	"encoding/json"
	"github.com/codegangsta/martini"
	"net/http"
	"strconv"

	. "github.com/rastasheep/todo/models"
)

var manager = NewTaskManager()
var task *Task

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var response struct{ Tasks []*Task }
	response.Tasks = manager.Tasks()
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "oops", http.StatusInternalServerError)
	}
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "oops", http.StatusInternalServerError)
	}
	task = nil
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var request struct{ Title string }

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newTask, err := NewTask(request.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	manager.Add(newTask)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task

	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "decode failed", http.StatusBadRequest)
		return
	}

	if newTask.Id != task.Id {
		http.Error(w, "inconsistent task ID", http.StatusBadRequest)
		return
	}
	manager.Add(&newTask)
	task = nil
}

func FindTask(w http.ResponseWriter, r *http.Request, params martini.Params) {
	id, err := strconv.ParseInt(params["id"], 10, 0)
	if err != nil {
		http.Error(w, "task ID is not a number", http.StatusBadRequest)
		return
	}

	var ok bool
	task, ok = manager.Find(id)
	if !ok {
		http.NotFound(w, r)
		return
	}
}
