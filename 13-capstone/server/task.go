package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

type Task struct {
	name   string
	points int
}

type TaskMaster struct {
	sync.Mutex
	tasks struct {
		active, completed map[string]Task
	}
}

func NewTaskMaster() *TaskMaster {
	tm := TaskMaster{}
	tm.tasks.active = map[string]Task{}
	tm.tasks.completed = map[string]Task{}
	return &tm
}

func (tm *TaskMaster) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "":
		fallthrough
	case "/":
		writeDocumentation(wr)
	case "new":
		tm.newToken(wr, req)

	default:
		if _, err := uuid.Parse(req.URL.Path); err != nil {
			wr.WriteHeader(http.StatusNotFound)
		} else {
			switch req.Method {
			case "GET":
				tm.taskStatus(wr, req.URL.Path)
			case "POST":
				tm.completeTask(wr, req.URL.Path)
			default:
			}
		}
	}
}

func (tm *TaskMaster) newToken(wr http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	if name == "" {
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	token := uuid.New().String()
	tm.Lock()
	tm.tasks.active[token] = Task{name: name}
	tm.Unlock()

	fmt.Fprintln(wr, token)
}

func (tm *TaskMaster) completeTask(wr http.ResponseWriter, token string) {
	tm.Lock()
	defer tm.Unlock()
	_, active := tm.tasks.active[token]
	_, completed := tm.tasks.completed[token]
	switch {
	case active:
		tm.tasks.completed[token] = tm.tasks.active[token]
		delete(tm.tasks.active, token)
	case completed:
		wr.WriteHeader(http.StatusGone)
	default:
		wr.WriteHeader(http.StatusNotFound)
	}
}

func (tm *TaskMaster) taskStatus(wr http.ResponseWriter, token string) {
	tm.Lock()
	defer tm.Unlock()
	_, active := tm.tasks.active[token]
	_, completed := tm.tasks.completed[token]
	switch {
	case active:
	case completed:
	}
}
