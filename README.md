## net-cat ![Static Badge](https://img.shields.io/badge/Go-v1.23.0-blue)

### Description

This project consists on recreating the NetCat in a Server-Client Architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.

### Usage

```bash
$ pwd
/home/username/.../net-cat
$ go build -o TCPChat main.go
$ ./TCPChat
Listening on the port :8989
$ ./TCPChat 2525
Listening on the port :2525
$ ./TCPChat 2525 localhost
[USAGE]: ./TCPChat $port
$
```

### Author(s)

-
-
- [Adam GONÇALVES](https://zone01normandie.org/git/agoncalv) (aka [Shasor#3755](https://discordapp.com/users/282816260075683841))
