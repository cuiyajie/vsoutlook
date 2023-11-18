package fixture

import (
	"os"
	"runtime"
	"strings"

	"vsoutlook.com/vsoutlook/models"
	"vsoutlook.com/vsoutlook/models/db"
)

func dropAllTables() {
	if runtime.GOOS != "darwin" {
		panic(`import-test-data will delete all data first. 
		to be executed on dev machines only.`)
	}
	sql := `SELECT 'DROP TABLE IF EXISTS "' || tablename || '" CASCADE;' 
	from pg_tables WHERE schemaname = 'public';`
	var results []string
	db.DB.Raw(sql).Scan(&results)
	for _, sql := range results {
		db.DB.Exec(sql)
	}
}

func createTables() {
	models.ReadyMaintainTestDB()
	models.Migrate("test")
}

func Create() {
	dropAllTables()
	createTables()

	data, _ := os.ReadFile("biz/fixture/data.sql")
	sql := string(data)
	result := db.DB.Exec(sql)
	if result.Error != nil {
		panic(result.Error)
	}
	// CreateUsers()
}

func CreateUsers() {
	uids := "alice bob clair dean eli"
	for _, uid := range strings.Fields(uids) {
		u := models.User{
			ModelId: models.ModelId{
				ID: uid,
			},
		}
		models.Save(&u)
	}

	projIDs := strings.Fields("proj1 proj2 proj3")
	for _, id := range projIDs {
		models.Save(&models.Project{ID: id})
	}
}
