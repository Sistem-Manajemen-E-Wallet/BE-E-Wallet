package responses

type MapResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func WebJSONResponse(msg string, data interface{}) MapResponse {
	return MapResponse{
		Message: msg,
		Data:    data,
	}
}

type MapResponseMeta struct {
	Message string      `json:"message"`
	Meta    interface{} `json:"meta,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func WebJSONResponseMeta(msg string, meta interface{}, data interface{}) MapResponseMeta {
	return MapResponseMeta{
		Message: msg,
		Meta:    meta,
		Data:    data,
	}
}
