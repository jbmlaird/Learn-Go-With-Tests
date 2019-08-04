package poker

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func (l League) find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)

	if err != nil {
		err = fmt.Errorf("Incorrectly parsed JSON, error %v", err)
	}

	return league, err
}
