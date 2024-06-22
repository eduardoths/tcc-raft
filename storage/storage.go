package storage

import (
	"os"
	"sync"

	"github.com/cockroachdb/pebble"
	"github.com/eduardoths/tcc-raft/pkg/logger"
)

type Storage struct {
	mu  sync.Mutex
	db  *pebble.DB
	log logger.Logger
}

type StorageSaveStruct struct {
	Key   string
	Value []byte
}

func NewStorage(name string) *Storage {
	l := logger.MakeLogger()
	db, err := pebble.Open("data/"+name, &pebble.Options{})
	if err != nil {
		l.Error(err, "failed to open storage")
		os.Exit(1)
	}
	s := &Storage{
		db:  db,
		log: l,
	}
	return s
}

func (s *Storage) Save(data StorageSaveStruct) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.db.Set([]byte(data.Key), data.Value, pebble.Sync)
}

func (s *Storage) Get(key string) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, closer, err := s.db.Get([]byte(key))
	if err != nil {
		return nil, err
	}

	defer closer.Close()
	return data, err
}

func (s *Storage) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.db.Delete([]byte(key), pebble.Sync)
}

func (s *Storage) Shutdown() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.db.Close(); err != nil {
		s.log.Error(err, "failed to shutdown storage")
		os.Exit(1)
	}
}
