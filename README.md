# GO-Chat!

### Telnet chat server in Golang
Project is build using Golang standard project layout taken from
[here](https://github.com/golang-standards/project-layout).

## List of all planned and existing functionalities
- [X] CLI with help
- [X] Telnet server:
    - [X] Base server
- [X] Chat:
    - [X] Base chat
    - [X] Support for rooms

## List of used libraries
* [cli](https://github.com/urfave/cli) to build nice cli
* [logrus](https://github.com/sirupsen/logrus) for logging
* [yaml](https://github.com/go-yaml/yaml) for configuration file

## Problems which I had to solve
* Golang - it was my first project in that language Understanding embedding
* Coding standard, programming in Golang, a structural language is completely
    different that coding in OOP like Python
* ([docs](https://golang.org/doc/effective_go.html#embedding)).  All Telnet
* commands should self register itself for easiness of future
  development. Init was super helpful to achieve that
  ([docs](https://golang.org/doc/effective_go.html#init)).
* I'm completely not used to use interfaces so I struggled dramatically with
  naming convention
