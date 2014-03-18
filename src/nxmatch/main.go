package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"
	"log"
	"fmt"
)

type Team struct {
	Name     string
	Location string
}

func NewTeam(name string) Team {
	return Team{Name: name, Location: ""}
}

type League struct {
	name  string
	teams []Team
}
type Score struct {
	HomeScore, VisitorScore int
	Home, Visitor           Team
}

type Match struct {
	Time     time.Time
	Scores   []Score
	Location string
}
type Season struct {
	matches []Match
}

func (s Score) AddPoints(hs, vs int) Score {
	return Score{Home: s.Home, Visitor: s.Visitor, HomeScore: hs, VisitorScore: vs}
}
func (m *Match) Add(s Score) {
	m.Scores = append(m.Scores, s)
}
func (m *Match) AddScore(hs, vs int) {
	s := m.Scores[0]
	m.Scores = append(m.Scores, s.AddPoints(hs, vs))

}

func NewMatch(l, v Team, time time.Time, where string) *Match {
	m := &Match{
		Time:     time,
		Scores:   make([]Score, 0),
		Location: where,
	}

	m.Add(Score{Home: l, Visitor: v})

	return m
}
func NewResult(h string, hs int, v string, vs int, when string, where string) *Match {
	time, e := time.Parse("Monday, January 02 2006, 03:04 PM MST", when)
	if e != nil {
		log.Print(e.Error())
	}
	m := NewMatch(NewTeam(h), NewTeam(v), time, where)
	m.AddScore(hs, vs)
	return m
}
func (m *Match) String() string {
	return fmt.Sprintf("%s %s at %s", m.Scores[len(m.Scores)-1], m.Time, m.Location)
}
func (s Score) String() string {
	return fmt.Sprintf("%s %d - %d %s", s.Home.Name, s.HomeScore, s.VisitorScore, s.Visitor.Name)
}

func main() {

	log.Println("Serving...")
	http.HandleFunc("/list", handler)
	http.ListenAndServe(":6060", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving %s to: %s",  r.RequestURI, r.RemoteAddr)
	results := []*Match{
		NewResult("BobCats", 101, "Bucks", 92, "Sunday, March 16 2014, 12:00 PM CST", "BMO Harris Bradley Center, Milwaukee, Wisconsin"),
		NewResult("Rockets", 104, "Heat", 103, "Sunday, March 16 2014, 02:30 PM CST", "AmericanAirlines Arena, Miami, Florida"),
		NewResult("Mavericks", 109, "Thunder", 86, "Sunday, March 16 2014, 06:00 PM CST", "Chesapeake Energy Arena, Oklahoma City, Oklahoma"),
		NewResult("Kings", 102, "Timberwolves", 104, "Sunday, March 16 2014, 06:00 PM CST", "Target Center, Minneapolis, Minnesota"),
		NewResult("Jazz", 104, "Spurs", 122, "Sunday, March 16 2014, 06:00 PM CST", "AT&T Center, San Antonio, Texas"),
		NewResult("Cavaliers", 80, "Clippers", 102, "Sunday, March 16 2014, 08:30 PM CST", "Staples Center, Los Angeles, California"),
		NewResult("Suns", 121, "Raptors", 113, "Sunday, March 16 2014, 12:00 PM CST", "Air Canada Centre, Toronto, Ontario"),
	}
	j, _ := json.Marshal(results)

	t, err := template.ParseFiles("webapp/src/templates/list.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, template.JS(j))

}
