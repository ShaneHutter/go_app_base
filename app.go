/*


Ternary example
	{ Foobar: ( map[bool]string{ true: "True" , false: "False" } )[ foobar ] }
*/

package main

import (
	"fmt"
	"strings"
	//"os"
	"log"
	"net/http"
	"html/template"
	"regexp"
	//"reflect"
)

type PageData struct {
	Title string
	Path string
	Foobar string
}

var staticDir string = "static/"
var templateDir string = "templates/"
var fileMatchRegex string = "[^_]+.*\\..*$"


func handler (w http.ResponseWriter , r *http.Request ){
	pageData := &PageData{ Path: r.URL.Path[ 1: ] , Foobar: "Foo Bar Baz" }
	templateFilename := templateDir + strings.ReplaceAll( pageData.Path , "/" , "_" )
	staticFilename := staticDir + pageData.Path
	isFile , err := regexp.MatchString( fileMatchRegex , templateFilename )
	if err != nil {
		panic( err )
	}
	if templateFilename == templateDir {
		// Index
		t , err := template.ParseFiles( templateFilename + "index" )
		if err != nil {
			// Error 404: No Index
			http.Error( w , err.Error() , http.StatusNotFound )
		} else { t.Execute( w , pageData ) }
	} else if ! isFile {
		t , err := template.ParseFiles( templateFilename )
		if err != nil {
			// Error 404
			http.Error( w , err.Error() , http.StatusNotFound )
			} else { t.Execute( w , pageData ) }
	} else {
		// Serve static files here
		fmt.Printf( "File: %s\n" , staticFilename )
		http.ServeFile( w , r , staticFilename )
	}
}


// Main
func main() {
	http.HandleFunc( "/" , handler )
	err := http.ListenAndServe( ":8080" , nil )
	if err != nil {
		log.Fatal( err.Error() )
	}
}