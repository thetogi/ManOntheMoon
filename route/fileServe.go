package route

import (
	"log"
	"net/http"
)

func SetupFileServer(ServerName string, mux *http.ServeMux) {

	log.Printf("Initializing File Server %v for static content", ServerName)

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

}
