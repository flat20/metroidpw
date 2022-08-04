package metroidpw_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/flat20/metroidpw"
)

func TestEncode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		setup  func(*metroidpw.GameData)
		expect string
	}{
		{
			setup: func(gd *metroidpw.GameData) {
				gd.SetMissiles(9)   // 9 Missiles
				gd.SetBit(71, true) // Swimsuit
			},
			expect: "000000 000020 00a000 000029",
		},
		{
			setup: func(gd *metroidpw.GameData) {
				gd.SetBit(24, true) // Has taken High Jump Boots
				gd.SetBit(73, true) // High Jump Boots
			},
			expect: "00000G 000000 0W0000 000003",
		},
		{
			setup: func(gd *metroidpw.GameData) {
				gd.SetGameTime(123)
			},
			expect: "000000 000000 001x00 00001x",
		},
	}

	for _, ts := range tests {
		var gd metroidpw.GameData
		ts.setup(&gd)
		encoded := metroidpw.Encode(&gd)
		if encoded.String() != ts.expect {
			t.Errorf("Was \"%s\" Expected: \"%s\"", encoded.String(), ts.expect)
		}
	}
}

func TestDecode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		password string
		expect   []byte
	}{
		{
			// https://metroid.fandom.com/wiki/List_of_Metroid_passwords
			// From NES emulator:
			password: "DRAGON BALL Z Dragon Ball z",
			expect:   []byte{26, 217, 72, 48, 185, 101, 42, 191, 241, 155, 172, 149, 101, 137, 114, 95, 255, 253},
		},
		{
			// https://metroid.fandom.com/wiki/List_of_Metroid_passwords
			password: "X----- --N?WO dV-Gm9 W01GMI",
			expect:   []byte{255, 255, 255, 255, 255, 235, 253, 3, 19, 191, 250, 24, 19, 0, 0, 176, 5, 146},
		},
	}

	for _, ts := range tests {
		pw, err := metroidpw.ParsePassword(ts.password)
		if err != nil {
			t.Error(err)

			continue
		}

		gd, err := metroidpw.Decode(&pw)
		if err != nil {
			t.Error(err)

			continue
		}

		if !bytes.Equal(ts.expect, gd[:]) {
			t.Errorf("GameData bytes don't match emulator %v != %v", ts.expect, gd[:])
		}
	}

	// Test broken checksum
	pw, err := metroidpw.ParsePassword("000000 000000 000000 000001")
	if err != nil {
		t.Error("Unable to create broken password")
	}
	_, err = metroidpw.Decode(&pw)
	if !errors.Is(err, metroidpw.ErrChecksum) {
		t.Error("Password should have checksum error")
	}
}

func TestTemp(t *testing.T) {
	t.Parallel()
	var data metroidpw.GameData

	data.SetMissiles(123)
	if data.Missiles() != 123 {
		t.Error("Incorrect Missiles")
	}

	pw := metroidpw.Encode(&data)
	if _, err := metroidpw.Decode(&pw); err != nil {
		t.Error(err)
	}
}

func TestBits(t *testing.T) {
	t.Parallel()

	var gd metroidpw.GameData
	gd.SetBit(9, true)
	if gd.GetBit(9) != true {
		t.Error("Bit not true")
	}
	gd.SetBit(9, false)
	if gd.GetBit(9) != false {
		t.Error("Bit not false")
	}
}
