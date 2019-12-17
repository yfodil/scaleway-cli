package printer

import (
	"fmt"
	"io"
	"strings"

	"github.com/scaleway/scaleway-cli/internal/human"
)

// FormatterType defines an formatter format.
type PrinterType string

// String returns the formatter format converted in a string.
func (o *PrinterType) String() string {
	return string(*o)
}

// Set sets the formatter format from a string.
func (o *PrinterType) Set(v string) error {
	*o = PrinterType(v)
	return nil
}

// Type returns the FormatterType string type.
func (o *PrinterType) Type() string {
	return "PrinterType"
}

var (
	// JSON defines a JSON formatter.
	JSON = PrinterType("json")

	// Human defines a human readable formatted formatter.
	Human = PrinterType("human")
)

// Printer contains all the formatter logic for a given format.
type Printer interface {
	Print(data interface{}, opt *human.MarshalOpt) error
}

// New returns an initialized formatter corresponding to a given FormatterType.
func New(formatterType PrinterType, writer io.Writer, errorWriter io.Writer) (Printer, error) {
	if formatterType == "" {
		formatterType = Human
	}

	formatterType = PrinterType(strings.ToLower(string(formatterType)))

	switch formatterType {
	case JSON:
		return &jsonPrinter{
			Writer:      writer,
			ErrorWriter: errorWriter,
		}, nil
	case Human:
		return &humanPrinter{
			Writer:      writer,
			ErrorWriter: errorWriter,
		}, nil
	default:
		return nil, fmt.Errorf("invalid format: %s", formatterType)
	}
}
