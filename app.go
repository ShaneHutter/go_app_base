package main

import (
	//"github.com/kluctl/go-jinja2"
	"fmt"
	"strings"
	"os"
	"log"
	"net/http"
)

type PageTemplate struct {
	Filename string
	Body []byte
}

var staticDir string = "static/"

// Loads a file, use this as a start to load a Jinja template
func loadPageTemplate( filename string ) ( *PageTemplate , error ){
	body , err := os.ReadFile( filename )
	if err != nil {
		return nil , err 
	}
	return &PageTemplate{ Filename: filename , Body: body } , nil
}

func handler (w http.ResponseWriter , r *http.Request ){
	templateFilename := staticDir + strings.ReplaceAll( r.URL.Path[ 1: ] , "/" , "_" )
	template , err := loadPageTemplate( templateFilename )
	if err != nil {
		panic( err )
	}
	fmt.Fprintf( w , string( template.Body ) )
}


// Main
func main() {
	http.HandleFunc( "/" , handler )
	log.Fatal( http.ListenAndServe( ":8080" , nil )	 )
}