package mage

import (
	"encoding/json"
	"google.golang.org/appengine"
	"google.golang.org/appengine/blobstore"
	"html/template"
	"io"
	"net/http"
)

// Renders a GO HTML template
type TemplateRenderer struct {
	Template     *template.Template
	TemplateName string
	Data         interface{}
}

func (renderer *TemplateRenderer) Render(w http.ResponseWriter) error {
	return renderer.Template.ExecuteTemplate(w, renderer.TemplateName, renderer.Data)
}

// Returns the data as JSON object(s)
type JSONRenderer struct {
	Data interface{}
}

func (renderer *JSONRenderer) Render(w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(renderer.Data)
}

// Renders plain text
type TextRenderer struct {
	Data string
}

func (renderer *TextRenderer) Render(w http.ResponseWriter) error {
	_, err := io.WriteString(w, renderer.Data)
	return err
}

// Renders a file as returned from the BlobStore
type BlobRenderer struct {
	Data appengine.BlobKey
}

func (renderer *BlobRenderer) Render(w http.ResponseWriter) error {
	blobstore.Send(w, renderer.Data)
	return nil
}

type ErrorRenderer struct {
	Data error
}

func (renderer *ErrorRenderer) Render(w http.ResponseWriter) error {
	_, err := io.WriteString(w, renderer.Data.Error())
	return err
}
