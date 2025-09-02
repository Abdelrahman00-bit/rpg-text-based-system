/*this is the worior file here I will type all methods and all data about worior*/

package characters

import (
	"RPG/items"
	"errors"
	"fmt"
)

// this the OOP of worior
type Warrior struct {
    Name   string
    Base   Stats
    Weapon items.Weapon
    Armor  items.Armor
}

func NewWarrior(name string) *Warrior {
    return &Warrior{
        Name: name,
        Base: Stats{Health: 100, Strength: 20, Defense: 10, Magic: 5},
    }
}

func (w *Warrior) GetName() string { return w.Name }

func (w *Warrior) GetStats() Stats {
    stats := w.Base
    if w.Weapon != nil {
        stats.Strength += w.Weapon.GetDamageBonus()
    }
    if w.Armor != nil {
        stats.Defense += w.Armor.GetDefenseBonus()
    }
    return stats
}

func (w *Warrior) EquipWeapon(weapon items.Weapon) error {
    if weapon.GetType() != "sword" {
        return errors.New("warriors can only equip swords")
    }
    w.Weapon = weapon
    return nil
}

func (w *Warrior) EquipArmor(armor items.Armor) error {
    if armor.GetType() == "robe" {
        return errors.New("warriors cannot equip robes")
    }
    w.Armor = armor
    return nil
}

func (w *Warrior) ShowStats() {
    stats := w.GetStats()
    fmt.Printf("Character: %s (Warrior)\n", w.Name)
    fmt.Printf("Health: %d\nStrength: %d\nDefense: %d\nMagic: %d\n",
        stats.Health, stats.Strength, stats.Defense, stats.Magic)

    if w.Weapon != nil {
        fmt.Printf("Equipped Weapon: %s (+%d strength)\n",
            w.Weapon.GetName(), w.Weapon.GetDamageBonus())
    }
    if w.Armor != nil {
        fmt.Printf("Equipped Armor: %s (+%d defense)\n",
            w.Armor.GetName(), w.Armor.GetDefenseBonus())
    }
}
