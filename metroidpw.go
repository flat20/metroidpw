package metroidpw

import (
	"errors"
)

var ErrChecksum = errors.New("password has an incorrect checksum")

// encode6 converts 18 8-bit values in to 24 6-bit values.
func encode6(gd *GameData) (pw Password) {
	for i := 0; i < 6; i++ {
		gi := i * 3
		pi := i * 4
		pw[pi+0] = gd[gi] >> 2
		pw[pi+1] = ((gd[gi] & 0x3) << 4) | (gd[gi+1] >> 4)
		pw[pi+2] = ((gd[gi+1] & 0xf) << 2) | (gd[gi+2] >> 6)
		pw[pi+3] = gd[gi+2] & 0x3f
	}

	return
}

// decode6 converts 24 6-bit value Password data to a 18 8-bit value GameData.
func decode6(pw *Password) (gd GameData) {
	for i := 0; i < 6; i++ {
		pi := i * 4
		gi := i * 3
		gd[gi+0] = (pw[pi] << 2) | (pw[pi+1] >> 4)
		gd[gi+1] = (pw[pi+1] << 4) | (pw[pi+2] >> 2)
		gd[gi+2] = (pw[pi+2] << 6) | (pw[pi+3])
	}

	return
}

// rotateRight rotates the GameData right the given number of times.
func rotateRight(gd *GameData, count uint8) {
	var carry byte = 1
	var carryTemp byte
	var i uint8

	for i = 0; i < count; i++ {
		temp := gd[0]

		for j := 0; j < 16; j++ {
			carryTemp = gd[j] & 0x1
			gd[j] >>= 1
			gd[j] |= (carry & 0x1) << 7
			carry = carryTemp
		}

		carryTemp = temp & 0x1
		temp >>= 1
		temp |= (carry & 0x1) << 7
		carry = carryTemp

		gd[0] = temp
	}
}

// rotateLeft rotates the GameData left the given number of times.
func rotateLeft(gd *GameData, count uint8) {
	var carry byte = 1
	var carryTemp byte
	var i uint8

	for i = 0; i < count; i++ {
		temp := gd[15]

		for j := 15; j >= 0; j-- {
			carryTemp = (gd[j] & 0x80) >> 7
			gd[j] = ((gd[j] << 1) & 0xff) | (carry & 0x1)
			carry = carryTemp
		}

		carryTemp = (temp & 0x80) >> 7
		temp = ((temp << 1) & 0xff) | (carry & 0x1)
		carry = carryTemp

		gd[15] = temp
	}
}

// Encode GameData in to a Metroid Password.
func Encode(gd *GameData) Password {
	// calculate and set the checksum byte
	checksum := gd.CalcChecksum()
	gd.SetChecksum(checksum)

	count := gd.Shift()
	rotateRight(gd, count)

	pw := encode6(gd)

	return pw
}

// Decode Password in to GameData. Incorrect checksums return errChecksum but also return the parsed GameData
// so that checksum inconsistencies can be repaired.
func Decode(pw *Password) (GameData, error) {
	gd := decode6(pw)

	count := gd[shiftByte]
	rotateLeft(&gd, count)

	if gd.Checksum() != gd.CalcChecksum() {
		return gd, ErrChecksum
	}

	return gd, nil
}

// func tobinary(data []byte) {
// 	for _, b := range data {
// 		fmt.Printf("%08b ", b) // prints 00000000 11111101
// 	}
// 	fmt.Printf("\n")
// }
