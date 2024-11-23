package main

import "fmt"

func main() {
    var health, strength, attack int
    fmt.Println("Enter Player Stats:")
    fmt.Print("Health: ")
    fmt.Scan(&health)
    fmt.Print("Strength: ")
    fmt.Scan(&strength)
    fmt.Print("Attack: ")
    fmt.Scan(&attack)

    fmt.Printf("Player Stats - Health: %d, Strength: %d, Attack: %d\n", health, strength, attack)
}
