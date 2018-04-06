package handlersTest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"../../main/handlers"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetBlogs(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/blogs/testusername")
	if assert.NoError(t, handlers.Mock_Get_Success_Blogs(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"testLogs\"", rec.Body.String())
	}
}

//TODO: Add more test cases for handlers
