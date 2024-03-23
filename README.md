# WebModeAnalyzer

WebModeAnalyzer is a Go-based project that analyzes a list of websites and determines whether they are in Dark Mode or Light Mode. The results are then stored in a JSON file.

## Requirements

- Go 1.16 or higher

## Installation

Clone this repository into your local Go workspace:

```bash
git clone https://github.com/yourusername/WebModeAnalyzer.git
```

Navigate into the project directory:

```bash
cd WebModeAnalyzer
```

## Usage

Run the main program:

```bash
go run main.go
```

The program reads a list of URLs from a file named `urls.json`, scrapes the HTML content of each page, tries to determine whether the page is in dark or light mode, and writes the results to a file named `output.json`. Each entry in `output.json` contains the URL of the page and the determined mode.
