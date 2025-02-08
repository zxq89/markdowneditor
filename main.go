package main

import (
        "fmt"
        "html/template"
        "net/http"
)

func main() {

        http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                //http.ServeFile(w, r, "./static/index.html")
                tmpl, err := template.ParseFiles("./static/index.html")
                if err != nil {
                        http.Error(w, fmt.Sprintf("Error parsing template: %v", err), http.StatusInternalServerError)
                        return
                }
                err = tmpl.Execute(w, nil)
                if err != nil {
                        http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
                }
        })

        port := ":18080"
        fmt.Println("Server is running on http://localhost" + port)
        err := http.ListenAndServe(port, nil)
        if err != nil {
                fmt.Println("Error starting server: ", err)
        }
}
