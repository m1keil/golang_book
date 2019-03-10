package main

/*
 Add additional handlers so that clients can create, read, update, and delete database entries. For example, a request
 of the form /update?item=socks&price=6 will update the price of an item in the inventory and report an error if the
 item does not exist or if the price is invalid. (Warning: this change introduces concurrent variable updates.)
*/

/*
 Change the handler for /list to print its output as an HTML table, not text. You may find the html/template package
 (§4.6) useful.
*/

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var mu sync.Mutex

type database map[string]dollars
type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

func (db database) list(w http.ResponseWriter, req *http.Request) {
	template.Must(template.New("price list").Parse(`
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Table sort</title>
	</head>
	<body>
		<table>
			<tr>
				<th>Item</th>
				<th>Price</th>
			</tr>
			{{range $key, $value := .}}
			<tr>
				<td>{{ $key }}</td>
				<td>{{ $value }}</td>
			</tr>
			{{end}}
		</table>
	</body>
</html>`)).Execute(w, db)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item must be provided")
		return
	}

	price, err :=  strconv.Atoi(req.URL.Query().Get("price"))
	if err != nil || price <= 0 {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "Price must be a positive number")
		return
	}

	if _, exists := db[item]; exists {
		w.WriteHeader(http.StatusConflict) // 409
		fmt.Fprintf(w, "item already exists: %q\n", item)
		return
	}
	mu.Lock()
	db[item] = dollars(price)
	mu.Unlock()

	fmt.Fprintf(w, "%s: %s\n", item, db[item])
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item must be provided")
		return
	}

	price, err :=  strconv.Atoi(req.URL.Query().Get("price"))
	if err != nil || price <= 0 {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "Price must be a positive number")
		return
	}

	if _, exists := db[item]; !exists {
		w.WriteHeader(http.StatusConflict) // 409
		fmt.Fprintf(w, "item doesn't exists: %q\n", item)
		return
	}

	mu.Lock()
	db[item] = dollars(price)
	mu.Unlock()

	fmt.Fprintf(w, "%s: %s\n", item, db[item])
}


func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item must be provided")
		return
	}

	if _, exists := db[item]; !exists {
		w.WriteHeader(http.StatusConflict) // 409
		fmt.Fprintf(w, "item doesn't exists: %q\n", item)
		return
	}

	mu.Lock()
	delete(db, item)
	mu.Unlock()
}