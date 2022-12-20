package models

type Todo struct {
	ID          int64  `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Done        bool   `json:"done" type:"bool"`
	CreatedAt   string `json:"created_at"`
}

type EntityValidatorError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type Response struct {
	Error   bool `json:"error"`
	Message any  `json:"message"`
}
