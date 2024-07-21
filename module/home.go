packagemodule
import(
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	
)


func HomeGaysdisal(w http.ResponseWriter, r *http.Request) {
	// Buat response dalam bentuk string
	Response := fmt.Sprintf("Gaysdisal AI", "8081")


	// Konversi response ke JSON
	jsonResponse, err := json.Marshal(Response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
\
	// Set header Content-Type
	w.Header().Set("Content-Type", "application/json")
\
	// Tulis response ke http.ResponseWriter
	w.Write(jsonResponse)
}


// NotFound handles 404 errors and provides a button to go back home
func NotFound(respw http.ResponseWriter, req *http.Request) {
	respw.WriteHeader(http.StatusNotFound)
	respw.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(respw, `
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>404 Not Found</title>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    text-align: center;
                    margin-top: 50px;
                }
                .container {
                    max-width: 600px;
                    margin: auto;
                }
                .button {
                    display: inline-block;
                    margin-top: 20px;
                    padding: 10px 20px;
                    font-size: 16px;
                    color: #fff;
                    background-color: #007bff;
                    text-decoration: none;
                    border-radius: 5px;
                }
                .button:hover {
                    background-color: #0056b3;
                }
            </style>
        </head>
        <body>
            <div class="container">
                <h1>404 - Not Found</h1>
                <p>The page you are looking for does not exist.</p>
                <a href="http://cerdas-buatan.projsonal.online/fe/" class="button">Home</a>
            </div>
        </body>
        </html>
    `)
}

