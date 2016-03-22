package main

import (
    "database/sql"
    "fmt"
    "html/template"
    "log"
    "net/http"
    _"github.com/go-sql-driver/mysql" 	
)

func login(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {
	
	
       r.ParseForm()
        // logic part of log in
		field1 := r.Form["username"]
		field2 := r.Form["password"]
      //  fmt.Fprintf(w,"%s", field1)
      //  fmt.Fprintf(w,"%s", field2)
		
		db, err := sql.Open("mysql", "root:@/ravikiran")
		if err != nil {
			 panic(err.Error())
		}
    defer db.Close()
    stmtIns, err :=db.Prepare("INSERT INTO go_ex VALUES(?, ?)")
	if err != nil {
     panic(err.Error())
        }
	defer stmtIns.Close()

		
		 _, err = stmtIns.Exec(field1[0], field2[0])
			if err != nil {
				 panic(err.Error())
			}
			fmt.Fprintf(w,"Inserted Successfully")
    }
}


func main() {
  //  http.HandleFunc("/", sayhelloName) // setting router rule
    http.HandleFunc("/login", login)
    err := http.ListenAndServe("localhost:9999", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}