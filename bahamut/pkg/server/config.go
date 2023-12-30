package server

import "net/http"

func ConfigServer() http.Server {
	return http.Server{
		Addr: ":80",
	}
}

func ConfigHandler() *http.ServeMux {
	mux := http.NewServeMux()

	return mux
}
