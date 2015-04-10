package app

import (
)

// APIServer runs a strayard api server.
type APIServer struct {
	ExternalHost string
	EnableProfiling bool
}

// NewAPIServer creates a new APIServer object with default parameters
func NewAPIServer() *APIServer {
    s := &APIServer{
    }
    return s
}

// Run runs the specified APIServer. This should never exit.
func (s *APIServer) Run(_ []string) error {
    return nil
}
