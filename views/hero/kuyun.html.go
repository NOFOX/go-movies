// Code generated by hero.
// source: C:\E\project\golang\src\github_s\go_movies\views\hero\kuyun.html
// DO NOT EDIT!
package template

import (
	"bytes"

	"github.com/shiyanhui/hero"
)

func Kuyun(show map[string]interface{}, buffer *bytes.Buffer) {
	buffer.WriteString(`

<iframe src="`)
	hero.EscapeHTML(show["play_url"].(string), buffer)
	buffer.WriteString(`" allowfullscreen="true" webkitallowfullscreen="true" mozallowfullscreen="true" marginwidth="0" marginheight="0" border="0" scrolling="no" frameborder="0" topmargin="0" width="100%" height="100%"></iframe>
`)

}