package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const message = "Hello World"

//for 1st Query
type CountAll struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Division string `json:"division"`
	Email    string `json:"email"`
}

type allData struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Division string `json:"division"`
	Email    string `json:"email"`
}

var new_id int = 0
var new_name string = ""
var new_comp string = ""
var new_sal string = ""

//for 2nd Query
type CountAll1 struct {
	Name string `json:"name"`
	Pos  string `json:"pos"`
}

type allData1 struct {
	Name string `json:"name"`
	Pos  string `json:"pos"`
}

var new_name1 string = ""
var new_pos1 string = ""

//for 3rd Query
type CountAll2 struct {
	Division1 string `json:"division"`
	Emp       int    `json:"emp"`
}

type allData2 struct {
	Division1 string `json:"division"`
	Emp       int    `json:"emp"`
}

var new_devi2 string = ""
var new_emp2 int = 0

//for 4th Query
type CountAll3 struct {
	Division2 string `json:"division"`
	Emp1      int    `json:"emp"`
}

type allData3 struct {
	Division2 string `json:"division"`
	Emp1      int    `json:"emp"`
}

var new_devi3 string = ""
var new_emp3 int = 0

//for 5th Query
type CountAll4 struct {
	Division4 string `json:"division"`
	Emp4      int    `json:"emp"`
}

type allData4 struct {
	Division4 string `json:"division"`
	Emp4      int    `json:"emp"`
}

var new_devi4 string = ""
var new_emp4 int = 0

func Query1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	db, err := sql.Open("mysql", "root:ContourSoftware123@tcp(localhost:3306)/info")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Connected to Database")

	//fetching results from Database
	//var results *sql.Rows
	results, err := db.Query("Select * from employees")

	if err != nil {
		log.Fatal(err)
	}

	//assigning struct to a single variable
	var emp_data []allData

	for results.Next() {
		var count_all CountAll
		err = results.Scan(&count_all.Id, &count_all.Name, &count_all.Division, &count_all.Email)
		if err != nil {
			log.Fatal(err)
		}

		new_id = count_all.Id
		new_name = count_all.Name
		new_comp = count_all.Division
		new_sal = count_all.Email

		p := allData{Id: new_id, Name: new_name, Division: new_comp, Email: new_sal}
		emp_data = append(emp_data, p)
	}

	t, _ := template.ParseFiles("table.html")
	t.Execute(w, emp_data)

}

func Query2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	db, err := sql.Open("mysql", "root:ContourSoftware123@tcp(localhost:3306)/info")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Database")
	defer db.Close()

	results, err := db.Query("select name, position from employees inner join status on employees.emp_id = status.emp_id")
	if err != nil {
		log.Fatal(err)
	}

	var emp_data []allData1
	for results.Next() {
		var count_all1 CountAll1
		err = results.Scan(&count_all1.Name, &count_all1.Pos)
		if err != nil {
			log.Fatal(err)
		}

		new_name1 = count_all1.Name
		new_pos1 = count_all1.Pos

		p := allData1{Name: new_name1, Pos: new_pos1}
		emp_data = append(emp_data, p)
	}
	//p := ToHtml{Id: All_id, Name: All_name, Division: All_comp, Email: All_sal}
	t, _ := template.ParseFiles("table1.html")
	t.Execute(w, emp_data)
}

func Query3(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	db, err := sql.Open("mysql", "root:ContourSoftware123@tcp(localhost:3306)/info")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Database")
	defer db.Close()

	results, err := db.Query("select division, count(name) as 'Employees' from employees group by (division)")
	if err != nil {
		log.Fatal(err)
	}
	var emp_data []allData2
	for results.Next() {
		var count_all2 CountAll2
		err = results.Scan(&count_all2.Division1, &count_all2.Emp)
		if err != nil {
			log.Fatal(err)
		}

		new_devi2 = count_all2.Division1
		new_emp2 = count_all2.Emp

		p := allData2{Division1: new_devi2, Emp: new_emp2}
		emp_data = append(emp_data, p)
	}
	//p := ToHtml{Id: All_id, Name: All_name, Division: All_comp, Email: All_sal}
	t, _ := template.ParseFiles("table2.html")
	t.Execute(w, emp_data)
}

func Query4(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	db, err := sql.Open("mysql", "root:ContourSoftware123@tcp(localhost:3306)/info")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Database")
	defer db.Close()

	results, err := db.Query("select division, count(position) as 'Trainees' from employees inner join status on employees.emp_id = status.emp_id where position='trainee' group by (division)")
	if err != nil {
		log.Fatal(err)
	}

	var emp_data []allData3
	for results.Next() {
		var count_all3 CountAll3
		err = results.Scan(&count_all3.Division2, &count_all3.Emp1)
		if err != nil {
			log.Fatal(err)
		}

		new_devi3 = count_all3.Division2
		new_emp3 = count_all3.Emp1

		p := allData3{Division2: new_devi3, Emp1: new_emp3}
		emp_data = append(emp_data, p)
	}
	//p := ToHtml{Id: All_id, Name: All_name, Division: All_comp, Email: All_sal}
	t, _ := template.ParseFiles("table3.html")
	t.Execute(w, emp_data)
}

func Query5(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	db, err := sql.Open("mysql", "root:ContourSoftware123@tcp(localhost:3306)/info")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Database")
	defer db.Close()

	results, err := db.Query("select division, count(position) as 'Trainees' from employees inner join status on employees.emp_id = status.emp_id where position='full-time' group by (division)")
	if err != nil {
		log.Fatal(err)
	}

	var emp_data []allData4
	for results.Next() {
		var count_all4 CountAll4
		err = results.Scan(&count_all4.Division4, &count_all4.Emp4)
		if err != nil {
			log.Fatal(err)
		}

		new_devi4 = count_all4.Division4
		new_emp4 = count_all4.Emp4

		p := allData4{Division4: new_devi4, Emp4: new_emp4}
		emp_data = append(emp_data, p)
	}
	//p := ToHtml{Id: All_id, Name: All_name, Division: All_comp, Email: All_sal}
	t, _ := template.ParseFiles("table4.html")
	t.Execute(w, emp_data)
}

func func1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
	fmt.Fprint(w, "<h1>Welcome to HTTP server HTML Tag!</h1>")
}

func main() {
	fmt.Println("Server started")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func1)
	mux.HandleFunc("/data1", Query1)
	mux.HandleFunc("/data2", Query2)
	mux.HandleFunc("/data3", Query3)
	mux.HandleFunc("/data4", Query4)
	mux.HandleFunc("/data5", Query5)
	err := http.ListenAndServe(":7070", mux)
	if err != nil {
		log.Fatal(err)
	}
}
