package cyoa

import (
	"encoding/json"
	"io"
	"net/http"
)

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type Handler struct {
	s Story
}

func (h handler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	//tplate := template.Must(template.New("").Parse(handlerTempalte))

}

func JsonReader(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		panic(err)
	}
	return story, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

var handlerTempalte = `
<!DOCTYPE html>
<html>

<head>
    <meta charset='utf-8'>
    <title>Azazel's Adventure</title>
    <style>
        body {
            background-color: #f4f1ea;
            font-family: 'Georgia', serif;
            color: #333;
            line-height: 1.6;
            padding: 20px;
            max-width: 800px;
            margin: auto;
        }

        h1 {
            text-align: center;
            font-size: 3em;
            margin-bottom: 0.5em;
            color: #5D4157;
        }

        p {
            text-indent: 50px;
            margin-bottom: 1.5em;
            font-size: 1.2em;
        }

        ul {
            list-style-type: none;
            padding: 0;
        }

        li {
            margin-bottom: 1em;
        }

        a {
            text-decoration: none;
            background-color: #8FB9A8;
            color: white;
            padding: 10px 15px;
            border-radius: 5px;
            font-size: 1.1em;
            transition: background-color 0.3s;
        }

        a:hover {
            background-color: #82a393;
        }

        /* Add a subtle box-shadow to give a book-like feel */
        .story-container {
            background-color: #fff;
            padding: 40px;
            box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
        }
    </style>
</head>

<body>
    <div class="story-container">
        <h1>{{.Title}}</h1>
        {{range .Paragraphs}}
        <p>{{.}}</p>
        {{end}}
        <ul>
            {{range .Options}}
            <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
            {{end}}
        </ul>
    </div>
</body>

</html>`
