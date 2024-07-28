package route
import (
	"github.com/kimseokgis/backend-ai/controller"
	"github.com/kimseokgis/backend-ai/module"
	"net/http"
)

func MenuRoutes(menuService *module.MenuService) {
	http.HandleFunc("/menu/rename", controller.RenameMenuHandler(menuService))
	http.HandleFunc("/menu/archive", controller.ArchiveMenuHandler(menuService))
	http.HandleFunc("/menu/add", controller.AddMenuHandler(menuService))
}
