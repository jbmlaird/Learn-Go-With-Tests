package server

import (
	"encoding/json"
	"fmt"
	"io"
)

func NewLeague(rdr io.Reader) ([]Player, error){
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)

	if err != nil {
		err = fmt.Errorf("Incorrectly parsed JSON, error %v", err)
	}

	return league, err
}
