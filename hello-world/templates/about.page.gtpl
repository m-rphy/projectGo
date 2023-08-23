{{template "base" .}}

{{define "content"}}
  <div class="container">
    <div class="row">
      <div class="col">
        <h1> This is the about page </h1>
        <p> ...and some more text </p>

        <p> This came from the template: {{index .StringMap "test"}} </p>
        <p>

          {{if ne (index .StringMap "remoteIP") ""}}
            Your remote IP address is: {{index .StringMap "remoteIP"}}
          {{else}}
            I don't know your IP address. Please visit the <a href="/">Home Page</a>
          {{end}}
          
        </p>
      </div>
    </div>
  </div>
{{ end }}
 