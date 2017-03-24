// This file is automatically generated by qtc from "page.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line templates/page.qtpl:1
package templates

//line templates/page.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/page.qtpl:1
import "github.com/yanzay/teslo/templates"

//line templates/page.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/page.qtpl:3
func StreamPage(qw422016 *qt422016.Writer, state State) {
	//line templates/page.qtpl:3
	qw422016.N().S(`
<!DOCTYPE html>
<html lang="en">
<head>
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
</head>
<body id="app">
    <nav class="navbar navbar-inverse">
      <div class="container">
        <div class="navbar-header">
          <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="navbar-brand" href="#">Hugo Hosting</a>
        </div>
        <div id="navbar" class="collapse navbar-collapse">
          <ul class="nav navbar-nav">
            <li class="active"><a href="#">Home</a></li>
            <li><a href="#about">About</a></li>
            <li><a href="#contact">Contact</a></li>
          </ul>
        </div><!--/.nav-collapse -->
      </div>
    </nav>

    <div class="container">
      <div>
        `)
	//line templates/page.qtpl:33
	if state.Login == "" {
		//line templates/page.qtpl:33
		qw422016.N().S(`
          <a href="`)
		//line templates/page.qtpl:34
		qw422016.E().S(state.GithubURL)
		//line templates/page.qtpl:34
		qw422016.N().S(`">Log in with github</a>
        `)
		//line templates/page.qtpl:35
	} else {
		//line templates/page.qtpl:35
		qw422016.N().S(`
          <p>Logged in as `)
		//line templates/page.qtpl:36
		qw422016.E().S(state.Login)
		//line templates/page.qtpl:36
		qw422016.N().S(`</p>
          `)
		//line templates/page.qtpl:37
		StreamProjects(qw422016, state)
		//line templates/page.qtpl:37
		qw422016.N().S(`
        `)
		//line templates/page.qtpl:38
	}
	//line templates/page.qtpl:38
	qw422016.N().S(`
      </div>

    </div>
 `)
	//line templates/page.qtpl:42
	templates.StreamJS(qw422016)
	//line templates/page.qtpl:42
	qw422016.N().S(`
</body>
`)
//line templates/page.qtpl:44
}

//line templates/page.qtpl:44
func WritePage(qq422016 qtio422016.Writer, state State) {
	//line templates/page.qtpl:44
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/page.qtpl:44
	StreamPage(qw422016, state)
	//line templates/page.qtpl:44
	qt422016.ReleaseWriter(qw422016)
//line templates/page.qtpl:44
}

//line templates/page.qtpl:44
func Page(state State) string {
	//line templates/page.qtpl:44
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/page.qtpl:44
	WritePage(qb422016, state)
	//line templates/page.qtpl:44
	qs422016 := string(qb422016.B)
	//line templates/page.qtpl:44
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/page.qtpl:44
	return qs422016
//line templates/page.qtpl:44
}

//line templates/page.qtpl:46
func StreamProjects(qw422016 *qt422016.Writer, state State) {
	//line templates/page.qtpl:46
	qw422016.N().S(`
<div>`)
	//line templates/page.qtpl:47
	StreamAddProject(qw422016)
	//line templates/page.qtpl:47
	qw422016.N().S(`</div>
<div>`)
	//line templates/page.qtpl:48
	StreamProjectList(qw422016, state)
	//line templates/page.qtpl:48
	qw422016.N().S(`</div>
`)
//line templates/page.qtpl:49
}

//line templates/page.qtpl:49
func WriteProjects(qq422016 qtio422016.Writer, state State) {
	//line templates/page.qtpl:49
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/page.qtpl:49
	StreamProjects(qw422016, state)
	//line templates/page.qtpl:49
	qt422016.ReleaseWriter(qw422016)
//line templates/page.qtpl:49
}

//line templates/page.qtpl:49
func Projects(state State) string {
	//line templates/page.qtpl:49
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/page.qtpl:49
	WriteProjects(qb422016, state)
	//line templates/page.qtpl:49
	qs422016 := string(qb422016.B)
	//line templates/page.qtpl:49
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/page.qtpl:49
	return qs422016
//line templates/page.qtpl:49
}

//line templates/page.qtpl:51
func StreamAddProject(qw422016 *qt422016.Writer) {
	//line templates/page.qtpl:51
	qw422016.N().S(`
<form id="addproject">
<input type="text" name="url" placeholder="Repository URL">
<input type="text" name="domain" placeholder="Domain">
<input type="submit" value="Add project">
</form>
`)
//line templates/page.qtpl:57
}

//line templates/page.qtpl:57
func WriteAddProject(qq422016 qtio422016.Writer) {
	//line templates/page.qtpl:57
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/page.qtpl:57
	StreamAddProject(qw422016)
	//line templates/page.qtpl:57
	qt422016.ReleaseWriter(qw422016)
//line templates/page.qtpl:57
}

//line templates/page.qtpl:57
func AddProject() string {
	//line templates/page.qtpl:57
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/page.qtpl:57
	WriteAddProject(qb422016)
	//line templates/page.qtpl:57
	qs422016 := string(qb422016.B)
	//line templates/page.qtpl:57
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/page.qtpl:57
	return qs422016
//line templates/page.qtpl:57
}

//line templates/page.qtpl:59
func StreamProjectList(qw422016 *qt422016.Writer, state State) {
	//line templates/page.qtpl:59
	qw422016.N().S(`
<table id="projects" class="table table-hover">
<thead>
  <tr>
    <th>URL</th>
    <th>Domain</th>
    <th>Deploy</th>
    <th>Auto-deploy</th>
  </tr>
</thead>
`)
	//line templates/page.qtpl:69
	for _, project := range state.Projects {
		//line templates/page.qtpl:69
		qw422016.N().S(`
  <tr>
    <td>`)
		//line templates/page.qtpl:71
		qw422016.E().S(project.URL)
		//line templates/page.qtpl:71
		qw422016.N().S(`</td>
    <td>`)
		//line templates/page.qtpl:72
		qw422016.E().S(project.Domain)
		//line templates/page.qtpl:72
		qw422016.N().S(`</td>
    <td><a href="javascript:void(0);" class="btn btn-default">Deploy</a></td>
    <td><input type="checkbox" data-id="`)
		//line templates/page.qtpl:74
		qw422016.E().S(project.ID)
		//line templates/page.qtpl:74
		qw422016.N().S(`" `)
		//line templates/page.qtpl:74
		if project.AutoDeploy {
			//line templates/page.qtpl:74
			qw422016.N().S(`checked="checked"`)
			//line templates/page.qtpl:74
		}
		//line templates/page.qtpl:74
		qw422016.N().S(`></td>
  </tr>
`)
		//line templates/page.qtpl:76
	}
	//line templates/page.qtpl:76
	qw422016.N().S(`
</table>
`)
//line templates/page.qtpl:78
}

//line templates/page.qtpl:78
func WriteProjectList(qq422016 qtio422016.Writer, state State) {
	//line templates/page.qtpl:78
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/page.qtpl:78
	StreamProjectList(qw422016, state)
	//line templates/page.qtpl:78
	qt422016.ReleaseWriter(qw422016)
//line templates/page.qtpl:78
}

//line templates/page.qtpl:78
func ProjectList(state State) string {
	//line templates/page.qtpl:78
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/page.qtpl:78
	WriteProjectList(qb422016, state)
	//line templates/page.qtpl:78
	qs422016 := string(qb422016.B)
	//line templates/page.qtpl:78
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/page.qtpl:78
	return qs422016
//line templates/page.qtpl:78
}
