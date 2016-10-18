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
