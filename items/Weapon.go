package items

type Weapon interface {
	Item
	GetDamageBonus() int
	GetType() string // "sword", "staff", "bow"
}

type Sword struct {
	Name        string
	DamageBonus int
}

func (s Sword) GetName() string     { return s.Name }
func (s Sword) GetDamageBonus() int { return s.DamageBonus }
func (s Sword) GetType() string     { return "sword" }

type Staff struct {
	Name        string
	DamageBonus int
}

func (st Staff) GetName() string     { return st.Name }
func (st Staff) GetDamageBonus() int { return st.DamageBonus }
func (st Staff) GetType() string     { return "staff" }

type Bow struct {
	Name        string
	DamageBonus int
}

func (b Bow) GetName() string     { return b.Name }
func (b Bow) GetDamageBonus() int { return b.DamageBonus }
func (b Bow) GetType() string     { return "bow" }
