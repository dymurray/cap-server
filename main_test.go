package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNulecules(t *testing.T) {
	req, err := http.NewRequest("GET", "/nulecules", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Nulecules)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Log(status)
	}

	expected := `{"nulecules":["apache-centos7-atomicapp","etherpad-centos7-atomicapp","flask-redis-centos7-atomicapp","gitlab-centos7-atomicapp","gocounter-scratch-atomicapp","guestbookgo-atomicapp","helloapache","mariadb-centos7-atomicapp","mariadb-fedora-atomicapp","mongodb-centos7-atomicapp","postgresql-centos7-atomicapp","redis-centos7-atomicapp","skydns-atomicapp","wordpress-centos7-atomicapp"]}`

	if strings.Trim(rr.Body.String(), "\n\t ") != expected {
		t.Errorf("handler returned unexpected body: got [%v] want [%v]",
			strings.Trim(rr.Body.String(), "\n\t "), expected)
	}
}

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/static/index.html", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IndexHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Log(status)
	}

	expected := `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>cap_ui</title>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/latest/css/bootstrap.min.css">
</head>
<body>
  <div id="main"></div>
  <script src="/static/cap_ui.js"></script>
</body>
</html>`

	actual := strings.TrimSpace(rr.Body.String())

	if actual != expected {
		t.Errorf("handler returned unexpected body: got [%v] want [%v]", actual, expected)
	}
}

func TestWrapScriptCmd(t *testing.T) {
	wrappedCmd := wrapScriptCmd("ls")
	if wrappedCmd != "\"ls\"" {
		t.Errorf("wrap returned %v expected %v", wrappedCmd, "\"ls\"")
	}
}

func TestMainGoDir(t *testing.T) {
	output := mainGoDir()
	if "." != output {
		t.Errorf("wrap returned %v expected %v", output, ".")
	}
}
