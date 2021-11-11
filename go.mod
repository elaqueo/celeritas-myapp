module myapp

go 1.17

replace github.com/elaqueo/celeritas => ../celeritas

require (
	github.com/elaqueo/celeritas v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi/v5 v5.0.5
)

require (
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/CloudyKit/jet/v6 v6.1.0 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
)
