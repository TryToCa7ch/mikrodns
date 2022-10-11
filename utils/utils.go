package utils

import (
	"errors"
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

	ErrNotFound     = errors.New("not found")
	ErrEmptyRecList = errors.New("Empty Static DNS List")
)

type DnsRecord struct {
	Id, Address, Host, Disabled string
}

func Dial() (*routeros.Client, error) {
	fmt.Printf(Info("\nConnecting to: %s as %s\n\n"), Address, Username)
	return routeros.Dial(Address, Username, Password)
}

func AddDnsRecord(c *routeros.Client, hostname string, address string) string {
	command := fmt.Sprintf("/ip/dns/static/add =name=%s =address=%s", hostname, address)
	_, err := c.RunArgs(strings.Split(command, " "))
	if err != nil {
		log.Print(err)
		return err.Error()
	}
	return "Added"
}

func GetAllDnsRecords(c *routeros.Client) ([]DnsRecord, error) {
	r, _ := c.Run("/ip/dns/static/print")
	record_list := []DnsRecord{}
	if len(r.Re) != 0 {
		for _, re := range r.Re {
			var record DnsRecord
			record.Id = re.Map[".id"]
			record.Address = re.Map["name"]
			record.Host = re.Map["address"]
			record.Disabled = re.Map["disabled"]
			record_list = append(record_list, record)
		}
		return record_list, nil
	}
	return record_list, fmt.Errorf("%w", ErrEmptyRecList)
}

func GetDnsRecord(c *routeros.Client, id string) (DnsRecord, error) {
	r, _ := c.Run("/ip/dns/static/print")
	var recToReturn DnsRecord
	for _, re := range r.Re {
		if re.Map[".id"][1:] == id {
			recToReturn.Id = re.Map[".id"]
			recToReturn.Address = re.Map["name"]
			recToReturn.Host = re.Map["address"]
			recToReturn.Disabled = re.Map["disabled"]
			return recToReturn, nil
		}
	}
	return recToReturn, fmt.Errorf("%q: %w", id, ErrNotFound)
}
