package theme

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// type Element struct {
// 	Type       string                 `json:"type"`
// 	Properties map[string]string      `json:"properties"`
// 	Features   map[string]interface{} `json:"feature"`
// }

type Element interface{}

type Theme struct {
	Title      string            `json:"title"`
	Background map[string]string `json:"background"`
	Elements   []Element         `json:"elements"`
}

func ParseFile(filepath string) (theme Theme) {

	themeFile, err := os.Open(filepath)
	defer themeFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	themeData, _ := ioutil.ReadAll(themeFile)
	if err := json.Unmarshal(themeData, &theme); err != nil {
		fmt.Println(err)
		return
	}

	return
}

func (theme *Theme) genLayout() {

	var Boxes []Box
	for _, element := range theme.Elements {

		var box Box
		boxes = append(boxes, box.FromThemeElement(element))
	}
}
