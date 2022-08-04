package metroidpw

import "encoding/binary"

const (
	BitMaruMariTaken    uint = 0 // Items taken of the map
	BitBombsTaken       uint = 6
	BitVariaTaken       uint = 11
	BitHighJumpTaken    uint = 24
	BitScrewAttackTaken uint = 26

	BitBombs       uint = 72 // Samus weapons enabled
	BitHighJump    uint = 73
	BitLongBeam    uint = 74
	BitScrewAttack uint = 75
	BitMaruMari    uint = 76
	BitWaveBeam    uint = 77
	BitIceBeam     uint = 78

	BitHasSwimsuit uint = 71

	BitStartNorfair uint = 64 // Map starting positions
	BitStartKraid   uint = 65 // All 3 off = Brinstar
	BitStartRidley  uint = 66 // All 3 on = Tourian
)

// GameData holds the raw game state data.
type GameData [18]byte

// CalcChecksum returns a checksum calculated over the first 17 bytes of the GameData.
func (gd GameData) CalcChecksum() (checksum byte) {
	for _, v := range gd[:17] {
		checksum += v
	}

	return
}

func (gd *GameData) SetChecksum(checksum byte) {
	gd[checksumByte] = checksum
}

// Checksum returns the current checksum byte stored in the GameData.
func (gd *GameData) Checksum() byte {
	return gd[checksumByte]
}

func (gd *GameData) SetGameTime(time uint32) {
	binary.LittleEndian.PutUint32(gd[gameTimeByte:], time)
}

// GameTime returns a 24-bit elapsed time. There's overflow data in the last byte
// which is included in the checksum calculation, so we return all 32-bits of data.
func (gd *GameData) GameTime() uint32 {
	return binary.LittleEndian.Uint32(gd[gameTimeByte:])
}

func (gd *GameData) SetMissiles(count uint8) {
	gd[missilesByte] = count
}

func (gd *GameData) Missiles() uint8 {
	return gd[missilesByte]
}

func (gd *GameData) SetShift(value uint8) {
	gd[shiftByte] = value
}

// Shift byte is used to determine how much to rotate the GameData when encoding/decoding.
func (gd *GameData) Shift() uint8 {
	return gd[shiftByte]
}

// SetBit sets a bit on GameData.
func (gd *GameData) SetBit(bit uint, value bool) {
	if bit >= 18*8 {
		panic("bit out of bounds")
	}

	i := bit / 8

	if value {
		gd[i] |= (1 << (bit % 8))
	} else {
		gd[i] &= ^(1 << (bit % 8))
	}
}

// GetBit returns a bit from GameData.
func (gd *GameData) GetBit(bit uint) bool {
	if bit >= 18*8 {
		panic("bit out of bounds")
	}

	i := bit / 8

	return (gd[i] & (1 << (bit % 8))) > 0
}
