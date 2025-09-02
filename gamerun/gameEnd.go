package gamerun

import (
	"RPG/characters"
	"fmt"
)

func gameEnd(char characters.CharacterMethods) {


	for {
		fmt.Println("Okay, adventurer. You're good to go!\n you can press 's' to see your stats or wirte 'exit' to exit from the game simple right?")
		var lastOption string
		fmt.Scanln(&lastOption)
		switch lastOption {

		case "s":
			char.ShowStats()

		case "exit":
			return

		default:
			fmt.Println("worng option please type 's' for stats or 'exit' to exit")
		}
	}

}