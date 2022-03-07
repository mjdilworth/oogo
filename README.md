# oogo
## introduction
I will use this to refresh and deepen my understanding of go

## some useful tips
dont forget to creat server certs

open TCP ports
netstat -anvp tcp | awk 'NR<3 || /LISTEN/'

or

lsof -i -P | grep LISTEN | grep :$PORT

## Cross compilation
env GOOS=linux GOARCH=arm64 go build -o filename_arm64
env GOOS=darwin GOARCH=arm64 go build -o filename_arm64
env GOOS=linux GOARCH=amd64 go build -o filename_arm64

etc.


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


## artifactory

this wont run on my rasperbbry pi as it is too old

#### To determine your distribution, run lsb_release -c or cat /etc/os-release
# Example:echo "deb https://releases.jfrog.io/artifactory/artifactory-pro-debs xenial main" | sudo tee -a /etc/apt/sources.list;
wget -qO - https://releases.jfrog.io/artifactory/api/gpg/key/public | sudo apt-key add -;
echo "deb https://releases.jfrog.io/artifactory/artifactory-debs {distribution} main" | sudo tee -a /etc/apt/sources.list;
sudo apt-get update && sudo apt-get install jfrog-artifactory-jcr


HOWEVER - this doe swork

https://gabrieltanner.org/blog/docker-registry

use docker compose

version: '3'

services:
  registry:
    image: registry:2
    ports:
    - "5000:5000"


    andi can list whats in the registry

     curl localhost:5000/v2/_catalog

     