package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body []byte
}

// Saves a page
func ( p *Page) save() error {
	filename := p.Title + ".jinja"
	return os.WriteFile( filename , p.Body , 0600 )
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


// Main
func main() {
	p1 := &Page{ Title: "foobar" , Body: []byte( "Foo Bar Baz" ) }
	p1.save()
	p2 , _ := loadJinja( "foobar" )
	fmt.Println( string( p2.Body ) )
}