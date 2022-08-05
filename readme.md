# Metroid Password Tool

`GameData` contains game state information which can be manipulated before encoding it in to a Metroid password.

`Password` is the encoded Metroid password which can be output as a formatted password string.

[metroid_pw_tester_lua](lua/metroid_pw_tester.lua) is a simple LUA script for the FCEUX NES emulator which can be used to test Metroid passwords and their decoded output in the emulator.

## References

- [Metroid NES ROM disassembly](https://github.com/nmikstas/metroid-disassembly/)
- [Password format analysis and how-to](https://games.technoplaza.net/mpg/password.txt)
- [Metroid(NES) Password Generator source code](https://github.com/jdratlif/mpg)
- [True Peace In Space online password generator tool](https://www.truepeacein.space/)
- [True Peace In Space source code](https://github.com/alexras/truepeacein.space)

The 18 bytes of Metroid game data would look something like this if one wanted to unsafe.Pointer over the bytes:

```go
type data struct { // 18 bytes
	mapItemsBits     [8]uint8
	samusHasItemBits uint8 
	missiles         uint8 
	gameAge          uint32 // only 24-bits used, but overflows to last byte
	bossKillBits     uint8 
	shiftByte        uint8 
	checksum         uint8 
}
```
