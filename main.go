package main

import (
	"fmt"
	"geo-shop-auth/internal/api"
	pb "geo-shop-auth/internal/api/gen/authpb"
	"geo-shop-auth/internal/application/common"
	"geo-shop-auth/internal/application/services"
	"geo-shop-auth/internal/infrastructure/postgres"
	postgresrep "geo-shop-auth/internal/infrastructure/postgres/repositories"
	"geo-shop-auth/internal/infrastructure/redis"
	redisrep "geo-shop-auth/internal/infrastructure/redis/repositories"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg := parseConfig()

	postgresDB, err := postgres.NewPostgresDB(cfg.Postgres)
	if err != nil {
		panic(err)
	}
	userRepository := postgresrep.NewUserRepository(postgresDB)

	redisClient := redis.NewRedis(cfg.Redis)
	tokenRepository := redisrep.NewRedisTokenRepository(redisClient)

	jwtOptions := common.JWTOptions{}
	tokenService := services.NewTokenService(tokenRepository, jwtOptions)
	passwordService := services.NewPasswordService()

	server := api.NewServer(userRepository, tokenService, passwordService)

	ln, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, server)

	if err = grpcServer.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func parseConfig() *Config {
	viper.SetConfigName("config.local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	return &cfg
}

type Config struct {
	Port     string
	Postgres postgres.Options
	Redis    redis.Options
}
