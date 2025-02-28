package handlers

type GameHandler struct {
	deps *HandlersDeps
}

func NewGameHandler(deps *HandlersDeps) *GameHandler {
	return &GameHandler{deps: deps}
}
