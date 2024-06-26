package config

import (
	"sync"

	"github.com/google/uuid"

	"github.com/siddhant-vij/YouTube-Video-Aggregator/database"
)

type ApiConfig struct {
	DatabaseURL        string
	VerifyEndpoint     string
	ResourceServerPort string
	ChannelBaseURL     string
	DBQueries          *database.Queries
	Mutex              *sync.RWMutex
	UserId             uuid.UUID
	AuthStatus         string
}

func (config *ApiConfig) GetAuthStatus() string {
	return config.AuthStatus
}
