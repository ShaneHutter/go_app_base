# Go App Base
## This is an experimental base template for building a Go web app

Features I'm looking to achieve:
* Session Management
    * https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/06.2.html
* Brotli compression
    * br encoding working, currently 100% content is br
* Minify served content
    * https://pkg.go.dev/github.com/tdewolff/minify/html#example-Minify
* Server side Websocket setup
    * https://pkg.go.dev/github.com/googollee/go-socket.io
* Load YAML files for app config
    * https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64
* Connect to DB cluster (psql)
    * https://github.com/lib/pq
* Connect to Redis Node/Cluster
    * https://redis.io/docs/clients/go/
