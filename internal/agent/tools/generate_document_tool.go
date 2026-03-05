package tools

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/tealeg/xlsx"
	"google.golang.org/adk/tool"
	"google.golang.org/genai"
)

type GenerateDocumentsArgs struct {
	Content  string `json:"content" jsonschema:"The content to be converted into a document."`
	MimeType string `json:"mime_type" jsonschema:"The MIME type of the document to be generated. Enabled MIME types include 'application/pdf' for PDF documents, 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' for Excel files, 'text/csv' for CSV files, 'text/plain' for plain text files"`
	FileName string `json:"file_name" jsonschema:"The desired name of the generated document, including the appropriate file extension, e.g., 'report.pdf', 'notes.txt'."`
}
type GenerateDocumentResponse struct {
	StatusCode int     `json:"status_code"`
	Message    string  `json:"message"`
	FilePath   *string `json:"file_path,omitempty"`
}

func GenerateDocumentsToolFunc(tctx tool.Context, args GenerateDocumentsArgs) (GenerateDocumentResponse, error) {
	fmt.Printf("Tool: Generating document for content: '%s' with MIME type: '%s'\n", args.Content, args.MimeType)
	var data []byte
	var err error
	switch args.MimeType {
	case "application/pdf":
		data, err = pdfGenerator([]byte(args.Content))
		if err != nil {
			log.Printf("Error generating PDF: %v", err)
			return GenerateDocumentResponse{
				StatusCode: 500,
				Message:    "Failed to generate PDF document " + err.Error(),
			}, nil
		}
	case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":
		data, err = excelGenerator([]byte(args.Content))
		if err != nil {
			log.Printf("Error generating Excel: %v", err)
			return GenerateDocumentResponse{
				StatusCode: 500,
				Message:    "Failed to generate XLSX document " + err.Error(),
			}, nil
		}
	case "text/csv":
		data, err = excelGenerator([]byte(args.Content))
		if err != nil {
			log.Printf("Error generating Excel: %v", err)
			return GenerateDocumentResponse{
				StatusCode: 500,
				Message:    "Failed to generate CSV document " + err.Error(),
			}, nil
		}
	case "text/plain":
		data = []byte(args.Content)
	default:
		return GenerateDocumentResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("unsupported MIME type: %s", args.MimeType),
		}, nil
	}
	response, err := tctx.Artifacts().Save(
		tctx, args.FileName, &genai.Part{
			InlineData: &genai.Blob{
				MIMEType: args.MimeType,
				Data:     data,
			},
		},
	)
	if err != nil {
		log.Printf("Error generating document: %v", err)
		return GenerateDocumentResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("Error generating document: %v", err),
		}, nil
	}
	version := response.Version
	filePath := fmt.Sprintf("%s/%s/%s/%s/%d", tctx.AppName(), tctx.UserID(), tctx.SessionID(), args.FileName, version)
	return GenerateDocumentResponse{
		StatusCode: 200,
		FilePath:   &filePath,
		Message:    "Document generated with success",
	}, nil
}

func excelGenerator(data []byte) ([]byte, error) {
	content := string(data)
	lines := strings.Split(content, "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("no content to generate Excel")
	}

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return nil, fmt.Errorf("error creating Excel sheet: %v", err)
	}

	for _, line := range lines {
		row := sheet.AddRow()
		cells := strings.Split(line, ",") // Assuming CSV format for simplicity
		for _, cellValue := range cells {
			cell := row.AddCell()
			cell.Value = strings.TrimSpace(cellValue)
		}
	}

	var buf bytes.Buffer
	err = file.Write(&buf)
	if err != nil {
		return nil, fmt.Errorf("error generating Excel: %v", err)
	}
	return buf.Bytes(), nil
}

func pdfGenerator(data []byte) ([]byte, error) {
	content := string(data)

	pdf := gofpdf.New("P", "mm", "A4", "")
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.AddPage()
	pdf.SetMargins(15, 15, 15)

	lines := strings.Split(content, "\n") //OR?
	fmt.Printf("Cantidad de líneas a procesar: %d\n", len(lines))
	for _, line := range lines {
		line = strings.TrimRight(line, " ")

		switch {
		case strings.HasPrefix(line, "# "):
			// H1
			pdf.SetFont("Arial", "B", 20)
			pdf.SetTextColor(30, 30, 30)
			pdf.MultiCell(0, 10, tr(strings.TrimPrefix(line, "# ")), "", "L", false)
			pdf.Ln(3)

		case strings.HasPrefix(line, "## "):
			// H2
			pdf.SetFont("Arial", "B", 15)
			pdf.SetTextColor(50, 50, 50)
			pdf.MultiCell(0, 9, tr(strings.TrimPrefix(line, "## ")), "", "L", false)
			pdf.Ln(2)

		case strings.HasPrefix(line, "### "):
			// H3
			pdf.SetFont("Arial", "B", 12)
			pdf.SetTextColor(70, 70, 70)
			pdf.MultiCell(0, 8, tr(strings.TrimPrefix(line, "### ")), "", "L", false)
			pdf.Ln(1)

		case strings.HasPrefix(line, "* ") || strings.HasPrefix(line, "- "):
			// Bullet point
			pdf.SetFont("Arial", "", 11)
			pdf.SetTextColor(0, 0, 0)
			text := strings.TrimPrefix(strings.TrimPrefix(line, "* "), "- ")
			text = cleanInlineMarkdown(text)
			pdf.MultiCell(0, 7, tr("  • "+text), "", "L", false)

		case strings.HasPrefix(line, "---"):
			pdf.Ln(2)
			pdf.SetDrawColor(180, 180, 180)
			x, y := pdf.GetX(), pdf.GetY()
			pdf.Line(x, y, x+180, y)
			pdf.Ln(4)

		case line == "":
			pdf.Ln(4)

		default:
			// Párrafo normal
			pdf.SetFont("Arial", "", 11)
			pdf.SetTextColor(0, 0, 0)
			pdf.MultiCell(0, 7, tr(cleanInlineMarkdown(line)), "", "L", false)
		}
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, fmt.Errorf("error generating PDF: %v", err)
	}
	return buf.Bytes(), nil
}

func cleanInlineMarkdown(text string) string {
	for strings.Contains(text, "**") {
		start := strings.Index(text, "**")
		end := strings.Index(text[start+2:], "**")
		if end == -1 {
			break
		}
		inner := text[start+2 : start+2+end]
		text = text[:start] + inner + text[start+2+end+2:]
	}
	for strings.Contains(text, "*") {
		start := strings.Index(text, "*")
		end := strings.Index(text[start+1:], "*")
		if end == -1 {
			break
		}
		inner := text[start+1 : start+1+end]
		text = text[:start] + inner + text[start+1+end+1:]
	}
	return text
}
