package main

// go standard library is like really powerful
import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMx, hasSPF, SPFRecord, hasDMARC, dmarcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	// var instantiation and if statement
	if err := scanner.Err(); err != nil {
		log.Fatal("ERROR: Could not read the input: %v\n", err)
	}
}

func checkDomain(domain string) {

	var hasMx, hasSPF, hasDMARC bool
	var SPFRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("ERROR: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMx = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	// this is kinda like enumeration for loop in python
	// the _ is needed but that is normally the index
	// record is the actual element
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			SPFRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc" + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMx, hasSPF, SPFRecord, hasDMARC, dmarcRecord)

}
