# ipc-lab
Playing with protocols and encryption for IPC transports.

UDS

````sh
# socat
$ socat UNIX-LISTEN:/tmp/test.sock,fork -
$ echo hi | socat - UNIX-CONNECT:/tmp/test.sock
# or nc -U /tmp/test.sock
#
$ go run uds/server/server.go
$ go run uds/client/client.go
````

# todos
- [ ] https://pkg.go.dev/net/rpc
- [ ] tls
