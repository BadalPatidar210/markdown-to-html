package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	boldPattern   = regexp.MustCompile(`\*\*(.*?)\*\*`)
	italicPattern = regexp.MustCompile(`\*(.*?)\*`)
	linkPattern   = regexp.MustCompile(`\[([^\]]+)\]\((https?://[^\)]+)\)`)
)

const htmlHeader = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>My Document</title>
    <style>
        body {
            font-family: sans-serif;
            max-width: 800px;
            margin: 40px auto;
            padding: 0 20px;
            line-height: 1.6;
            color: #333;
        }
        h1, h2, h3 { color: #111; }
        a { color: #0070f3; }
        ul { padding-left: 20px; }
        li { margin: 4px 0; }
    </style>
</head>
<body>
`
const htmlFooter = `
</body>
</html>`

func main() {
	if len(os.Args) < 2 {
		fmt.Println("❌ No file provided!")
		fmt.Println()
		fmt.Println("Usage:   go run main.go <file.md>")
		fmt.Println("Example: go run main.go input.md")
		return
	}
	inputFile := os.Args[1]

	// Check it ends in .md
	if !strings.HasSuffix(inputFile, ".md") {
		fmt.Println("❌ File must be a .md file, got:", inputFile)
		fmt.Println("Example: go run main.go input.md")
		return
	}
	outputFile := strings.Replace(inputFile, ".md", ".html", 1)

	html, err := convert("input.md")
	if err != nil {
		fmt.Println("Error converting markdown:", err)
		return
	}

	err = os.WriteFile(outputFile, []byte(html), 0644)
	if err != nil {
		fmt.Println("Error writing HTML file:", err)
		return
	}
	fmt.Println("✅ Converted:", inputFile, "→", outputFile)
	fmt.Println("📂 Open", outputFile, "in your browser to see the result!")
}

func applyInline(line string) string {
	line = linkPattern.ReplaceAllString(line, `<a href="$2">$1</a>`)
	line = boldPattern.ReplaceAllString(line, "<strong>$1</strong>")
	line = italicPattern.ReplaceAllString(line, "<em>$1</em>")
	return line
}

func convertLine(line string) string {
	if strings.HasPrefix(line, "### ") {
		content := strings.TrimPrefix(line, "### ")
		return "<h3>" + applyInline(content) + "</h3>"
	}

	if strings.HasPrefix(line, "## ") {
		content := strings.TrimPrefix(line, "## ")
		return "<h2>" + applyInline(content) + "</h2>"
	}

	if strings.HasPrefix(line, "# ") {
		content := strings.TrimPrefix(line, "# ")
		return "<h1>" + applyInline(content) + "</h1>"
	}

	return applyInline(line)
}

func convert(inputFile string) (string, error) {
	file, err := os.Open("input.md")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", fmt.Errorf("could not open file: %w", err)
	}

	defer file.Close()

	var b strings.Builder
	b.WriteString(htmlHeader)

	scanner := bufio.NewScanner(file)
	insideList := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "- ") {
			if !insideList {
				// fmt.Println("<ul>")
				b.WriteString("<ul>\n")
				insideList = true
			}
			content := strings.TrimPrefix(line, "- ")
			b.WriteString("<li>" + applyInline(content) + "</li>\n")
			// fmt.Println("<li>" + applyInline(content) + "</li>")
		} else {
			if insideList {
				// fmt.Println("</ul>")
				b.WriteString("</ul>\n")
				insideList = false
			}

			result := convertLine(line)

			if result != "" && !strings.HasPrefix(result, "<h") {
				result = "<p>" + result + "</p>"
			}
			b.WriteString(result + "\n")
		}
	}

	if insideList {
		// fmt.Println("</ul>")
		b.WriteString("</ul>\n")
	}
	b.WriteString(htmlFooter)
	return b.String(), nil
}
