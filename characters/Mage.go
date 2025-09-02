/*this is the Mage file here I will type all methods and all data about Mage*/
package characters

import (
	"RPG/items"
	"errors"
	"fmt"
)

// this the OOP of Mage
type Mage struct {
    Name   string
    Base   Stats
    Weapon items.Weapon
    Armor  items.Armor
}

func NewMage(name string) *Mage {
    return &Mage{
        Name: name,
        Base: Stats{Health: 50, Strength: 5, Defense: 30, Magic: 50},
    }
}

func (m *Mage) GetName() string { return m.Name }

func (m *Mage) GetStats() Stats {
    stats := m.Base
    if m.Weapon != nil {
        stats.Strength += m.Weapon.GetDamageBonus()
    }
    if m.Armor != nil {
        stats.Defense += m.Armor.GetDefenseBonus()
    }
    return stats
}

func (m *Mage) EquipWeapon(weapon items.Weapon) error {
    if weapon.GetType() != "staff" {
        return errors.New("mages can only equip swords")
    }
    m.Weapon = weapon
    return nil
}

func (m *Mage) EquipArmor(armor items.Armor) error {
    if armor.GetType() == "plate" {
        return errors.New("mages cannot equip robes")
    }
    m.Armor = armor
    return nil
}

func (m *Mage) ShowStats() {
    stats := m.GetStats()
    fmt.Printf("Character: %s (Mage)\n", m.Name)
    fmt.Printf("Health: %d\nStrength: %d\nDefense: %d\nMagic: %d\n",
        stats.Health, stats.Strength, stats.Defense, stats.Magic)

    if m.Weapon != nil {
        fmt.Printf("Equipped Weapon: %s (+%d strength)\n",
            m.Weapon.GetName(), m.Weapon.GetDamageBonus())
    }
    if m.Armor != nil {
        fmt.Printf("Equipped Armor: %s (+%d defense)\n",
            m.Armor.GetName(), m.Armor.GetDefenseBonus())
    }
}
