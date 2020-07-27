package main

import (
	"testing"
	"github.com/gin-gonic/gin"
	"os"
	"net/http"
	"net/http/httptest"
	"fmt"
)

var tmpArticleList []article

func TestMain(m *testing.M)  {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

func testHTTPResponse(t *testing.T,r *gin.Engine,req *http.Request,f func(w *httptest.ResponseRecorder)bool)  {
	w := httptest.NewRecorder()
	fmt.Println(r)
	r.ServeHTTP(w,req)
	fmt.Println(r)
	if !f(w){
		t.Fail()
	}
}

func saveLists()  {
	tmpArticleList = articleList
}
func restoreLists()  {
	articleList = tmpArticleList
}