package app
import(
	"app/internal/storage"
	"log"
)
func buildTestDB(){
	db, err := ConnectPgsqlTest()
	if err != nil{
		log.Fatalf(err.Error())
	}
	var poststore = storage.NewPostgresStorage(db)
	poststore.BuildDevDB()
}