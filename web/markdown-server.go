package main

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type MarkdownFile struct {
	Name     string
	Path     string
	Category string
}

type PageData struct {
	Title   string
	Content template.HTML
	Files   []MarkdownFile
}

var rootDir string
var templates *template.Template

func mustGlob(pattern string) []string {
	matches, _ := filepath.Glob(pattern)
	return matches
}

func init() {
	var err error
	rootDir, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	templatePath := "templates/*.html"
	if _, err := filepath.Glob(templatePath); err != nil || len(mustGlob(templatePath)) == 0 {
		templatePath = "web/templates/*.html"
	}
	templates = template.Must(template.New("").ParseGlob(templatePath))

	staticPath := "static"
	if _, err := os.Stat(staticPath); os.IsNotExist(err) {
		staticPath = "web/static"
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/view/", viewHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticPath))))

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	files := getAllMarkdownFiles()
	data := PageData{
		Title: "Go Programming - Markdown Documentation",
		Files: files,
	}
	templates.ExecuteTemplate(w, "index.html", data)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/view/")
	if path == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fullPath := filepath.Join(rootDir, path)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	htmlContent := markdownToHTML(content)
	files := getAllMarkdownFiles()

	data := PageData{
		Title:   filepath.Base(path),
		Content: template.HTML(htmlContent),
		Files:   files,
	}

	templates.ExecuteTemplate(w, "view.html", data)
}

func markdownToHTML(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func getAllMarkdownFiles() []MarkdownFile {
	var files []MarkdownFile
	
	filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		
		if d.IsDir() {
			name := d.Name()
			if name == ".git" || name == "bin" || name == "out" || name == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}
		
		if strings.HasSuffix(d.Name(), ".md") {
			relPath, _ := filepath.Rel(rootDir, path)
			category := getCategoryFromPath(relPath)
			
			files = append(files, MarkdownFile{
				Name:     d.Name(),
				Path:     relPath,
				Category: category,
			})
		}
		
		return nil
	})
	
	return files
}

func getCategoryFromPath(path string) string {
	dir := filepath.Dir(path)
	if dir == "." {
		return "Root"
	}
	parts := strings.Split(dir, string(filepath.Separator))
	return strings.Title(parts[0])
}
