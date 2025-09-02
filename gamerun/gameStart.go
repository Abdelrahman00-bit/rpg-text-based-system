package gamerun

import (
	"RPG/characters"
	"RPG/items"
	"fmt"
)

func GameStart() characters.CharacterMethods {

	// the start point of the game: user choose his char Name and char type
	var charName string
	fmt.Println("Choose your character name: ")

	fmt.Scanln(&charName)

	var char characters.CharacterMethods

	for {

		fmt.Println("Choose your character:")
		fmt.Println("1. Warrior")
		fmt.Println("2. Mage")
		fmt.Println("3. Archer")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			char = characters.NewWarrior(charName)
		case 2:
			char = characters.NewMage(charName)
		case 3:
			char = characters.NewArcher(charName)
		default:
			fmt.Println("Invalid choice")
			continue
		}
		break
	}

		// now the  stage 2: user will chose his char weapon.
	for {

		fmt.Println("Choose a weapon:")
		fmt.Println("1. Sword")
		fmt.Println("2. Staff")
		fmt.Println("3. Bow")

		var wChoice int
		fmt.Scanln(&wChoice)

		var weapon items.Weapon

		switch wChoice {
		case 1:
			weapon = &items.Sword{Name: "Irown Sword", DamageBonus: 10}
		case 2:
			weapon = &items.Staff{Name: "Magic Staff", DamageBonus: 5}
		case 3:
			weapon = &items.Bow{Name: "Wooden Bow", DamageBonus: 30}
		}

		if err := char.EquipWeapon(weapon); err != nil {
			fmt.Println(err)
			continue
		}
		break
	}

	// Choose armor
	for {
		fmt.Println("Choose an armor:")
		fmt.Println("1. Plate Armor")
		fmt.Println("2. Magic Robe")
		fmt.Println("3. Leather Armor")
		var aChoice int
		fmt.Scanln(&aChoice)

		var armor items.Armor
		switch aChoice {
		case 1:
			armor = &items.PlateArmor{Name: "Plate Armor", DefenseBonus: 10}
		case 2:
			armor = &items.Robe{Name: "Magic Robe", DefenseBonus: 8}
		case 3:
			armor = &items.LeatherArmor{Name: "Leather Armor", DefenseBonus: 5}
		}

		if err := char.EquipArmor(armor); err != nil {
			fmt.Println(err)
			continue
		}
		 break} 
		 
		 return char
}