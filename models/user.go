package models

//Aca definimos como va a funcionar el flujo de datos en nuestra aplicacion

type User struct { // estructura para el usuario. Los struct en Go son como las clases en otros lenguajes
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
