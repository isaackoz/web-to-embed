package generateFile

import (
	"html/template"
	"os"
	"path/filepath"

	"github.com/isaackoz/web-to-embed/convert"
)

type templateData struct {
	Assets         []convert.Asset
	IncludeProgmem bool
}

type Options struct {
	Progmem   bool
	OutputDir string
}

func Generate(assets []convert.Asset, options Options) error {

	err := os.MkdirAll(options.OutputDir, os.ModePerm)
	if err != nil {
		return err
	}
	tmpl, err := template.New("cppHeader").Parse(cppHeaderTemplate)
	if err != nil {
		return err
	}

	outputPath := filepath.Join(options.OutputDir, "output.h")

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	data := templateData{
		Assets:         assets,
		IncludeProgmem: options.Progmem,
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}

	return nil
}

const cppHeaderTemplate = `#pragma once
namespace static_files {

    struct file {
        const char *path;
        uint32_t size;
        const char *type;
        const uint8_t *contents;
    };

    {{- range .Assets }}
    {{- if $.IncludeProgmem }}
    const uint32_t f_{{ .NormalizedName }}_size PROGMEM = {{ .Size }};
    const uint8_t f_{{ .NormalizedName }}_contents[] PROGMEM = {
    {{- else }}
    const uint32_t f_{{ .NormalizedName }}_size = {{ .Size }};
    const uint8_t f_{{ .NormalizedName }}_contents[] = {
    {{- end }}
        {{ .Contents}}
    };
    {{- end }}

    const file files[] {{- if $.IncludeProgmem}} PROGMEM {{- end}} = {
        {{- range .Assets }}
        { .path = "{{ .Path }}", .size = f_{{ .NormalizedName }}_size, .type = "{{ .MimeType }}", .contents = f_{{ .NormalizedName }}_contents },
        {{- end }}
    };

    const uint8_t num_of_files = sizeof(files) / sizeof(file);
}
`
