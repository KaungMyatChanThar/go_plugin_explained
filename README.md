# go_plugin_explained

I try to explain modular programming in go as simple as possible.

- Create shared object libraries

  `go build -buildmode=plugin -o dev/dev.so dev/persona.go`

  `go build -buildmode=plugin -o monk/monk.so monk/persona.go`

- We are ready to go

  `go run plug_person.go dev`
