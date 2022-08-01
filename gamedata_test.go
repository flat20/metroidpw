package metroidpw_test

import (
	"bytes"
	"testing"

	"github.com/flat20/metroidpw/v2"
)

func TestSetBit(t *testing.T) {
	t.Parallel()

	var gd metroidpw.GameData
	gd.SetBit(17, true)
	if gd[2] != 2 {
		t.Errorf("SetBit not 2")
	}
	gd.SetBit(18, false)
	if gd[2] != 2 {
		t.Errorf("SetBit not still 2")
	}
	gd.SetBit(17, false)
	if gd[2] != 0 {
		t.Errorf("SetBit not zero")
	}

	// Test out of bounds panics.
	testSetPanic := func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("SetBit out of bounds did not panic")
			}
		}()
		gd.SetBit(9000, true)
	}
	testSetPanic()

	testGetPanic := func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("GetBit out of bounds did not panic")
			}
		}()
		gd.GetBit(9000)
	}
	testGetPanic()
}

func TestGameData(t *testing.T) {
	t.Parallel()
	var gd metroidpw.GameData

	gd.SetGameTime(123)
	if gd.GameTime() != 123 {
		t.Error("Incorrect GameTime")
	}

	gd.SetMissiles(123)
	if gd.Missiles() != 123 {
		t.Error("Incorrect Missiles")
	}

	gd.SetShift(123)
	if gd.Shift() != 123 {
		t.Error("Incorrect Shift")
	}

	// Test Encode
	var encdata metroidpw.GameData
	encdata.SetGameTime(123)

	pw := metroidpw.Encode(&encdata)
	if pw.String() != "000000 000000 001x00 00001x" {
		t.Errorf("Incorrect Encode Password: %s", pw.String())
	}

	decoded, err := metroidpw.Decode(&pw)
	if err != nil {
		t.Errorf("Unable to decode Password. %s", err)
	}

	// Not checking checksum byte since it's only set during encoding.
	if !bytes.Equal(decoded[:17], encdata[:17]) {
		t.Logf("Decoded not matching original")
	}
}
