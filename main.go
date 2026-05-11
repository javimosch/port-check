package main
import ("fmt";"net";"os";"strconv";"strings";"time")
func main() {
	if len(os.Args) < 2 { fmt.Fprintln(os.Stderr,"Usage: port-check <host:port> [timeout_secs]"); os.Exit(1) }
	target := os.Args[1]
	if !strings.Contains(target, ":") { target = "localhost:" + target }
	timeout := 3
	if len(os.Args) > 2 { if n,err:=strconv.Atoi(os.Args[2]); err==nil {timeout=n} }
	conn, err := net.DialTimeout("tcp", target, time.Duration(timeout)*time.Second)
	if err != nil { fmt.Printf("CLOSED (%v)\n", err); os.Exit(1) }
	conn.Close()
	fmt.Println("OPEN")
}
