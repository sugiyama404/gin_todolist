package testtool

import (
	"app/cmd/driver"
	"app/cmd/infrastructure/dao"
	"net/http/httptest"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

func SetupRouter() (*gin.Engine, *gorm.DB) {
	conn := dao.ConnectDB()
	dao.InitDB(conn)
	engine := driver.Router(conn)
	return engine, conn
}

func NewTestServer() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}

func NewTestServerWithRouter() (*gin.Engine, *httptest.ResponseRecorder, *gorm.DB) {
	router, conn := SetupRouter()
	w := NewTestServer()
	return router, w, conn
}
