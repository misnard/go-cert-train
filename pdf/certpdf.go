package pdf

import (
	"fmt"
	"go-cert-train/cert"
	"os"
	"path"

	"github.com/jung-kurt/gofpdf"
)

type PdfSaver struct {
	OutputDir string
}

func New(outputdir string) (*PdfSaver, error) {
	var p *PdfSaver
	err := os.MkdirAll(outputdir, os.ModePerm)
	if err != nil {
		return p, err
	}

	p = &PdfSaver{
		OutputDir: outputdir,
	}

	return p, nil
}

func (p *PdfSaver) Save(cert cert.Cert) error {
	pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4", "")
	pdf.SetTitle(cert.LabelTitle, false)
	pdf.AddPage()

	//Background
	background(pdf)

	// --
	// Header
	header(pdf, &cert)
	pdf.Ln(30)

	// --
	// Body
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelPresented, "C")
	pdf.Ln(30)

	// Body - student name
	pdf.SetFont("Times", "B", 40)
	pdf.WriteAligned(0, 50, cert.Name, "C")
	pdf.Ln(30)

	// Body - Participation
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, cert.LabelParticipation, "C")
	pdf.Ln(30)

	// Body - Date
	pdf.SetFont("Helvetica", "I", 15)
	pdf.WriteAligned(0, 50, cert.LabelDate, "C")

	// --
	// Footer
	footer(pdf)

	//save file
	filename := fmt.Sprintf("%v.pdf", cert.LabelTitle)
	fmt.Printf("%v", p.OutputDir)
	savePath := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(savePath)
	if err != nil {
		return err
	}

	fmt.Printf("Saved certificate to '%v'\n", savePath)

	return nil
}

func background(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	pageWidth, pageHeight := pdf.GetPageSize()

	pdf.ImageOptions("img/border.png",
		0, 0,
		pageWidth, pageHeight,
		false, opts, 0, "")
}

func header(pdf *gofpdf.Fpdf, c *cert.Cert) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}

	margin := 30.0
	x := 0.0
	imageWidth := 30.0
	filename := "img/goopher.png"
	pdf.ImageOptions(
		filename,
		x+margin, 20,
		imageWidth, 0,
		false, opts, 0, "",
	)

	pageWidth, _ := pdf.GetPageSize()
	x = pageWidth - imageWidth
	pdf.ImageOptions(filename,
		x-margin, 20,
		imageWidth, 0,
		false, opts, 0, "",
	)
	pdf.SetFont("Helvetica", "", 40)
	pdf.WriteAligned(0, 50, c.LabelCompletion, "C")
}

func footer(pdf *gofpdf.Fpdf) {
	opts := gofpdf.ImageOptions{
		ImageType: "png",
	}
	filename := "img/stamp.png"

	pageWidth, pageHeight := pdf.GetPageSize()
	imageWidth := 60.0
	marginBottom := 60.0
	marginRight := 80.0

	pdf.ImageOptions(
		filename,
		pageWidth-marginRight, pageHeight-marginBottom,
		imageWidth, 0,
		false, opts, 0, "",
	)

}
