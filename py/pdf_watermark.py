import sys
import os
import PyPDF2

def is_valid_pdf(file_path):
    try:
        with open(file_path, "rb") as pdf_file:
            reader = PyPDF2.PdfReader(pdf_file)
            return True
    except PyPDF2.utils.PdfReadError:
        return False

def add_watermark(input_pdf, watermark_text, output_pdf):
    watermark = PyPDF2.PageObject.create_watermark(watermark_text)

    with open(input_pdf, "rb") as input_file:
        pdf_reader = PyPDF2.PdfReader(input_file)
        pdf_writer = PyPDF2.PdfWriter()

        for page_num in range(len(pdf_reader.pages)):
            page = pdf_reader.pages[page_num]
            page.merge_page(watermark)
            pdf_writer.add_page(page)

        with open(output_pdf, "wb") as output_file:
            pdf_writer.write(output_file)

if __name__ == "__main__":
    if len(sys.argv) < 3:
        print("Usage: python pdf_watermark.py input.pdf \"watermark text\" [output.pdf]")
        sys.exit()

    input_pdf = sys.argv[1]
    watermark_text = sys.argv[2]
    output_pdf = "wb-" + os.path.basename(input_pdf)

    if len(sys.argv) > 3:
        output_pdf = sys.argv[3]

    if not is_valid_pdf(input_pdf):
        print("Error: Invalid PDF file.")
        sys.exit()

    add_watermark(input_pdf, watermark_text, output_pdf)
    print(f"Watermarked PDF saved as {output_pdf}")
