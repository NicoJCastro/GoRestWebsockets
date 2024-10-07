package repository

import (
	"context"
	"go/rest-ws/models"
)

//Inversion de dependencias para el repositorio de usuarios. Usando abstracciones

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	InsertPost(ctx context.Context, post *models.Post) error
	GetPostById(ctx context.Context, id string) (*models.Post, error)
	UpdatePost(ctx context.Context, post *models.Post) error
	DeletePost(ctx context.Context, id string, userId string) error
	ListPost(ctx context.Context, page uint64) ([]*models.Post, error)
	Close() error
}

// Inyeccion de dependencias para el repositorio de usuarios usando abstracciones y una funcion de creacion de repositorios de usuarios
//que recibe un contexto y una cadena de conexion a la base de datos y devuelve un repositorio de usuarios
//

var implementation Repository // Acada se puede acceder a cualquier BD que implemente UserRepository

func SetRepository(repository Repository) { // Implementacion iguala a la bd que se le pase como parametro en el main
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error { // Inserta un usuario en la bd
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) { // Obtiene un usuario por id
	return implementation.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) { // Obtiene un usuario por email
	return implementation.GetUserByEmail(ctx, email)
}

func Close() error { // Cierra la conexion a la bd
	return implementation.Close()
}

func InsertPost(ctx context.Context, post *models.Post) error { // Inserta un post en la bd
	return implementation.InsertPost(ctx, post)
}

func GetPostById(ctx context.Context, id string) (*models.Post, error) { // Obtiene un post por id
	return implementation.GetPostById(ctx, id)
}

func UpdatePost(ctx context.Context, post *models.Post) error { // Actualiza un post en la bd
	return implementation.UpdatePost(ctx, post)
}

func DeletePost(ctx context.Context, id string, userId string) error { // Elimina un post en la bd
	return implementation.DeletePost(ctx, id, userId)
}

func ListPost(ctx context.Context, page uint64) ([]*models.Post, error) { // Lista los post en la bd
	return implementation.ListPost(ctx, page)
}
