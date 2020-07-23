package dice

import (
	"fmt"
	"testing"
)

func TestSetFaces(t *testing.T) {
	faces := 6
	d := Dice{}
	d.SetFaces(faces)

	expect := faces
	actual := d.Faces

	if expect != actual {
		t.Errorf("expected: %d, actual: %d", expect, actual)
	}
}

func TestGetFaces(t *testing.T) {
	faces := 6
	d := Dice{}
	d.Faces = faces

	expect := faces
	actual := d.GetFaces()

	if expect != actual {
		t.Errorf("expected: %d, actual: %d", expect, actual)
	}
}

// Dice型が作成されることのテスト
func TestNewDice01(t *testing.T) {
	d := NewDice(6)
	dice := new(Dice)

	expect := fmt.Sprintf("%T", dice)
	actual := fmt.Sprintf("%T", d)

	if expect != actual {
		t.Errorf("expected: %s, actual: %s", expect, actual)
	}
}

// facesで指定した値がFacesになることのテスト
func TestNewDice02(t *testing.T) {
	faces := 12
	d := NewDice(faces)

	expect := faces
	actual := d.Faces

	if expect != actual {
		t.Errorf("expected: %d, actual: %d", expect, actual)
	}
}

// facesで0以下の値を指定した場合, Facesが6になることのテスト
func TestNewDice03(t *testing.T) {
	faces := -12
	d := NewDice(faces)

	expect := 6
	actual := d.Faces

	if expect != actual {
		t.Errorf("expected: %d, actual: %d", expect, actual)
	}
}

// Roll結果が0以上faces以下になることのテスト
func TestRoll01(t *testing.T) {
	faces := 6
	d := NewDice(faces)
	number := d.Roll()

	if number > faces || number <= 0 {
		t.Errorf("invalid roll result: %d", number)
	}
}

// Cheat結果がnumberと等しくなることのテスト
func TestCheat01(t *testing.T) {
	faces := 6
	number := 6
	d := NewDice(faces)

	expect := number
	actual := d.Cheat(number)

	if expect != actual {
		t.Errorf("expected: %d, actual: %d", expect, actual)
	}
}

// numberで無効な値を指定した場合, Cheat結果がfacesと等しくなることのテスト
func TestCheat02(t *testing.T) {
	faces := 6
	number := -6
	d := NewDice(faces)
	d.Cheat(number)

	expect := d.Faces
	actual := d.Cheat(number)

	if expect != actual {
		t.Errorf("expected: %d, actual: %d", expect, actual)
	}
}
