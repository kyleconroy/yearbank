package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Println("Listening on :3001...")
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func slugify(in string) string {
	in = strings.Replace(in, " ", "-", -1)
	in = strings.Replace(in, ".", "", -1)
	in = strings.Replace(in, ":", "", -1)
	in = strings.Replace(in, "'", "", -1)
	return strings.ToLower(in)
}

type sectionData struct {
	Photos []string
	Name   string
}

func run() error {
	sections := map[string]string{
		"31": "Class of '07",
		"52": "Class of '08",
		"53": "Class of '09",
		"54": "Class of '10",

		"1":  "Spirit Week",
		"2":  "Style",
		"3":  "Zen Garden",
		"4":  "Costumes",
		"5":  "Homecoming",
		"6":  "U.S. History",
		"8":  "Art",
		"9":  "Air Band",
		"10": "Eating Contest",
		"11": "Book Club",
		"12": "Ally",
		"13": "Around Campus",
		"14": "Cooking",
		"15": "Cool",
		"16": "Theater",
		"17": "Special Education",
		"18": "Digital Music",
		"19": "Rotary Youth Exchange",
		"20": "RYLA",
		"21": "Unknown: A",
		"22": "Model U.N.",
		"23": "Yearbook",
		"24": "Key Club",
		"25": "National Honor Society",
		"26": "Unknown: B",
		"27": "S.M.I.L.E",
		"28": "French",
		"29": "Vegan",
		"30": "Kids",
		"32": "Teachers and Staff",
		"36": "Dance Team",
		"37": "Football",
		"38": "Golf",
		"39": "ROTC",
		"40": "Ski Team",
		"41": "Nordic Skiing",
		"42": "Soccer",
		"43": "Softball",
		"44": "Swimming",
		"45": "Tennis",
		"46": "Track & Field",
		"47": "Volleyball",
		"48": "Wrestling",
		"49": "Cross Country",
		"50": "Choir",
		"51": "Band",
	}

	in := "images"

	sectionTmpl := template.Must(template.New("").Parse(section))

	files, err := ioutil.ReadDir(in)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		images, err := ioutil.ReadDir(filepath.Join(in, file.Name()))
		if err != nil {
			return err
		}
		name := file.Name()
		if name == "7" {
			continue
		}
		if name == "33" {
			continue
		}
		if name == "34" {
			continue
		}
		if name == "35" {
			continue
		}
		if rename, ok := sections[file.Name()]; ok {
			name = rename
		}
		data := sectionData{
			Photos: []string{},
			Name:   name,
		}
		for _, image := range images {
			if !strings.HasSuffix(strings.ToLower(image.Name()), ".jpg") {
				continue
			}
			data.Photos = append(data.Photos, filepath.Join(file.Name(), image.Name()))
		}

		sectionDir := filepath.Join("sections", slugify(name))
		sectionHTML := filepath.Join(sectionDir, "index.html")
		os.MkdirAll(sectionDir, 0700)
		f, err := os.Create(sectionHTML)
		if err != nil {
			return err
		}
		if err := sectionTmpl.Execute(f, data); err != nil {
			return err
		}
		f.Close()
	}

	// Serve
	return nil
}

