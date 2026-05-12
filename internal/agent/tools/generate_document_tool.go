package tools

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/gomutex/godocx"
	"github.com/gomutex/godocx/docx"
	"github.com/gomutex/godocx/wml/stypes"
	"github.com/jung-kurt/gofpdf"
	"github.com/tealeg/xlsx"
	"google.golang.org/adk/artifact"
	"google.golang.org/adk/tool"
	"google.golang.org/genai"
)

type GenerateDocumentsArgs struct {
	Blocks   []Block `json:"blocks" jsonschema:"CRITICAL: Array of structured content blocks. You MUST decompose ALL content into typed blocks. FORBIDDEN in any text field: '#', '##', '**', '*', '-', '>', backticks, or ANY other markdown syntax. Violations will cause rendering errors. Block types and rules: (1) 'h1','h2','h3' = section headings, text field required, plain text only. (2) 'paragraph' = body text, text field required, plain text only. (3) 'bullet' = ONE bullet item per block, text field required, plain text only, do NOT use '-' or '*' prefix. (4) 'divider' = horizontal separator, no text field. (5) 'table' = structured data, requires headers array and rows array, no text field. CORRECT example: [{\"type\":\"h1\",\"text\":\"Annual Report\"},{\"type\":\"paragraph\",\"text\":\"This report covers Q1 results.\"},{\"type\":\"bullet\",\"text\":\"Revenue increased by 12 percent\"},{\"type\":\"table\",\"headers\":[\"Region\",\"Sales\"],\"rows\":[[\"North\",\"120k\"],[\"South\",\"98k\"]]}]. WRONG example: [{\"type\":\"paragraph\",\"text\":\"## Title\\n**bold** and - bullet\"}]"`
	MimeType string  `json:"mime_type" jsonschema:"MIME type of the document. Allowed values: 'application/pdf', 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet', 'text/csv', 'text/plain', 'application/vnd.openxmlformats-officedocument.wordprocessingml.document'"`
	FileName string  `json:"file_name" jsonschema:"File name with extension, e.g. 'report.pdf', 'data.xlsx'"`
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

	pageWidth, pageHeight := r.pdf.GetPageSize()
	margins, _, bottom, _ := r.pdf.GetMargins()
	tableWidth := pageWidth - 2*margins
	colWidth := tableWidth / float64(len(headers))
	lineHeight := 6.0
	headerLineHeight := 6.0
	usableHeight := pageHeight - bottom - margins

	calcLines := func(text string, width float64, bold bool) int {
		style := ""
		if bold {
			style = "B"
		}
		r.pdf.SetFont("Arial", style, 10)
		lines := r.pdf.SplitLines([]byte(r.tr(text)), width-2)
		if len(lines) == 0 {
			return 1
		}
		return len(lines)
	}

	maxHeaderLines := 1
	for _, h := range headers {
		n := calcLines(h, colWidth, true)
		if n > maxHeaderLines {
			maxHeaderLines = n
		}
	}
	headerHeight := float64(maxHeaderLines)*headerLineHeight + 2

	drawHeaders := func() {
		r.pdf.SetFont("Arial", "B", 10)
		r.pdf.SetFillColor(79, 70, 229)
		r.pdf.SetTextColor(255, 255, 255)

		startX, startY := r.pdf.GetX(), r.pdf.GetY()
		for j, h := range headers {
			x := startX + float64(j)*colWidth
			r.pdf.SetXY(x, startY)
			r.pdf.Rect(x, startY, colWidth, headerHeight, "F")
			r.pdf.Rect(x, startY, colWidth, headerHeight, "D")
			r.pdf.SetXY(x+1, startY+1)
			r.pdf.MultiCell(colWidth-2, headerLineHeight, r.tr(h), "", "C", false)
		}
		r.pdf.SetXY(startX, startY+headerHeight)
	}

	spaceLeft := usableHeight - r.pdf.GetY()
	if spaceLeft < headerHeight+lineHeight*2 {
		r.pdf.AddPage()
	}

	drawHeaders()

	// ── Filas ──
	r.pdf.SetFont("Arial", "", 9)
	r.pdf.SetTextColor(0, 0, 0)

	for i, row := range rows {
		maxLines := 1
		for j, cell := range row {
			if j < len(headers) {
				n := calcLines(cell, colWidth, false)
				if n > maxLines {
					maxLines = n
				}
			}
		}
		rowHeight := float64(maxLines) * lineHeight

		if r.pdf.GetY()+rowHeight > usableHeight {
			r.pdf.AddPage()
			drawHeaders()
			r.pdf.SetFont("Arial", "", 9)
			r.pdf.SetTextColor(0, 0, 0)
		}

		if i%2 == 0 {
			r.pdf.SetFillColor(245, 245, 255)
		} else {
			r.pdf.SetFillColor(255, 255, 255)
		}

		startX, startY := r.pdf.GetX(), r.pdf.GetY()
		for j, cell := range row {
			if j >= len(headers) {
				continue
			}
			x := startX + float64(j)*colWidth
			r.pdf.SetXY(x, startY)
			r.pdf.Rect(x, startY, colWidth, rowHeight, "F")
			r.pdf.Rect(x, startY, colWidth, rowHeight, "D")
			r.pdf.SetXY(x+1, startY+1)
			r.pdf.MultiCell(colWidth-2, lineHeight, r.tr(cell), "", "L", false)
		}

		r.pdf.SetXY(startX, startY+rowHeight)
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

type DocxRenderer struct {
	dox *docx.RootDoc
}

func NewDocxRenderer() *DocxRenderer {
	document, err := godocx.NewDocument()
	if err != nil {
		log.Fatalf("Error creating DOCX document: %v", err)
	}
	return &DocxRenderer{dox: document}
}

func (r *DocxRenderer) Heading(level int, text string) {
	switch level {
	case 1:
		_, err := r.dox.AddHeading(text, 0)
		if err != nil {
			log.Printf("Error adding heading: %v", err)
		}
	case 2:
		_, err := r.dox.AddHeading(text, 1)
		if err != nil {
			log.Printf("Error adding heading: %v", err)
		}
	case 3:
		_, err := r.dox.AddHeading(text, 3)
		if err != nil {
			log.Printf("Error adding heading: %v", err)
		}
	default:
		_, err := r.dox.AddHeading(text, 1)
		if err != nil {
			log.Printf("Error adding heading: %v", err)
		}
	}
}

func (r *DocxRenderer) Paragraph(text string) {
	p := r.dox.AddParagraph(text)
	p.Style("Italic")
}

func (r *DocxRenderer) Bullet(text string) {
	p := r.dox.AddParagraph(text)
	p.Style("List Bullet")
	p.Style("Italic")
}

func (r *DocxRenderer) Divider() {
	breakType := stypes.BreakTypeTextWrapping
	r.dox.AddParagraph("").AddRun().AddBreak(&breakType)
}

func (r *DocxRenderer) Table(headers []string, rows [][]string) {
	table := r.dox.AddTable()
	headerRow := table.AddRow()
	for _, h := range headers {
		cell := headerRow.AddCell()
		cell.AddParagraph(h).Style("Heading 3")
	}
	for _, row := range rows {
		excelRow := table.AddRow()
		for _, val := range row {
			cell := excelRow.AddCell()
			cell.AddParagraph(val)
		}
	}
}

func (r *DocxRenderer) Bytes() ([]byte, error) {
	var buf bytes.Buffer
	err := r.dox.Write(&buf)
	if err != nil {
		return nil, fmt.Errorf("failed to write docx to buffer: %w", err)
	}

	return buf.Bytes(), nil
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
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document": func() Render {
			return NewDocxRenderer()
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
	log.Printf("[generate_document] Starting upload - file: %s, size: %d bytes", args.FileName, len(data))
	var response *artifact.SaveResponse
	var uploadErr error

	for attempt := 0; attempt < 3; attempt++ {
		response, uploadErr = tctx.Artifacts().Save(
			tctx, args.FileName, &genai.Part{
				InlineData: &genai.Blob{
					MIMEType: args.MimeType,
					Data:     data,
				},
			},
		)
		if uploadErr == nil {
			break
		}
		log.Printf("[generate_document] Upload attempt %d failed - file: %s, error: %v",
			attempt+1, args.FileName, uploadErr)
	}

	log.Printf("[generate_document] Save() completed - err: %v", err)
	if uploadErr != nil {
		return GenerateDocumentResponse{
			StatusCode: 500,
			Message:    "Failed to upload document, please try again",
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
