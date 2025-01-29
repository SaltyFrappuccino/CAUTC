# CAUTC - Content Analysis and URL Testing CLI Tool 🚀

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/SaltyFrappuccino/CAUTC) <br/>
[![License](https://img.shields.io/badge/license-GNU%20GPLv3-blue)](https://www.gnu.org/licenses/gpl-3.0.html) <br/>
![GitHub repo size](https://img.shields.io/github/repo-size/SaltyFrappuccino/CAUTC) <br/>
![GitHub last commit](https://img.shields.io/github/last-commit/SaltyFrappuccino/CAUTC) <br/>
[![Contributions Welcome](https://img.shields.io/badge/contributions-welcome-orange.svg)](https://github.com/your-repo/CAUTC/issues) <br/>

CAUTC is a command-line tool designed for analyzing the content size and response time of URLs. It supports multiple content size units, processes files with lists of URLs, and optionally saves results to a file in various formats.

## 🌟 Features

- Content size analysis: Measure the size of content in bytes, kilobytes, megabytes, or characters.
- Concurrent processing: Fast and efficient processing of multiple URLs.
- Error handling: Gracefully handle invalid or unreachable URLs.
- File integration: Read URLs from a file and save results to a new file.
- Customizable output: Choose size units for processing (bytes, KB, MB, or characters).
- Depth control: Specify the depth of content downloading for nested resources.
- Export options: Save results in TXT, JSON, or CSV formats.

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
./cautc.exe --path=<file-path> --size=<unit> --save=<true/false> --depth=<depth> --export=<format>
```

### Arguments

| Argument     | Required | Default   | Description                                                                 |
|--------------|----------|-----------|-----------------------------------------------------------------------------|
| `--path`     | Yes      |           | Path to the file containing URLs (relative or absolute).                   |
| `--size`     | No       | `bytes`   | Unit of content size measurement (`bytes`, `kb`, `mb`, or `chars`).        |
| `--save`     | No       | `false`   | Save results to a file in the same directory as the input file (`true`).   |
| `--depth`    | No       | `1`       | Depth of content downloading for nested resources (default: 1).            |
| `--export`   | No       | `txt`     | Export type for results (`txt`, `json`, or `csv`).                         |

---

## 🛠️ Examples

### Basic Execution

To process a file containing URLs, run:
```bash
./cautc.exe --path=sites.txt
```

### Custom Size Unit

To measure content size in kilobytes, use the following:
```bash
./cautc.exe --path=sites.txt --size=kb
```

### Save Results to a File

To save results to a file while processing content size in characters:
```bash
./cautc.exe --path=sites.txt --size=chars --save=true
```

### Specify Download Depth

To download content with a depth of 2:
```bash
./cautc.exe --path=sites.txt --depth=2
```

### Export Results in JSON Format

To save results in JSON format:
```bash
./cautc.exe --path=sites.txt --export=json --save=true
```

---

## 📄 Output Example

### Console Output

Example output printed to the console:
```
URL: https://example.com - Size: 150 KB - Time: 234ms  
URL: https://another-site.org - Error downloading content  
URL: https://valid-site.net - Size: 1.2 MB - Time: 1.3s
```

### File Output (TXT)

When saving to results.txt, the output might look like this:
```
https://example.com - 150 KB - 234ms  
https://another-site.org - Error downloading content - 0ms  
https://valid-site.net - 1.2 MB - 1.3s
```

### File Output (JSON)

When saving to results.json, the output might look like this:
```json
[
   {
      "URL": "https://google.com",
      "Size": 20,
      "Duration": 348120600
   },
   {
      "URL": "https://youtube.com",
      "Size": 562,
      "Duration": 386938600
   },
   {
      "URL": "https://facebook.com",
      "Size": -1,
      "Duration": 10011235000
   }
]
```

### File Output (CSV)

When saving to results.csv, the output might look like this:
```csv
url,size,time  
https://example.com,150 KB,234ms  
https://another-site.org,Error downloading content,0ms  
https://valid-site.net,1.2 MB,1.3s
```

---

## 📂 Project Structure

The project is organized as follows:

- `main.go`: Entry point for the application.
- `site_processor.go`: Core logic for processing and analyzing URLs.
- `link_processor.go`: Handles file reading and URL extraction.
- `sites.txt`: Example input file.

---

## 🚧 Roadmap

- ~~Add JSON/CSV export options.~~
- ~~Improve error messages and logging.~~
- Add support for HTTP headers analysis.
- Implement retries for failed URLs.

---

## 🤝 Contributing

We welcome contributions! Feel free to fork the project, open an issue, or submit a pull request.

---

## 📜 License

This project is licensed under the **GNU GPLv3** License. See the LICENSE file for details.