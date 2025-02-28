package handlers

type UserHandler struct {
	deps *HandlersDeps
}

func NewUserHandler(deps *HandlersDeps) *UserHandler {
	return &UserHandler{deps: deps}
}
