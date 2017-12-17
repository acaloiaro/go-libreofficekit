package libreofficekit

import (
	"testing"
	"os"
	"time"
)

const (
	DefaultLibreOfficePath  = "/usr/lib/libreoffice/program/"
	DocumentThatDoesntExist = "testdata/kittens.docx"
	SampleDocument          = "testdata/sample.docx"
	SaveDocumentPath        = "/tmp/out.docx"
	SaveDocumentFormat      = "docx"
)

func TestInvalidOfficePath(t *testing.T) {
	_, err := NewOffice("/etc/passwd")
	if err == nil {
		t.Fail()
	}
}

func TestValidOfficePath(t *testing.T) {
	_, err := NewOffice(DefaultLibreOfficePath)
	if err != nil {
		t.Fail()
	}
}

func TestGetOfficeErrorMessage(t *testing.T) {
	office, _ := NewOffice(DefaultLibreOfficePath)
	office.LoadDocument(DocumentThatDoesntExist)
	message := office.GetError()
	if len(message) == 0 {
		t.Fail()
	}
}

func TestLoadDocumentThatDoesntExist(t *testing.T) {
	office, _ := NewOffice(DefaultLibreOfficePath)
	_, err := office.LoadDocument(DocumentThatDoesntExist)
	if err == nil {
		t.Fail()
	}
}

func TestSuccessLoadDocument(t *testing.T) {
	office, _ := NewOffice(DefaultLibreOfficePath)
	_, err := office.LoadDocument(SampleDocument)
	if err != nil {
		t.Fail()
	}
}

func TestSuccessLoadDocumentSafe(t *testing.T) {
	office, _ := NewOffice(DefaultLibreOfficePath)
	_, err := office.LoadDocumentSafe(SampleDocument, time.Duration(1 * time.Second))
	if err != nil {
		t.Fail()
	}
}

func TestSuccessLoadDocumentSafeFailure(t *testing.T) {
	office, _ := NewOffice(DefaultLibreOfficePath)
	_, err := office.LoadDocumentSafe(SampleDocument, time.Duration(100 * time.Millisecond))

	if err == nil {
		t.Fail()
	}
}

func TestSuccessfulLoadAndSaveDocument(t *testing.T) {
	office, _ := NewOffice(DefaultLibreOfficePath)
	doc, err := office.LoadDocumentSafe(SampleDocument, time.Duration(1 * time.Second))
	if err != nil {
		t.Fail()
	}

	err = doc.SaveAs(SaveDocumentPath, SaveDocumentFormat, "")
	if err != nil {
		t.Fail()
	}

	defer os.Remove(SaveDocumentPath)
}

func TestGetPartPageRectangles(t *testing.T) {
	office, _ := NewOffice(DefaultLibreOfficePath)
	document, _ := office.LoadDocumentSafe(SampleDocument, time.Duration(1 * time.Second))
	rectangles := document.GetPartPageRectangles()
	if len(rectangles) != 2 {
		t.Fail()
	}
}

func TestGetParts(t *testing.T) {
	office, _ := NewOffice(DefaultLibreOfficePath)
	document, _ := office.LoadDocumentSafe(SampleDocument, time.Duration(1 * time.Second))
	parts := document.GetParts()
	if parts != 2 {
		t.Fail()
	}
}

func TestGetTileMode(t *testing.T) {
	office, _ := NewOffice(DefaultLibreOfficePath)
	document, _ := office.LoadDocumentSafe(SampleDocument, time.Duration(1 * time.Second))
	mode := document.GetTileMode()
	if mode != RGBATilemode && mode != BGRATilemode {
		t.Fail()
	}
}

func TestGetType(t *testing.T) {
	office, _ := NewOffice(DefaultLibreOfficePath)
	document, _ := office.LoadDocumentSafe(SampleDocument, time.Duration(1 * time.Second))
	documentType := document.GetType()
	if documentType != TextDocument {
		t.Fail()
	}
}

func TestTextSelection(t *testing.T) {
	office, _ := NewOffice(DefaultLibreOfficePath)
	document, _ := office.LoadDocumentSafe(SampleDocument, time.Duration(1 * time.Second))
	rectangle := document.GetPartPageRectangles()[0]
	document.SetTextSelection(SetGraphicSelectionStart, rectangle.Min.X, rectangle.Min.Y)
	document.SetTextSelection(SetGraphicSelectionEnd, rectangle.Max.X, rectangle.Max.Y)
	plaintext := document.GetTextSelection("text/plain;charset=utf-8")
	if len(plaintext) < 1000 {
		t.Fail()
	}
	document.ResetTextSelection()
	plaintext = document.GetTextSelection("text/plain;charset=utf-8")
	if len(plaintext) != 0 {
		t.Fail()
	}
}
