This is the resumé of {{.Name}} {{.Surname}}

Name: {{.Name}}
Surname: {{.Surname}}
Phone No.: {{.PhoneNo}}

{{ padleft "aaa" "pippo" }} {{ template "inner.tpl" . }}

{{if .Developer -}} Role: Developer {{- end}}
{{if .SysAdmin}} Role: SysAdmin {{- end}}
{{range .Emails -}} Email: 
    Description: {{ .Name }}
    Address: {{.Address}}
{{end}}
