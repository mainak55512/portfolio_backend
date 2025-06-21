package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type Entity struct {
	ProjectName        string `json:"name"`
	ProjectDescription string `json:"description"`
	ResourcePath       string `json:"path"`
	UsedTech           string `json:"tech"`
}

type Blogs struct {
	Title     string `json:"title"`
	URL       string `json:"url"`
	Thumbnail string `json:"thumbnail"`
}

func loadEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}

func readJson(filename string) ([]byte, error) {
	err := loadEnv()
	if err != nil {
		return nil, err
	}
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	jsonFile, err := os.ReadFile(wd + "/" + filename)
	if err != nil {
		return nil, err
	}
	return jsonFile, nil
}

func getEntities(w http.ResponseWriter, r *http.Request) {
	allowedURL := os.Getenv("CLIENT_URL")
	w.Header().Set("Access-Control-Allow-Origin", allowedURL)
	entityJson, err := readJson("entity.json")
	if err != nil {
		log.Fatal("JSON Error! Cannot read file")
	}
	var entityList []Entity
	err = json.Unmarshal(entityJson, &entityList)
	if err != nil {
		log.Fatalf("Unmarshal error! %s", err)
	}
	jsonObj, err := json.MarshalIndent(entityList, "", "  ")
	if err != nil {
		log.Fatal("JSON Error! Cannot parse JSON")
	}
	fmt.Fprintf(w, "%s", string(jsonObj))
}

func getBlogs(w http.ResponseWriter, r *http.Request) {
	allowedURL := os.Getenv("CLIENT_URL")
	w.Header().Set("Access-Control-Allow-Origin", allowedURL)
	blogJson, err := readJson("blogs.json")
	if err != nil {
		log.Fatal("JSON Error! Cannot read file")
	}
	var blogList []Blogs
	err = json.Unmarshal(blogJson, &blogList)
	if err != nil {
		log.Fatalf("Unmarshal error! %s", err)
	}
	jsonObj, err := json.MarshalIndent(blogList, "", "  ")
	if err != nil {
		log.Fatal("JSON Error! Cannot parse JSON")
	}
	fmt.Fprintf(w, "%s", string(jsonObj))
}

func main() {
	http.HandleFunc("/api/entities", getEntities)
	http.HandleFunc("/api/blogs", getBlogs)
	log.Println("Server running at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
