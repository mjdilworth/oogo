#oogo#
##introduction##
I will use this to refresh and deepen my understanding of go


##Structure##

- i am looking to keep to a minimal main func
- vendor is the place to put in any 3rd party libraies and pin them to specif version - vendoring


├── LICENSE
├── README.md
├── clientlib
│   ├── ooclient.go
│   └── ooclient_test.go
├── cmd
│   ├── oogo-client
│   └── oogo-server
├── config.go
├── config_test.go
├── go.mod
├── internal
│   └── auth
│       ├── auth.go
│       └── auth_test.go
├── servicelib
│   ├── ooservice.go
│   └── ooservice_test.go
└── vendor



i plan to cover topics such as:
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
