package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"strings"
)

func TestShowIndexPageUnauthenticated(t *testing.T)  {
	r := getRouter(true)
	r.GET("/",showIndexPage)
	req,_ := http.NewRequest("GET","/",nil)
	testHTTPResponse(t,r,req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p,err := ioutil.ReadAll(w.Body)
		pageOk := err == nil && strings.Index(string(p),"<title>Home Page</title>") > 0
		return statusOK && pageOk
	})
}