const section = `<!DOCTYPE html>
<html lang="en">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <style>
/*! normalize.css v8.0.1 | MIT License | github.com/necolas/normalize.css */

/* Document
   ========================================================================== */

/**
 * 1. Correct the line height in all browsers.
 * 2. Prevent adjustments of font size after orientation changes in iOS.
 */

html {
  line-height: 1.15; /* 1 */
  -webkit-text-size-adjust: 100%; /* 2 */
}

/* Sections
   ========================================================================== */

/**
 * Remove the margin in all browsers.
 */

body {
  margin: 0;
}

/**
 * Render the main element consistently in IE.
 */

main {
  display: block;
}

/**
 * Correct the font size and margin on h1 elements within section and
 * article contexts in Chrome, Firefox, and Safari.
 */

h1 {
  font-size: 2em;
  margin: 0.67em 0;
}

/* Grouping content
   ========================================================================== */

/**
 * 1. Add the correct box sizing in Firefox.
 * 2. Show the overflow in Edge and IE.
 */

hr {
  box-sizing: content-box; /* 1 */
  height: 0; /* 1 */
  overflow: visible; /* 2 */
}

/**
 * 1. Correct the inheritance and scaling of font size in all browsers.
 * 2. Correct the odd em font sizing in all browsers.
 */

pre {
  font-family: monospace, monospace; /* 1 */
  font-size: 1em; /* 2 */
}

/* Text-level semantics
   ========================================================================== */

/**
 * Remove the gray background on active links in IE 10.
 */

a {
  background-color: transparent;
}

/**
 * 1. Remove the bottom border in Chrome 57-
 * 2. Add the correct text decoration in Chrome, Edge, IE, Opera, and Safari.
 */

abbr[title] {
  border-bottom: none; /* 1 */
  text-decoration: underline; /* 2 */
  text-decoration: underline dotted; /* 2 */
}

/**
 * Add the correct font weight in Chrome, Edge, and Safari.
 */

b,
strong {
  font-weight: bolder;
}

/**
 * 1. Correct the inheritance and scaling of font size in all browsers.
 * 2. Correct the odd em font sizing in all browsers.
 */

code,
kbd,
samp {
  font-family: monospace, monospace; /* 1 */
  font-size: 1em; /* 2 */
}

/**
 * Add the correct font size in all browsers.
 */

small {
  font-size: 80%;
}

/**
 * Prevent sub and sup elements from affecting the line height in
 * all browsers.
 */

sub,
sup {
  font-size: 75%;
  line-height: 0;
  position: relative;
  vertical-align: baseline;
}

sub {
  bottom: -0.25em;
}

sup {
  top: -0.5em;
}

/* Embedded content
   ========================================================================== */

/**
 * Remove the border on images inside links in IE 10.
 */

img {
  border-style: none;
}

/* Forms
   ========================================================================== */

/**
 * 1. Change the font styles in all browsers.
 * 2. Remove the margin in Firefox and Safari.
 */

button,
input,
optgroup,
select,
textarea {
  font-family: inherit; /* 1 */
  font-size: 100%; /* 1 */
  line-height: 1.15; /* 1 */
  margin: 0; /* 2 */
}

/**
 * Show the overflow in IE.
 * 1. Show the overflow in Edge.
 */

button,
input { /* 1 */
  overflow: visible;
}

/**
 * Remove the inheritance of text transform in Edge, Firefox, and IE.
 * 1. Remove the inheritance of text transform in Firefox.
 */

button,
select { /* 1 */
  text-transform: none;
}

/**
 * Correct the inability to style clickable types in iOS and Safari.
 */

button,
[type="button"],
[type="reset"],
[type="submit"] {
  -webkit-appearance: button;
}

/**
 * Remove the inner border and padding in Firefox.
 */

button::-moz-focus-inner,
[type="button"]::-moz-focus-inner,
[type="reset"]::-moz-focus-inner,
[type="submit"]::-moz-focus-inner {
  border-style: none;
  padding: 0;
}

/**
 * Restore the focus styles unset by the previous rule.
 */

button:-moz-focusring,
[type="button"]:-moz-focusring,
[type="reset"]:-moz-focusring,
[type="submit"]:-moz-focusring {
  outline: 1px dotted ButtonText;
}

/**
 * Correct the padding in Firefox.
 */

fieldset {
  padding: 0.35em 0.75em 0.625em;
}

/**
 * 1. Correct the text wrapping in Edge and IE.
 * 2. Correct the color inheritance from fieldset elements in IE.
 * 3. Remove the padding so developers are not caught out when they zero out
 *    fieldset elements in all browsers.
 */

legend {
  box-sizing: border-box; /* 1 */
  color: inherit; /* 2 */
  display: table; /* 1 */
  max-width: 100%; /* 1 */
  padding: 0; /* 3 */
  white-space: normal; /* 1 */
}

/**
 * Add the correct vertical alignment in Chrome, Firefox, and Opera.
 */

progress {
  vertical-align: baseline;
}

/**
 * Remove the default vertical scrollbar in IE 10+.
 */

textarea {
  overflow: auto;
}

/**
 * 1. Add the correct box sizing in IE 10.
 * 2. Remove the padding in IE 10.
 */

[type="checkbox"],
[type="radio"] {
  box-sizing: border-box; /* 1 */
  padding: 0; /* 2 */
}

/**
 * Correct the cursor style of increment and decrement buttons in Chrome.
 */

[type="number"]::-webkit-inner-spin-button,
[type="number"]::-webkit-outer-spin-button {
  height: auto;
}

/**
 * 1. Correct the odd appearance in Chrome and Safari.
 * 2. Correct the outline style in Safari.
 */

[type="search"] {
  -webkit-appearance: textfield; /* 1 */
  outline-offset: -2px; /* 2 */
}

/**
 * Remove the inner padding in Chrome and Safari on macOS.
 */

[type="search"]::-webkit-search-decoration {
  -webkit-appearance: none;
}

/**
 * 1. Correct the inability to style clickable types in iOS and Safari.
 * 2. Change font properties to inherit in Safari.
 */

::-webkit-file-upload-button {
  -webkit-appearance: button; /* 1 */
  font: inherit; /* 2 */
}

/* Interactive
   ========================================================================== */

/*
 * Add the correct display in Edge, IE 10+, and Firefox.
 */

details {
  display: block;
}

/*
 * Add the correct display in all browsers.
 */

summary {
  display: list-item;
}

/* Misc
   ========================================================================== */

/**
 * Add the correct display in IE 10+.
 */

template {
  display: none;
}

/**
 * Add the correct display in IE 10.
 */

[hidden] {
  display: none;
}

body {
  background: rgb(28, 30, 33);
  color: rgb(228, 230, 235);
  font-family: -apple-system, BlinkMacSystemFont,
    "Segoe UI", "Roboto", "Oxygen",
    "Ubuntu", "Cantarell", "Fira Sans",
    "Droid Sans", "Helvetica Neue", sans-serif;
}

section {
  max-width: 1200px;
  margin: auto;
}

div.album {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr 1fr;
  column-gap: 15px;
  row-gap: 10px;
}

div.album a {
  align-self: center;
}

div.album a img {
  max-width: 100%;
  height: auto;
}
  </style>
  <body>
    <section>
      <h1><a href="/">STHS Yearbook 2006-07</a></h1>
      <h2>{{.Name}}</h2>
      <div class="album">
        {{range .Photos}}
        <a href="/images/{{.}}">
          <img src="/images/{{.}}" />
        </a>
        {{end}}
      </div>
    </section>
  </body>
</html>`
