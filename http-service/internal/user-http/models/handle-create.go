package models

type ErrorJSON struct {
	Error string `json:"error"`
}

type HandleCreateJSON struct {
	Surname    string  `json:"surname"`
	Name       *string `json:"name,omitempty"`
	Secondname *string `json:"secondname,omitempty"`
	Age        *int32  `json:"age,omitempty"`
}
