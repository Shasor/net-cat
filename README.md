## net-cat

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Shasor/net-cat)

### Description

This project consists on recreating the NetCat in a Server-Client Architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.

### Start the Server

In order to run the project correctly, your terminal <ins>MUST</ins> be initialized in the root folder (net-cat), and proceed as follows:

```bash
$ git clone $url # $url = link to my project (github, gitea...)
$ cd net-cat/
$ make build
$ ./TCPChat
Listening on the port :8989
$ ./TCPChat 2525
Listening on the port :2525
$
```

### Author(s)

-
-
- [Adam GONÇALVES](https://zone01normandie.org/git/agoncalv) (aka [Shasor#3755](https://discordapp.com/users/282816260075683841))
