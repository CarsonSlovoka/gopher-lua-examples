package main

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"log"
)

func main() {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(`print("hello")`); err != nil {
		panic(err)
	}

	// 定義自己的Module (Package)
	myMod := L.RegisterModule("myModule", map[string]lua.LGFunction{
		"logPrint": func(l *lua.LState) int {
			top := L.GetTop()
			for i := 1; i <= top; i++ {
				log.Print(L.ToStringMeta(L.Get(i)).String())
				if i != top {
					fmt.Print("\t")
				}
			}
			fmt.Println("")
			return 0
		},
	})
	L.Push(myMod) // 記得要Push不然無法使用
	if err := L.DoString(`myModule.logPrint("hello")`); err != nil {
		panic(err)
	}

	// dynamic import
	L.PreloadModule("Math2", func(l *lua.LState) int {
		lTable := l.SetFuncs(l.NewTable(), map[string]lua.LGFunction{
			"add": func(l *lua.LState) int {
				x := l.ToInt(1)
				y := l.ToInt(2)

				l.Push(lua.LNumber(x + y))
				return 1
			},
			"sub": func(l *lua.LState) int {
				x := l.ToInt(1)
				y := l.ToInt(2)

				l.Push(lua.LNumber(x - y))
				return 1
			},
		})
		l.Push(lTable)
		return 1
	})
	if err := L.DoFile("preloadModule_test.lua"); err != nil {
		panic(err)
	}
}
