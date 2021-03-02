package main

import (
	"fmt"
	"math/rand"
	"time"
)

var locationMap = map[string]*Location{
	"Bridge":      {"You are on the bridge of a spaceship sitting in the Captain's chair.", []string{"Ready Room", "Turbo Lift"}, []string{"alienAttack"}},
	"Ready Room":  {"The Captain's ready room.", []string{"Bridge"}, []string{}},
	"Turbo Lift":  {"A Turbo Lift that takes you anywhere in the ship.", []string{"Bridge", "Lounge", "Engineering"}, []string{"android"}},
	"Engineering": {"You are in engineering where you see the star drive", []string{"Turbo Lift"}, []string{"alienAttack"}},
	"Lounge":      {"You are in the lounge, you feel very relaxed", []string{"Turbo Lift"}, []string{"relaxing"}},
}

var evts = map[string]*Event{
	"alienAttack":     {Chance: 20, Description: "An alien beams in front of you and shoots you with a ray gun.", Health: -50, Evt: "doctorTreatment"},
	"doctorTreatment": {Chance: 10, Description: "The doctor rushes in and inject you with a health boost.", Health: +30, Evt: ""},
	"android":         {Chance: 50, Description: "Data is in the turbo lift and says hi to you", Health: 0, Evt: ""},
	"relaxing":        {Chance: 100, Description: "In the lounge you are so relaxed that your health improves.", Health: +10, Evt: ""},
}

func (e *Event) ProcessEvent() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	if e.Chance >= r1.Intn(100) {
		hp := e.Health
		if e.Type == "Combat" {
			fmt.Println("Combat Event")
		}
		fmt.Printf("\t%s\n", e.Description)
		if e.Evt != "" {
			hp = hp + evts[e.Evt].ProcessEvent()
		}

		return hp
	}

	return 0
}

func main() {
	g := &Game{Health: 100, Welcome: "Welcome to the Starship Enterprise"}
	g.Play()
}
