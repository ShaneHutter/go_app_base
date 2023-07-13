# Go App Base
## This is an experimental base template for building a Go web app that includes features I commonly use in Python/Flask applications

Features to replicate:
* Session Management
* ~~render Jinja into html/css/js~~
    * This is causing issues, look for other rendering methods?
    * Try html/template first.  This seems to be standard for Golang apps.
* G-Zip/Brotli compression
* Server side Websocket setup
* Load YAML files for app config
* Connect to DB cluster (psql)
* Connect to Redis Node/Cluster