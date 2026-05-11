# port-check
Check if a TCP port is open on a remote host. Written in Go.
```bash
curl -LO https://github.com/javimosch/port-check/releases/latest/download/port-check-linux-amd64
chmod +x port-check-linux-amd64
sudo mv port-check-linux-amd64 /usr/local/bin/port-check
```
Usage: `port-check google.com 80` → `OPEN`
`port-check localhost 9999` → `CLOSED`
