package main

import (
	"fmt"
	"github.com/AleksK1NG/api-mc/config"
	"github.com/AleksK1NG/api-mc/internal/db/postgres"
	"github.com/AleksK1NG/api-mc/internal/db/redis"
	"github.com/AleksK1NG/api-mc/internal/logger"
	"github.com/AleksK1NG/api-mc/internal/server"
	"go.uber.org/zap"
	"log"
)

const (
	configPath = "./config/config-docker"
)

func main() {
	log.Println("Starting auth server")

	l, err := logger.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		l.Fatal("fatal", zap.String("LoadConfig", err.Error()))
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		l.Fatal("fatal", zap.String("ParseConfig", err.Error()))
	}

	psqlDB, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		l.Fatal("", zap.String("init psql", err.Error()))
	}
	defer psqlDB.Close()

	l.Info("Postgres connected", zap.String("DB Status: %#v", fmt.Sprintf("%#v", psqlDB.Stats())))
	redisConn := redis.NewRedisClient(cfg)
	l.Info("Redis connected")

	s := server.NewServer(cfg, l, psqlDB, redisConn)
	log.Fatal(s.Run())
}
