//go:build pdfium_experimental
// +build pdfium_experimental

package shared_tests

import (
	"io/ioutil"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/klippa-app/go-pdfium/references"
	"github.com/klippa-app/go-pdfium/requests"
)

var _ = Describe("text", func() {
	BeforeEach(func() {
		Locker.Lock()
	})

	AfterEach(func() {
		Locker.Unlock()
	})

	Context("a normal PDF file", func() {
		var doc references.FPDF_DOCUMENT

		BeforeEach(func() {
			pdfData, err := ioutil.ReadFile(TestDataPath + "/testdata/test.pdf")
			Expect(err).To(BeNil())

			newDoc, err := PdfiumInstance.FPDF_LoadMemDocument(&requests.FPDF_LoadMemDocument{
				Data: &pdfData,
			})
			Expect(err).To(BeNil())

			doc = newDoc.Document
		})

		AfterEach(func() {
			FPDF_CloseDocument, err := PdfiumInstance.FPDF_CloseDocument(&requests.FPDF_CloseDocument{
				Document: doc,
			})
			Expect(err).To(BeNil())
			Expect(FPDF_CloseDocument).To(Not(BeNil()))
		})

		When("is opened", func() {
			Context("when the structured page text is requested", func() {
				Context("when PixelPositions is enabled", func() {
					It("returns the correct font information", func() {
						pageTextStructured, err := PdfiumInstance.GetPageTextStructured(&requests.GetPageTextStructured{
							Page: requests.Page{
								ByIndex: &requests.PageByIndex{
									Document: doc,
									Index:    0,
								},
							},
							CollectFontInformation: true,
						})
						Expect(err).To(BeNil())
						Expect(pageTextStructured).To(Or(loadStructuredText(pageTextStructured, TestDataPath+"/testdata/text_experimental_"+TestType+"_testpdf_experimental_with_font_information.json", TestDataPath+"/testdata/text_experimental_"+TestType+"_testpdf_experimental_with_font_information_7019.json")...))
					})

					It("when dynamic tolerance is used for a specific PDF", func() {
						// This test case is designed to verify the dynamic tolerance feature
						// implemented in GetPageTextStructured for rectangle font information.
						//
						// To properly test this, 'test_dynamic_tolerance.pdf' should be
						// crafted or chosen carefully. It needs to contain text rectangles where:
						// a) The first character's font information can be found with an initial
						//    low tolerance (e.g., 1.0).
						// b) The first character's font information is initially missed with a low
						//    tolerance but found when the tolerance is dynamically increased.
						// c) The first character's font information is missed even with the
						//    maximum allowed tolerance.
						//
						// The corresponding JSON file,
						// 'text_experimental_"+TestType+"_testpdf_dynamic_tolerance.json',
						// will need to be created or adjusted based on the actual output of this
						// test when run with the aforementioned PDF. The JSON should reflect
						// the expected font information (or lack thereof) for each scenario.

						pdfDataDynamic, err := ioutil.ReadFile(TestDataPath + "/testdata/test_dynamic_tolerance.pdf")
						Expect(err).To(BeNil()) // Handle error if PDF is missing, though test setup implies it.

						newDocDynamic, err := PdfiumInstance.FPDF_LoadMemDocument(&requests.FPDF_LoadMemDocument{
							Data: &pdfDataDynamic,
						})
						Expect(err).To(BeNil())
						Expect(newDocDynamic).To(Not(BeNil()))
						// Ensure the document is closed after the test.
						defer func() {
							_, closeErr := PdfiumInstance.FPDF_CloseDocument(&requests.FPDF_CloseDocument{
								Document: newDocDynamic.Document,
							})
							Expect(closeErr).To(BeNil())
						}()

						pageTextStructured, err := PdfiumInstance.GetPageTextStructured(&requests.GetPageTextStructured{
							Page: requests.Page{
								ByIndex: &requests.PageByIndex{
									Document: newDocDynamic.Document,
									Index:    0, // Assuming the test cases are on the first page.
								},
							},
							CollectFontInformation: true,
						})
						Expect(err).To(BeNil())
						// The expected JSON needs to be created based on the output of a PDF
						// specifically designed for testing dynamic tolerance.
						// Note: Using Or(...) with a single path is equivalent to a direct match,
						// but if multiple valid JSONs were possible, this structure supports it.
						Expect(pageTextStructured).To(Or(loadStructuredText(pageTextStructured, TestDataPath+"/testdata/text_experimental_"+TestType+"_testpdf_dynamic_tolerance.json")...))
					})

					Context("and PixelPositions is enabled", func() {
						It("returns the correct font information", func() {
							pageTextStructured, err := PdfiumInstance.GetPageTextStructured(&requests.GetPageTextStructured{
								Page: requests.Page{
									ByIndex: &requests.PageByIndex{
										Document: doc,
										Index:    0,
									},
								},
								CollectFontInformation: true,
								PixelPositions: requests.GetPageTextStructuredPixelPositions{
									Calculate: true,
									Width:     3000,
									Height:    3000,
								},
							})
							Expect(err).To(BeNil())
							Expect(pageTextStructured).To(Or(loadStructuredText(pageTextStructured, TestDataPath+"/testdata/text_experimental_"+TestType+"_testpdf_experimental_with_font_information_and_pixel_positions.json", TestDataPath+"/testdata/text_experimental_"+TestType+"_testpdf_experimental_with_font_information_and_pixel_positions_7019.json")...))
						})
					})
				})
			})
		})
	})
})
