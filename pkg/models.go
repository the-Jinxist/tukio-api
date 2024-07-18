package pkg

type GenericResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type DataResponse struct {
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

type ResponseParams struct {
	NextCursor string `json:"next_cursor"`
}
