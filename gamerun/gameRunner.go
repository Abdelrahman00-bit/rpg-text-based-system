// this packge was made to control the runnig of the system

/*
I've divided the code into two files:

1. GameStart file runs the character creation, selects the character type, weapons and armor.

2. GameEnd   One runs the final loop, which allows you to exit the game or view player data.

3. I call the tow funcs in this 'GameRunner' func
*/
package gamerun

func GameRunner() {

	char := GameStart()
	gameEnd(char)
}