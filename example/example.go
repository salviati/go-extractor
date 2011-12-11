package main

import(
	"fmt"
	"extractor"
	"flag"
	"log"
)

var filename = flag.String("f", "", "Input file")

func main() {
	flag.Parse()

	if *filename == "" { log.Fatal("Need an input file.") }

	ex := extractor.New(extractor.OPTION_DEFAULT_POLICY)
	meta := ex.File(*filename)

	for _, m := range meta {
		fmt.Println("mime type:", string(m.DataMimeType))
		fmt.Println("meta type:", extractor.MetaTypeToString(m.Type))
		fmt.Println("plugin name:", string(m.PluginName))
		if m.Format == extractor.METAFORMAT_C_STRING ||
			m.Format == extractor.METAFORMAT_UTF8 {
			fmt.Println("data:", string(m.Data))
		}
		fmt.Println("")
	}
}
