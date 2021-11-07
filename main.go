package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/haccer/available"
)

var (
	domainName, listDomain, webhookURL string
)

func checkSSL(domain string, webhookURL string) {
	if !available.Domain(domain) {
		conn, err := tls.Dial("tcp", domain+":443", nil)
		if err != nil {
			panic("Server doesn't support SSL certificate err: " + err.Error())
		}
		err = conn.VerifyHostname(domain)
		if err != nil {
			panic("Hostname doesn't match with certificate: " + err.Error())
		}
		expiry := conn.ConnectionState().PeerCertificates[0].NotAfter
		date := expiry
		currentTime := time.Now()
		expiredDays := int(date.Sub(currentTime).Hours() / 24)
		fmt.Printf("Domain: %s\nIssuer: %s\nExpiry Date: %v\nDays: %v day\n=================\n", domain, conn.ConnectionState().PeerCertificates[0].Issuer, date.Format(time.RFC850), expiredDays)
		if expiredDays < 30 {
			m := map[string]string{"content": "Hello @everyone, domain " + domain + " mau expired nih tanggal " + date.Format(time.RFC822)}
			r, w := io.Pipe()
			go func() {
				json.NewEncoder(w).Encode(m)
				w.Close()
			}()
			http.Post(webhookURL, "application/json", r)
		}
	}
}
func main() {
	flag.StringVar(&webhookURL, "webhook", "", "Your Webhook URL")
	flag.StringVar(&domainName, "d", "", "Type your domain like domain.com")
	flag.StringVar(&listDomain, "L", "", "List file your domain")
	flag.Parse()
	if len(domainName) != 0 && len(webhookURL) != 0 {
		checkSSL(domainName, webhookURL)
	} else if len(listDomain) != 0 && len(webhookURL) != 0 {
		file, err := os.Open(string(listDomain))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			checkSSL(scanner.Text(), webhookURL)
		}
	}
}
