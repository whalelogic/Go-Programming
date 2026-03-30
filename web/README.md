# Markdown Web Server

This web server serves all markdown files in the repository in a clean, browsable format.

## Features

- 📄 Automatically discovers all `.md` files in the repository
- 🎨 Clean, responsive UI with syntax highlighting
- 📂 Organized by directory/category
- 🔍 Easy navigation between files
- 📱 Mobile-friendly design

## Usage

From the repository root, run:

```bash
go run web/markdown-server.go
```

Then open your browser to: **http://localhost:8080**

## Structure

- `markdown-server.go` - Main server application
- `templates/` - HTML templates
  - `index.html` - Homepage listing all markdown files
  - `view.html` - Individual markdown file viewer
- `static/` - Static assets
  - `style.css` - Stylesheet for the entire site

## Dependencies

- `github.com/gomarkdown/markdown` - Markdown to HTML conversion

The server uses only Go standard library packages (`net/http`, `html/template`) plus the markdown parser.
