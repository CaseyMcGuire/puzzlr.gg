package views

import (
	"fmt"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	"maragu.dev/gomponents/html"
)

func ReactPage(title string, bundleName string) Node {
	return HTML5(
		HTML5Props{
			Title:    title,
			Language: "en",
			Head: []Node{
				html.Meta(Attr("charset", "UTF-8")),
				html.Meta(
					Attr("name", "viewport"),
					Attr("content", "initial-scale=1.0, maximum-scale=1.0, width=device-width"),
				),
				html.Link(
					Attr("rel", "stylesheet"),
					Attr("href", fmt.Sprintf("/assets/bundles/%s.stylex.css", bundleName)),
				),
				html.Script(
					Attr("type", "importmap"),
					Raw(`{
              "imports": {
                "react": "https://esm.sh/react@18.3.1",
                "react-dom": "https://esm.sh/react-dom@18.3.1"
              }
            }`),
				),
				html.Style(
					`
              * {
                margin: 0;
                padding: 0;
                box-sizing: border-box;
              }
 
              html, body, #root {
                height: 100%;
              }
        `),
			},
			Body: []Node{
				html.Div(Attr("id", "root")),
				html.Script(
					Attr("src", fmt.Sprintf("/assets/bundles/%s.bundle.js", bundleName)),
					Attr("type", "module"),
				),
			},
		})
}
