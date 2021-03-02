package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	Welcome         string
	Health          int
	CurrentLocation string
}

type Location struct {
	Description string
	Transitions []string
	Events      []string
}

type Event struct {
	Type        string
	Chance      int
	Description string
	Health      int
	Evt         string
}

type Character struct {
	Name    string
	Health  int
	Evasion int
	Alive   bool
	Speed   int
	Weap    int
	Npc     bool
}

type Weapon struct {
	Min  int
	Max  int
	Name string
}

func (w *Weapon) Fire() int {
	return w.Min + rand.Intn(w.Max-w.Min)
}

func (p *Character) Equip(w int) {
	p.Weap = w
}

func (p *Character) Attack() int {
	return Weaps[p.Weap].Fire()
}

type Players []Character

func (slice Players) Len() int {
	return len(slice)
}

func (slice Players) Less(i, j int) bool {
	return slice[i].Speed > slice[j].Speed
}

func (slice Players) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
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

func (g *Game) ProcessEvents(events []string) {
	for _, evtName := range events {
		g.Health += evts[evtName].ProcessEvent()
	}
}

func (g *Game) Play() {
	CurrentLocation := locationMap["Bridge"]
	fmt.Println(g.Welcome)
	for {
		fmt.Println(CurrentLocation.Description)
		g.ProcessEvents(CurrentLocation.Events)
		if g.Health <= 0 {
			fmt.Println("You are dead. Game Over.")
			return
		}

		fmt.Printf("Health: %d\n", g.Health)
		fmt.Println("You can go these places:")
		fmt.Println("\t0 - Quit")
		for index, loc := range CurrentLocation.Transitions {
			fmt.Printf("\t%d - %s\n", index+1, loc)
		}
		i := 0
		for i < 1 || i > len(CurrentLocation.Transitions) {
			fmt.Println("Where do you want to go?")
			fmt.Scan(&i)
		}
		newLoc := i - 1
		CurrentLocation = locationMap[CurrentLocation.Transitions[newLoc]]
	}
}
