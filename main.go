// +build js,wasm

package main

import (
	"encoding/csv"

	"syscall/js"
	"fmt"
	"strings"
	"bytes"
	"unsafe"
	"reflect"

	"github.com/tealeg/xlsx"
)

func generateXLSXFromCSV(csvContent string, delimiter string) (error, bytes.Buffer) {
	var buf bytes.Buffer

	reader := csv.NewReader(strings.NewReader(csvContent))
	if len(delimiter) > 0 {
		reader.Comma = rune(delimiter[0])
	} else {
		reader.Comma = rune(';')
	}
	xlsxFile := xlsx.NewFile()
	sheet, xlxsErr := xlsxFile.AddSheet("sheet1")
	if xlxsErr != nil {
		return xlxsErr, buf
	}
	fields, readerErr := reader.Read()
	for readerErr == nil {
		row := sheet.AddRow()
		for _, field := range fields {
			cell := row.AddCell()
			cell.Value = field
		}
		fields, readerErr = reader.Read()
	}
	if readerErr != nil {
		fmt.Printf(readerErr.Error())
	}

	
	err := xlsxFile.Write(&buf)
	return err, buf
}

func main() {
	fileInput := js.Global().Get("dataFile").String()
	err, val := generateXLSXFromCSV(fileInput, ";")
	if err != nil {
		fmt.Println("err: ",err)
	}

	out := val.Bytes()
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&out))
	ptr := uintptr(unsafe.Pointer(hdr.Data))
	js.Global().Call("makeExcelFile", ptr, len(out))
}
