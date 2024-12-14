package main

import (
	"encoding/csv"
	"fmt"
	"net/smtp"
	"os"
	"strconv"
	"strings"
)

func sendEmail(report string) {
	from := os.Getenv("EMAIL")
	password := os.Getenv("APP_PASS")
	to := "rs206987@gmail.com"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	msg := []byte("To: " + to + "\r\n" +
		"Subject: Web Crawler Report\r\n" +
		"\r\n" +
		report + "\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		fmt.Printf("failed to send email: %v\n", err)
	} else {
		fmt.Println("Email sent successfully")
	}
}

func generateReport(pages map[string]int, baseURL string) string {
	report := "=============================\n"
	report += fmt.Sprintf("  REPORT for %s\n", baseURL)
	report += "=============================\n"

	internalLinks, externalLinks := 0, 0
	sortedPages := sortPages(pages)
	for _, page := range sortedPages {
		if strings.Contains(page.URL, baseURL) {
			internalLinks += page.Count
		} else {
			externalLinks += page.Count
		}
		report += fmt.Sprintf("Found %d links to %s\n", page.Count, page.URL)
	}

	report += fmt.Sprintf("\nTotal internal links: %d\n", internalLinks)
	report += fmt.Sprintf("Total external links: %d\n", externalLinks)
	return report
}

func saveReportAsCSV(pages map[string]int, baseURL string) error {
    file, err := os.Create("report.csv")
    if err != nil {
        return fmt.Errorf("could not create CSV file: %w", err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    // Write header
    if err := writer.Write([]string{"URL", "Count"}); err != nil {
        return fmt.Errorf("could not write header to CSV: %w", err)
    }

    for url, count := range pages {
        if err := writer.Write([]string{url, strconv.Itoa(count)}); err != nil {
            return fmt.Errorf("could not write row to CSV: %w", err)
        }
    }

    return nil
}