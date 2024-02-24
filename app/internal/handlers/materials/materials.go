package materialshandler

import(
	"app/internal/storage"
	"app/internal/utils"
	"github.com/labstack/echo/v4"
)
type MaterialsHandler struct {
	store storage.Storage
}

func New(store storage.Storage) *MaterialsHandler {
	return &MaterialsHandler{
		store: store,
	}
}
//Process:
//User logs in, get the groups that they are in the and the materials their groups have
//search the database for every material they have
//Then query the databae for those materials
//The database returns what? maybe the filepath, maybe the files itself
//Then maybe get the file by querying the file server
//Then display the file here. Prob just the raw bc i dont do frontend.
func (materials *MaterialsHandler) GetMaterials(c echo.Context) error{
	
	return nil
}