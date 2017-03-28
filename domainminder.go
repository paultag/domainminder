package main

import (
	"fmt"
	"os"

	"bufio"

	"crypto/tls"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		host := scanner.Text()
		conn, err := tls.Dial("tcp", host, nil)
		if err != nil {
			fmt.Printf("# domainminder error: %s. Fix me!\n", err)
			continue
		}
		defer conn.Close()

		for _, chain := range conn.ConnectionState().VerifiedChains {
			leaf := chain[0]

			fmt.Printf(
				"REM %s TAG cert MSG TLS Certificate for %s expires.%%\n",
				leaf.NotAfter.Format("2006-01-02"),
				leaf.Subject.CommonName,
			)
		}
	}
}
