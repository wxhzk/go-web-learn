package main

import "fmt"
import "log"
import "net/http"
import "html/template"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/html/login.html")
		if err != nil {
			http.Error(w, "Server error", 500)
			return
		}
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		fmt.Println("Username:", r.FormValue("username"), " Password:", r.FormValue("passwd"))
		w.Write([]byte("Hello"))
	} else {
		http.NotFound(w, r)
	}
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	t, err := template.ParseFiles("template/html/404.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

func main() {
	http.Handle("/js/", http.FileServer(http.Dir("template")))
	http.Handle("/css/", http.FileServer(http.Dir("template")))
	http.Handle("/img/", http.FileServer(http.Dir("template")))

	//http.HandleFunc("/admin/", adminHandler)
	http.HandleFunc("/login/", LoginHandler)
	//http.HandleFunc("/ajax/", ajaxHandler)
	http.HandleFunc("/", NotFoundHandler)

	fmt.Println("Listen and serving on 9090...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
