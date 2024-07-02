package api

import (
	"context"
	"investify/api/middleware"
	"investify/config"
	"investify/db/adapters"
	"investify/db/aws"
	db "investify/db/sqlc"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store db.Store
	s3    *aws.S3Service
	// store  *db.SQLStore
	router *gin.Engine
}

// runs on a specific address
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

// NewHTTPServer creates a new HTTP server and sets up routing

func NewHTTPServer() *Server {
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())
	pgConn := adapters.InitDb(config.EnvVars.DBSource)

	// if pgConn != nil {
	// 	log.Fatal("Connected to database")
	// }

	store := db.NewStore(pgConn)

	ctx := context.Background()
	s3Service := aws.NewS3Service(ctx)
	if s3Service == nil || s3Service.Client == nil {
		log.Fatalf("Failed to initialize S3Service")
	}

	server := &Server{

		router: router,
		store:  store,
		s3:     s3Service,
	}

	// corsConfig := cors.DefaultConfig()
	log.Println("Setting up server...")

	SetupRouter(server)

	return server
}
