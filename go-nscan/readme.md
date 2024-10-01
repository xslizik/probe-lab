### Nscan.go
Custom Network Scanner that can perform icmp ping sweep, and detect ports based on hostname or ips provided, that supports concurency.

### Build using
```bash
go mod init project-layout
go get github.com/prometheus-community/pro-bing
go mod tidy
go build nscan.go
```

### Test without compiling
```bash
go run .\nscan.go -r 192.168.56.1,10.10.20.5 -n scanme.nmap.org -b
```