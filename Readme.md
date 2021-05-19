# Make it go
0. Pull the repository, navigate to the home directory, and open a terminal.
1. Run `docker-compose up` from the main directory to get a mysql instance running on 3306
2. Navigate to the 'mysqlscanner' directory and create a binary with `go build scanner`
3. Run `./scanner <url> <port>` e.g.`./scanner 127.0.0.1 3306`


# Basics
- This is enough to tell you whether or not a mysql server is running and what version it is running (if it is at all).

- [There is more information you can get from the initial handshake](https://www.oreilly.com/library/view/understanding-mysql-internals/0596009577/ch04.html) if you parse the application packet. This includes:
  - Whether the server is in transaction or autocommit mode.
  - The default character set code.
  - [The server capabilities bit mask](https://www.oreilly.com/library/view/understanding-mysql-internals/0596009577/ch04.html#orm9780596009571-CHP-4-TABLE-5)


# For Fun
You can install `ngrep` with `brew install ngrep` and then run `sudo ngrep -x -q -d lo0 '' 'port 3306'` to see the packets being sent back and fourth. I used it to make sure that I was hitting the server and to inspect the packets being sent.