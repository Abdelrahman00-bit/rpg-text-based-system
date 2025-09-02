/*
this file where I write the base and reueasble characters
info then i can use it again with diffrent type of char
*/
package characters

import "RPG/items"

// this is the states struct it will represent the char attributes
type Stats struct {
	Health   int
	Strength int
	Defense  int
	Magic    int
}


// Character defines what every character must be able to do
type CharacterMethods interface {
    GetName() string                    // to get the name of the character
    GetStats() Stats                   //  to get the stats of the charcter
    EquipWeapon(w items.Weapon) error //   to equip a weapon
    EquipArmor(a items.Armor) error  //    to equip an armor
    ShowStats()                     //     show character info
}

