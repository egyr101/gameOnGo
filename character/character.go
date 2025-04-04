package character

import (
	"fmt"
	"gameongo/logic"
)

type Character struct {
	hp      float32
	attack  float32
	defense float32
	damage  float32
	speed   float32
	name    string
}

func CreateCharacter(hp float32, attack float32, defense float32, speed float32, name string) Character {
	character := Character{hp, attack, defense, 1, speed, name}
	// if defense > 0 && defense <= 1 { //!!!
	// 	return character
	// } else {
	// 	panic(" value Def is most big")
	// }
	return character
}
func main() {
	fmt.Println(logic.E.Name)
}
