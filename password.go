package metroidpw

import (
	"errors"
	"strings"
)

const (
	missilesByte = 10
	gameTimeByte = 11
	shiftByte    = 16
	checksumByte = 17
	spaceValue   = 0xFF
	alpha        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz?- "
)

var ErrParse = errors.New("unable to parse password")

// Password holds the 24 byte encoded password data.
type Password [24]byte

// String returns a Metroid formatted password string.
func (p Password) String() string {
	conv := func(b []byte) (s string) {
		for _, v := range b {
			if v == spaceValue {
				v = 64
			}
			s += string(alpha[v])
		}

		return
	}

	return strings.Join([]string{
		conv(p[0:6]),
		conv(p[6:12]),
		conv(p[12:18]),
		conv(p[18:24]),
	}, " ")
}

// ParsePassword attempts to read the Metroid password input string and return a Password
// returns ErrParse if unable to parse.
func ParsePassword(password string) (Password, error) {
	var parsedPw Password

	password = strings.TrimSpace(password)

	// It has formatting spaces and we need to truncate. Might regex this,
	// although code wouldn't be any easier to read
	if len(password) == 27 && password[6] == ' ' && password[13] == ' ' && password[20] == ' ' {
		password = password[:6] + password[7:13] + password[14:20] + password[21:]
	}

	for i := 0; i < 24; i++ {
		if password[i] == ' ' {
			parsedPw[i] = spaceValue
		} else {
			pos := strings.IndexByte(alpha, password[i])
			if pos < 0 {
				return parsedPw, ErrParse
			}

			parsedPw[i] = byte(pos)
		}
	}

	return parsedPw, nil
}
