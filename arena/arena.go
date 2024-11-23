package main

import "fmt"

type Arena struct{}

func (a Arena) Fight(player1, player2 *Player) {
    fmt.Printf("Fight between %s and %s begins!\n", player1.Name, player2.Name)

    for player1.Health > 0 && player2.Health > 0 {
        if player1.Health < player2.Health {
            a.performAttack(player1, player2)
        } else {
            a.performAttack(player2, player1)
        }
    }

    if player1.Health <= 0 {
        fmt.Printf("%s has won the fight!\n", player2.Name)
    } else {
        fmt.Printf("%s has won the fight!\n", player1.Name)
    }
}

func (a Arena) performAttack(attacker, defender *Player) {
    attackRoll := RollDice()
    defenseRoll := RollDice()

    attackDamage := attacker.Attack * attackRoll
    defenseStrength := defender.Strength * defenseRoll

    damageToDefender := attackDamage - defenseStrength
    if damageToDefender > 0 {
        defender.Health -= damageToDefender
        fmt.Printf("%s attacks %s: Damage dealt = %d, %s's health = %d\n",
            attacker.Name, defender.Name, damageToDefender, defender.Name, defender.Health)
    } else {
        fmt.Printf("%s defends successfully against %s!\n", defender.Name, attacker.Name)
    }
}
