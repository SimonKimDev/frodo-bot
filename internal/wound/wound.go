package wound

import (
	"sync"
)

// Tracker records which server (guild) have Frodo stabbed. Monsters...
// It's safe for concurrent use.
type Tracker struct {
	mu      sync.Mutex
	stabbed map[string]bool
}

func New() *Tracker {
	return &Tracker{stabbed: make(map[string]bool)}
}

// Stab marks guildID as stabbed return false if already stabbed
// Because it wouldn't make sense to stab Frodo multiple times.
// Come on, he's already wounded.
func (t *Tracker) Stab(guildID string) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.stabbed[guildID] {
		return false
	}

	t.stabbed[guildID] = true
	return true
}

// IsStabbed reports whether Frodo was already stabbed or not
func (t *Tracker) IsStabbed(guildID string) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.stabbed[guildID]
}
