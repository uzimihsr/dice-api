package dice

import (
	"math/rand"
	"time"
)

type DiceInterface interface {
	Roll() int
	Cheat(int) int
	GetFaces() int
	SetFaces(int)
}

type Dice struct {
	Faces int
}

func (d *Dice) GetFaces() int {
	return d.Faces
}

func (d *Dice) SetFaces(faces int) {
	d.Faces = faces
}

func (d *Dice) Roll() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(d.Faces) + 1
}

func (d *Dice) Cheat(n int) int {
	if n <= 0 || n >= d.Faces {
		n = d.Faces
	}
	return n
}

func NewDice(faces int) *Dice {
	if faces <= 0 {
		faces = 6
	}
	d := new(Dice)
	d.Faces = faces
	return d
}
