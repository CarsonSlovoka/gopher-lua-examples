--[[
載入我們定義的go函數
--]]

local m = require("Math2") -- 我們在PreloadModule中所定義的名稱: https://github.com/CarsonSlovoka/gopher-lua-examples/blob/cd03ce4/examples/custom-func/main.go#L36-L55
local a = m.add(1,2)
print(a)
local b = m.sub(1,2)
myModule.logPrint(b) -- 這一個是我們之前於RegisterModule所定義的Module: https://github.com/CarsonSlovoka/gopher-lua-examples/blob/cd03ce4/examples/custom-func/main.go#L17-L30
