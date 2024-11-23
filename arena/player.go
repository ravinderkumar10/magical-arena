package main

type Player struct {
    Name     string
    Health   int
    Strength int
    Attack   int
}

func NewPlayer(name string, health, strength, attack int) Player {
    return Player{
        Name:     name,
        Health:   health,
        Strength: strength,
        Attack:   attack,
    }
}
