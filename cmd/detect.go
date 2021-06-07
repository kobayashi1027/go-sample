// 参考: https://text.baldanders.info/golang/detecting-character-encoding/

package cmd

import (
	"fmt"

	"github.com/saintfish/chardet"
	"github.com/spf13/cobra"
)

// CharEncode is type of character encoding
type CharEncode int

const (
	// CharUnknown is unknown character
	CharUnknown CharEncode = iota
	// CharUTF8 is UTF-8
	CharUTF8
	// CharISO8859_1 is ISO-8859-1
	CharISO8859_1
	// CharShiftJIS is Shift-JIS
	CharShiftJIS
	// CharEUCJP is EUC-JP
	CharEUCJP
	// CharISO2022JP is ISO-2022-JP
	CharISO2022JP
)

// DetectCharEncode returns character encoding
func detectCharEncode(body []byte) CharEncode {
	det := chardet.NewTextDetector()
	res, err := det.DetectBest(body)
	if err != nil {
		return CharUnknown
	}
	return typeOfCharEncode(res.Charset)
}

func typeOfCharEncode(charset string) CharEncode {
	switch charset {
	case "UTF-8":
		return CharUTF8
	case "ISO-8859-1":
		return CharISO8859_1
	case "EUC-JP":
		return CharEUCJP
	case "ISO-2022-JP":
		return CharISO2022JP
	default:
		return CharUnknown
	}
}

// Main logic
func Detect(cmd *cobra.Command, args []string) {
	fmt.Println("detect is running")
	fmt.Println(args)
}
