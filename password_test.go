package metroidpw_test

import (
	"errors"
	"testing"

	"github.com/flat20/metroidpw"
)

func TestParsePassword(t *testing.T) {
	t.Parallel()

	_, err := metroidpw.ParsePassword("00!000 000000 000000 000000")
	if !errors.Is(err, metroidpw.ErrParse) {
		t.Error("Password should not parse")
	}

	// #nosec G101
	pass := "DRAGON BALL Z Dragon Ball z"
	// pass := "X----- --N?WO dV-Gm9 W01GMI"
	pw, err := metroidpw.ParsePassword(pass)
	if err != nil {
		t.Error("Password should parse")
	}
	if pw.String() != pass {
		t.Errorf("Password.String() did not match. Was %s expected %s", pass, pw.String())
	}
}
