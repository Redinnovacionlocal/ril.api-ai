package tools

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/jung-kurt/gofpdf"
	"github.com/tealeg/xlsx"
	"google.golang.org/adk/tool"
	"google.golang.org/genai"
)

type GenerateDocumentsArgs struct {
	Blocks   []Block `json:"blocks" jsonschema:"Array of structured content blocks. Each block MUST have a 'type' field. 'text' is required for types: h1, h2, h3, paragraph, bullet. 'text' is NOT required for types: table, divider. For 'table' type use 'headers' (string array) and 'rows' (array of string arrays) instead of 'text'. NEVER use markdown syntax inside text fields."`
	MimeType string  `json:"mime_type" jsonschema:"The MIME type of the document to be generated. Enabled MIME types include 'application/pdf' for PDF documents, 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' for Excel files, 'text/csv' for CSV files, 'text/plain' for plain text files"`
	FileName string  `json:"file_name" jsonschema:"The desired name of the generated document, including the appropriate file extension, e.g., 'report.pdf', 'notes.txt'."`
}

type Block struct {
	Type    string     `json:"type"`
	Text    string     `json:"text,omitempty"`
	Headers []string   `json:"headers,omitempty"`
	Rows    [][]string `json:"rows,omitempty"`
}
type GenerateDocumentResponse struct {
	StatusCode int     `json:"status_code"`
	Message    string  `json:"message"`
	FilePath   *string `json:"file_path,omitempty"`
	FileCdn    *string `json:"file_cdn,omitempty"`
}

type Render interface {
	Heading(level int, text string)
	Paragraph(text string)
	Bullet(text string)
	Divider()
	Table(headers []string, rows [][]string)
	Bytes() ([]byte, error)
}

type PdfRenderer struct {
	pdf *gofpdf.Fpdf
	tr  func(string) string
}

func NewPdfRenderer() *PdfRenderer {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(15, 15, 15)
	pdf.AddPage() // ← esto faltaba
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	return &PdfRenderer{pdf: pdf, tr: tr}
}

func (r *PdfRenderer) Bytes() ([]byte, error) {
	var buf bytes.Buffer
	err := r.pdf.Output(&buf)
	return buf.Bytes(), err
}

func (r *PdfRenderer) Heading(level int, text string) {
	switch level {
	case 1:
		r.pdf.SetFont("Arial", "B", 20)
		r.pdf.SetTextColor(30, 30, 30)
		r.pdf.MultiCell(0, 10, r.tr(text), "", "L", false)
		r.pdf.Ln(3)
	case 2:
		r.pdf.SetFont("Arial", "B", 15)
		r.pdf.SetTextColor(50, 50, 50)
		r.pdf.MultiCell(0, 9, r.tr(text), "", "L", false)
		r.pdf.Ln(2)
	case 3:
		r.pdf.SetFont("Arial", "B", 12)
		r.pdf.SetTextColor(70, 70, 70)
		r.pdf.MultiCell(0, 8, r.tr(text), "", "L", false)
		r.pdf.Ln(1)
	default:
		r.pdf.SetFont("Arial", "B", 15)
		r.pdf.SetTextColor(50, 50, 50)
		r.pdf.MultiCell(0, 9, r.tr(text), "", "L", false)
		r.pdf.Ln(2)
	}
}

func (r *PdfRenderer) Paragraph(text string) {
	r.pdf.SetFont("Arial", "", 11)
	r.pdf.SetTextColor(0, 0, 0)
	r.pdf.MultiCell(0, 7, r.tr(text), "", "L", false)
}

func (r *PdfRenderer) Bullet(text string) {
	r.pdf.SetFont("Arial", "", 11)
	r.pdf.SetTextColor(0, 0, 0)
	r.pdf.MultiCell(0, 7, r.tr("."+
		text), "", "L", false)
}

func (r *PdfRenderer) Divider() {
	r.pdf.Ln(2)
	r.pdf.SetDrawColor(180, 180, 180)
	x, y := r.pdf.GetX(), r.pdf.GetY()
	r.pdf.Line(x, y, x+180, y)
	r.pdf.Ln(4)
}

