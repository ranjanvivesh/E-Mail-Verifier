package main

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
	fmt.Println("domain,hasMX,hasSPF,sprRecord,hasDMARC,dmarcRecord")

	for scanner.Scan(){
		checkDomain(scanner.Text())
	}

	if err := scanner.Err();err != nil {
		log.Fatal("Error:Could not read from the input \n",err)
	}
}

func checkDomain(domain string) {

	var hasMX,hasSPF,hasDMARC bool
	var spfRecord,dmarcRecord string
	mxRecord,err := net.LookupMX(domain)

	if err != nil{
		log.Fatal(err)
	}

	if len(mxRecord) >0 {
		fmt.Println(mxRecord)
	}

	txtRecords , err := net.LookupTXT(domain)

	
	if err != nil{
		log.Fatal(err)
	}

	for _,record := range txtRecords{
		if strings.HasPrefix(record,"v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords,err:= net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Fatal(err)
	}

	for _,record := range dmarcRecords{
		if strings.HasPrefix(record,"v=DMARC1"){
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("DOMAIN : %v\n hasMX : %v \n hasSPF : %v \n spfRECORD: %v \n hasDMARC : %v \n dmarcRECORD : %v \n",domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord)

}