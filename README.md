# Go App Base
## This is an experimental base template for building a Go web app

Features I'm looking to achieve:
* Session Management
* Brotli compression
    * br encoding working, currently 100% content is br
* Minify served content
    * https://pkg.go.dev/github.com/tdewolff/minify/html#example-Minify
* Server side Websocket setup
* Load YAML files for app config
* Connect to DB cluster (psql)
* Connect to Redis Node/Cluster