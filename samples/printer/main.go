package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

type TMainForm struct {
	*vcl.TForm
	CbbPrinters *vcl.TComboBox
	Btn1        *vcl.TButton
}

var (
	mainForm *TMainForm
)

func main() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&mainForm)
	vcl.Application.Run()
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {

	f.Btn1 = vcl.NewButton(f)
	f.Btn1.SetParent(f)
	f.Btn1.SetBounds(10, 10, 88, 28)
	f.Btn1.SetCaption("打印")
	f.Btn1.SetOnClick(f.OnButtonClick)

	f.CbbPrinters = vcl.NewComboBox(f)
	f.CbbPrinters.SetParent(f)
	f.CbbPrinters.SetBounds(f.Btn1.Left(), f.Btn1.Top()+f.Btn1.Height()+20, 200, f.CbbPrinters.Height())
	f.CbbPrinters.Items().Assign(vcl.Printer.Printers())
	f.CbbPrinters.SetOnChange(f.OnCbbChange)

	// Printer.Orientation 方向

}

func (f *TMainForm) OnButtonClick(sender vcl.IObject) {

	if f.CbbPrinters.ItemIndex() == -1 {
		vcl.ShowMessage("先设置一个打印机.")
		return
	}

	// 注意：由于打印机的Canvas与系统的DPI有关，所以很多都需要通过换算
	// vcl.Screen.PixelsPerInch()

	rx := float64(vcl.Screen.PixelsPerInch()) / 96.0
	fmt.Println("vcl.Screen.PixelsPerInch(): ", vcl.Screen.PixelsPerInch(), ", rx: ", rx)
	//vcl.Printer.SetTitle("标题啊。。。。 ")

	vcl.Printer.BeginDoc()
	defer vcl.Printer.EndDoc()

	canvas := vcl.Printer.Canvas()
	//canvas.Brush().SetColor(colors.ClWhite)
	canvas.Brush().SetStyle(types.BsClear)

	font := canvas.Font()
	font.SetName("微软雅黑")
	font.SetSize(16)
	font.SetColor(colors.ClGreen)

	canvas.TextOut(0, 0, "这是一个测试")

	// 这里画个图片

	jpgImg := vcl.NewJPEGImage()
	stream := vcl.NewMemoryStreamFromBytes(testImgBytes)
	stream.SetPosition(0)
	jpgImg.LoadFromStream(stream)
	canvas.Draw(200, 200, jpgImg)
	//canvas.StretchDraw(types.TRect{10, 10, int32(float64(jpgImg.Width()) * rx), int32(float64(jpgImg.Height()) * rx)}, jpgImg)
	stream.Free()
	jpgImg.Free()

	// 新建一页
	vcl.Printer.NewPage()
	canvas.Font().SetSize(12)
	canvas.Font().SetColor(colors.ClBlack)
	// 这里实际要通过相关的api获取打印机canvas大小，而且那东西与dpi有关
	r := types.TRect{0, 0, 2000, 2000}
	canvas.TextRect3(&r,
		"这是一段文字，只是用来做测试。这是一段文字，只是用来做测试。这是一段文字，只是用来做测试。这是一段文字，只是用来做测试。这是一段文字，只是用来做测试。",
		rtl.Include(0, types.TfCenter, types.TfWordBreak, types.TfVerticalCenter))
}

func (f *TMainForm) OnCbbChange(sender vcl.IObject) {
	if f.CbbPrinters.ItemIndex() != -1 {
		vcl.Printer.SetPrinterIndex(f.CbbPrinters.ItemIndex())
	}
}

