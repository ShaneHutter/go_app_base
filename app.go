/*


Ternary example
	{ Foobar: ( map[bool]string{ true: "True" , false: "False" } )[ foobar ] }
*/

package main

import (
	//"fmt"
	"strings"
	//"os"
	"log"
	"net/http"
	"html/template"
	"regexp"
)

type PageData struct {
	Title string
	Foobar string
}

var staticDir string = "static/"
var templateDir string = "templates/"
var fileMatchRegex string = "[^_]+.*\\..*$"

func handler (w http.ResponseWriter , r *http.Request ){
	templateFilename := staticDir + strings.ReplaceAll( r.URL.Path[ 1: ] , "/" , "_" )
	isFile , err := regexp.MatchString( fileMatchRegex , templateFilename )
	if err != nil {
		panic( err )
	}
	if ! isFile {
		t , err := template.ParseFiles( templateFilename )
		if err != nil {
			panic( err )
		}
		t.Execute( w , &PageData{ 
			Title: templateFilename, 
			})
		}
	}


// Main
func main() {
	http.HandleFunc( "/" , handler )
	log.Fatal( http.ListenAndServe( ":8080" , nil )	 )
}