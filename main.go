package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cloudflare/cloudflare-go"
)

const (
	// Service to get public IP address. Alternatives: https://ifconfig.me/ip, https://api.ipify.org
	ipEchoService = "https://icanhazip.com"
	recordType    = "A" // We are updating IPv4 address records
)

func main() {
	log.Println("Starting Cloudflare DDNS Client...")

	// --- Configuration ---
	// Read configuration from environment variables
	// Best practice: Use a dedicated config file or flags for production
	apiToken := os.Getenv("CF_API_TOKEN")
	zoneName := os.Getenv("CF_ZONE_NAME")   // e.g., "example.com"
	recordName := os.Getenv("CF_RECORD_NAME") // e.g., "home.example.com" or "example.com"

	if apiToken == "" || zoneName == "" || recordName == "" {
		log.Fatalln("Error: Missing required environment variables: CF_API_TOKEN, CF_ZONE_NAME, CF_RECORD_NAME")
	}

	// --- Get Current Public IP ---
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // Context with timeout
	defer cancel()

	publicIP, err := getCurrentPublicIP(ctx)
	if err != nil {
		log.Fatalf("Error getting public IP: %v\n", err)
	}
	log.Printf("Current public IP: %s\n", publicIP)

	// --- Cloudflare API Interaction ---
	api, err := cloudflare.NewWithAPIToken(apiToken)
	if err != nil {
		log.Fatalf("Error creating Cloudflare API client: %v\n", err)
	}

	// --- Find Zone ID ---
	zoneID, err := api.ZoneIDByName(zoneName)
	if err != nil {
		log.Fatalf("Error finding Zone ID for '%s': %v\n", zoneName, err)
	}
	log.Printf("Found Zone ID for '%s': %s\n", zoneName, zoneID)

	// --- Find DNS Record ---
	// We need to list records and find the one matching the name and type
	records, _, err := api.ListDNSRecords(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.ListDNSRecordsParams{
		Type: recordType,
		Name: recordName,
	})
	if err != nil {
		log.Fatalf("Error listing DNS records for '%s': %v\n", recordName, err)
	}

	if len(records) == 0 {
		log.Fatalf("Error: No '%s' record found with name '%s' in zone '%s'\n", recordType, recordName, zoneName)
	}
	if len(records) > 1 {
		log.Printf("Warning: Found multiple '%s' records for '%s'. Using the first one found (ID: %s).\n", recordType, recordName, records[0].ID)
	}

	targetRecord := records[0] // Use the first record found
	log.Printf("Found DNS Record: ID=%s, Name=%s, Type=%s, Content=%s, Proxied=%t\n",
		targetRecord.ID, targetRecord.Name, targetRecord.Type, targetRecord.Content, *targetRecord.Proxied)

	// --- Compare and Update (if necessary) ---
	if targetRecord.Content == publicIP {
		log.Printf("DNS record IP (%s) matches current public IP (%s). No update needed.\n", targetRecord.Content, publicIP)
		log.Println("DDNS check complete.")
		return // Exit cleanly
	}

	log.Printf("DNS record IP (%s) differs from current public IP (%s). Updating...\n", targetRecord.Content, publicIP)

	// Construct the update parameters
	// Keep Proxied status and TTL the same, just update Content (the IP)
	updateParams := cloudflare.UpdateDNSRecordParams{
		ID:      targetRecord.ID,
		Type:    targetRecord.Type,
		Name:    targetRecord.Name,
		Content: publicIP,
		TTL:     targetRecord.TTL,     // Keep original TTL (1 = auto)
		Proxied: targetRecord.Proxied, // Keep original proxied status
	}

	// Perform the update
	_, err = api.UpdateDNSRecord(ctx, cloudflare.ZoneIdentifier(zoneID), updateParams)
	if err != nil {
		log.Fatalf("Error updating DNS record '%s' (ID: %s): %v\n", targetRecord.Name, targetRecord.ID, err)
	}

	log.Printf("Successfully updated DNS record '%s' to IP %s\n", targetRecord.Name, publicIP)
	log.Println("DDNS update complete.")
}

// getCurrentPublicIP fetches the current public IPv4 address from an external service.
func getCurrentPublicIP(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", ipEchoService, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{Timeout: 10 * time.Second} // Ensure HTTP client respects timeout
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to get response from %s: %w", ipEchoService, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status code from %s: %d", ipEchoService, resp.StatusCode)
	}

	ipBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Trim whitespace, as some services might add a newline
	ip := string(bytes.TrimSpace(ipBytes))

	// Basic validation (optional, but good practice)
	if !isValidIPv4(ip) {
		return "", fmt.Errorf("received invalid IP address format: '%s'", ip)
	}

	return ip, nil
}

// isValidIPv4 checks if the string is a valid IPv4 address.
// Note: This is a simple check, net.ParseIP is more robust but returns net.IP
func isValidIPv4(ip string) bool {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return false
	}
	// Further checks could be added here (e.g., each part is 0-255)
	// For simplicity, we rely on the echo service providing a valid format.
	return true
}
