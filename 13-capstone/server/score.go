package main

import (
	"html/template"
	"net/http"
	"sort"
	"time"
)

type Score struct {
	name   string
	points int
	ts     time.Time
}

func (s *Score) Name() string      { return s.name }
func (s *Score) Points() int       { return s.points }
func (s *Score) Timestamp() string { return s.ts.Format("3:04:05PM") }

type ScoreKeeper map[string]Score

func (sk ScoreKeeper) UpdateScore(t *Task) {
	s := Score{t.name, t.points, time.Now()}
	currentScore := sk[s.name]
	if s.points > currentScore.points {
		sk[t.name] = s
	}
}

type scores []Score

func (s scores) Len() int      { return len(s) }
func (s scores) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s scores) Less(i, j int) bool {
	if s[i].points == s[j].points {
		return s[i].ts.Before(s[j].ts)
	}
	return s[i].points > s[j].points
}

var t = template.Must(template.ParseFiles("status.html"))

func (sk ScoreKeeper) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	var t = template.Must(template.ParseFiles("status.html"))
	ss := scores{}
	for _, s := range sk {
		ss = append(ss, s)
	}
	sort.Sort(ss)

	t.Execute(wr, ss)
}
