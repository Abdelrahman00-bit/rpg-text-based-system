package main

import (
	"RPG/characters"
	"RPG/items"
	"fmt"
)

func main(){
// the start point of the game: user chose his char type
var char characters.CharacterMethods

for {

fmt.Println("Choose your character:")
fmt.Println("1. Warrior")
fmt.Println("2. Mage")
fmt.Println("3. Archer")

var choice int
fmt.Scanln(&choice)


if choice == 1 {
    char = characters.NewWarrior("Player1")
} else if choice == 2 {
    char = characters.NewMage("Player1")
} else if choice == 3 {
    char = characters.NewArcher("Player1")
} else {
    fmt.Println("Invalid choice")
	continue
}
 break
   }
// now the  stage 2: user will chose his char weapon.
fmt.Println("Choose a weapon:")
fmt.Println("1. Sword")
fmt.Println("2. Staff")
fmt.Println("3. Bow")

var wChoice int
fmt.Scanln(&wChoice)

var weapon items.Weapon

switch wChoice {
case 1:
    weapon = items.Sword{Name: "Irown Sword"}
case 2:
    weapon = items.Staff{Name: "Magic Staff"}
case 3:
    weapon = items.Bow{Name: "Wooden Bow"}
}

if err := char.EquipWeapon(weapon); err != nil {
        fmt.Println(err)
    }

    // Choose armor
    fmt.Println("Choose an armor:")
    fmt.Println("1. Plate Armor")
    fmt.Println("2. Magic Robe")
    fmt.Println("3. Leather Armor")
    var aChoice int
    fmt.Scanln(&aChoice)

    var armor items.Armor
    switch aChoice {
    case 1:
        armor = items.PlateArmor{Name: "Plate Armor", DefenseBonus: 5}
    case 2:
        armor = items.Robe{Name: "Magic Robe", DefenseBonus: 2}
    case 3:
        armor = items.LeatherArmor{Name: "Leather Armor", DefenseBonus: 3}
    }

    if err := char.EquipArmor(armor); err != nil {
        fmt.Println(err)
    }

    // Show stats
    char.ShowStats()

}

