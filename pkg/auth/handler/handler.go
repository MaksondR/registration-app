package handler

import "registration-app/pkg/helper"

type Request struct {
	Token string
}

type Response struct {
	Message string
}

type Handler struct {}

func (h *Handler) Validate (request Request, response *Response) error {
	valid := helper.JwtValidate(request.Token)

	if valid {
		response.Message = "Token is valid"
	} else {
		response.Message = "Token is invalid"
	}

	return nil
}
