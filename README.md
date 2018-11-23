# golang-grpc
My Golang gRPC boilerplate

## service test case
Bulubulu Planet is the most remote planet in the universe (relative to earth), scientist has found life thriving on it surface. Surprisingly bulubulueans understand our language but our writings, they only use letter "u" as vowel and they write "b" and "l" is swapped one another. Help scientist to communicate with them by translating our writings so bulubulueans can understand.

## install dependency
```sh
~$ dep ensure
```

## start server
```sh
$ go build -o server .
$ ./server
```
Option flag: `-p {int PORT}` (default `8080`)

## start client
```sh
/client$ go build .
/client$ ./client -t "This is human speaking"
```
Option flag: `-b {string HOST}` (default `localhost:8080`), `-t {string TEXTTOTRANSLATE}` (default `""`)
