package controllers_test

import (
	"testing"

	"github.com/Yutan0423/go-medium-level/controllers"
	"github.com/Yutan0423/go-medium-level/controllers/testdata"
	_ "github.com/go-sql-driver/mysql"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
