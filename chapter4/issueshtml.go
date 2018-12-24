package main

import "html/template"
//import "time"
import "gotest/chapter4/github"
import "os"
import "log"

var templ = `<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='textalign: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Items}}
<tr>
<td><a href='{{.HTMLURL}}'>{{.Number}}</td>
<td>{{.State}}</td>
<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>`

/*
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours()/24)
}
*/

var report = template.Must(template.New("issuelist").Parse(templ))

func main() {
	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}