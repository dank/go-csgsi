package csgsi

import(
	"encoding/json"
	"io"
)

func parseJson(reader io.Reader, int interface{}) error {
	return json.NewDecoder(reader).Decode(int)
}