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

