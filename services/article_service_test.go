package services_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/Yutan0423/go-medium-level/services"

	_ "github.com/go-sql-driver/mysql"
)

// レシーバとして扱うためのMyAppService構造体
var aSer *services.MyAppService

func TestMain(m *testing.M) {
	user := "docker"
	pass := "docker"
	database := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", user, pass, database)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	aSer = services.NewMyAppService(db)
	m.Run()
}

func BenchmarkGetArticleService(b *testing.B) {
	articleID := 1
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := aSer.GetArticleService(articleID)
		if err != nil {
			b.Error(err)
			break
		}
	}
}
