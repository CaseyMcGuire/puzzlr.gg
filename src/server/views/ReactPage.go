package views

import (
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
					Attr("href", "/assets/bundles/stylex.css"),
				),
				html.Script(
					Attr("type", "importmap"),
					Raw(`{
              "imports": {
                "react": "https://esm.sh/react@19.2.3",
                "react-dom": "https://esm.sh/react-dom@19.2.3",
                "react-dom/client": "https://esm.sh/react-dom@19.2.3/client",
                "react/jsx-runtime": "https://esm.sh/react@19.2.3/jsx-runtime",
                "react/jsx-dev-runtime": "https://esm.sh/react@19.2.3/jsx-dev-runtime"
              }
            }`),
				),
			},
			Body: []Node{
				html.Div(Attr("id", "root")),
				html.Script(
					Attr("src", "/assets/bundles/index.bundle.js"),
					Attr("type", "module"),
				),
			},
		})
}
