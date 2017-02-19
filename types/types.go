package types

import (
	"time"
)

// Room represents a shared playlist among users. Each room has a current song playing, and when that song was started.
type Room struct {
	ID       string
	People   []*User
	Playlist []*Song

	Start time.Time
	Song  *Song
}

// User represents a user.
type User struct {
	ID string
}

// Song represents a unique song and its duration.
type Song struct {
	ID       string
	Duration time.Duration
}
