package main

import "math/rand"

func RollDice() int {
    return rand.Intn(6) + 1 // Random number between 1 and 6
}
