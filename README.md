# CAUTC - Content Analysis and URL Testing CLI Tool 🚀

[![Go Version](https://img.shields.io/github/go-mod/go-version/golang/go)](https://golang.org/)  
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](https://github.com/SaltyFrappuccino/CAUTC/actions)  
[![License](https://img.shields.io/badge/license-GNU%20GPLv3-blue)](https://www.gnu.org/licenses/gpl-3.0.html)
[![Contributions Welcome](https://img.shields.io/badge/contributions-welcome-orange.svg)](https://github.com/your-repo/CAUTC/issues)

CAUTC is a command-line tool designed for analyzing the content size and response time of URLs. It supports multiple content size units, processes files with lists of URLs, and optionally saves results to a file.

## 🌟 Features

- Content size analysis: Measure the size of content in bytes, kilobytes, megabytes, or characters.
- Concurrent processing: Fast and efficient processing of multiple URLs.
- Error handling: Gracefully handle invalid or unreachable URLs.
- File integration: Read URLs from a file and save results to a new file.
- Customizable output: Choose size units for processing (bytes, KB, MB, or characters).

---

## 📦 Installation

1. Clone the repository:  
   Run the following in your terminal:
   ```bash
   git clone https://github.com/SaltyFrappuccino/CAUTC.git
   cd CAUTC
   ```

2. Build the binary:
    ```bash 
   go build -o cautc.exe
    ```

---

## 🚀 Usage

Run the tool with the following options:  
```bash
./cautc.exe --path=<file-path> --size=<unit> --save=<true/false>
```

### Arguments

| Argument     | Required | Default   | Description                                                                 |
|--------------|----------|-----------|-----------------------------------------------------------------------------|
| `--path`     | Yes      |           | Path to the file containing URLs (relative or absolute).                   |
| `--size`     | No       | `bytes`   | Unit of content size measurement (`bytes`, `kb`, `mb`, or `chars`).        |
| `--save`     | No       | `false`   | Save results to a file in the same directory as the input file (`true`).   |

---

## 🛠️ Examples

### Basic Execution

To process a file containing URLs, run:  
```bash
./cautc.exe --path=urls.txt
```

### Custom Size Unit

To measure content size in kilobytes, use the following:  
```bash
./cautc.exe --path=urls.txt --size=kb
````

### Save Results to a File

To save results to a file while processing content size in characters:  
```bash
./cautc.exe --path=urls.txt --size=chars --save=true
```

---

## 📄 Output Example

### Console Output

Example output printed to the console:  

URL: https://example.com - Size: 150 KB - Time: 234ms  
URL: https://another-site.org - Error downloading content  
URL: https://valid-site.net - Size: 1.2 MB - Time: 1.3s

### File Output

When saving to results.txt, the output might look like this:  
https://example.com - 150 KB - 234ms  
https://another-site.org - Error downloading content - 0ms  
https://valid-site.net - 1.2 MB - 1.3s

---

## 📂 Project Structure

The project is organized as follows:

- main.go: Entry point for the application.
- site_processor.go: Core logic for processing and analyzing URLs.
- link_processor.go: Handles file reading and URL extraction.
- sites.txt: Example input file.

---

## 🚧 Roadmap

- Add support for HTTP headers analysis.
- Implement retries for failed URLs.
- Add JSON/CSV export options.
- Improve error messages and logging.

---

## 🤝 Contributing

We welcome contributions! Feel free to fork the project, open an issue, or submit a pull request. 
<br/>Please ensure your code adheres to the project's coding standards.

---

## 📜 License

This project is licensed under the **GNU GPLv3** License. See the LICENSE file for details.