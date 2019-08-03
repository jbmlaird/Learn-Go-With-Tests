package server

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(player string) int {
	league := f.GetLeague()

	for _, name := range league {
		if name.Name == player {
			return name.Wins
		}
	}

	return 0
}
