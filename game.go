package csgsi

import(
	"net/http"
	"io"
)

// Game ...
type Game struct {
	// Channel for game state data.
	Channel chan State
}

// Returns a new Game object.
func New(size int) *Game{
	return &Game{Channel: make(chan State, size)}
}

// Starts listening to address provided.
// If a POST request is received, it sends it through
// the Game.Channel channel.
func (gs *Game) Listen(addr string) error {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			state := &State{}
			if err := parseJson(req.Body, &state); err != nil {
				io.WriteString(res, "bad")
			}

			gs.Channel <- *state

			io.WriteString(res, "ok")
		}
	})

	if err := http.ListenAndServe(addr, nil); err != nil {
		return err
	}

	return nil
}