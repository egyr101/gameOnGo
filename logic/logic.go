package logic

const firstTypeValueChange = 0.25
const secondTypeValueChange = 0.5

var ValueChange = map[string]float32{"hp": firstTypeValueChange, "attack": secondTypeValueChange, "defense": secondTypeValueChange, "damage": firstTypeValueChange, "speed": firstTypeValueChange}

type buff string
type debuff string

type StateCharacter struct {
	isSpeedUp   int8
	isAttackUp  int8
	isHpUp      int8
	isDefenseUp int8
	isDamageUp  int8
}

type Effect struct {
	Name        string
	StateTarget string
	Impact      float32
}

var E = Effect{"defenseUp", "defense", 0.5}
