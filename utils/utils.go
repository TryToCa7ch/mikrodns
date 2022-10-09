package utils

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-routeros/routeros"
)

var (
	Address  = os.Getenv("MIKROTIK_HOST")
	Username = os.Getenv("MIKROTIK_USER")
	Password = os.Getenv("MIKROTIK_PASS")
	Tls      = os.Getenv("MIKROTIK_TLS")
)

type DnsRecord struct {
	Address, Host, Disabled string
}

func Dial() (*routeros.Client, error) {
	return routeros.Dial(Address, Username, Password)
}

func AddDnsRecord(c *routeros.Client, hostname string, address string) string {
	command := fmt.Sprintf("/ip/dns/static/add =name=%s =address=%s", hostname, address)
	r, err := c.RunArgs(strings.Split(command, " "))
	log.Print(r)
	if err != nil {
		log.Print(err)
		return err.Error()
	}
	return "Added"
}

func GetAllDnsRecords(c *routeros.Client) []DnsRecord {
	r, _ := c.Run("/ip/dns/static/print")
	record_list := []DnsRecord{}
	for _, re := range r.Re {
		var record DnsRecord
		record.Address = re.Map["name"]
		record.Host = re.Map["address"]
		record.Disabled = re.Map["disabled"]
		record_list = append(record_list, record)
	}
	return record_list
}
