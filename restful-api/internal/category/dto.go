package category

type CreateRequest struct {
	Name string `json:"name" validate:"required,min=3,max=200"`
}

type Response struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateRequest struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,min=3,max=200"`
}

type HandlerResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
