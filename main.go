package main

import (
	"context"
	"go/rest-ws/handlers"
	"go/rest-ws/middleware"
	"go/rest-ws/server"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	//Creamos un servidor nuevo

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseURL: DATABASE_URL,
	})
	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}
	log.Println("Server created successfully", s)

	s.Start(BinderRoutes)

}

func BinderRoutes(s server.Server, r *mux.Router) {

	//Evitamos el middleware para algunos endpoints
	api := r.PathPrefix("/api/v1").Subrouter()

	//Middleware para la autenticacion.
	//Modificamos r por api para que el middleware solo aplique a los endpoints de la api. A las que tenga el prefijo /api/v1
	api.Use(middleware.CheckAuthMiddleware(s))

	//Enpoints
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)

	// Endpoint para el registro de usuarios (signup)
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)

	// Endpoint para el login de usuarios
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)

	// Endpoint Middleware
	api.HandleFunc("/me", handlers.MeHandler(s)).Methods(http.MethodGet)

	// Endpoint para insertar un post
	api.HandleFunc("/posts", handlers.InsertPostHandler(s)).Methods(http.MethodPost)

	// Endpoint para obtener un post por id (Read)
	r.HandleFunc("/posts/{id}", handlers.GetPostByIdHandler(s)).Methods(http.MethodGet)

	// Endpoint para actualizar un post por id (Update)
	api.HandleFunc("/posts/{id}", handlers.UpdatePostHandler(s)).Methods(http.MethodPut)

	// Endpoint para eliminar un post por id (Delete)
	api.HandleFunc("/posts/{id}", handlers.DeletePostHandler(s)).Methods(http.MethodDelete)

	// Endpoint para listar los post (List)
	r.HandleFunc("/posts", handlers.ListPostHandler(s)).Methods(http.MethodGet)

	// Endpoint para el websocket
	r.HandleFunc("/ws", s.Hub().HandleWebSocket)

}
