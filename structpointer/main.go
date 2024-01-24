package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

type Gamer struct {
	Name  string
	Games []string
}

func (student *Student) Graduate() {
	student.Name = student.Name + " S.T"
}

func (gamer *Gamer) AddGame(addedGame string) {
	gamer.Games = append(gamer.Games, addedGame)
}

func main() {
	// student := Student{1, "Jeremy", 2.98}
	// fmt.Println(student.Name)

	// student.Graduate()

	// fmt.Println(student.Name)
	gamer := Gamer{Name: "Jeremy"}

	gamer.AddGame("Valorant")
	gamer.AddGame("Hogwarts Legacy")
	gamer.AddGame("Lethal Company")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}
