package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// Создаю структуру для хранения полученных от whois данных

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
	// Получаю информацию о домене
	domain := "yandex.ru"
	var whoisInfo WhoisInfo

	cmd := exec.Command("./whois.exe", domain)
	// Проверяю, получил ли мы информацию о домене
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	whoisInfo = parseWhoisOutput(string(output))
	
	// Массив байт возвращаю методом json.Marshal() В строку проеобразовал с помощью string() передав в него массив байт.

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
		switch {
		case strings.HasPrefix(line, "Domain Name:"):
			result.DomainName = strings.TrimSpace(strings.TrimPrefix(line, "Domain Name:"))
		case strings.HasPrefix(line, "Registrar:"):
			result.Registrar = strings.TrimSpace(strings.TrimPrefix(line, "Registrar:"))
		case strings.HasPrefix(line, "Creation Date:"):
			result.CreationDate = strings.TrimSpace(strings.TrimPrefix(line, "Creation Date:"))
		case strings.HasPrefix(line, "Expiration Date:"):
			result.ExpirationDate = strings.TrimSpace(strings.TrimPrefix(line, "Expiration Date:"))
		case strings.HasPrefix(line, "Updated Date:"):
			result.UpdatedDate = strings.TrimSpace(strings.TrimPrefix(line, "Updated Date:"))
		case strings.HasPrefix(line, "Name Server:"):
			result.NameServer = strings.TrimSpace(strings.TrimPrefix(line, "Name Server:"))
		case strings.HasPrefix(line, "Registrar WHOIS Server:"):
			result.RegistrarWhois = strings.TrimSpace(strings.TrimPrefix(line, "Registrar WHOIS Server:"))
		case strings.HasPrefix(line, "Registrar URL:"):
			result.RegistrarURL = strings.TrimSpace(strings.TrimPrefix(line, "Registrar URL:"))
		case strings.HasPrefix(line, "Registrant Name:"):
			result.RegistrantName = strings.TrimSpace(strings.TrimPrefix(line, "Registrant Name:"))
		case strings.HasPrefix(line, "Registrant Organization:"):
			result.RegistrantOrganization = strings.TrimSpace(strings.TrimPrefix(line, "Registrant Organization:"))
		case strings.HasPrefix(line, "domain:"):
			result.DomainName = strings.TrimSpace(strings.TrimPrefix(line, "domain:"))
		case strings.HasPrefix(line, "registrar:"):
			result.Registrar = strings.TrimSpace(strings.TrimPrefix(line, "registrar:"))
		case strings.HasPrefix(line, "created:"):
			result.CreationDate = strings.TrimSpace(strings.TrimPrefix(line, "created:"))
		case strings.HasPrefix(line, "free-date:"):
			result.ExpirationDate = strings.TrimSpace(strings.TrimPrefix(line, "free-date:"))
		case strings.HasPrefix(line, "paid-till:"):
			result.UpdatedDate = strings.TrimSpace(strings.TrimPrefix(line, "paid-till:"))
		case strings.HasPrefix(line, "nserver:"):
			result.NameServer = strings.TrimSpace(strings.TrimPrefix(line, "nserver:"))
		case strings.HasPrefix(line, "source:"):
			result.RegistrarWhois = strings.TrimSpace(strings.TrimPrefix(line, "source:"))
		case strings.HasPrefix(line, "admin-contact:"):
			result.RegistrarURL = strings.TrimSpace(strings.TrimPrefix(line, "admin-contact:"))
		case strings.HasPrefix(line, "org:"):
			result.RegistrantName = strings.TrimSpace(strings.TrimPrefix(line, "org:"))
		case strings.HasPrefix(line, "org:"):
			result.RegistrantOrganization = strings.TrimSpace(strings.TrimPrefix(line, "org:"))
		}
	}

	return result
}
