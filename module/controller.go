packagemodule
import(
	"encoding/json"
	"fmt"
//   	"encoding/json"
//  	"fmt"
	"net/http"
	"strconv"
	"time"
	
)


func HomeMakmur(w http.ResponseWriter, r *http.Request) {
	Response := fmt.Sprintf("Gaysdisal AI %s", "8080")
	response, err := json.Marshal(Response)
	if err != nil {
		http.Error(w, "Internal server error: JSON marshaling failed", http.StatusInternalServerError)
		return
	}
	w.Write(response)
	return
}

// RenameMenuHandler handles renaming a menu
func RenameMenuHandler(s *MenuService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var request struct {
            ID      string `json:"id"`
            NewName string `json:"new_name"`
        }
        if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        id, err := primitive.ObjectIDFromHex(request.ID)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        if err := s.RenameMenu(r.Context(), id, request.NewName); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(model.Response{Status: true, Message: "Menu renamed successfully"})
    }
}

// ArchiveMenuHandler handles moving a menu to the archive
func ArchiveMenuHandler(s *MenuService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var request struct {
            ID string `json:"id"`
        }
        if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        id, err := primitive.ObjectIDFromHex(request.ID)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        if err := s.ArchiveMenu(r.Context(), id); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(model.Response{Status: true, Message: "Menu moved to archive successfully"})
    }
}

// AddMenuHandler handles adding a new menu
func AddMenuHandler(s *MenuService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var menu model.Menu
        if err := json.NewDecoder(r.Body).Decode(&menu); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        menu, err := s.AddMenu(r.Context(), menu)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(model.Response{Status: true, Message: "New menu added successfully", Data: menu})
    }
}


// NotFound handles 404 errors
func NotFound(respw http.ResponseWriter, req *http.Request) {
	var resp model.Response
	resp.Message = "Not Found"
	helper.WriteJSON(respw, http.StatusNotFound, resp)
}




// func NotFound(respw http.ResponseWriter, req *http.Request) {
// 	respw.WriteHeader(http.StatusNotFound)
// 	respw.Header().Set("Content-Type", "text/html")
// 	fmt.Fprintln(respw, `
//         <!DOCTYPE html>
//         <html lang="en">
//         <head>
//             <meta charset="UTF-8">
//             <meta name="viewport" content="width=device-width, initial-scale=1.0">
//             <title>404 Not Found</title>
//             <style>
//                 body {
//                     font-family: Arial, sans-serif;
//                     text-align: center;
//                     margin-top: 50px;
//                 }
//                 .container {
//                     max-width: 600px;
//                     margin: auto;
//                 }
//                 .button {
//                     display: inline-block;
//                     margin-top: 20px;
//                     padding: 10px 20px;
//                     font-size: 16px;
//                     color: #fff;
//                     background-color: #007bff;
//                     text-decoration: none;
//                     border-radius: 5px;
//                 }
//                 .button:hover {
//                     background-color: #0056b3;
//                 }
//             </style>
//         </head>
//         <body>
//             <div class="container">
//                 <h1>404 - Not Found</h1>
//                 <p>The page you are looking for does not exist.</p>
//                 <a href="http://cerdas-buatan.projsonal.online/fe/" class="button">Home</a>
//             </div>
//         </body>
//         </html>
//     `)
// }

