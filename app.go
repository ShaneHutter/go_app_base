package main

import (
	"fmt"
	"strings"
	"os"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body []byte
}

// Loads a file, use this as a start to load a Jinja template
func loadJinja( title string ) ( *Page , error ){
	filename := title + ".jinja"
	body , err := os.ReadFile( filename )
	if err != nil {
		return nil , err 
	}
	return &Page{ Title: title , Body: body } , nil
}

func handler (w http.ResponseWriter , r *http.Request ){
	title := strings.ReplaceAll( r.URL.Path[ 1: ] , "/" , "_" )
	jinja , _ := loadJinja( title )
	fmt.Fprintf( w , string( jinja.Body ) )
}


// Main
func main() {
	http.HandleFunc( "/" , handler )
	log.Fatal( http.ListenAndServe( ":8080" , nil )	 )
}