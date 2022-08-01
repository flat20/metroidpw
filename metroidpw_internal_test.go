package metroidpw

import (
	"testing"
)

func TestRotate(t *testing.T) {
	t.Parallel()
	var gd GameData

	gd.SetGameTime(123)
	rotateRight(&gd, 2)

	if gd.GameTime() == 123 {
		t.Error("GameData not rotated")
	}

	rotateLeft(&gd, 2)

	if gd.GameTime() != 123 {
		t.Error("Incorrect GameTime after rotating")
	}
}
