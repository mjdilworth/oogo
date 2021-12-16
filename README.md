# oogo
## introduction
I will use this to refresh and deepen my understanding of go
## Structure

- i am looking to keep to a minimal main func
- vendor is the place to put in any 3rd party libraies and pin them to specif version - vendoring

```
.
├── LICENSE
├── README.md
├── cmd
│   └── oogo-client
│       ├── main.go
│       └── oogo-client
├── config.go
├── config_test.go
├── design
│   ├── architecture.md
│   ├── oo-go.drawio
│   ├── oo-go.md
│   └── oo-go.svg
├── go.mod
├── internal
│   └── auth
│       ├── auth.go
│       └── auth_test.go
├── main.go
├── serviceOne
│   ├── serviceOne.go
│   └── serviceOne_test.go
├── serviceThree
│   ├── serviceThree.go
│   └── serviceThree_test.go
├── serviceTwo
│   ├── serviceTwo.go
│   └── serviceTwo_test.go
└── vendor
```


cover topics such as:
- modules
- packages
- vendoring
- godoc
- encapsulation
- polymorphism
- inheritance
- unit testing
- interfaces
- channels
- multi-threading
- build pipelines
- disrtoless containers
- functions as a service
