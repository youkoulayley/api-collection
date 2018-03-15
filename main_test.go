package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/youkoulayley/api-collection/bootstrap"
	"github.com/youkoulayley/api-collection/databases/migrations"
)

func TestMain(m *testing.M) {
	c := bootstrap.GetConf("conf_test.json")
	bootstrap.InitLogs(c)
	bootstrap.OpenDB(c)
	migrations.LaunchMigrations()

	router := InitializeRouter()
	http.ListenAndServe(":"+c.Port, router)

	code := m.Run()

	os.Exit(code)
}

func clearTable() {
	bootstrap.Db().Exec("DELETE * FROM paintcans")
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/paintcans", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
