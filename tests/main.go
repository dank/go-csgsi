package main

import (
	".."
	"fmt"
)

func main()  {
	game := csgsi.New(0)

	go func() {
		for state := range game.Channel {
			for weapon := range state.Player.Weapons {
				weapon := state.Player.Weapons[weapon]
				if(weapon.State == "active") {
					fmt.Printf("%s %d/%d\n", weapon.Name, weapon.Ammo_clip, weapon.Ammo_reserve) // => weapon_glock 20/120
				}
			}
		}
	}()

	game.Listen(":3000")
}