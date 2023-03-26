package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

type WhoisInfo struct {
	DomainName             string
	Registrar              string
	CreationDate           string
	ExpirationDate         string
	UpdatedDate            string
	NameServer             string
	RegistrarWhois         string
	RegistrarURL           string
	RegistrantName         string
	RegistrantOrganization string
}

func main() {
	cmd := exec.Command("./whois.exe", "vk.com")

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	whoisInfo := parseWhoisOutput(string(output))

	jsonBytes, err := json.MarshalIndent(whoisInfo, "", "  ")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Println(string(jsonBytes))
}

func parseWhoisOutput(output string) WhoisInfo {
	result := WhoisInfo{}

	lines := strings.Split(output, "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "Domain Name:") {
			result.DomainName = strings.TrimSpace(strings.TrimPrefix(line, "Domain Name:"))
		} else if strings.HasPrefix(line, "Registrar:") {
			result.Registrar = strings.TrimSpace(strings.TrimPrefix(line, "Registrar:"))
		} else if strings.HasPrefix(line, "Creation Date:") {
			result.CreationDate = strings.TrimSpace(strings.TrimPrefix(line, "Creation Date:"))
		} else if strings.HasPrefix(line, "Expiration Date:") {
			result.ExpirationDate = strings.TrimSpace(strings.TrimPrefix(line, "Expiration Date:"))
		} else if strings.HasPrefix(line, "Updated Date:") {
			result.UpdatedDate = strings.TrimSpace(strings.TrimPrefix(line, "Updated Date:"))
		} else if strings.HasPrefix(line, "Name Server:") {
			result.NameServer = strings.TrimSpace(strings.TrimPrefix(line, "Name Server:"))
		} else if strings.HasPrefix(line, "Registrar WHOIS Server:") {
			result.RegistrarWhois = strings.TrimSpace(strings.TrimPrefix(line, "Registrar WHOIS Server:"))
		} else if strings.HasPrefix(line, "Registrar URL:") {
			result.RegistrarURL = strings.TrimSpace(strings.TrimPrefix(line, "Registrar URL:"))
		} else if strings.HasPrefix(line, "Registrant Name:") {
			result.RegistrantName = strings.TrimSpace(strings.TrimPrefix(line, "Registrant Name:"))
		} else if strings.HasPrefix(line, "Registrant Organization:") {
			result.RegistrantOrganization = strings.TrimSpace(strings.TrimPrefix(line, "Registrant Organization:"))
		}
	}

	return result
}
