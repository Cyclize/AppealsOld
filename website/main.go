package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tidwall/gjson"
)

type Page struct {
	UUID   string
	Name   string
	Date   string
	Reason string
	Status string
}

func main() {

	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "templates/error.html") })

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://luckynetwork.net", http.StatusFound)
	})

	r.Handle("/{uuid}", userHandler)
	r.Handle("/{uuid}/submit", submitHandler)
	r.Handle("/staff", reviewHandler)
	r.Handle("/staff/{uuid}", reviewHandler)
	r.Handle("/staff/{uuid}/status", statusHandler)
	r.Handle("/archive/{uuid}", archiveHandler)

	http.ListenAndServe(":30016", r)
	log.Print("Running website on port 30016")
}

var userHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	p := loadPage(uuid)

	switch p.Status {
	case "PENDING":
		renderTemplate(w, "user", p)
	case "SUBMITTED":
		renderTemplate(w, "archive", p)
	case "":
		renderTemplate(w, "error", p)
	default:
		renderTemplate(w, "inactive", p)
		deleteAppeal(p.UUID)
	}
})

var archiveHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	p := loadArchive(uuid)

	renderTemplate(w, "archive", p)
})

var reviewHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	p := loadPage(uuid)

	renderTemplate(w, "review", p)
})

var statusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	reason := r.FormValue("reason")
	status := r.FormValue("status")
	p := loadPage(uuid)

	type Payload struct {
		Status string `json:"status"`
		Date   string `json:"date"`
		Reason string `json:"reason"`
	}

	data := Payload{
		status,
		"NOW()",
		reason,
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Data could not be translated to JSON! %s", err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("PATCH", "SUPABASE_API_URL/appeals?uuid=eq."+uuid, body)

	if err != nil {
		log.Fatalf("Could not make new request! %s", err)
	}

	req.Header.Set("Apikey", "SUPABASE_API_KEY")
	req.Header.Set("Authorization", "Bearer SUPABASE_API_KEY")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Prefer", "return=representation")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to run request! %s", err)
	}
	defer resp.Body.Close()

	archiveAppeal(p, status)
	http.Redirect(w, r, "/staff/"+uuid, http.StatusFound)
})

var submitHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	reason := r.FormValue("reason")

	type Payload struct {
		Status string `json:"status"`
		Reason string `json:"reason"`
	}

	data := Payload{
		"SUBMITTED",
		reason,
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Data could not be translated to JSON! %s", err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("PATCH", "SUPABASE_API_URL/appeals?uuid=eq."+uuid, body)
	if err != nil {
		log.Fatalf("Could not make new request! %s", err)
	}
	req.Header.Set("Apikey", "SUPABASE_API_KEY")
	req.Header.Set("Authorization", "Bearer SUPABASE_API_KEY")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Prefer", "return=representation")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to run request! %s", err)
	}
	defer resp.Body.Close()

	http.Redirect(w, r, "/"+uuid, http.StatusFound)
})

func loadPage(uuid string) *Page {
	req, err := http.NewRequest("GET", "SUPABASE_API_URL/appeals?uuid=eq."+uuid+"&select=*", nil)
	if err != nil {
		log.Fatalf("Could not make new request! %s", err)
	}
	req.Header.Set("Apikey", "SUPABASE_API_KEY")
	req.Header.Set("Authorization", "Bearer SUPABASE_API_KEY")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Prefer", "return=representation")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to run request! %s", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	name := gjson.Get(bodyString, "0.name").String()
	date := gjson.Get(bodyString, "0.date").String()
	reason := gjson.Get(bodyString, "0.reason").String()
	status := gjson.Get(bodyString, "0.status").String()

	return &Page{UUID: uuid, Name: name, Date: date, Reason: reason, Status: status}
}

func loadArchive(uuid string) *Page {
	req, err := http.NewRequest("GET", "SUPABASE_API_URL/archive?uuid=eq."+uuid+"&select=*", nil)
	if err != nil {
		log.Fatalf("Could not make new request! %s", err)
	}
	req.Header.Set("Apikey", "SUPABASE_API_KEY")
	req.Header.Set("Authorization", "Bearer SUPABASE_API_KEY")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Prefer", "return=representation")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to run request! %s", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)

	name := gjson.Get(bodyString, "0.name").String()
	date := gjson.Get(bodyString, "0.date").String()
	reason := gjson.Get(bodyString, "0.reason").String()
	status := gjson.Get(bodyString, "0.status").String()

	return &Page{UUID: uuid, Name: name, Date: date, Reason: reason, Status: status}
}

func deleteAppeal(uuid string) {
	req, err := http.NewRequest("DELETE", "SUPABASE_API_URL/appeals?uuid=eq."+uuid+"&select=*", nil)
	if err != nil {
		log.Fatalf("Could not make new request! %s", err)
	}
	req.Header.Set("Apikey", "SUPABASE_API_KEY")
	req.Header.Set("Authorization", "Bearer SUPABASE_API_KEY")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Prefer", "return=representation")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to run request! %s", err)
	}
	defer resp.Body.Close()
}

func archiveAppeal(p *Page, status string) {
	type Payload struct {
		UUID   string `json:"uuid"`
		Name   string `json:"name"`
		Reason string `json:"reason"`
		Status string `json:"status"`
	}

	data := Payload{
		p.UUID,
		p.Name,
		p.Reason,
		status,
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Data could not be translated to JSON! %s", err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "SUPABASE_API_URL/archive", body)
	if err != nil {
		log.Fatalf("Could not make new request! %s", err)
	}
	req.Header.Set("Apikey", "SUPABASE_API_KEY")
	req.Header.Set("Authorization", "Bearer SUPABASE_API_KEY")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Prefer", "return=representation")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to run request! %s", err)
	}
	defer resp.Body.Close()
}

var templates = template.Must(template.ParseFiles(
	"templates/user.html",
	"templates/inactive.html",
	"templates/archive.html",
	"templates/error.html",
	"templates/staff.html",
	"templates/review.html",
))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
