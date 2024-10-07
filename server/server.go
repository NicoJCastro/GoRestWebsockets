package server

import (
	"context"
	"errors"
	"go/rest-ws/database"
	"go/rest-ws/repository"
	"go/rest-ws/websocket"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// struct para la configuracion del servidor

type Config struct {
	Port        string // puerto del servidor
	JWTSecret   string // clave secreta para JWT
	DatabaseURL string // URL de la base de datos
}

// Interfaz para el servidor. Recordar que las interfaces en Go son implicitas y no se declaran explicitamente como en otros lenguajes
// En las interfaces se definen los metodos que debe implementar una estructura para cumplir con la interfaz
type Server interface {
	Config() *Config // metodo para obtener la configuracion del servidor
	Hub() *websocket.Hub
}

// Broker para el servidor: es un componente que se encarga de la comunicacion entre los diferentes componentes del servidor

type Broker struct {
	config *Config    // configuracion del servidor
	router mux.Router // router para el servidor
	hub    *websocket.Hub
}

// Metodo para obtener la configuracion del servidor ReciberFunction
func (b *Broker) Config() *Config {
	return b.config
}

// Metodo para obtener el Hub del servidor
func (b *Broker) Hub() *websocket.Hub {
	return b.hub
}

// Constructor para el servidor

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	// Revisamos nuestra configuracion
	if config.Port == "" {
		return nil, errors.New("port is required")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("jwt_secret is required")
	}

	if config.DatabaseURL == "" {
		return nil, errors.New("database_url is required")
	}

	// Instancia del Broker
	broker := &Broker{
		config: config,
		router: *mux.NewRouter(),
		hub:    websocket.NewHub(),
	}

	// Retornamos el broker

	return broker, nil
}

// Metodo para ejecutar el Broker

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = *mux.NewRouter()
	binder(b, &b.router) // llamamos a la funcion binder que se encarga de enlazar los diferentes componentes del servidor
	handler := cors.Default().Handler(&b.router)
	repo, err := database.NewPostgresRepository(b.Config().DatabaseURL)
	if err != nil {
		log.Fatal("NewPostgresRepository:", err)
	}

	go b.hub.Run()

	repository.SetRepository(repo)
	log.Println("Server is running on port", b.Config().Port)
	if err := http.ListenAndServe(b.config.Port, handler); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
