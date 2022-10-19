package store

import (
	"fmt"
	"strings"
	"testing"
)

// TestStore...
func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper() // we say that it is test method

	config := NewConfig()
	config.DatabaseURL = databaseURL
	s := New(config)

	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.DB.Exec(fmt.Sprintf("TRUNCATE %s CASCAD", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}
		s.Close()
	}
}
