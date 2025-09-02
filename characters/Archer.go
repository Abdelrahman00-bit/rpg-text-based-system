/*this is the Archer file here I will type all methods and all data about Archer*/
package characters

import (
	"RPG/items"
	"errors"
	"fmt"
)

// this the OOP of Archer
type Archer struct {
    Name   string
    Base   Stats
    Weapon items.Weapon
    Armor  items.Armor
}

func NewArcher(name string) *Archer {
    return &Archer{
        Name: name,
        Base: Stats{Health: 100, Strength: 35, Defense: 10, Magic: 0},
    }
}

func (a *Archer) GetName() string { return a.Name }

func (a *Archer) GetStats() Stats {
    stats := a.Base
    if a.Weapon != nil {
        stats.Strength += a.Weapon.GetDamageBonus()
    }
    if a.Armor != nil {
        stats.Defense += a.Armor.GetDefenseBonus()
    }
    return stats
}

func (a *Archer) EquipWeapon(weapon items.Weapon) error {
    if weapon.GetType() != "bow" {
        return errors.New("archers can only equip swords")
    }
    a.Weapon = weapon
    return nil
}

func (a *Archer) EquipArmor(armor items.Armor) error {
    if armor.GetType() == "robe" {
        return errors.New("archers cannot equip robes")
    }
    a.Armor = armor
    return nil
}

func (a *Archer) ShowStats() {
    stats := a.GetStats()
    fmt.Printf("Character: %s (Archer)\n", a.Name)
    fmt.Printf("Health: %d\nStrength: %d\nDefense: %d\nMagic: %d\n",
        stats.Health, stats.Strength, stats.Defense, stats.Magic)

    if a.Weapon != nil {
        fmt.Printf("Equipped Weapon: %s (+%d strength)\n",
            a.Weapon.GetName(), a.Weapon.GetDamageBonus())
    }
    if a.Armor != nil {
        fmt.Printf("Equipped Armor: %s (+%d defense)\n",
            a.Armor.GetName(), a.Armor.GetDefenseBonus())
    }
}
