package webserver

type WebServerConfig struct {
	Port int
}

func GetWebServerConfig() (*WebServerConfig, error) {
	return &WebServerConfig{
		Port: 8080,
	}, nil
}
