# go_plugin_explained

I try to explain modular programming in go as simple as possible. Read my post in Burmese
[here](https://medium.com/@kaungmyatchanthar/%E1%80%96%E1%80%BC%E1%80%AF%E1%80%90%E1%80%BA%E1%80%9C%E1%80%BD%E1%80%9A%E1%80%BA-%E1%80%90%E1%80%95%E1%80%BA%E1%80%9C%E1%80%BD%E1%80%9A%E1%80%BA-go-3aa01b8450e4)

- Create shared object libraries

  `go build -buildmode=plugin -o dev/dev.so dev/persona.go`

  `go build -buildmode=plugin -o monk/monk.so monk/persona.go`

- We are ready to go

  `go run plug_person.go dev`
