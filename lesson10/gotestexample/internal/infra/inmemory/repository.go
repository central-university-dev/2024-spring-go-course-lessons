package inmemory

import (
	"context"
	"fmt"
	"sync"

	"gotestexample/internal/common"
)

type Repository struct {
	store map[string]string
	mu    sync.RWMutex
}

func New() *Repository {
	return &Repository{
		store: make(map[string]string),
	}
}

func (r *Repository) set(key, val string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.store[key] = val
}

func (r *Repository) get(key string) (string, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	v, ok := r.store[key]

	return v, ok
}

func (r *Repository) Save(_ context.Context, id, longURL string) error {
	r.set(id, longURL)

	return nil
}

func (r *Repository) Get(_ context.Context, id string) (string, error) {
	v, ok := r.get(id)
	if !ok {
		return "", fmt.Errorf("%w: store has no key", common.ErrNotFound)
	}

	return v, nil
}