func (r *PdfRenderer) Table(headers []string, rows [][]string) {
	if len(headers) == 0 {
		return
	}

	pageWidth, _ := r.pdf.GetPageSize()
	margins, _, _, _ := r.pdf.GetMargins()
	tableWidth := pageWidth - 2*margins
	colWidth := tableWidth / float64(len(headers))

	// Headers
	r.pdf.SetFont("Arial", "B", 10)
	r.pdf.SetFillColor(79, 70, 229) // indigo
	r.pdf.SetTextColor(255, 255, 255)
	for _, h := range headers {
		r.pdf.CellFormat(colWidth, 8, r.tr(h), "1", 0, "C", true, 0, "")
	}
	r.pdf.Ln(-1)

	// Rows
	r.pdf.SetFont("Arial", "", 9)
	r.pdf.SetTextColor(0, 0, 0)
	for i, row := range rows {
		// Filas alternadas
		if i%2 == 0 {
			r.pdf.SetFillColor(245, 245, 255)
		} else {
			r.pdf.SetFillColor(255, 255, 255)
		}
		for j, cell := range row {
			if j < len(headers) {
				r.pdf.CellFormat(colWidth, 7, r.tr(cell), "1", 0, "L", true, 0, "")
			}
		}
		r.pdf.Ln(-1)
	}

	r.pdf.Ln(4)
}

type ExcelRenderer struct {
	file        *xlsx.File
	sheet       *xlsx.Sheet
	headerStyle *xlsx.Style
}

func NewExcelRenderer() *ExcelRenderer {
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Sheet1")

	// Estilo para headers
	headerStyle := xlsx.NewStyle()
	headerStyle.Font.Bold = true
	headerStyle.Fill.FgColor = "FF4F46E5"
	headerStyle.Fill.PatternType = "solid"
	headerStyle.Font.Color = "FFFFFFFF"

	return &ExcelRenderer{file: file, sheet: sheet, headerStyle: headerStyle}
}

func (r *ExcelRenderer) Heading(level int, text string) {
	row := r.sheet.AddRow()
	cell := row.AddCell()
	cell.Value = text

	style := xlsx.NewStyle()
	style.Font.Bold = true
	style.Font.Size = 14 - level
	style.Fill.FgColor = "FF4F46E5"
	style.Fill.PatternType = "solid"
	style.Font.Color = "FFFFFFFF"
	cell.SetStyle(style)

	r.sheet.AddRow()
}

func (r *ExcelRenderer) Paragraph(text string) {
	row := r.sheet.AddRow()
	cell := row.AddCell()
	cell.Value = text
}

func (r *ExcelRenderer) Bullet(text string) {
	row := r.sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "• " + text
}

func (r *ExcelRenderer) Divider() {
	r.sheet.AddRow()
}

func (r *ExcelRenderer) Table(headers []string, rows [][]string) {
	// Headers con estilo
	headerRow := r.sheet.AddRow()
	for _, h := range headers {
		cell := headerRow.AddCell()
		cell.Value = h
		cell.SetStyle(r.headerStyle)
	}

	// Filas alternadas
	for i, row := range rows {
		excelRow := r.sheet.AddRow()
		for _, val := range row {
			cell := excelRow.AddCell()
			cell.Value = val

			// Filas alternadas para legibilidad
			if i%2 == 0 {
				style := xlsx.NewStyle()
				style.Fill.FgColor = "FFF5F5FF"
				style.Fill.PatternType = "solid"
				cell.SetStyle(style)
			}
		}
	}

	r.sheet.AddRow()
}

func (r *ExcelRenderer) Bytes() ([]byte, error) {
	var buf bytes.Buffer
	err := r.file.Write(&buf)
	return buf.Bytes(), err
}

func renderBlocks(blocks []Block, r Render) ([]byte, error) {
	for _, block := range blocks {
		switch block.Type {
		case "h1", "h2", "h3":
			level := int(block.Type[1] - '0')
			r.Heading(level, block.Text)
		case "paragraph":
			r.Paragraph(block.Text)
		case "bullet":
			r.Bullet(block.Text)
		case "divider":
			r.Divider()
		case "table":
			r.Table(block.Headers, block.Rows)
		}
	}
	return r.Bytes()
}

func GenerateDocumentsToolFunc(tctx tool.Context, args GenerateDocumentsArgs) (GenerateDocumentResponse, error) {
	renderers := map[string]func() Render{
		"application/pdf": func() Render {
			return NewPdfRenderer()
		},
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": func() Render {
			return NewExcelRenderer()
		},
		"text/csv": func() Render {
			return NewExcelRenderer()
		},
	}
	factory, ok := renderers[args.MimeType]
	if !ok {
		return GenerateDocumentResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("Unsupported MIME type: %s", args.MimeType),
		}, nil
	}
	data, err := renderBlocks(args.Blocks, factory())
	if err != nil {
		log.Printf("Error rendering document: %v", err)
		return GenerateDocumentResponse{
			StatusCode: 500,
			Message:    fmt.Sprintf("Error rendering document: %v", err),
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
	fileCdn := os.Getenv("ARTIFACT_BUCKET_URL")
	return GenerateDocumentResponse{
		StatusCode: 200,
		FilePath:   &filePath,
		FileCdn:    &fileCdn,
		Message:    "Document generated with success",
	}, nil
}
