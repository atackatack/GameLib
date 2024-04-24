package server

import (
	"GAMELIB/internal/storage/minio"
	"GAMELIB/internal/storage/postgres"
	"GAMELIB/pkg/config"
	"context"
	"github.com/forbiddencoding/howlongtobeat"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	Storage    *postgres.Storage
	Minio      *minio.Client
	HLTB       *howlongtobeat.Client
}

func NewServer(cfg *config.Config) (*Server, error) {
	postgresClient, err := postgres.NewStorage(cfg.Storage)
	if err != nil {
		return nil, err
	}

	minioClient, err := minio.NewStorage(cfg.Minio)
	if err != nil {
		return nil, err
	}

	hltbClient, err := howlongtobeat.New()
	if err != nil {
		return nil, err
	}

	return &Server{
		Storage: &postgresClient,
		Minio:   &minioClient,
		HLTB:    hltbClient,
	}, nil
}

func (s *Server) Start(Port string) error {
	billingRouter := s.configureRoutes()

	s.httpServer = &http.Server{
		Addr:    ":" + Port,
		Handler: billingRouter,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
