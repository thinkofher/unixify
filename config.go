package main

var appName = "unixify"
var version = "0.1.0"
var author = "Beniamin Dudek"
var email = "beniamin.dudek@yahoo.com"

var usageRaw = `USAGE:
    {{appName}} [OPTIONS]

INFO:
    Simple portable command line interface tool
    for changing file and folder names to be more
    unixy.

    Example:

    * "Some File"         -> "some_file"
    * "sOme File"         -> "some_file"
    * "SOMEFILE"          -> "some_file"
    * ".some hidden file" -> ".some_hidden_file"

AUTHOR:
    {{author}} <{{email}}>

VERSION:
    {{version}}

OPTIONS:
    -dirs
        include directories
    -hidden
        include hidden files
    -interactive
        asks before every rename
    -verbose
        print current actions
    -yes
        skip asking to rename at startup
`
