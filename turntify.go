package turntify

import (
	"github.com/eddiezane/turntify/store"
	"github.com/eddiezane/turntify/types"
)

// Turntify defines...
type Turntify struct {
	store store.Store
}

// NewTurntify creates a new turntify instance.
func NewTurntify(store store.Store) *Turntify {
	return &Turntify{store}
}

// CreateRoom creates a new room in the loca store.
func (t *Turntify) CreateRoom(id string) (*types.Room, error) {
	r := &types.Room{ID: id}
	return r, t.store.WriteRoom(r)
}

// GetRoom gets a room.
func (t *Turntify) GetRoom(id string) (*types.Room, error) {
	return t.store.ReadRoom(id)
}

// AddSong adds a new song to the room.
func (t *Turntify) AddSong(room string, song string) error {
	r, err := t.store.ReadRoom(room)
	if err != nil {
		return err
	}

	r.Playlist = append(r.Playlist, &types.Song{ID: song})
	return t.store.WriteRoom(r)
}
