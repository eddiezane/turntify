package store

import (
	"sync"

	"github.com/eddiezane/turntify/types"
)

// ErrNotFound is return if the resource is not found.
type ErrNotFound struct{}

func (e ErrNotFound) Error() string {
	return "not found"
}

// Store TODO(wlynch): docs
type Store interface {
	ReadRoom(id string) (*types.Room, error)
	WriteRoom(room *types.Room) error
}

// LocalStore defines a local in-memory store.
type LocalStore struct {
	mu *sync.RWMutex

	rooms map[string]*types.Room
}

// NewLocalStore creates a new local store.
func NewLocalStore() *LocalStore {
	m := make(map[string]*types.Room)
	return &LocalStore{
		mu:    &sync.RWMutex{},
		rooms: m,
	}
}

// ReadRoom returns a room from the store, else ErrNotFound.
func (s *LocalStore) ReadRoom(id string) (*types.Room, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	r, ok := s.rooms[id]
	if !ok {
		return nil, ErrNotFound{}
	}
	return r, nil
}

// WriteRoom writes a room to the store.
func (s *LocalStore) WriteRoom(room *types.Room) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.rooms[room.ID] = room
	return nil
}
