package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type countAll struct {
	Id       int    `json:"id"`
	Division string `json:"division"`
}

type countAll1 struct {
	Pos string `json:"pos"`
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", processor)
	http.HandleFunc("/validate", validation)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Connected to Server")
	tpl.ExecuteTemplate(w, "index.gohtml", nil) //here I will print the form so no need to pass any data
}

func validation(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	fname := r.FormValue("firster")
	pswd := r.FormValue("laster")

	fmt.Println(fname)
	fmt.Println(pswd)

	db, err := sql.Open("mysql", "root:ContourSoftware123@tcp(localhost:3306)/info")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Connected to Database")

	/*sqlQuery:= "select * from credentials where"
	sqlQuery+=" name="+fname+" AND password="+pswd*/

	results, err := db.Query("select * from credentials where name=? AND password=?", fname, pswd)
	if err != nil {
		log.Fatal(err)
	}
	for results.Next() { //first for loop
		new_res, err := db.Query("select emp_id, division from employees where name=?", fname)
		if err != nil {
			log.Fatal(err)
		}
		for new_res.Next() { //second for loop
			var cnt countAll
			err := new_res.Scan(&cnt.Id, &cnt.Division)
			if err != nil {
				log.Fatal(err)
			}
			output1 := cnt.Id
			output2 := cnt.Division
			//trying for getting third data
			new_res1, err := db.Query("select position from status where emp_id=?", output1)
			if err != nil {
				log.Fatal(err)
			}
			for new_res1.Next() {
				var cnt1 countAll1
				err := new_res1.Scan(&cnt1.Pos)
				if err != nil {
					log.Fatal(err)
				}
				empPos := cnt1.Pos

				//struct for sending data
				d := struct {
					Name     string
					Division string
					Position string
				}{
					Name:     fname,
					Division: output2,
					Position: empPos,
				}

				//sending data to .gohtml file
				tpl.ExecuteTemplate(w, "processor.gohtml", d)
			}
		}
	}
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fname := r.FormValue("firster")
	lname := r.FormValue("laster")

	d := struct {
		First string
		Last  string
	}{
		First: fname,
		Last:  lname,
	}

	tpl.ExecuteTemplate(w, "processor.gohtml", d)
}
