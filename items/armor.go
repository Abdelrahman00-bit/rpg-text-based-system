package items

type Armor interface {
	Item
	GetDefenseBonus() int
	GetType() string // "plate", "robe", "leather"
}

type PlateArmor struct {
	Name         string
	DefenseBonus int
}

func (p PlateArmor) GetName() string      { return p.Name }
func (p PlateArmor) GetDefenseBonus() int { return p.DefenseBonus }
func (p PlateArmor) GetType() string      { return "plate" }

type Robe struct {
	Name         string
	DefenseBonus int
}

func (r Robe) GetName() string      { return r.Name }
func (r Robe) GetDefenseBonus() int { return r.DefenseBonus }
func (r Robe) GetType() string      { return "robe" }

type LeatherArmor struct {
	Name         string
	DefenseBonus int
}

func (l LeatherArmor) GetName() string      { return l.Name }
func (l LeatherArmor) GetDefenseBonus() int { return l.DefenseBonus }
func (l LeatherArmor) GetType() string      { return "leather" }
