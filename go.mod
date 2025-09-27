module sportsstore

go 1.23.3

toolchain go1.24.7

require platform v1.0.0

require (
	github.com/gorilla/securecookie v1.1.2 // indirect
	github.com/gorilla/sessions v1.4.0 // indirect
)

replace platform v1.0.0 => ../platform
