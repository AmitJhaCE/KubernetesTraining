package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    return localAddr.IP
}

//Handler
func handler(w http.ResponseWriter, r *http.Request) {
	localAddr := GetOutboundIP()
	fmt.Fprintln(w, "Hello, 世界.", "Host IP address is:", localAddr)
}

func main() {
	http.HandleFunc("/go", handler)
	fmt.Println("Running demo app. Press Ctrl+C to exit...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
