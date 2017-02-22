package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/google/uuid"
)

type Task struct {
	name    string
	points  int
	current struct {
		name string
		Problem
		got bool
	}
}

type TaskMaster struct {
	sync.Mutex
	active   map[string]*Task
	problems map[string]ProblemGenerator
	next     map[string]string
	sk       ScoreKeeper
}

func NewTaskMaster() *TaskMaster {
	tm := TaskMaster{
		active:   map[string]*Task{},
		problems: map[string]ProblemGenerator{},
		next:     map[string]string{},
		sk:       ScoreKeeper{},
	}
	return &tm
}

func (tm *TaskMaster) Mux() http.Handler {
	// replace with a muxer
	mux := http.NewServeMux()
	mux.HandleFunc("/", tm.postID)
	mux.HandleFunc("/new", tm.newToken)
	for key := range tm.problems {
		path := "/" + key + "/"
		mux.Handle(path, http.StripPrefix(path, tm.handleTask(key)))
	}

	return mux
}

func (tm *TaskMaster) postID(wr http.ResponseWriter, req *http.Request) {
	path := strings.TrimPrefix(req.URL.Path, "/")
	if path == "" {
		writeDocumentation(wr, req)
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	uuid, err := uuid.Parse(path)
	if err != nil {
		wr.WriteHeader(http.StatusNotFound)
		return
	}
	token := uuid.String()

	// parse out ID
	tm.Lock()
	defer tm.Unlock()

	task, ok := tm.active[token]
	if !ok || task.current.name != "" {
		wr.WriteHeader(http.StatusNotFound)
		return
	}

	tm.nextTask(token, task)
}

func (tm *TaskMaster) nextTask(token string, task *Task) {
	task.points++
	tm.sk.UpdateScore(task)

	next := tm.next[task.current.name]
	if next == "" {
		delete(tm.active, token)
		return
	}

	task.current.got = false
	task.current.name = next
	task.current.Problem = tm.problems[next].Generate()
}

// Responsible for points
func (tm *TaskMaster) handleTask(name string) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {
		path := strings.TrimPrefix(req.URL.Path, "/")
		log.Println(name, path)
		if path == "" {
			tm.problems[name].Documentation(wr)
			return
		}

		uuid, err := uuid.Parse(path)
		if err != nil {
			wr.WriteHeader(http.StatusNotFound)
			return
		}
		token := uuid.String()

		// parse out ID
		tm.Lock()
		defer tm.Unlock()

		task, ok := tm.active[token]
		if !ok || task.current.name != name {
			wr.WriteHeader(http.StatusNotFound)
			return
		}

		switch req.Method {
		case "GET":
			if !task.current.got {
				task.points++
				task.current.got = true
				tm.sk.UpdateScore(task)
			}
			task.current.Send(wr)
		case "POST":
			defer req.Body.Close()
			if !task.current.Validate(req.Body) {
				wr.WriteHeader(http.StatusNotAcceptable)
				return
			}
			tm.nextTask(token, task)

		default:
			tm.problems[name].Documentation(wr)
			wr.WriteHeader(http.StatusBadRequest)
		}
	}
}

func (tm *TaskMaster) newToken(wr http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	if name == "" {
		writeDocumentation(wr, req)
		wr.WriteHeader(http.StatusBadRequest)
		return
	}

	token := uuid.New().String()
	tm.Lock()
	tm.active[token] = &Task{name: name, points: 1}
	tm.Unlock()

	fmt.Fprint(wr, token)
}
