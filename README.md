# go-csgsi

A library written in Go for Counter-Strike: Global Offensive's Game State Integration.

## About

The `go-csgsi` package listens on a specified port waiting for a HTTP POST request by CS:GO. Once a request is received, it is parsed and ready for the end user.

This information can be continuously retrieved and analyzed using a [Channel](https://tour.golang.org/concurrency/2). (See usage & example)

[Read more about Valve's Game State Integration](https://developer.valvesoftware.com/wiki/Counter-Strike:_Global_Offensive_Game_State_Integration)

## Installation

Via [Go](https://golang.org/):

```
$ go get -u github.com/dank/go-csgsi
```

# Usage

To initialize the library:
```go
// size int - Size of the channel
game := csgsi.New(10)
```

To receive updates, read from the channel provided: `game.Channel`:

```go
// Example 1 - Looping (blocking)
for state := range game.Channel {
	fmt.Println(state.Player.SomeData)
}
// End

// Example 2 - Wait for a state object to come in (blocking)
state := <-game.Channel
fmt.Println(state.Player.SomeData)
// End

// Example 3 - User creates a func and runs it in a Go routine (thread) (non-blocking)
go func() {
    for state := range game.Channel {
		fmt.Println(state.Player.SomeData)
	}
}()
// My non-blocked code can go here!!!
// End
```

To start listening to an address:
```go
// address string - TCP address to listen on
game.Listen(":3000") // localhost:3000
game.Listen("8.8.4.4:3000") // 8.8.4.4:3000
```

Full struct layout:
```
State
    .Provider
        .Name; string
        .AppId; int
        .Version; int
        .SteamId; string
        .Timestamp; float32
    .Map
        .Name; string
        .Phase; string
        .Round; int
        .Team_ct
            .Score; int
        .Team_t
            .Score; int
    .Round
        .Phase; string
        .Win_team; string
        .Bomb; string
    .Player // Note: Once you're dead, this will become the player you are spectating
        .SteamId; string
        .Name; string
        .Name; string
        .Team; string
        .Activity; string
        .State
            .Health; int
            .Armor; int
            .Helmet; bool
            .Flashed; int
            .Smoked; int
            .Burning; int
            .Money; int
            .Round_kills; int
            .Round_killhs; int
        .Weapons
            []; string - weapon_0, weapon_1, weapon_2 ...
                .Name; string
                .PaintKit; string
                .Type; string
                .State; string
                .Ammo_clip; int
                .Ammo_clip_max; int
                .Ammo_reserve; int
        .Match_stats
            .Kills; int
            .Assists; int
            .Deaths; int
            .Mvps; int
            .Score; int
        .AllPlayers
            []; string - steamid64 ...
                .Player; Player
        .Previously
                .State; State
        .Added
            .State; State
        .Auth
            .Token; string
```
## Example

Prints weapon stats (name, ammo) when weapon is active:
```go
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
```

You'll need to create a `gamestate_integration_*.cfg` file under `csgo/cfg` in your CS:GO installation directory. For example, `gamestate_integration_consolesample.cfg`:

```
"Console Sample v.1"
{
    "uri" "http://127.0.0.1:3000"
    "timeout" "5.0"
    "buffer"  "0.1"
    "throttle" "0.5"
    "heartbeat" "60.0"
    "auth"
    {
        "token" "SuperSecretPassword"
    }
    "data"
    {
        "provider"            "1"
        "map"                 "1"
        "round"               "1"
        "player_id"           "1"
        "player_state"        "1"
        "player_weapons"      "1"
        "player_match_stats"  "1"
    }
}
```

[from Valve's wiki.](https://developer.valvesoftware.com/wiki/Counter-Strike:_Global_Offensive_Game_State_Integration)

## License

```
The MIT License (MIT)

Copyright (c) 2016 Dan Kyung

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
