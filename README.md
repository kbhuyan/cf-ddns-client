# Cloudflare Dynamic DNS Client (Go)

[![Go Report Card](https://goreportcard.com/badge/github.com/kbhuyan/cf-ddns-client)](https://goreportcard.com/report/github.com/kbhuyan/cf-ddns-client)
[![Build Status](https://github.com/kbhuyan/cf-ddns-client/actions/workflows/go.yml/badge.svg)](https://github.com/kbhuyan/cf-ddns-client/actions/workflows/go.yml) <!-- Optional: Setup GitHub Actions for Go -->
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A simple, efficient Dynamic DNS (DDNS) client written in Go. It automatically updates a specified Cloudflare DNS record (Type `A`) with the machine's current public IPv4 address.

## Motivation

Many Internet Service Providers (ISPs) assign dynamic public IP addresses to residential and some business connections. This makes it challenging to reliably connect to services hosted on these networks (e.g., home servers, NAS, cameras).

This DDNS client solves this problem by periodically checking the machine's public IP address and updating a corresponding DNS record in Cloudflare. This ensures that your domain or subdomain always points to your network's current IP address.

## Features

*   **Public IP Detection:** Automatically determines the current public IPv4 address using an external service (`icanhazip.com`).
*   **Cloudflare Integration:** Uses the official Cloudflare Go SDK to interact with the API securely via API Tokens.
*   **Targeted Updates:** Updates a specific `A` record within a specified Cloudflare zone.
*   **Idempotent:** Only performs an update if the detected public IP differs from the one currently set in the DNS record, minimizing unnecessary API calls.
*   **Preserves Settings:** Keeps the existing Time-To-Live (TTL) and Proxied (Orange Cloud) status of the DNS record during updates.
*   **Configuration:** Easily configured using environment variables.
*   **Cross-Platform:** Compiles and runs on Linux, macOS, and Windows.
*   **Lightweight:** Minimal dependencies and resource footprint.
*   **Docker Support:** Ready to be containerized using Docker.

## Prerequisites

Before you begin, ensure you have the following:

1.  **Go:** Version 1.23.4 or later installed (required for building from source). ([Download Go](https://golang.org/dl/))
2.  **Docker:** Required if you plan to use the Docker instructions. ([Install Docker](https://docs.docker.com/get-docker/))
3.  **Cloudflare Account:** An active Cloudflare account.
4.  **Domain Managed by Cloudflare:** The domain for which you want to update DNS records must be managed by Cloudflare.
5.  **Cloudflare API Token:**
    *   Generate a token with **Edit zone DNS** permissions.
    *   Go to Cloudflare Dashboard -> My Profile -> API Tokens -> Create Token.
    *   Use the "Edit zone DNS" template.
    *   Permissions needed: `Zone:Zone:Read`, `Zone:DNS:Edit`.
    *   Zone Resources: Select the specific zone(s) this token should manage (e.g., `Include -> Specific zone -> yourdomain.com`).
    *   **Important:** Securely store the generated token. You will only see it once.
6.  **Existing DNS Record:** An existing `A` record in your Cloudflare zone that you wish to keep updated (e.g., `subdomain.yourdomain.com` or `yourdomain.com`). This client currently *updates* existing records, it does not *create* them if missing.

## Installation

### Option 1: From Source (Recommended for Development)

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/kbhuyan/cf-ddns-client.git
    cd cf-ddns-client
    ```
2.  **Build the executable:**
    ```bash
    go build -o cf-ddns-client .
    ```
    This will create the `cf-ddns-client` (or `cf-ddns-client.exe` on Windows) executable in the current directory.

### Option 2: From Releases (Recommended for Users)

Check the [Releases](https://github.com/kbhuyan/cf-ddns-client/releases) page for pre-compiled binaries for your operating system and architecture. Download the appropriate binary and place it in a convenient location (e.g., `/usr/local/bin` on Linux/macOS).

### Option 3: Using Docker (See Docker Support section below)

## Configuration

The client is configured using environment variables:

*   `CF_API_TOKEN` (Required): Your Cloudflare API Token (see Prerequisites).
    *   **Security:** Treat this token like a password. Do not commit it directly into your code or share it publicly.
*   `CF_ZONE_NAME` (Required): The name of the Cloudflare zone (your domain) that contains the DNS record you want to update.
    *   Example: `example.com`
*   `CF_RECORD_NAME` (Required): The full DNS name of the `A` record to update.
    *   Example (Subdomain): `home.example.com`
    *   Example (Root Domain): `example.com`

## Usage (Direct Execution)

1.  **Set Environment Variables:**
    *   **Linux/macOS:**
        ```bash
        export CF_API_TOKEN="YOUR_CLOUDFLARE_API_TOKEN"
        export CF_ZONE_NAME="yourdomain.com"
        export CF_RECORD_NAME="subdomain.yourdomain.com"
        ```
    *   **Windows (Command Prompt):**
        ```cmd
        set CF_API_TOKEN=YOUR_CLOUDFLARE_API_TOKEN
        set CF_ZONE_NAME=yourdomain.com
        set CF_RECORD_NAME=subdomain.yourdomain.com
        ```
    *   **Windows (PowerShell):**
        ```powershell
        $env:CF_API_TOKEN="YOUR_CLOUDFLARE_API_TOKEN"
        $env:CF_ZONE_NAME="yourdomain.com"
        $env:CF_RECORD_NAME="subdomain.yourdomain.com"
        ```
    *(Replace placeholder values with your actual configuration)*

2.  **Run the client:**
    *   If built from source: `./cf-ddns-client`
    *   If downloaded from releases: `/path/to/cf-ddns-client`

The client will output logs to standard output indicating the steps taken (fetching IP, checking DNS, updating if necessary).

## Docker Support

You can easily build and run this client as a Docker container. This is useful for isolating the application and its dependencies, and for deploying in containerized environments.

1.  **Build the Docker Image:**
    Navigate to the project root directory (where the `Dockerfile` is located) and run:
    ```bash
    docker build -t your-dockerhub-username/cf-ddns-client:latest .
    # Or simply:
    # docker build -t cf-ddns-client .
    ```
    *(Replace `your-dockerhub-username` if you plan to push the image to Docker Hub)*

2.  **Run the Docker Container:**
    Use `docker run` to execute the client within a container. Pass the required configuration as environment variables using the `-e` flag:
    ```bash
    docker run --rm \
      -e CF_API_TOKEN="YOUR_CLOUDFLARE_API_TOKEN" \
      -e CF_ZONE_NAME="yourdomain.com" \
      -e CF_RECORD_NAME="subdomain.yourdomain.com" \
      cf-ddns-client
    ```
    *   `--rm`: Automatically removes the container when it exits.
    *   Replace placeholder values with your actual configuration.

    For persistent storage of logs (optional), you could mount a volume:
    ```bash
    docker run --rm \
      -e CF_API_TOKEN="YOUR_CLOUDFLARE_API_TOKEN" \
      -e CF_ZONE_NAME="yourdomain.com" \
      -e CF_RECORD_NAME="subdomain.yourdomain.com" \
      -v /path/on/host/logs:/app/logs \
      cf-ddns-client > /path/on/host/logs/cf-ddns-client.log 2>&1
    ```
    *(Note: The current application logs to stdout/stderr. Redirecting the container's output (`> ... 2>&1`) is a common way to capture logs when running manually or via simple schedulers like cron)*.

## Scheduling Execution

For the DDNS client to be effective, it needs to run periodically (e.g., every 5-15 minutes). Choose the method that best suits your environment:

### cron (Linux/macOS)

Edit your crontab: `crontab -e`

*   **For Direct Execution:**
    ```crontab
    */10 * * * * export CF_API_TOKEN='YOUR_TOKEN'; export CF_ZONE_NAME='yourdomain.com'; export CF_RECORD_NAME='subdomain.yourdomain.com'; /path/to/your/cf-ddns-client >> /var/log/cf-ddns-client.log 2>&1
    ```
*   **For Docker Execution:**
    ```crontab
    */10 * * * * /usr/bin/docker run --rm -e CF_API_TOKEN='YOUR_TOKEN' -e CF_ZONE_NAME='yourdomain.com' -e CF_RECORD_NAME='subdomain.yourdomain.com' cf-ddns-client >> /var/log/cf-ddns-client-docker.log 2>&1
    ```
    *   **Note:** Ensure the environment variables are correctly defined or passed. Make sure the `docker` command is in the PATH for the cron user, or use the full path (e.g., `/usr/bin/docker`).

### systemd (Linux)

Create a systemd service unit file (e.g., `/etc/systemd/system/cf-ddns.service`) and a timer unit file (`/etc/systemd/system/cf-ddns.timer`). This provides more robust service management, especially for Docker containers. You can define environment variables directly within the service file using `Environment=` or `EnvironmentFile=`.

### Task Scheduler (Windows)

Use the built-in Task Scheduler application to create a task that runs the `cf-ddns-client.exe` executable or a `docker run` command at your desired interval. You can configure the task to load the necessary environment variables or pass them in the command arguments if running Docker.

## Acknowledgements

* Uses the cloudflare-go library for interacting with the Cloudflare API.
* Relies on the public IP detection service provided by icanhazip.com.

