# GO-Chat!

### Telnet chat server in Golang

Project is build using Golang standard project layout taken from [here](https://github.com/golang-standards/project-layout).

## List of all planned and existing functionalities
- [ ] CLI with help
- [ ] Telnet server:
    - [ ] Base server
    - [ ] TELNETS support
- [ ] Chat:
    - [ ] Base chat
    - [ ] Support for rooms 
- [ ] REST API:
    - [ ] Get all messages
    - [ ] Post new message
- [ ] Multi master server
- [ ] Automatic testing on GH

## List of used libraries
* [cli](https://github.com/urfave/cli) to build nice cli
* [logrus](https://github.com/sirupsen/logrus) for logging
* [go-telnet](https://github.com/reiver/go-telnet) for telnet server
* yaml

## Problems which I had to solve
* Understanding embedding ([docs](https://golang.org/doc/effective_go.html#embedding)).
* All Telnet commands should self register itself for easines of future
    development. Init was super helpful to achieve that ([docs](https://golang.org/doc/effective_go.html#init)).
