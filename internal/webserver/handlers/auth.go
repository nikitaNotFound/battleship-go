package handlers

type AuthHandler struct {
	deps *HandlersDeps
}

func NewAuthHandler(deps *HandlersDeps) *AuthHandler {
	return &AuthHandler{deps: deps}
}
