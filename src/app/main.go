package main

import (
	"html/template"
	"log"
	"fmt"
	"net/http"
	"app/mqtt"
)

type TemperatureHum struct {
	Title string
	DateTime []string
	Temperature []string
}

type LedStrip struct {
	State string
	Color string
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
	http.HandleFunc("/tempHum", temphumiHandler)
	http.HandleFunc("/ledstrip", ledStripHandler)
	http.Handle("/src/app/html/static/", http.StripPrefix("/src/app/html/static/",
		http.FileServer(http.Dir("src/app/html/static/"))))

	http.ListenAndServe(":8080", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fp := "src/app/html/templates/index.html"
	tmpl, err := template.ParseFiles(fp)

	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fp := "src/app/html/templates/login.html"

	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func temphumiHandler(w http.ResponseWriter, r *http.Request) {
	DateTime := []string {"20","29", "40"} //just for testing
	Temperature := []string {"18-06-26", "18-06-26", "18-06-26"} //just for testing
	page := TemperatureHum{"Temperature/Humidity", DateTime, Temperature}
	fp := "src/app/html/templates/temphum.html"

	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal(err)
	}
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	fp := "src/app/html/templates/dashboard.html"

	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}

}

func ledStripHandler(w http.ResponseWriter, r *http.Request) {
	State := "true" // will be read from db later 
	Color := "blue" // will be read from db later
	page := LedStrip{State, Color}
	fp := "src/app/html/templates/ledstrip.html"

	topic := "ledStrip"
	ledStrip_state := r.FormValue("light1")
	ledStrip_color := r.FormValue("ledStrip_Color")

	if ledStrip_state == "true" {
		fmt.Println("ledStrip_state is true")
		mqtt.ChangeState("on",topic)
	} else {
		fmt.Println("ledStrip_state is false")
		mqtt.ChangeState("off",topic)
	}

	if ledStrip_color != "" {
		mqtt.ChangeColor(ledStrip_color,topic)
	}

	fmt.Println("The selected color is :", ledStrip_color)

	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, page)
	if err != nil {
		log.Fatal(err)
	}
}