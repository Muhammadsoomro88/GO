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

var All_id []int = []int{}
var All_name []string = []string{}
var All_comp []string = []string{}
var All_sal []string = []string{}

var new_id int = 0
var new_name string = ""
var new_comp string = ""
var new_sal string = ""

type ToHtml struct {
	Id       []int
	Name     []string
	Division []string
	Email    []string
}

//for 2nd Query
type CountAll1 struct {
	Name string `json:"name"`
	Pos  string `json:"pos"`
}

type allData1 struct {
	Name string `json:"name"`
	Pos  string `json:"pos"`
}

var All_name1 []string = []string{}
var All_pos1 []string = []string{}

type ToHtml1 struct {
	Name []string
	Pos  []string
}

//for 3rd Query
type CountAll2 struct {
	Division1 string `json:"division"`
	Emp       int    `json:"emp"`
}

type allData2 struct {
	Division1 string `json:"division"`
	Emp       int    `json:"emp"`
}

var All_division []string = []string{}
var All_emp []int = []int{}

type ToHtml2 struct {
	Division1 []string
	Emp       []int
}

//for 4th Query
type CountAll3 struct {
	Division2 string `json:"division"`
	Emp1      int    `json:"emp"`
}

type allData3 struct {
	Division2 string `json:"division"`
	Emp1      int    `json:"emp"`
}

var All_division1 []string = []string{}
var All_emp1 []int = []int{}

type ToHtml3 struct {
	Division2 []string
	Emp1      []int
}

//for 5th Query
type CountAll4 struct {
	Division4 string `json:"division"`
	Emp4      int    `json:"emp"`
}

type allData4 struct {
	Division4 string `json:"division"`
	Emp4      int    `json:"emp"`
}

var All_division4 []string = []string{}
var All_emp4 []int = []int{}

type ToHtml4 struct {
	Division4 []string
	Emp4      []int
}

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
	//var count_all CountAll
	for results.Next() {
		var count_all CountAll
		err = results.Scan(&count_all.Id, &count_all.Name, &count_all.Division, &count_all.Email)
		if err != nil {
			log.Fatal(err)
		}

		All_id = append(All_id, count_all.Id)
		All_name = append(All_name, count_all.Name)
		All_comp = append(All_comp, count_all.Division)
		All_sal = append(All_sal, count_all.Email)

		new_id = count_all.Id
		new_name = count_all.Name
		new_comp = count_all.Division
		new_sal = count_all.Email

	}
	allData := ToHtml{Id: All_id, Name: All_name, Division: All_comp, Email: All_sal}
	//p := ToHtml{Id: All_id, Name: All_name, Division: All_comp, Email: All_sal}
	t, _ := template.ParseFiles("table.html")
	t.Execute(w, allData)
	//fmt.Println(count_all)
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
	for results.Next() {
		var count_all1 CountAll1
		err = results.Scan(&count_all1.Name, &count_all1.Pos)
		if err != nil {
			log.Fatal(err)
		}

		All_name1 = append(All_name1, count_all1.Name)
		All_pos1 = append(All_pos1, count_all1.Pos)

	}
	allData1 := ToHtml1{Name: All_name1, Pos: All_pos1}
	//p := ToHtml{Id: All_id, Name: All_name, Division: All_comp, Email: All_sal}
	t, _ := template.ParseFiles("table1.html")
	t.Execute(w, allData1)
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
	for results.Next() {
		var count_all2 CountAll2
		err = results.Scan(&count_all2.Division1, &count_all2.Emp)
		if err != nil {
			log.Fatal(err)
		}

		All_division = append(All_division, count_all2.Division1)
		All_emp = append(All_emp, count_all2.Emp)

	}
	allData2 := ToHtml2{Division1: All_division, Emp: All_emp}
	//p := ToHtml{Id: All_id, Name: All_name, Division: All_comp, Email: All_sal}
	t, _ := template.ParseFiles("table2.html")
	t.Execute(w, allData2)
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
	for results.Next() {
		var count_all3 CountAll3
		err = results.Scan(&count_all3.Division2, &count_all3.Emp1)
		if err != nil {
			log.Fatal(err)
		}

		All_division1 = append(All_division1, count_all3.Division2)
		All_emp1 = append(All_emp1, count_all3.Emp1)

	}
	allData3 := ToHtml3{Division2: All_division1, Emp1: All_emp1}
	//p := ToHtml{Id: All_id, Name: All_name, Division: All_comp, Email: All_sal}
	t, _ := template.ParseFiles("table3.html")
	t.Execute(w, allData3)
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
	for results.Next() {
		var count_all4 CountAll4
		err = results.Scan(&count_all4.Division4, &count_all4.Emp4)
		if err != nil {
			log.Fatal(err)
		}

		All_division4 = append(All_division4, count_all4.Division4)
		All_emp4 = append(All_emp4, count_all4.Emp4)

	}
	allData4 := ToHtml4{Division4: All_division4, Emp4: All_emp4}
	//p := ToHtml{Id: All_id, Name: All_name, Division: All_comp, Email: All_sal}
	t, _ := template.ParseFiles("table4.html")
	t.Execute(w, allData4)
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
