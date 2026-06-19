# 📝 Markdown to HTML Converter
A simple command-line tool written in Go that converts `.md` (Markdown) files into `.html` files you can open in any browser.
Built as a beginner Go learning project — no external dependencies, just the Go standard library.
---
## 🚀 Getting Started
### Prerequisites
Make sure you have Go installed:
```bash
go version
# should print something like: go version go1.21.0 darwin/arm64
```
If not, download it from [https://go.dev/dl](https://go.dev/dl)
### Clone / Download
```bash
# If using git:
git clone <your-repo-url>
cd markdown-to-html
# Or just navigate to the project folder:
cd path/to/markdown-to-html
```
---
## 📦 Project Structure
```
markdown-to-html/
├── main.go       ← the converter (all logic lives here)
├── input.md      ← your test markdown file
├── input.html    ← generated output (created after running)
├── go.mod        ← Go module file
└── README.md     ← you are here
```
---
## ▶️ How to Run
```bash
go run main.go <your-file.md>
```
### Example:
```bash
go run main.go input.md
```
Output:
```
✅ Converted: input.md → input.html
📂 Open input.html in your browser!
```
Then open `input.html` in your browser to see the result.
---
## ⚙️ Build as a Binary (Optional)
Instead of using `go run` every time, you can compile it into a standalone executable:
```bash
go build -o mdtohtml main.go
```
Then use it directly:
```bash
./mdtohtml input.md
./mdtohtml notes.md
./mdtohtml readme.md
```
---
## ✅ Supported Markdown Features
|
 Markdown Syntax 
|
 Output HTML 
|
|
----------------
|
-------------
|
|
`# Heading 1`
|
`<h1>Heading 1</h1>`
|
|
`## Heading 2`
|
`<h2>Heading 2</h2>`
|
|
`### Heading 3`
|
`<h3>Heading 3</h3>`
|
|
`**bold text**`
|
`<strong>bold text</strong>`
|
|
`*italic text*`
|
`<em>italic text</em>`
|
|
`- list item`
|
`<ul><li>list item</li></ul>`
|
|
`[text](https://url)`
|
`<a href="https://url">text</a>`
|
|
 Normal text 
|
`<p>Normal text</p>`
|
---
## 🧪 Example Input & Output
**input.md**
```markdown
# Hello World
## About This Tool
This converter is written in **Go** and supports *basic* markdown.
## Useful Links
- Visit [Go official site](https://go.dev)
- Check [Go by Example](https://gobyexample.com)
## Summary
A simple tool, built to **learn Go** step by step.
```
**input.html** (opened in browser)
```
Hello World              ← big heading
About This Tool          ← smaller heading
This converter is written in bold Go and supports italic basic markdown.
• Visit Go official site (clickable link)
• Check Go by Example   (clickable link)
```
---
## ❌ Error Handling
The tool gives clear messages for common mistakes:
```bash
# No file provided
$ go run main.go
❌ No file provided!
Usage:   go run main.go <file.md>
Example: go run main.go input.md
# Wrong file type
$ go run main.go notes.txt
❌ File must be a .md file, got: notes.txt
# File does not exist
$ go run main.go ghost.md
❌ Error: could not open file: open ghost.md: no such file or directory
```
---
## 🧠 What I Learned Building This
This project covers these Go concepts:
|
 Concept 
|
 Where Used 
|
|
---------
|
-----------
|
|
`package`
, 
`import`
, 
`func main()`
|
 Program entry point 
|
|
 Variables (
`string`
, 
`int`
, 
`bool`
) 
|
 Storing data 
|
|
 Structs 
|
 Grouping related data 
|
|
 Slices & 
`for range`
 loops 
|
 Processing lines 
|
|
 Functions & multiple return values 
|
`convert()`
, 
`convertLine()`
|
|
`os.Open`
, 
`bufio.Scanner`
|
 Reading file line by line 
|
|
`strings.HasPrefix / HasSuffix / TrimPrefix`
|
 Detecting markdown patterns 
|
|
`regexp.MustCompile`
, 
`ReplaceAllString`
|
 Bold, italic, link conversion 
|
|
 State variables 
|
 Tracking open 
`<ul>`
 lists 
|
|
`strings.Builder`
|
 Collecting HTML output 
|
|
`os.WriteFile`
, 
`[]byte()`
|
 Writing output file 
|
|
`os.Args`
|
 Reading CLI arguments 
|
|
 Error handling (
`if err != nil`
) 
|
 Graceful failures 
|
---
## 🚀 Ideas to Extend This Project
Once comfortable, try adding these features:
- [ ] `` `code` `` → `<code>` inline code blocks
- [ ] `---` → `<hr>` horizontal divider
- [ ] Numbered lists `1.` → `<ol><li>`
- [ ] Custom output filename with a flag (`--output result.html`)
- [ ] Convert an entire folder of `.md` files at once
- [ ] Dark mode toggle in the generated HTML
---
## 📄 License
Free to use and modify. Built for learning purposes.
