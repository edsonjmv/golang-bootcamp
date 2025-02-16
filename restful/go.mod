module main.go

go 1.18

require (
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	restful/controllers v0.0.0-00010101000000-000000000000 // indirect
	restful/global v0.0.0-00010101000000-000000000000 // indirect
)

replace restful/controllers => ./controllers

replace restful/global => ./global
