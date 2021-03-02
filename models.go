package main

import "fmt"

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
