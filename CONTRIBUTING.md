# Contributing to Go Cloudflare DDNS Client

First off, thank you for considering contributing! We welcome any help, whether it's reporting a bug, proposing a feature, improving documentation, or writing code.

Please take a moment to review this document to ensure a smooth and effective contribution process for everyone involved.

## Code of Conduct

This project and everyone participating in it is governed by our [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. Please report unacceptable behavior to the project maintainers. *(Note: You will need to create a CODE_OF_CONDUCT.md file, often using a template like the Contributor Covenant)*.

## How Can I Contribute?

There are many ways to contribute:

*   **Reporting Bugs:** If you find a bug, please report it by opening an issue.
*   **Suggesting Enhancements:** Have an idea for a new feature or an improvement to an existing one? Open an issue to discuss it.
*   **Writing Code:** Submit Pull Requests with bug fixes or new features.
*   **Improving Documentation:** Found a typo or think something could be clearer in the README or code comments? Submit a Pull Request.
*   **Answering Questions:** Help other users by answering questions in the Issues section.

## Reporting Bugs

Before submitting a bug report, please:

1.  **Check the [Issues](https://github.com/YOUR_USERNAME/cf-ddns-client/issues)** to see if the bug has already been reported.
2.  **Ensure you are using the latest version** of the client.
3.  **Try to reproduce the issue** reliably.

If the bug hasn't been reported, please [open a new issue](https://github.com/YOUR_USERNAME/cf-ddns-client/issues/new) and provide the following information:

*   **A clear and descriptive title.**
*   **Steps to reproduce the bug.**
*   **What you expected to happen.**
*   **What actually happened.** Include relevant logs or error messages (please anonymize sensitive information like API tokens).
*   **Your environment:** Operating System, Go version (if building from source), client version (if using a release).

## Suggesting Enhancements

1.  **Check the [Issues](https://github.com/YOUR_USERNAME/cf-ddns-client/issues)** to see if your idea has already been suggested or discussed.
2.  **Open a new issue** to propose your enhancement.
3.  **Explain the motivation:** Why is this enhancement needed? What problem does it solve?
4.  **Describe the proposed solution:** How would it work? Provide details and examples if possible.

## Pull Request Process

Ready to contribute code or documentation? Great!

1.  **Fork the repository** to your own GitHub account.
2.  **Clone your fork** locally: `git clone https://github.com/YOUR_GITHUB_USERNAME/cf-ddns-client.git`
3.  **Create a new branch** for your changes: `git checkout -b feature/your-feature-name` or `fix/short-bug-description`. Use a descriptive name.
4.  **Make your changes:** Write your code or update documentation.
    *   Follow Go standard coding conventions. Run `gofmt` or `goimports` on your code.
    *   Add comments to explain complex logic.
    *   Ensure your changes work correctly.
    *   **(Optional but Recommended)** Add tests for any new functionality or bug fixes. Run tests using `go test ./...`.
5.  **Update Documentation:** If your changes affect usage, configuration, or behavior, update the `README.md` or other relevant documentation.
6.  **Commit your changes** with a clear and concise commit message. Follow conventional commit messages if possible (e.g., `feat: Add IPv6 support`, `fix: Correctly handle API rate limits`, `docs: Update configuration instructions`).
    ```bash
    git add .
    git commit -m "feat: Describe your change here"
    ```
7.  **Push your changes** to your fork: `git push origin feature/your-feature-name`
8.  **Open a Pull Request (PR)** against the `main` branch of the original `YOUR_USERNAME/cf-ddns-client` repository.
    *   Provide a clear title and description for your PR.
    *   Explain the 'what' and 'why' of your changes.
    *   Link to any relevant issues (e.g., `Closes #123`).
9.  **Code Review:** Project maintainers will review your PR. Be prepared to discuss your changes and make adjustments based on feedback.
10. **Merge:** Once approved, your PR will be merged. Thank you for your contribution!

## Development Setup

To set up the project for local development:

1.  **Install Go:** Ensure you have Go (version 1.18 or later) installed.
2.  **Clone the repository** (as described in the Pull Request Process section).
3.  **Navigate to the project directory:** `cd cf-ddns-client`
4.  **Install dependencies:** `go mod download` (or `go mod tidy`)
5.  **Build the project:** `go build -o cf-ddns-client .`
6.  **Run tests:** `go test ./...`

## Style Guides

*   **Go Code:** Follow the standard Go formatting guidelines. Use `gofmt` or `goimports` to format your code before committing. Adhere to effective Go principles.
*   **Git Commit Messages:** Write clear and concise commit messages. Consider using [Conventional Commits](https://www.conventionalcommits.org/).
*   **Documentation:** Use Markdown for documentation files.

## Questions?

If you have questions about contributing, feel free to open an issue and label it as a `question`.

Thank you for helping make this project better!
