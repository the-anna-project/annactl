// Package os implements spec.FSService and provides a real OS
// bound file system implementation.
package os

import (
	"io/ioutil"
	"os"

	servicespec "github.com/the-anna-project/spec/service"
)

// New creates a new OS file system service.
func New() servicespec.FSService {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.ServiceCollection

	// Settings.

	metadata map[string]string
}

func (s *service) Boot() {
	id, err := s.Service().ID().New()
	if err != nil {
		panic(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"kind": "os",
		"name": "file-system",
		"type": "service",
	}
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) ReadFile(filename string) ([]byte, error) {
	s.Service().Log().Line("func", "ReadFile")

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, maskAny(err)
	}

	return bytes, nil
}

func (s *service) Service() servicespec.ServiceCollection {
	return s.serviceCollection
}

func (s *service) SetServiceCollection(sc servicespec.ServiceCollection) {
	s.serviceCollection = sc
}

func (s *service) WriteFile(filename string, bytes []byte, perm os.FileMode) error {
	s.Service().Log().Line("func", "WriteFile")

	err := ioutil.WriteFile(filename, bytes, perm)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
