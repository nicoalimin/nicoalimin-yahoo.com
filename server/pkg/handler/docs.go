package handler

import (
	"io/ioutil"
	"net/http"
)

// Docs is the HTTP handler to serve the Swagger documentation
type Docs struct {
	DocsPathIndex string
	DocsPathYaml  string
}

// GetDocsPage serves the Swagger HTML page to the users
func (d *Docs) GetDocsPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	docsBytes, err := ioutil.ReadFile(d.DocsPathIndex)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(docsBytes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetSwaggerFile returns the swagger.yaml that is used to render the Swagger docs page
func (d *Docs) GetSwaggerFile(w http.ResponseWriter, r *http.Request) {
	_, err := ioutil.ReadFile(d.DocsPathYaml)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	http.ServeFile(w, r, d.DocsPathYaml)
}
