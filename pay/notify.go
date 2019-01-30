package pay

import (
	"bytes"
	"encoding/xml"
	"strings"
)

func Success() string {
	params := map[string]string{
		"return_code": "SUCCESS",
		"return_msg":  "ok",
	}
	return mapToXml(params)
}

func Fail(errMsg string) string {
	params := map[string]string{
		"return_code": "FAIL",
		"return_msg":  errMsg,
	}
	return mapToXml(params)
}

func mapToXml(params map[string]string) string {
	var buf bytes.Buffer
	buf.WriteString(`<xml>`)
	for k, v := range params {
		buf.WriteString(`<`)
		buf.WriteString(k)
		buf.WriteString(`><![CDATA[`)
		buf.WriteString(v)
		buf.WriteString(`]]></`)
		buf.WriteString(k)
		buf.WriteString(`>`)
	}
	buf.WriteString(`</xml>`)

	return buf.String()
}

func XmlToMap(data []byte) map[string]string {

	params := make(map[string]string)
	decoder := xml.NewDecoder(bytes.NewReader(data))

	var (
		key   string
		value string
	)
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement: // 开始标签
			key = token.Name.Local
		case xml.CharData: // 标签内容
			content := string([]byte(token))
			value = content
		}
		if key != "xml" && strings.TrimSpace(value) != ""  {
			params[key] = value
		}
	}
	return params
}
