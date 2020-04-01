package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetDocsPage_Success(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/docs", nil)

	docsHandler := Docs{
		DocsPathIndex: "../../docs/index.html",
	}

	docsHandler.GetDocsPage(w, r)
	if res := w.Result().StatusCode; res != http.StatusOK {
		t.Errorf("expected status code doesn't match actual. \nGot: %d, \nWant: %d", res, http.StatusOK)
	}
}

func Test_GetDocsPage_Failure(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/docs", nil)

	docsHandler := Docs{
		DocsPathIndex: "invalid-file",
	}

	docsHandler.GetDocsPage(w, r)
	if res := w.Result().StatusCode; res != http.StatusInternalServerError {
		t.Errorf("expected status code doesn't match actual. \nGot: %d, \nWant: %d", res, http.StatusInternalServerError)
	}
}

func Test_GetSwaggerFile_Success(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/docs/swagger.yaml", nil)

	docsHandler := Docs{
		DocsPathYaml: "../../docs/index.html",
	}

	docsHandler.GetSwaggerFile(w, r)
	if res := w.Result().StatusCode; res != http.StatusOK {
		t.Errorf("expected status code doesn't match actual. \nGot: %d, \nWant: %d", res, http.StatusOK)
	}
}

func Test_GetSwaggerFile_Failure(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/docs/swagger.yaml", nil)

	docsHandler := Docs{
		DocsPathYaml: "invalid-file",
	}

	docsHandler.GetSwaggerFile(w, r)
	if res := w.Result().StatusCode; res != http.StatusInternalServerError {
		t.Errorf("expected status code doesn't match actual. \nGot: %d, \nWant: %d", res, http.StatusInternalServerError)
	}
}
