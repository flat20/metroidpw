-- FCEUX script for Metroid
-- 
-- Writes the password to NES memory so there's no need to enter it manually.
-- Just press start on password screen and the password from memory will be read by Metroid.
-- The script then prints out the resulting bytes decoded by Metroid in the Emu console.

-- Declare and set variables or functions if needed
local gamedata = {}
memory.registerwrite(0x6988, 18, function(addr, size, value)
    if size ~= 1 then
        emu.print("no clue")
        return
    end

    -- +1 because lua starts at 1..
    gamedata[addr+1 - 0x6988] = value
    emu.print("gamedata", gamedata)

end)

-- Strip formatted spaces and convert to table of 24 bytes
local function parsePassword(pws)

    local alpha = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz?- "

    local pw = {}

    -- loop password
    for i=1, #pws do
        local c = pws:sub(i,i)
        if i == 7 or i == 14 or i == 21 then
            -- skip if a formatting space
        elseif c == " " then
            -- 0xFF for space char
            table.insert(pw, 0xFF)
        else
            -- store index as the pw byte
            idx = string.find(alpha, c)
            idx = idx - 1
            if idx < 0 then
                -- incorrect char entered
                return nil
            end
            table.insert(pw, idx)
        end
    end

    return pw

end

local pw = parsePassword("X----- --N?WO dV-Gm9 W01GMI")
-- local pw = parsePassword("DRAGON BALL Z Dragon Ball z")

-- Write password
addr = 0x699A
for i, v in ipairs(pw) do
    memory.writebyte(addr + (i-1), v)
end