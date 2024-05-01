package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Profile struct {
	Name       string
	Age        int
	Level      int
	Department string
	Origin     string
}

func (p *Profile) save() {
	err := os.MkdirAll("./profiles", os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating profiles directory: %s\n", err)
	}
	f, _ := os.Create("./profiles/" + p.Name + ".json")
	fmt.Println("Created!")

	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	// turns the file into an encoder
	encoder := json.NewEncoder(f)
	// re-adjust the indentation
	encoder.SetIndent("", "  ")
	//encodes the struct into the json file
	err = encoder.Encode(p)
	if err != nil {
		return
	}
	fmt.Println("Saved!")

}

func (p *Profile) update(name string) {
	err := os.Remove("./profiles/" + name + ".json")
	if err != nil {
		fmt.Printf("File with name: %s not delected", name)
	} else {
		fmt.Println("Delected!")
	}

	f, _ := os.Create("./profiles/" + p.Name + ".json")

	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	// turns the file into an encoder
	encoder := json.NewEncoder(f)
	// re-adjust the indentation
	encoder.SetIndent("", "  ")
	//encodes the struct into the json file
	encoder.Encode(p)
	fmt.Println("Updated!")
}

func viewUsers() []string {
	var profileNames []string

	file, _ := os.ReadDir("profiles/")
	for _, f := range file {
		profileNames = append(profileNames, f.Name()[:len(f.Name())-5])
	}
	return profileNames
}

func viewJson(name string) (*Profile, error) {
	// reads the json file
	file, err := os.ReadFile("profiles/" + name + ".json")
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return nil, err
	}
	var profile Profile

	// parses the json into the struct format
	_ = json.Unmarshal(file, &profile)

	return &Profile{
		Name:       profile.Name,
		Age:        profile.Age,
		Level:      profile.Level,
		Department: profile.Department,
		Origin:     profile.Origin,
	}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Profile) {
	t, _ := template.ParseFiles(tmpl + ".html")
	_ = t.Execute(w, p)
}

func viewAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		profileNames := viewUsers()
		t, _ := template.ParseFiles("templates/viewUsers.html")
		_ = t.Execute(w, profileNames)

	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	//this finds every thing after the '/view/'
	header := r.URL.Path[len("/Users/"):]
	if len(header) == 0 {
		//renderTemplate(w, "templates/404", nil)
		//return
		http.Redirect(w, r, "/Users", http.StatusFound)
	}
	parts := strings.Split(header, "/")
	name, condition := parts[len(parts)-2], parts[len(parts)-1]

	if condition == "edit" {
		editPorfile(w, r)
	} else if condition == "profile" {
		//this loads the .txt page
		p, err := viewJson(name)
		if err != nil {
			//http.Error(w, err.Error(), http.StatusInternalServerError)
			editPorfile(w, r)
			renderTemplate(w, "templates/404", nil)
			return
		}
		renderTemplate(w, "templates/user", p)
	}
}

func creatProfile(w http.ResponseWriter, r *http.Request) {
	age, _ := strconv.Atoi(r.FormValue("Age"))
	level, _ := strconv.Atoi(r.FormValue("Level"))
	if r.Method == "GET" {
		//editPorfile(w, r)
		renderTemplate(w, "templates/create", nil)
	} else {
		form := &Profile{
			Name:       r.FormValue("Name"),
			Age:        age,
			Level:      level,
			Department: r.FormValue("Department"),
			Origin:     r.FormValue("Origin"),
		}
		form.save()
		http.Redirect(w, r, "/Users/"+form.Name+"/profile", http.StatusFound)
	}
}

func editPorfile(w http.ResponseWriter, r *http.Request) {
	age, _ := strconv.Atoi(r.FormValue("Age"))
	level, _ := strconv.Atoi(r.FormValue("Level"))
	parts := strings.Split(r.URL.Path, "/")
	user, _ := parts[len(parts)-2], parts[len(parts)-1]
	if r.Method == "GET" {
		p, err := viewJson(user)
		if err != nil {
			editPorfile(w, r)
			renderTemplate(w, "templates/404", nil)
			return
		}
		renderTemplate(w, "templates/edit", p)
	} else {
		form := &Profile{
			Name:       r.FormValue("Name"),
			Age:        age,
			Level:      level,
			Department: r.FormValue("Department"),
			Origin:     r.FormValue("Origin"),
		}
		form.update(r.FormValue("Old_Name"))
		http.Redirect(w, r, "/Users/"+form.Name+"/profile", http.StatusFound)
	}
}

func urlHandler() {
	http.HandleFunc("/Users", viewAllUsers)
	http.HandleFunc("/Users/", viewHandler)
	http.HandleFunc("/create", creatProfile)
	http.HandleFunc("/update", editPorfile)
	//http.HandleFunc("/User/edit", creatProfile)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func main() {
	urlHandler()
}
