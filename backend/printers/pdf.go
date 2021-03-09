package printers

import (
	"bz.moh.epi/godatatools/age"
	"bz.moh.epi/godatatools/models"
	"fmt"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"strconv"
)

// PdfPrinter generates a pdf of the lab test result.
func PdfPrinter(test models.LabTest) (pdf.Maroto, error) {
	m := pdf.NewMaroto(consts.Portrait, consts.Letter)
	m.SetBorder(false)

	if test.DateSampleTaken == nil {
		return m, fmt.Errorf("PdfPrinter: can not print a test(%s) result without date sample taken", test.ID)
	}

	if test.DateOfResult == nil {
		return m, fmt.Errorf("PdfPrinter: can not print a test(%s) result without a date of result", test.ID)
	}

	person := test.Person
	fullName := person.FirstName

	fullName = fmt.Sprintf("%s %s", fullName, person.LastName)
	ageAtTest := age.AgeAt(person.Dob, *test.DateSampleTaken)

	grayColor := getGrayColor()
	whiteColor := color.NewWhite()
	header(m)
	m.SetBackgroundColor(grayColor)
	m.Row(10, func() {
		m.Col(3, func() {
			m.Text("Test Details", props.Text{
				Top:   3,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
		m.ColSpace(9)
	})
	m.SetBackgroundColor(whiteColor)
	m.Row(4, func() {})

	m.Row(10, func() {
		m.SetBorder(true)
		m.Col(3, func() {
			m.Text("Name", props.Text{
				Top:   3,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
		m.Col(3, func() {
			m.Text(fullName, props.Text{
				Top:   3,
				Size:  12,
				Align: consts.Center,
			})
		})
		m.Col(3, func() {
			m.Text("Date Sample Taken", props.Text{
				Top:   3,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
		m.Col(3, func() {
			m.Text(test.DateSampleTaken.Format("2006-01-02"), props.Text{
				Top:   3,
				Size:  12,
				Align: consts.Center,
			})
		})
		m.SetBorder(false)
	})

	m.Row(10, func() {
		m.SetBorder(true)
		m.Col(3, func() {
			m.Text("Gender", props.Text{
				Top:   3,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
		m.Col(3, func() {
			m.Text(person.Gender, props.Text{
				Top:   3,
				Size:  12,
				Align: consts.Center,
			})
		})
		m.Col(3, func() {
			m.Text("Date of Result", props.Text{
				Top:   3,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
		m.Col(3, func() {
			m.Text(test.DateOfResult.Format("2006-01-02"), props.Text{
				Top:   3,
				Size:  12,
				Align: consts.Center,
			})
		})
		m.SetBorder(false)
	})

	m.Row(10, func() {
		m.SetBorder(true)
		m.Col(3, func() {
			m.Text("Age", props.Text{
				Top:   3,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
		m.Col(3, func() {
			m.Text(strconv.Itoa(ageAtTest), props.Text{
				Top:   3,
				Size:  12,
				Align: consts.Center,
			})
		})
		m.Col(3, func() {
			m.Text("Date of Birth", props.Text{
				Top:   3,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
		m.Col(3, func() {
			m.Text(person.Dob.Format("2006-01-02"), props.Text{
				Top:   3,
				Size:  12,
				Align: consts.Center,
			})
		})
		m.SetBorder(false)
	})

	m.Row(10, func() {})

	m.Row(10, func() {
		m.SetBorder(true)
		m.Col(3, func() {
			m.Text("Test Name", props.Text{
				Top:   3,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Middle,
			})
		})
		m.Col(3, func() {
			m.Text(test.TestType, props.Text{
				Top:   3,
				Size:  12,
				Align: consts.Middle,
			})
		})

		m.SetBorder(true)
		m.Col(3, func() {
			m.Text("Test Result", props.Text{
				Top:   3,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Middle,
			})
		})
		m.Col(3, func() {
			m.Text(test.Result, props.Text{
				Top:   3,
				Size:  12,
				Align: consts.Middle,
			})
		})

		m.SetBorder(false)
	})
	return m, nil
}

func header(m pdf.Maroto) {

	m.RegisterHeader(func() {
		m.Row(12, func() {
			m.Col(12, func() {
				m.Text("Ministry of Health & Wellness", props.Text{
					Top:         4,
					Size:        24,
					Extrapolate: true,
					Align:       consts.Center,
				})

			})
		})

		m.Row(4, func() {
			m.Col(12, func() {
				m.Text("Third Floor, East Block Building", props.Text{
					Top:   4,
					Size:  8,
					Align: consts.Center,
				})
			})
		})
		m.Row(4, func() {
			m.Col(12, func() {
				m.Text("Belmopan, Belize, Central America", props.Text{
					Top:   4,
					Size:  8,
					Align: consts.Center,
				})
			})
		})

		m.Line(10)

		m.Row(4, func() {
			m.Col(4, func() {
				m.Text("dhsmohw@health.gob.bz", props.Text{
					Top:         12,
					Size:        8,
					Align:       consts.Left,
					Extrapolate: true,
				})
			})
			m.ColSpace(4)
			m.Col(4, func() {
				m.Text(" Phone: +(501) 822-2363/2325/2497", props.Text{
					Top:         12,
					Size:        8,
					Align:       consts.Right,
					Extrapolate: true,
				})
			})
		})
	})
	m.Row(10, func() {})
}

func getGrayColor() color.Color {
	return color.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}
}
