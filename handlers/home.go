package handlers

import (
	"encoding/json"
	"go/rest-ws/server"
	"net/http"
)

// Se encarga procesar la peticion de la ruta principal

type HomeResponse struct { // lo usamos para devolver al cliente con la respuesta que le queremos dar
	Message string `json:"message"` // mensaje que queremos devolver
	Status  bool   `json:"status"`  // estado de la peticion
}

func HomeHandler(s server.Server) http.HandlerFunc { // recibe un servidor y devuelve un handler
	return func(w http.ResponseWriter, r *http.Request) { // recibe un response writer y un request
		w.Header().Set("Content-Type", "application/json") // seteamos el header de la respuesta
		w.WriteHeader(http.StatusOK)                       // seteamos el status code de la respuesta
		json.NewEncoder(w).Encode(HomeResponse{            // codificamos la respuesta en formato JSON})
			Message: "Welcome to the API",
			Status:  true,
		})
	}
}
