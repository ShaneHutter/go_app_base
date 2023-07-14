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
	//"reflect"
	//"io"
	"bytes"
	brotli "github.com/google/brotli/go/cbrotli"
)

type PageData struct {
	Title string
	Path string
	Foobar string
}


var staticDir string = "static/"
var templateDir string = "templates/"
var fileMatchRegex string = "[^_]+.*\\..*$"
var listenAddress string = ":8080"
var brotlyOpts brotli.WriterOptions


func handler (w http.ResponseWriter , r *http.Request ){
	// Set compression on writer
	brotlyOpts.Quality = 9
	brotlyOpts.LGWin = 0 
	w.Header().Set( "Content-Encoding" , "br" )
	w.Header().Set( "Accept-Encoding" , "br")
	//bw := brotli.NewWriter( w , brotli.WriterOptions{ Quality: 9 , LGWin: 0 } )
	//defer bw.Close()

	// Handle response
	pageData := &PageData{ Path: r.URL.Path[ 1: ] , Foobar: "Foo Bar Baz" }
	templateFilename := templateDir + strings.ReplaceAll( pageData.Path , "/" , "_" )
	staticFilename := staticDir + pageData.Path
	isFile , err := regexp.MatchString( fileMatchRegex , templateFilename )
	if err != nil {
		panic( err )
	}

	if templateFilename == templateDir {
		// Index
		tmpl , err := template.ParseFiles( templateFilename + "index" )
		if err != nil {
			// Error 404: No Index
			http.Error( w , err.Error() , http.StatusNotFound )
		} else { 
			var templateBuffer bytes.Buffer
			tmpl.Execute( &templateBuffer , pageData ) 
			bt , err := brotli.Encode( templateBuffer.Bytes() , brotlyOpts )
			if err != nil {
				panic( err )
			} else { w.Write( bt ) }
		}

	} else if ! isFile {
		tmpl , err := template.ParseFiles( templateFilename )
		if err != nil {
			// Error 404
			http.Error( w , err.Error() , http.StatusNotFound )
		} else { 
			var templateBuffer bytes.Buffer
			tmpl.Execute( &templateBuffer , pageData ) 
			bt , err := brotli.Encode( templateBuffer.Bytes() , brotlyOpts )
			if err != nil {
				panic( err )
			} else { w.Write( bt ) }
		}
	} else {
		// Serve static files here

		// Old, no compression file serving
		/*
		fmt.Printf( "File: %s\n" , staticFilename )
		http.ServeFile( w , r , staticFilename )
		*/

		/*
			 For Brotli compression, you will need to load, compress, and 
			 write the file to http.ResponseWriter

			 It may make more sense to brotli compress all static files,
			 and serve them as such.
		*/
		http.ServeFile( w , r , staticFilename + ".br" )


	}
	//bw.Close()
}


// Main
func main() {
	//http.HandleFunc( "/" , handler )
	err := http.ListenAndServe( listenAddress , http.HandlerFunc( handler ) )
	if err != nil {
		log.Fatal( err.Error() )
	}
}