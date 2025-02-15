package app

import (
	"user-service/handlers"
	"user-service/middleware"
	"user-service/store"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router      *gin.Engine
	UserHandler *handlers.UserHandler
}

func NewServer(store store.UserStore) *Server {
	// Initialize Gin router
	router := gin.Default()

	// Initialize UserHandler with injected store
	userHandler := handlers.NewUserHandler(store)

	s := &Server{
		Router:      router,
		UserHandler: userHandler,
	}

	s.routes()
	return s
}

func (s *Server) routes() {
	// public endpoints
	public := s.Router.Group("/api/v1/auth")
	{
		public.POST("/signup", s.UserHandler.SignUp)
		public.POST("/login", s.UserHandler.SignIn)
	}
	// protected endpoints
	protected := s.Router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// Protected endpoints go here
	}
}
