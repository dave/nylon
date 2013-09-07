package store

import "strings"

func Get(id string) *strings.Reader {
	s := ""
	switch id {
	case "root":
		s = `
<html>
	<head></head>
	<body>
		<h1>Hello world</h1>
		<my-tag id="Root">
			<p>Here's a normal paragraph</p>
			<my-bio id="Me"></my-bio>
			<p>Here's another paragraph</p>
		</my-tag>
	</body>
</html>
`
	case "bio":
		s = `<span>This is my bio!</span><span>Here's a second span</span>`
	}
	return strings.NewReader(s)
}
