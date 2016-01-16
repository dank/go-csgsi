package csgsi

type State struct {
	Provider *provider
	Map *csmap
	Round *round
	Player *player
	AllPlayers map[string]*player // allplayers_*: steamid64 ...
	Previously *State
	Added *State
	Auth *auth
}

// provider
type provider struct {
	Name string
	AppId int
	Version int
	SteamId string
	Timestamp float32
}

// map
type csmap struct {
	Name string
	Phase string
	Round int
	Team_ct *team
	Team_t *team
}

// round
type round struct {
	Phase string
	Win_team string
	Bomb string
}

// player_id
type player struct {
	SteamId string
	Name string
	Team string
	Activity string
	State *playerState
	Weapons map[string]*weapon
	Match_stats *playerMatchStats
}

// win_team
type team struct {
	Score int
}

// player_state
type playerState struct {
	Health int
	Armor int
	Helmet bool
	Flashed int
	Smoked int
	Burning int
	Money int
	Round_kills int
	Round_killhs int
}

// player_weapons: weapon_0, weapon_1, weapon_2 ...
type weapon struct {
	Name string
	PaintKit string
	Type string
	State string
	Ammo_clip int
	Ammo_clip_max int
	Ammo_reserve int
}

// player_match_stats
type playerMatchStats struct {
	Kills int
	Assists int
	Deaths int
	Mvps int
	Score int
}

type auth struct {
	Token string
}