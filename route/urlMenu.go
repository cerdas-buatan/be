package route
import (
	"github.com/cerdas-buatan/be/module"
	"net/http"
)

func MenuRoutes(menuService *module.MenuService) {
	http.HandleFunc("/menu/rename", module.RenameMenuHandler(menuService))
	http.HandleFunc("/menu/archive", module.ArchiveMenuHandler(menuService))
	http.HandleFunc("/menu/add", module.AddMenuHandler(menuService))
}