// 一个图片字节, jpg格式
var (
	testImgBytes = []byte{
		0xFF, 0xD8, 0xFF, 0xE1, 0x00, 0x9C, 0x45, 0x78, 0x69, 0x66, 0x00, 0x00, 0x49, 0x49, 0x2A, 0x00,
		0x08, 0x00, 0x00, 0x00, 0x07, 0x00, 0x32, 0x01, 0x02, 0x00, 0x14, 0x00, 0x00, 0x00, 0x62, 0x00,
		0x00, 0x00, 0x08, 0x92, 0x03, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x01,
		0x03, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x03, 0x00, 0x01, 0x00,
		0x00, 0x00, 0xC4, 0x00, 0x00, 0x00, 0x07, 0x92, 0x03, 0x00, 0x01, 0x00, 0x00, 0x00, 0xFF, 0xFF,
		0xFF, 0xFF, 0x00, 0x01, 0x03, 0x00, 0x01, 0x00, 0x00, 0x00, 0xC6, 0x00, 0x00, 0x00, 0x69, 0x87,
		0x04, 0x00, 0x01, 0x00, 0x00, 0x00, 0x76, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x32, 0x30,
		0x31, 0x36, 0x3A, 0x30, 0x35, 0x3A, 0x32, 0x33, 0x20, 0x31, 0x37, 0x3A, 0x35, 0x32, 0x3A, 0x33,
		0x36, 0x00, 0x02, 0x00, 0x01, 0x02, 0x04, 0x00, 0x01, 0x00, 0x00, 0x00, 0x94, 0x00, 0x00, 0x00,
		0x02, 0x02, 0x04, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xFF, 0xDB, 0x00, 0x43, 0x00, 0x06, 0x04, 0x05, 0x06, 0x05, 0x04, 0x06, 0x06, 0x05, 0x06, 0x07,
		0x07, 0x06, 0x08, 0x0A, 0x10, 0x0A, 0x0A, 0x09, 0x09, 0x0A, 0x14, 0x0E, 0x0F, 0x0C, 0x10, 0x17,
		0x14, 0x18, 0x18, 0x17, 0x14, 0x16, 0x16, 0x1A, 0x1D, 0x25, 0x1F, 0x1A, 0x1B, 0x23, 0x1C, 0x16,
		0x16, 0x20, 0x2C, 0x20, 0x23, 0x26, 0x27, 0x29, 0x2A, 0x29, 0x19, 0x1F, 0x2D, 0x30, 0x2D, 0x28,
		0x30, 0x25, 0x28, 0x29, 0x28, 0xFF, 0xDB, 0x00, 0x43, 0x01, 0x07, 0x07, 0x07, 0x0A, 0x08, 0x0A,
		0x13, 0x0A, 0x0A, 0x13, 0x28, 0x1A, 0x16, 0x1A, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28,
		0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28,
		0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28,
		0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0x28, 0xFF, 0xC2, 0x00, 0x11, 0x08, 0x00,
		0xC4, 0x00, 0xC6, 0x03, 0x01, 0x22, 0x00, 0x02, 0x11, 0x01, 0x03, 0x11, 0x01, 0xFF, 0xC4, 0x00,
		0x1B, 0x00, 0x01, 0x00, 0x02, 0x03, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x05, 0x06, 0x02, 0x03, 0x04, 0x01, 0x07, 0xFF, 0xC4, 0x00, 0x16, 0x01, 0x01,
		0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x01, 0x02, 0xFF, 0xDA, 0x00, 0x0C, 0x03, 0x01, 0x00, 0x02, 0x10, 0x03, 0x10, 0x00, 0x00, 0x01,
		0xFA, 0x88, 0x00, 0x00, 0x02, 0x37, 0xA8, 0xE8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0xA6,
		0x04, 0xB2, 0x2A, 0xBD, 0x04, 0xE6, 0x30, 0xF8, 0x1A, 0x26, 0x21, 0x2C, 0x67, 0x51, 0xC5, 0x1D,
		0xAF, 0x3D, 0xA0, 0x08, 0xFA, 0xC9, 0x76, 0x43, 0x49, 0x9B, 0x80, 0x00, 0x08, 0x49, 0xBA, 0xE1,
		0xB6, 0x0E, 0x5F, 0xA4, 0x91, 0xD7, 0xDF, 0x84, 0x53, 0x6E, 0x5E, 0x7B, 0x98, 0xA9, 0xDB, 0xF9,
		0xB5, 0x75, 0x56, 0xF8, 0x6D, 0xF2, 0x47, 0xCF, 0xD5, 0xEC, 0x5A, 0xBB, 0x81, 0xCF, 0xF2, 0x8F,
		0xAB, 0xD1, 0x88, 0xFF, 0x00, 0xA3, 0x7C, 0xF3, 0xE8, 0x47, 0x60, 0x00, 0x01, 0x5C, 0xB0, 0x46,
		0x11, 0x76, 0x9A, 0xA5, 0xBA, 0x0E, 0x48, 0x72, 0xC6, 0xC3, 0x86, 0xA4, 0x7C, 0xAC, 0xE4, 0x48,
		0x6D, 0xF7, 0x32, 0x3F, 0x1E, 0xE8, 0x32, 0xD4, 0xD2, 0x37, 0x7C, 0xD3, 0xE9, 0x70, 0xC5, 0x07,
		0xE8, 0x50, 0x56, 0x02, 0x48, 0x00, 0x01, 0xE6, 0x8E, 0x8C, 0x08, 0x2B, 0x05, 0x72, 0xC7, 0x99,
		0x5A, 0xEC, 0xD1, 0x60, 0x2A, 0xDD, 0x1C, 0xD3, 0x9A, 0xBB, 0xBA, 0x80, 0x0D, 0x50, 0x16, 0x4C,
		0x0E, 0x6C, 0xA0, 0x65, 0xCE, 0x9F, 0x9B, 0xFD, 0x3A, 0x82, 0x75, 0x59, 0x69, 0x56, 0x12, 0x7F,
		0x65, 0x32, 0xE6, 0x00, 0x01, 0x86, 0x65, 0x4E, 0xDB, 0x07, 0x24, 0x45, 0x58, 0x39, 0xFA, 0x0D,
		0x35, 0xAB, 0x5C, 0x39, 0x30, 0x00, 0x00, 0xC2, 0xAB, 0x6D, 0xC0, 0xE1, 0xCA, 0x06, 0xC6, 0x22,
		0x27, 0xC6, 0x19, 0xF3, 0xE1, 0x1D, 0x62, 0x80, 0xF3, 0xDF, 0x3D, 0x34, 0xC4, 0xCE, 0x44, 0x93,
		0x08, 0xB9, 0x40, 0x00, 0x00, 0x00, 0x0D, 0x35, 0x8B, 0x6F, 0x39, 0x9E, 0x75, 0x7F, 0x4D, 0x3A,
		0x3A, 0x25, 0x33, 0x25, 0x36, 0x56, 0xEC, 0x9A, 0xA0, 0x79, 0xE8, 0x30, 0xCF, 0xC2, 0xB7, 0x61,
		0xD1, 0x02, 0x5A, 0xC0, 0x00, 0x02, 0x2C, 0x94, 0x03, 0xCF, 0x44, 0x1C, 0xCD, 0x7E, 0xC0, 0x55,
		0xB3, 0xE9, 0xEB, 0x21, 0xAD, 0x95, 0x5B, 0x50, 0x00, 0x00, 0x20, 0x67, 0xB0, 0x21, 0x67, 0xE9,
		0xF6, 0x08, 0x90, 0x14, 0x05, 0x6B, 0xDD, 0xBE, 0xE2, 0x6C, 0x9A, 0xA6, 0xDB, 0x37, 0x77, 0x61,
		0x8C, 0x31, 0x94, 0xDC, 0x6E, 0xE8, 0xCF, 0x5C, 0x1E, 0xBA, 0xDF, 0x65, 0x88, 0x97, 0x00, 0x00,
		0x00, 0x39, 0xEA, 0x57, 0x5E, 0x5C, 0xCC, 0xBA, 0xA9, 0xD6, 0x6D, 0x5E, 0xB8, 0xD9, 0x28, 0x53,
		0x83, 0xAE, 0x7B, 0x8C, 0x8D, 0xE6, 0x9A, 0x80, 0x2C, 0x99, 0xEC, 0x89, 0x84, 0x3E, 0x89, 0x99,
		0x25, 0x77, 0x1A, 0xA0, 0x00, 0x00, 0x00, 0x01, 0xA6, 0x02, 0xCB, 0x81, 0x1B, 0xDF, 0x19, 0x1A,
		0x5B, 0x55, 0xF9, 0x33, 0x44, 0x34, 0xC4, 0x59, 0xBB, 0x9B, 0x9F, 0x66, 0x65, 0xC2, 0x16, 0x37,
		0x4E, 0xAD, 0xB7, 0x3F, 0x3D, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0F, 0x31, 0xCC, 0x45, 0xC7, 0xD9,
		0x04, 0x14, 0xD6, 0x63, 0xCD, 0x5B, 0x87, 0x3E, 0xDC, 0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x07, 0xFF, 0xC4, 0x00, 0x2A, 0x10, 0x00, 0x02, 0x02, 0x02, 0x00,
		0x05, 0x03, 0x04, 0x03, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x03, 0x02, 0x04, 0x00,
		0x05, 0x11, 0x12, 0x13, 0x14, 0x20, 0x10, 0x30, 0x31, 0x06, 0x21, 0x22, 0x23, 0x15, 0x40, 0x50,
		0x24, 0x35, 0xFF, 0xDA, 0x00, 0x08, 0x01, 0x01, 0x00, 0x01, 0x05, 0x02, 0xFF, 0x00, 0x05, 0x36,
		0xA2, 0xC0, 0xA9, 0x86, 0x47, 0xFA, 0xA6, 0x71, 0x19, 0x23, 0xCC, 0xBA, 0xC9, 0x31, 0x4A, 0x2C,
		0xF5, 0x7F, 0xA5, 0x23, 0xCB, 0x18, 0xDC, 0xB3, 0x28, 0xD7, 0xDA, 0xCE, 0xCE, 0x7F, 0x1E, 0xD6,
		0x64, 0x35, 0x75, 0x72, 0x7A, 0x75, 0x61, 0x75, 0x8A, 0x59, 0x5D, 0xD0, 0x72, 0xF3, 0xBA, 0x4F,
		0x3F, 0x1F, 0x0B, 0xEC, 0x2A, 0xA9, 0x73, 0x61, 0x2E, 0xDF, 0x61, 0xB0, 0x85, 0x2C, 0x5C, 0x84,
		0xE3, 0xE5, 0x6F, 0x65, 0x14, 0x35, 0x7B, 0x3A, 0xED, 0xCA, 0xBC, 0xD7, 0x05, 0x64, 0xC5, 0x2B,
		0xB5, 0x0E, 0x65, 0x68, 0x4B, 0x84, 0x70, 0x80, 0x45, 0x88, 0xFF, 0x00, 0x18, 0xD7, 0xBE, 0x2B,
		0x42, 0xF5, 0xDC, 0xC2, 0x8B, 0x8A, 0x9B, 0xEA, 0xF8, 0xC6, 0x6B, 0xBE, 0x85, 0x44, 0x6E, 0x91,
		0x04, 0x4E, 0xAC, 0xFA, 0x8A, 0xF2, 0x1F, 0x8E, 0xCB, 0x73, 0xC8, 0x2A, 0x6A, 0xAA, 0x0A, 0xB5,
		0x70, 0x8E, 0x38, 0x07, 0x0F, 0x57, 0x2A, 0x2D, 0x80, 0x32, 0x33, 0xF8, 0xCD, 0xC7, 0xEB, 0xC5,
		0x91, 0x28, 0xFA, 0x3C, 0x19, 0x2B, 0x61, 0x57, 0xB6, 0xD2, 0xD7, 0x6C, 0xAC, 0x5B, 0xA9, 0xCF,
		0xD1, 0xF2, 0xBC, 0x7A, 0x36, 0xEE, 0x7F, 0xD3, 0x72, 0x23, 0x80, 0xF1, 0x3F, 0x13, 0xAD, 0x1E,
		0xE3, 0x36, 0x2A, 0xEB, 0x54, 0xD2, 0xB7, 0xAB, 0x43, 0xD7, 0xEA, 0x31, 0xF9, 0xF5, 0x5A, 0xB5,
		0x6B, 0x6C, 0xAA, 0xC2, 0x7C, 0x47, 0xC5, 0xA4, 0x07, 0x2F, 0x48, 0x4B, 0x2C, 0xE5, 0x87, 0x41,
		0x2B, 0x4E, 0xD9, 0x33, 0x3C, 0x72, 0xD5, 0x90, 0xAC, 0x83, 0xDA, 0xDC, 0xE5, 0xB0, 0x05, 0x68,
		0x91, 0x09, 0x40, 0x8C, 0x39, 0xF4, 0xEF, 0xE2, 0xB9, 0xC8, 0x8C, 0x12, 0xE2, 0x73, 0x63, 0x55,
		0x4D, 0x42, 0x76, 0x48, 0xAC, 0xBA, 0x30, 0x02, 0x3E, 0x3C, 0x32, 0x63, 0x35, 0xC3, 0x96, 0xDE,
		0x6C, 0x87, 0x56, 0xDA, 0x35, 0xF0, 0x58, 0xAB, 0x6B, 0xB6, 0x6C, 0x55, 0x37, 0xB1, 0x2A, 0x8A,
		0xE3, 0xEA, 0x62, 0x0E, 0x48, 0x1A, 0x2E, 0x24, 0x4E, 0x2A, 0xF9, 0x97, 0x13, 0x1D, 0xDD, 0x59,
		0x17, 0x54, 0xD8, 0x41, 0x55, 0x6B, 0x5A, 0x5D, 0x8C, 0x1F, 0x1E, 0x24, 0x71, 0x15, 0x8F, 0x1D,
		0xA6, 0x5B, 0xFF, 0x00, 0xD3, 0xCF, 0xA8, 0x2B, 0x86, 0x22, 0x90, 0x87, 0x6F, 0xE2, 0x47, 0x1C,
		0xB1, 0x19, 0x51, 0xCA, 0xAE, 0x83, 0xA1, 0x9B, 0x9A, 0xDD, 0xAC, 0xF6, 0xD5, 0xA2, 0x8A, 0x42,
		0xFD, 0x1A, 0x48, 0xFA, 0x75, 0xCD, 0x7B, 0x3C, 0xA3, 0xFA, 0x76, 0xC3, 0x2E, 0x7D, 0xAF, 0xE4,
		0xE1, 0xCF, 0x1F, 0xA7, 0xA7, 0xFA, 0x3C, 0xB8, 0x63, 0xD5, 0x3A, 0x13, 0xA9, 0x66, 0x16, 0x57,
		0x3A, 0xAA, 0x9B, 0x1C, 0x98, 0x3A, 0x0A, 0xD4, 0x52, 0x59, 0x00, 0x01, 0xE1, 0x03, 0xC4, 0x66,
		0xC6, 0x1F, 0x8D, 0x76, 0x75, 0x17, 0x35, 0x46, 0x72, 0xF4, 0xA9, 0x57, 0xA0, 0xFF, 0x00, 0x32,
		0x3E, 0xD6, 0xD6, 0x68, 0x35, 0x4C, 0x0C, 0x87, 0xA3, 0x59, 0x15, 0x84, 0xD9, 0x53, 0xBC, 0x07,
		0xC6, 0x4E, 0x22, 0x51, 0x54, 0xBA, 0x0F, 0xF7, 0x25, 0x11, 0x20, 0x92, 0x75, 0xB6, 0x01, 0x04,
		0x63, 0x90, 0x36, 0x17, 0x6F, 0xA6, 0x3A, 0xEC, 0x89, 0x12, 0x1E, 0x83, 0xD6, 0xF5, 0x60, 0xE8,
		0xD1, 0xB3, 0xD7, 0x8F, 0xB8, 0xD8, 0x06, 0x41, 0x48, 0x9C, 0x58, 0xCE, 0xEA, 0x9E, 0x7D, 0x3E,
		0xC0, 0xD5, 0xDB, 0x48, 0xB0, 0x9F, 0xA7, 0x9D, 0xD6, 0xD7, 0xFA, 0x0F, 0x59, 0x0E, 0x22, 0xD9,
		0xED, 0x5A, 0xB9, 0x09, 0x47, 0xDB, 0x39, 0x28, 0x9E, 0xE3, 0xE4, 0x68, 0xC7, 0x05, 0x6B, 0xDD,
		0xF9, 0x6A, 0xBF, 0x45, 0xDF, 0x13, 0x8F, 0x50, 0x9C, 0x75, 0xAC, 0x35, 0xD9, 0xE6, 0x2F, 0x57,
		0x2D, 0xF0, 0xB1, 0xCB, 0xDC, 0x0F, 0xB8, 0xD7, 0xD7, 0xE8, 0x10, 0x88, 0x86, 0x37, 0xF1, 0xDB,
		0x79, 0xED, 0x2A, 0x77, 0x10, 0xD4, 0x5C, 0xEE, 0x55, 0xE3, 0xBD, 0x94, 0xA3, 0x48, 0x6A, 0x93,
		0x15, 0xEB, 0xAD, 0x75, 0x47, 0xA1, 0x3C, 0x04, 0x5D, 0x29, 0xB4, 0x7C, 0x49, 0x63, 0x38, 0x67,
		0x0E, 0xA6, 0xD3, 0xCE, 0x43, 0x8E, 0x6C, 0x95, 0x2A, 0xCC, 0xA7, 0x66, 0x36, 0x15, 0xE1, 0xBA,
		0x5F, 0x3E, 0xBF, 0xB8, 0xE1, 0x52, 0xCB, 0x0A, 0xA5, 0x09, 0x09, 0xC6, 0x52, 0xE5, 0x13, 0xB1,
		0x27, 0x65, 0x5A, 0xFC, 0x99, 0x66, 0xCA, 0xEB, 0xC0, 0x5A, 0xB5, 0x6B, 0x3B, 0x6B, 0xF9, 0xAF,
		0xA7, 0xDB, 0x0F, 0x61, 0xB1, 0xE2, 0x26, 0xA9, 0x6A, 0x98, 0x96, 0x06, 0x43, 0xD2, 0xED, 0xA1,
		0x55, 0x72, 0x45, 0xD7, 0xC7, 0x5F, 0x5C, 0xA6, 0xBD, 0x7A, 0x31, 0x59, 0xD5, 0x12, 0x86, 0x18,
		0x71, 0x3C, 0xA0, 0x66, 0xC6, 0xE4, 0x29, 0x26, 0xA5, 0x39, 0xBA, 0x51, 0x88, 0x1E, 0xE4, 0xE3,
		0xCC, 0x00, 0x9E, 0xB2, 0x55, 0xDF, 0x17, 0x2F, 0x26, 0x3A, 0x97, 0xB1, 0xB1, 0xFB, 0xE6, 0xDC,
		0xF4, 0x5A, 0x3E, 0x2F, 0xDE, 0x55, 0x38, 0xC3, 0x5C, 0xCB, 0xAD, 0xA4, 0xD9, 0x73, 0x7B, 0xB2,
		0x88, 0x39, 0x3A, 0xB2, 0xAF, 0x3A, 0x97, 0xA0, 0xFC, 0xE9, 0x81, 0x3F, 0x4B, 0xCC, 0xE8, 0xA3,
		0xB7, 0x7D, 0xC6, 0x6D, 0x36, 0x22, 0xA0, 0xD5, 0xEB, 0xB9, 0x25, 0x97, 0x17, 0xFB, 0x47, 0xBC,
		0x62, 0x08, 0xB7, 0x46, 0x16, 0x07, 0x75, 0x62, 0x9E, 0x57, 0xDB, 0xD4, 0x6E, 0x0B, 0x2A, 0x38,
		0xEB, 0xF5, 0x94, 0x25, 0xB0, 0x65, 0x8C, 0x3A, 0x50, 0x62, 0x2E, 0x5B, 0xAA, 0x3F, 0x9D, 0x4E,
		0x1B, 0x33, 0xD8, 0x4C, 0x0E, 0x03, 0xDE, 0x38, 0x62, 0x0E, 0x3A, 0x82, 0x1B, 0x87, 0x49, 0x4C,
		0xE5, 0x7D, 0x55, 0x55, 0x60, 0x00, 0x0E, 0x18, 0x62, 0x0E, 0x74, 0x86, 0x01, 0xC3, 0xFD, 0x8F,
		0xFF, 0xC4, 0x00, 0x14, 0x11, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x70, 0xFF, 0xDA, 0x00, 0x08, 0x01, 0x03, 0x01, 0x01, 0x3F, 0x01,
		0x29, 0xFF, 0xC4, 0x00, 0x18, 0x11, 0x00, 0x02, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x12, 0x40, 0x50, 0x60, 0xFF, 0xDA, 0x00, 0x08, 0x01,
		0x02, 0x01, 0x01, 0x3F, 0x01, 0x9A, 0x39, 0x15, 0xBB, 0xFF, 0xC4, 0x00, 0x39, 0x10, 0x00, 0x02,
		0x01, 0x02, 0x04, 0x03, 0x06, 0x03, 0x06, 0x05, 0x05, 0x01, 0x00, 0x00, 0x00, 0x00, 0x01, 0x02,
		0x03, 0x00, 0x11, 0x04, 0x12, 0x21, 0x31, 0x10, 0x41, 0x51, 0x13, 0x20, 0x22, 0x32, 0x61, 0x71,
		0x30, 0x42, 0x81, 0x14, 0x23, 0x33, 0x52, 0x91, 0xD1, 0x34, 0x40, 0x43, 0x50, 0xB1, 0x05, 0x24,
		0x72, 0xA1, 0xC1, 0xE1, 0xFF, 0xDA, 0x00, 0x08, 0x01, 0x01, 0x00, 0x06, 0x3F, 0x02, 0xFE, 0xC2,
		0xCD, 0xB4, 0x6A, 0x6D, 0x98, 0xF3, 0xA0, 0xCA, 0x6E, 0x0E, 0xC7, 0xF9, 0x6D, 0x58, 0x0F, 0xAD,
		0x35, 0x8D, 0xCD, 0x8E, 0xD5, 0x85, 0x59, 0x06, 0x58, 0xD1, 0x6E, 0xC0, 0xF5, 0xBD, 0x0C, 0xB1,
		0x38, 0x5E, 0x4C, 0x46, 0x9F, 0xC9, 0x13, 0xD2, 0xB3, 0xAC, 0x08, 0x50, 0xEA, 0x3C, 0x56, 0x36,
		0xAB, 0x41, 0x86, 0x6C, 0xE3, 0x43, 0x98, 0xE8, 0x3E, 0xBC, 0xEB, 0xFD, 0xC6, 0x29, 0xCF, 0xA4,
		0x7E, 0x11, 0x47, 0xC2, 0x58, 0xF3, 0xBB, 0x55, 0xE1, 0x92, 0x48, 0x5B, 0xAA, 0x9A, 0xB6, 0x21,
		0x7E, 0xD3, 0x0F, 0xE7, 0x51, 0xA8, 0xF7, 0x14, 0x1E, 0x36, 0x0C, 0xA7, 0x6B, 0x70, 0xCB, 0xDA,
		0xAE, 0x6E, 0x97, 0xEE, 0xCC, 0xEB, 0xE6, 0x55, 0xB8, 0xAF, 0xF4, 0xD6, 0x12, 0x00, 0xF2, 0xBA,
		0xE6, 0xB7, 0x31, 0x51, 0x76, 0x9B, 0x39, 0xB7, 0xB5, 0x02, 0xA6, 0xE0, 0xEC, 0x7B, 0xFD, 0x98,
		0x8A, 0x49, 0x58, 0x0B, 0x9C, 0x83, 0x6A, 0x28, 0x58, 0xC6, 0xDD, 0x1C, 0x5A, 0x84, 0x2A, 0x4A,
		0x43, 0x0E, 0x8E, 0x47, 0xCC, 0x7A, 0x0A, 0x08, 0x8A, 0x14, 0x0E, 0x94, 0xD7, 0x62, 0xA3, 0x7B,
		0x8A, 0x21, 0xE0, 0x39, 0x1D, 0x89, 0x12, 0x13, 0xA9, 0xF7, 0x1C, 0x2C, 0x76, 0xAE, 0xDA, 0x3F,
		0xE1, 0xA4, 0x3F, 0x78, 0xA3, 0xE5, 0xF5, 0x14, 0xF2, 0x13, 0xE1, 0x0B, 0x7A, 0x8D, 0x8A, 0xA1,
		0xCD, 0xE2, 0x7B, 0xEF, 0xFA, 0xD1, 0xC3, 0x4A, 0xD9, 0xAC, 0x2E, 0xAD, 0xD4, 0x77, 0x19, 0x5F,
		0xCA, 0x45, 0x8D, 0x21, 0xC2, 0xC5, 0x26, 0x55, 0x90, 0x0E, 0xD2, 0x43, 0xA7, 0xD2, 0xB0, 0x8D,
		0x2C, 0xCD, 0x88, 0x97, 0x38, 0x24, 0x7A, 0x7B, 0x52, 0xB0, 0x52, 0xA0, 0x8D, 0x88, 0xD4, 0x77,
		0xF1, 0x03, 0xF3, 0x20, 0x35, 0x2B, 0x3C, 0x61, 0xB4, 0xB0, 0xBF, 0x5A, 0x8D, 0x39, 0xDA, 0xE7,
		0xDF, 0x8E, 0x9C, 0x59, 0x1C, 0x5D, 0x58, 0x58, 0x8A, 0x18, 0x06, 0x26, 0xCA, 0xF7, 0xBF, 0x54,
		0xDF, 0x84, 0x18, 0x81, 0xBC, 0x2E, 0x2F, 0xEC, 0x77, 0xA0, 0x47, 0x3D, 0x78, 0xC8, 0x17, 0x72,
		0x34, 0xA0, 0xB2, 0x11, 0x9D, 0x18, 0x36, 0xFE, 0xB4, 0x64, 0xC0, 0x61, 0x95, 0x97, 0x2D, 0x9A,
		0x49, 0xB6, 0x06, 0x97, 0xB5, 0x65, 0x67, 0xE6, 0x57, 0x6E, 0xFE, 0x1E, 0x5D, 0x81, 0xBA, 0x1F,
		0xFC, 0xAC, 0x2C, 0x03, 0xCA, 0xBF, 0x7C, 0xDF, 0xF9, 0x56, 0xEF, 0xF6, 0xDF, 0x36, 0x5C, 0xA3,
		0x86, 0x21, 0x7A, 0xA9, 0xA8, 0x58, 0xEF, 0x6B, 0x77, 0x1F, 0x2D, 0xE5, 0x96, 0xD7, 0xD7, 0xCB,
		0x10, 0xFD, 0xEA, 0x08, 0x5F, 0x0C, 0xC9, 0x85, 0x16, 0xCC, 0x62, 0xD4, 0xB7, 0xBD, 0x29, 0x83,
		0xC8, 0x05, 0xAD, 0xD3, 0xBD, 0xAD, 0x32, 0x75, 0xFF, 0x00, 0xA3, 0x58, 0xB7, 0x7F, 0x32, 0xDA,
		0x3F, 0xD3, 0x81, 0x79, 0x1B, 0x2A, 0x8E, 0x74, 0xA0, 0xAB, 0xA8, 0x6D, 0x99, 0x85, 0xB8, 0x5B,
		0xE6, 0x20, 0x91, 0x50, 0xB0, 0xB9, 0x3D, 0x96, 0x6B, 0x0D, 0x89, 0xA0, 0x05, 0xCE, 0x55, 0xE7,
		0xF3, 0x13, 0xFB, 0x50, 0x57, 0x25, 0x8D, 0xB5, 0x35, 0xA6, 0xBC, 0x31, 0x11, 0xF2, 0x49, 0x2D,
		0xC2, 0xDC, 0x26, 0x0E, 0x44, 0x79, 0xC6, 0xAF, 0x4B, 0x04, 0x19, 0xF1, 0x24, 0x69, 0xE1, 0x1B,
		0xD6, 0x7E, 0xC0, 0x44, 0xCD, 0xB8, 0x1D, 0xF0, 0x6B, 0x1D, 0xD4, 0xB0, 0x3F, 0xF5, 0xC3, 0x05,
		0x19, 0x17, 0x17, 0x2D, 0x6F, 0x6A, 0x94, 0x12, 0xCE, 0xAF, 0xC9, 0x8D, 0x49, 0x85, 0x99, 0xAD,
		0x6D, 0x63, 0x27, 0x98, 0xA8, 0x59, 0xC5, 0x9A, 0x22, 0x41, 0x3C, 0x88, 0xA0, 0xAA, 0x2C, 0x06,
		0xDD, 0xCD, 0xAA, 0x47, 0xCB, 0x78, 0x24, 0x37, 0x6B, 0x6E, 0xA7, 0xF6, 0xAB, 0xA9, 0xBA, 0x91,
		0xA1, 0x1C, 0x0D, 0x8D, 0x8F, 0x5A, 0x82, 0x1C, 0xF2, 0x62, 0x67, 0x63, 0x7C, 0xA4, 0xE8, 0x05,
		0x03, 0xD8, 0x22, 0x11, 0x27, 0x63, 0x65, 0xDA, 0x9F, 0xB3, 0x60, 0xC1, 0x4D, 0x89, 0x15, 0xA6,
		0xBD, 0xF9, 0xC0, 0xE5, 0x18, 0xBF, 0xBD, 0xF8, 0x60, 0xFD, 0x9B, 0x82, 0xCA, 0x14, 0x16, 0x88,
		0xDF, 0xDC, 0x73, 0x14, 0x86, 0x21, 0x65, 0x22, 0xFD, 0xED, 0x76, 0xA3, 0x24, 0x40, 0xB6, 0x1F,
		0x76, 0x4F, 0xCB, 0xEA, 0x28, 0x3C, 0x6D, 0x75, 0x3C, 0x1F, 0x15, 0x16, 0x21, 0xC4, 0xF2, 0x9C,
		0xAA, 0xBD, 0x7D, 0x2B, 0x04, 0x92, 0x5F, 0x2F, 0x68, 0x3B, 0x43, 0xEF, 0xBD, 0x00, 0x24, 0x8C,
		0x20, 0x1A, 0x05, 0xAC, 0x54, 0x96, 0x61, 0x87, 0x66, 0xBC, 0x79, 0xBB, 0xE2, 0xC0, 0x85, 0x9D,
		0x75, 0xD7, 0x73, 0xC3, 0x06, 0xDF, 0xF2, 0x1C, 0x18, 0x1D, 0x88, 0xB5, 0x4D, 0x09, 0xDE, 0x19,
		0x0A, 0x7C, 0x03, 0x3E, 0x19, 0x4B, 0x44, 0xDF, 0x89, 0x18, 0xFF, 0x00, 0x22, 0x83, 0xC6, 0xD7,
		0x53, 0x4B, 0x23, 0x20, 0x2E, 0xBB, 0x13, 0xCA, 0x8A, 0xC8, 0xA1, 0x94, 0xF2, 0x22, 0xAE, 0xB8,
		0x64, 0xBD, 0x58, 0x0B, 0x0E, 0x9D, 0xF0, 0xE3, 0x42, 0x9A, 0xDF, 0xD3, 0x9D, 0x03, 0xCC, 0xF2,
		0xE9, 0x4A, 0x4E, 0xEA, 0x6E, 0x38, 0xE2, 0x9C, 0x1D, 0x25, 0x6B, 0xDB, 0xA7, 0xC1, 0xFB, 0x4C,
		0x20, 0xF6, 0x4C, 0x7E, 0xF5, 0x07, 0xF9, 0x14, 0x19, 0x48, 0x2A, 0x45, 0xC1, 0x1C, 0x6E, 0xEC,
		0x14, 0x75, 0x26, 0xD4, 0x7B, 0x39, 0x15, 0xAD, 0xD3, 0xBC, 0x41, 0x17, 0xAC, 0x84, 0xF8, 0x5F,
		0x6F, 0x7E, 0x9F, 0x14, 0x82, 0x2E, 0x0D, 0x76, 0x6F, 0xFC, 0x2C, 0x87, 0xC0, 0x7F, 0x21, 0xE9,
		0x57, 0x06, 0xFC, 0x24, 0xCE, 0x7C, 0x10, 0x58, 0x0F, 0x7E, 0x75, 0x0C, 0xF0, 0x0B, 0x2A, 0x35,
		0x9C, 0x0E, 0x94, 0x08, 0xD8, 0xF7, 0x86, 0xB6, 0x20, 0xDD, 0x48, 0xE4, 0x68, 0x86, 0x16, 0x75,
		0x36, 0x23, 0xE2, 0x95, 0x60, 0x08, 0x3C, 0x8D, 0x18, 0x93, 0x10, 0xCA, 0x3A, 0x2A, 0x68, 0x3E,
		0xB5, 0x9F, 0x3F, 0xDA, 0x62, 0x1E, 0x65, 0x23, 0xC4, 0x3D, 0xAB, 0x11, 0x20, 0xD9, 0xE5, 0x27,
		0x5A, 0x9D, 0x0E, 0xA1, 0x94, 0x8A, 0x40, 0xDE, 0x68, 0xFC, 0x07, 0xBE, 0xB3, 0x81, 0xE1, 0x1E,
		0x17, 0x1E, 0x9D, 0x68, 0x10, 0x6E, 0x0E, 0xDF, 0x16, 0xE5, 0x5E, 0xC3, 0x9D, 0xF4, 0xFD, 0x38,
		0x62, 0x14, 0x72, 0x95, 0xBF, 0xCD, 0x46, 0xB7, 0xD7, 0xC4, 0x18, 0x7A, 0xDE, 0xB1, 0x91, 0x72,
		0x63, 0xDA, 0x0E, 0xF9, 0x04, 0x5C, 0x11, 0x63, 0x4F, 0x84, 0x90, 0xE8, 0x3C, 0x51, 0x93, 0xCC,
		0x74, 0xF8, 0x1D, 0x98, 0x99, 0x4B, 0xF4, 0xEE, 0xF8, 0xFB, 0x61, 0xAE, 0x96, 0xF2, 0x9E, 0x18,
		0x9D, 0x6E, 0x1A, 0x42, 0xC2, 0x8B, 0x85, 0x01, 0x8E, 0xE6, 0xB0, 0xA4, 0x7C, 0xC8, 0x41, 0xF8,
		0x00, 0xA9, 0xCB, 0x2A, 0x6A, 0x8D, 0xD0, 0xD1, 0x0E, 0x2D, 0x24, 0x67, 0x2B, 0x0E, 0xF1, 0xCA,
		0x48, 0xB9, 0x00, 0xDB, 0xA5, 0x15, 0x51, 0x94, 0x91, 0xE6, 0xE7, 0x7A, 0x78, 0xE4, 0xFC, 0x58,
		0xB4, 0x6F, 0x5F, 0x5E, 0xE5, 0xD1, 0x94, 0xA5, 0xEC, 0x47, 0x31, 0xC3, 0x4F, 0xAF, 0x08, 0x6D,
		0xFD, 0x14, 0x24, 0xFD, 0x7E, 0x08, 0xC5, 0xC0, 0x35, 0x1A, 0x48, 0xA3, 0xE6, 0x14, 0xAE, 0x86,
		0xE0, 0x8B, 0xF7, 0x67, 0xEA, 0x06, 0x6A, 0x12, 0x8F, 0x17, 0x82, 0xF5, 0x0E, 0x20, 0x79, 0x90,
		0x0E, 0xD0, 0x0E, 0x6A, 0x7F, 0x6A, 0x05, 0x4D, 0xC1, 0x17, 0x06, 0xBD, 0x6A, 0xCA, 0xA4, 0x29,
		0xD0, 0x91, 0xBA, 0x9A, 0xCC, 0xD6, 0x2F, 0x6B, 0x12, 0x39, 0xD1, 0x69, 0x18, 0x00, 0x2A, 0xF0,
		0x22, 0xC3, 0x1F, 0x26, 0x71, 0x72, 0x7E, 0x95, 0xE1, 0xC6, 0xC6, 0xC3, 0xD5, 0x69, 0xD9, 0xDF,
		0x3C, 0x8E, 0x6E, 0xCC, 0x7E, 0x0E, 0xD7, 0xA3, 0x24, 0x77, 0x6C, 0x2B, 0x1F, 0x1A, 0xFE, 0x43,
		0xD4, 0x50, 0x75, 0x21, 0x81, 0x1A, 0x11, 0xCF, 0x8D, 0xEC, 0x59, 0x8E, 0x8A, 0xA3, 0x9D, 0x1C,
		0xD3, 0xA4, 0x21, 0xBE, 0x50, 0xB7, 0xA5, 0x89, 0x8E, 0x6C, 0xA2, 0xD7, 0xEB, 0x4F, 0xE2, 0x2C,
		0x18, 0x65, 0xB1, 0xE4, 0x2A, 0x5C, 0x2B, 0x1B, 0x85, 0xF1, 0x27, 0xFC, 0x7A, 0x55, 0xC9, 0xAB,
		0xDA, 0x8B, 0xBE, 0xFF, 0x00, 0x28, 0xEA, 0x68, 0x4F, 0x8C, 0xF1, 0x39, 0xF2, 0xA7, 0x24, 0xAF,
		0x5F, 0x88, 0x74, 0xF7, 0x14, 0x4C, 0x4A, 0x5F, 0x0A, 0x4D, 0xD9, 0x06, 0xF1, 0xFB, 0x7A, 0x52,
		0xBC, 0x6C, 0x19, 0x5B, 0x62, 0x38, 0x2D, 0xF6, 0x89, 0x6F, 0xF5, 0x3C, 0x2E, 0x38, 0x61, 0x26,
		0x5D, 0xD5, 0xF2, 0x9F, 0x63, 0xC3, 0xEF, 0x1B, 0xC4, 0x76, 0x51, 0xB9, 0xA5, 0x9F, 0x19, 0xA6,
		0x53, 0x75, 0x8F, 0x90, 0xA6, 0x8E, 0x4D, 0x64, 0x1A, 0xB1, 0x1B, 0x0E, 0x83, 0xE3, 0x99, 0x30,
		0x84, 0x29, 0x3A, 0xB2, 0x1F, 0x29, 0xFD, 0xAB, 0x2F, 0x92, 0x41, 0xBA, 0x36, 0xF4, 0x5F, 0x99,
		0x16, 0xE2, 0xCE, 0x10, 0xB1, 0x1F, 0x28, 0xE7, 0x50, 0x19, 0xE2, 0x11, 0x47, 0x19, 0xCD, 0x6B,
		0xEA, 0x4D, 0x2A, 0x20, 0xCF, 0x33, 0xE8, 0xAA, 0x3F, 0xCD, 0x7D, 0xA3, 0x12, 0x7B, 0x4C, 0x41,
		0xE6, 0x79, 0x7B, 0x70, 0x85, 0xC0, 0x06, 0xC7, 0x99, 0xB0, 0x1E, 0xBF, 0x1F, 0x5A, 0x19, 0xAE,
		0x08, 0xD9, 0x97, 0x42, 0x2A, 0xD3, 0x46, 0x71, 0x31, 0xFE, 0x78, 0xF7, 0x1E, 0xE2, 0xBF, 0x18,
		0x29, 0xE8, 0xD5, 0xA4, 0xA8, 0x7E, 0xB5, 0xE3, 0x9D, 0x05, 0x65, 0xC1, 0x42, 0x4D, 0xFF, 0x00,
		0xA8, 0xFA, 0x28, 0xFD, 0xEB, 0x3B, 0xCC, 0xC7, 0x13, 0x7B, 0xF6, 0x9D, 0x3F, 0xF9, 0x41, 0x65,
		0xC3, 0x76, 0xEA, 0x3E, 0x68, 0x7F, 0x6A, 0xF1, 0x45, 0x32, 0x9F, 0x54, 0xA4, 0xEC, 0xF0, 0x8D,
		0xD9, 0x83, 0xE7, 0x97, 0x41, 0xFA, 0x55, 0xBF, 0x90, 0xDA, 0xBC, 0x71, 0x29, 0xF7, 0xAF, 0xC2,
		0xB7, 0xB5, 0x5C, 0x42, 0xB7, 0xF5, 0xAB, 0x01, 0x61, 0xC3, 0x6A, 0xDA, 0xB4, 0xFE, 0xF1, 0xFF,
		0xC4, 0x00, 0x27, 0x10, 0x00, 0x01, 0x02, 0x05, 0x03, 0x04, 0x03, 0x01, 0x01, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x11, 0x10, 0x20, 0x21, 0x30, 0x41, 0x31, 0x40, 0x51, 0x61,
		0x71, 0x81, 0xA1, 0x50, 0xB1, 0xF1, 0xD1, 0xE1, 0xFF, 0xDA, 0x00, 0x08, 0x01, 0x01, 0x00, 0x01,
		0x3F, 0x21, 0xF8, 0x11, 0xAF, 0x70, 0x01, 0xC0, 0x02, 0x21, 0xBB, 0x0C, 0x01, 0x2A, 0x11, 0xFF,
		0x00, 0x42, 0x03, 0x66, 0x09, 0xA6, 0xC4, 0x0A, 0x00, 0x97, 0xF1, 0xC6, 0x14, 0x07, 0x27, 0xC7,
		0x42, 0x18, 0x81, 0x11, 0x94, 0x61, 0x1E, 0x60, 0xC9, 0x51, 0xD2, 0x08, 0x48, 0xE0, 0xC4, 0xA3,
		0x57, 0x40, 0x80, 0xCC, 0x84, 0x0B, 0x22, 0x94, 0x21, 0x01, 0x43, 0x36, 0x0F, 0x23, 0xA0, 0x30,
		0xC9, 0x0D, 0x60, 0x48, 0x0C, 0x38, 0x92, 0x0D, 0xA0, 0x15, 0xD2, 0x98, 0x01, 0x1C, 0xC0, 0x00,
		0xD0, 0x54, 0x24, 0x90, 0x3D, 0xA6, 0x00, 0x9D, 0xD8, 0x08, 0xCC, 0x45, 0x3C, 0x1C, 0x38, 0x5E,
		0x83, 0x28, 0x05, 0x38, 0x03, 0x0B, 0x24, 0x21, 0x25, 0x06, 0x13, 0x22, 0x00, 0x48, 0xE6, 0x55,
		0xC0, 0x0A, 0x20, 0x00, 0x26, 0x5E, 0x40, 0x00, 0x32, 0x00, 0xC1, 0x75, 0xA4, 0x86, 0xC4, 0x01,
		0xE2, 0x20, 0xD3, 0x36, 0xD8, 0x01, 0x03, 0x14, 0x28, 0x7F, 0x14, 0xB9, 0x20, 0xCB, 0xA5, 0x80,
		0x05, 0xF4, 0x0A, 0x04, 0xE6, 0x70, 0x21, 0x1E, 0x50, 0x02, 0x0C, 0x58, 0x20, 0xAA, 0x51, 0x81,
		0xC4, 0x0A, 0x0F, 0xF0, 0x60, 0x4C, 0x18, 0x22, 0x80, 0x48, 0x30, 0x80, 0x81, 0xD9, 0x62, 0x12,
		0x4E, 0x28, 0xF4, 0x60, 0x6A, 0x82, 0x7F, 0x3C, 0x91, 0xF2, 0xB0, 0x60, 0x81, 0x01, 0x02, 0x06,
		0xFB, 0x06, 0x23, 0x5F, 0x2A, 0xBA, 0x64, 0x67, 0x09, 0x64, 0xF4, 0x00, 0x80, 0x68, 0x01, 0x0C,
		0x58, 0x00, 0x50, 0x06, 0xB1, 0x80, 0x48, 0x27, 0x3D, 0x86, 0x08, 0x15, 0x2F, 0x59, 0x2E, 0x9E,
		0x94, 0x0A, 0xBB, 0x81, 0x32, 0x00, 0xA9, 0xE0, 0xA1, 0xA5, 0x4B, 0x27, 0x01, 0x54, 0x48, 0x4E,
		0x6E, 0x52, 0x07, 0xC8, 0x7E, 0x8A, 0xA6, 0xBA, 0x68, 0x02, 0x19, 0x0C, 0x8E, 0x61, 0x09, 0xC5,
		0x00, 0x80, 0x3E, 0x49, 0x00, 0x75, 0x12, 0x03, 0x1F, 0xA5, 0x4D, 0x87, 0x65, 0x9B, 0x0B, 0x80,
		0x88, 0x47, 0x48, 0x05, 0xC0, 0x17, 0x0C, 0x00, 0x0F, 0xA8, 0x21, 0xC4, 0x65, 0x50, 0x35, 0xE2,
		0x86, 0xA0, 0x65, 0x66, 0x03, 0x69, 0x52, 0x01, 0xC8, 0x80, 0x62, 0x01, 0xD6, 0x52, 0x76, 0xD6,
		0xC4, 0x05, 0x03, 0x20, 0x1E, 0xD4, 0x72, 0x20, 0x01, 0x14, 0x48, 0x0F, 0xDF, 0xD8, 0x3F, 0x93,
		0xE0, 0x1A, 0xA8, 0xBE, 0x40, 0x0C, 0x32, 0x60, 0x21, 0xD0, 0xF0, 0x12, 0x00, 0x66, 0x9E, 0x30,
		0x44, 0x2B, 0x21, 0x24, 0x09, 0xFF, 0x00, 0x88, 0x69, 0xF0, 0x7D, 0x97, 0x92, 0x10, 0x80, 0x10,
		0x00, 0x00, 0x05, 0x00, 0x0B, 0x14, 0x36, 0x80, 0x83, 0x11, 0xB0, 0x20, 0xEF, 0x89, 0x5B, 0x2F,
		0x81, 0x3E, 0x25, 0x4D, 0x00, 0x48, 0x6D, 0xB8, 0x78, 0x38, 0xEC, 0x66, 0xAC, 0x00, 0x82, 0xE4,
		0x8F, 0x6B, 0x20, 0x41, 0xA6, 0x80, 0x01, 0x62, 0x40, 0xB5, 0x57, 0x0A, 0x80, 0x1F, 0xA0, 0x2D,
		0x67, 0x00, 0x96, 0x9E, 0xC3, 0x78, 0x96, 0x28, 0x92, 0x82, 0x30, 0x88, 0x30, 0x09, 0x00, 0x00,
		0x0E, 0x20, 0x04, 0x03, 0xE4, 0x71, 0x41, 0x90, 0x07, 0xD2, 0x37, 0xD0, 0xBD, 0x80, 0x84, 0x33,
		0x88, 0x60, 0x80, 0x1E, 0x3E, 0x48, 0x04, 0x8F, 0x11, 0xA3, 0x35, 0x88, 0x28, 0x44, 0x0A, 0x06,
		0xA9, 0x31, 0x10, 0x03, 0xED, 0xA4, 0x6C, 0x8D, 0x20, 0x38, 0x80, 0x67, 0x30, 0x12, 0xE8, 0x41,
		0x99, 0x31, 0x44, 0x56, 0x11, 0x0C, 0x29, 0x81, 0x4A, 0xD8, 0xC4, 0x6E, 0x80, 0x99, 0x5C, 0x02,
		0x4F, 0xE4, 0xFC, 0x63, 0x23, 0x40, 0xC1, 0x6A, 0x21, 0x52, 0xB5, 0x2A, 0x20, 0x59, 0x11, 0x60,
		0xF9, 0x08, 0xDB, 0x02, 0x32, 0x54, 0x1B, 0x02, 0x02, 0x09, 0xE8, 0x14, 0x28, 0x23, 0xB2, 0x25,
		0x08, 0x40, 0xF6, 0x8B, 0x5D, 0xAB, 0xC5, 0x07, 0xE5, 0xC8, 0x24, 0xE6, 0xC5, 0x2A, 0xA3, 0x23,
		0xB1, 0x08, 0x84, 0xE2, 0x57, 0x6C, 0x91, 0x02, 0x3D, 0x29, 0x96, 0x09, 0xDE, 0x45, 0x7D, 0x81,
		0xF3, 0xB0, 0x01, 0x12, 0x3F, 0xFF, 0xDA, 0x00, 0x0C, 0x03, 0x01, 0x00, 0x02, 0x00, 0x03, 0x00,
		0x00, 0x00, 0x10, 0xF3, 0xCF, 0x3C, 0xF3, 0xCF, 0x3C, 0xF3, 0xCF, 0x3C, 0xF3, 0xCE, 0x3C, 0xA3,
		0x8F, 0xBC, 0xF3, 0x8F, 0x3C, 0xF3, 0xCB, 0x3C, 0xA9, 0xF3, 0x9E, 0x53, 0xC1, 0x28, 0xF3, 0xCF,
		0x24, 0x1B, 0xEE, 0x24, 0x63, 0x0F, 0x38, 0xE3, 0xCF, 0x3C, 0x61, 0x4B, 0x6C, 0x73, 0xC6, 0x1C,
		0xE1, 0x8F, 0x3C, 0xF3, 0x46, 0x34, 0x73, 0xCF, 0x3C, 0x63, 0x47, 0x1E, 0xF3, 0xCD, 0x2C, 0xF3,
		0xCF, 0x3C, 0xF3, 0xCB, 0x1C, 0x5B, 0xCF, 0x1C, 0xB1, 0x8F, 0x3C, 0xF3, 0xCF, 0x3C, 0xA2, 0x07,
		0x3C, 0xF3, 0xCA, 0x16, 0xF3, 0xC2, 0xD0, 0x83, 0x67, 0x1C, 0xF3, 0xCF, 0x3C, 0xA5, 0x8A, 0x28,
		0x12, 0xE6, 0x7C, 0xF3, 0xCF, 0x3C, 0xF3, 0xCA, 0x3C, 0xC3, 0x87, 0xE4, 0x53, 0xCF, 0x3C, 0xF3,
		0xCF, 0x3C, 0xB1, 0xCB, 0x1C, 0x33, 0xCF, 0x3C, 0xF3, 0xCF, 0x3C, 0xF3, 0xCF, 0x3C, 0xF3, 0xCF,
		0x3C, 0xF3, 0xFF, 0xC4, 0x00, 0x14, 0x11, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x70, 0xFF, 0xDA, 0x00, 0x08, 0x01, 0x03, 0x01, 0x01,
		0x3F, 0x10, 0x29, 0xFF, 0xC4, 0x00, 0x18, 0x11, 0x01, 0x00, 0x03, 0x01, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x11, 0x00, 0x40, 0x50, 0x20, 0xFF, 0xDA, 0x00,
		0x08, 0x01, 0x02, 0x01, 0x01, 0x3F, 0x10, 0xA4, 0xD6, 0xB1, 0xC2, 0x70, 0x1E, 0x8E, 0xC7, 0xFF,
		0xC4, 0x00, 0x28, 0x10, 0x00, 0x01, 0x02, 0x05, 0x04, 0x01, 0x04, 0x03, 0x01, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x11, 0x10, 0x20, 0x21, 0x30, 0x31, 0x41, 0x51, 0x61, 0x71,
		0x40, 0x50, 0x81, 0xC1, 0xF0, 0x91, 0xA1, 0xB1, 0xE1, 0xFF, 0xDA, 0x00, 0x08, 0x01, 0x01, 0x00,
		0x01, 0x3F, 0x10, 0xF4, 0x13, 0x1B, 0xE4, 0x70, 0x14, 0x90, 0x10, 0x10, 0xF9, 0x81, 0x00, 0x69,
		0x20, 0x02, 0x09, 0xAF, 0xE1, 0x95, 0x31, 0xAB, 0x69, 0x1A, 0x39, 0x87, 0xDE, 0xE4, 0x70, 0x88,
		0xEC, 0x82, 0x4D, 0x0C, 0x9A, 0x84, 0x50, 0xE5, 0x57, 0x20, 0x01, 0x02, 0x80, 0x01, 0x53, 0xBD,
		0xE6, 0x00, 0x3F, 0x6A, 0xD2, 0x30, 0x8C, 0xA2, 0xEC, 0x35, 0x4E, 0xC0, 0x6E, 0x82, 0x55, 0x20,
		0x58, 0xEB, 0x21, 0x91, 0x01, 0xA8, 0x09, 0x1E, 0x6A, 0x10, 0x80, 0x48, 0x37, 0x4A, 0x02, 0x40,
		0x83, 0x94, 0x08, 0x21, 0x00, 0xA0, 0xD2, 0x7E, 0xCE, 0x91, 0x34, 0x5F, 0x18, 0x04, 0x80, 0x19,
		0x54, 0x33, 0x4C, 0x7F, 0x87, 0xB6, 0xD0, 0x84, 0xCA, 0x92, 0x01, 0x81, 0x0D, 0x8E, 0x28, 0x41,
		0x52, 0x91, 0xA6, 0x10, 0x64, 0xA0, 0x46, 0x03, 0x5A, 0x94, 0x33, 0x12, 0x08, 0xA0, 0x03, 0x85,
		0x92, 0x41, 0x50, 0x23, 0x21, 0x80, 0x8D, 0xB0, 0xD3, 0x61, 0xCC, 0xC9, 0xE2, 0x30, 0xAC, 0x54,
		0xC7, 0xDC, 0xB2, 0x1C, 0xC0, 0x18, 0x2D, 0xE5, 0x8B, 0x20, 0xA5, 0x28, 0x0D, 0xD0, 0x74, 0x95,
		0x28, 0xF4, 0x93, 0x7A, 0x03, 0x0D, 0x76, 0x34, 0x8E, 0x71, 0xC9, 0x0A, 0xA6, 0x1D, 0x4D, 0x80,
		0x01, 0x00, 0x00, 0xE4, 0xB4, 0x44, 0x60, 0x01, 0x92, 0xDD, 0xF1, 0x38, 0x45, 0x1D, 0x00, 0xE4,
		0x93, 0x85, 0x44, 0x76, 0x2C, 0x0D, 0x0D, 0xE6, 0x02, 0x7B, 0x7B, 0x5D, 0x82, 0xCC, 0xDC, 0xE3,
		0x01, 0xA0, 0x69, 0xCA, 0x0B, 0x26, 0xD0, 0xCC, 0x93, 0x50, 0x20, 0x7A, 0x1A, 0xA0, 0x65, 0x0E,
		0xD0, 0x91, 0x14, 0xA0, 0x91, 0x06, 0x14, 0x80, 0x20, 0x48, 0x90, 0x3C, 0xC3, 0x84, 0x46, 0x01,
		0xF5, 0x1B, 0x43, 0x30, 0x46, 0x12, 0x0B, 0x30, 0x01, 0xAC, 0x90, 0xFC, 0x00, 0x00, 0x0A, 0x22,
		0x64, 0x68, 0xDE, 0x14, 0x49, 0x60, 0x80, 0x46, 0xF8, 0x18, 0x0F, 0x92, 0xA8, 0x13, 0xEF, 0xD0,
		0x84, 0x35, 0x2F, 0xB6, 0x34, 0x86, 0x4F, 0xE8, 0x10, 0xE8, 0x0C, 0x8E, 0x98, 0xB9, 0xBD, 0x02,
		0x42, 0x81, 0x36, 0x48, 0x8A, 0xB0, 0x29, 0x13, 0xA6, 0x68, 0x80, 0xE3, 0x18, 0x08, 0x1E, 0x92,
		0xC0, 0xA8, 0xE0, 0x78, 0x80, 0xA5, 0x18, 0x10, 0x70, 0x23, 0xE6, 0xC5, 0x52, 0x35, 0x66, 0xE1,
		0x48, 0x52, 0x8C, 0x04, 0xE8, 0x60, 0xFD, 0x44, 0xB6, 0x76, 0x20, 0x15, 0xEC, 0x66, 0x70, 0x2B,
		0xA1, 0xC0, 0x37, 0xB1, 0x80, 0xD7, 0xBE, 0xD8, 0x21, 0xA2, 0x23, 0x44, 0x90, 0x1A, 0x10, 0x02,
		0x5E, 0x22, 0x20, 0xA8, 0x01, 0x2E, 0xA8, 0x41, 0xF7, 0xC7, 0x02, 0x84, 0xE1, 0x70, 0x8A, 0x68,
		0x8F, 0x9A, 0x88, 0xA7, 0xC5, 0xC8, 0xE8, 0xC8, 0x98, 0x27, 0x29, 0xC0, 0x20, 0x44, 0x0A, 0xD5,
		0x08, 0x53, 0x60, 0x7B, 0xE8, 0x33, 0x6A, 0x80, 0x28, 0x0C, 0xDD, 0x02, 0x35, 0xE5, 0x05, 0x0E,
		0x41, 0x51, 0xE1, 0x0C, 0x02, 0x10, 0x46, 0x82, 0xE4, 0x00, 0x1C, 0x01, 0x9D, 0x88, 0xB7, 0x23,
		0x27, 0x0F, 0x53, 0x14, 0x37, 0xEC, 0x61, 0x2C, 0x44, 0x14, 0x43, 0x22, 0xAC, 0xF0, 0x60, 0x40,
		0x45, 0xA8, 0x05, 0x83, 0x9A, 0xC3, 0xFB, 0x9A, 0x88, 0x98, 0x67, 0x60, 0x0C, 0xF2, 0x03, 0xCC,
		0x10, 0xD5, 0x08, 0x1A, 0x6B, 0xAD, 0xE2, 0xEC, 0x82, 0xC0, 0x81, 0x00, 0x91, 0x36, 0xF1, 0x2A,
		0x2D, 0x08, 0x2E, 0x6A, 0x66, 0x52, 0xAD, 0x30, 0x29, 0xC4, 0xFD, 0x10, 0x6B, 0x80, 0xF7, 0x31,
		0x19, 0x49, 0x02, 0xC0, 0xED, 0x75, 0xAE, 0xF8, 0x41, 0x00, 0x78, 0x83, 0x18, 0x09, 0x50, 0x3E,
		0x12, 0x3B, 0xC1, 0x1F, 0x85, 0x68, 0x88, 0x2E, 0xA8, 0x31, 0x1A, 0x50, 0x01, 0x75, 0x06, 0x45,
		0x53, 0x45, 0x41, 0xA9, 0x45, 0x00, 0xC1, 0x97, 0x88, 0x03, 0x84, 0x17, 0xA1, 0x64, 0x18, 0x3B,
		0x50, 0x90, 0x0B, 0x49, 0x22, 0xC3, 0x15, 0x60, 0x26, 0x48, 0x2E, 0xC7, 0xFC, 0x87, 0xE0, 0x86,
		0xF1, 0xD0, 0xF4, 0x89, 0x5C, 0x08, 0x6B, 0xBC, 0x41, 0x28, 0x94, 0x80, 0x99, 0x1B, 0x1E, 0x17,
		0x0A, 0x38, 0x88, 0x72, 0xAD, 0xA7, 0x63, 0x5F, 0x66, 0x10, 0xCE, 0x9F, 0xBA, 0x90, 0x29, 0x30,
		0x99, 0xA9, 0x3E, 0xD0, 0x96, 0x0B, 0xF4, 0xFA, 0x6A, 0x0B, 0xFE, 0x8F, 0x02, 0x1A, 0xA3, 0x0A,
		0x27, 0xFA, 0x19, 0x95, 0x8C, 0x62, 0x08, 0x38, 0x1E, 0xB2, 0xBF, 0xFF, 0xD9}
)
