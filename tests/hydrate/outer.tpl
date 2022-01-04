This is the resumé of {{.Name}} {{.Surname}}

Name: {{.Name | red}}
Surname: {{.Surname | red}}
Phone No.: {{.PhoneNo}}

{{ template "inner.tpl" . }}

{{if .Developer -}} Role: Developer {{- end}}
{{if .SysAdmin}} Role: SysAdmin {{- end}}
{{range .Emails -}} Email: 
    Description: {{ .Name }}
    Address: {{.Address}}
{{end}}

Include script as quote:
{{ include "tests/included.sh" . "> " -}}
Include script with 2-spaces indentation:
{{ include "tests/included.sh" "  " -}}
Include script as is:
{{ include "tests/included.sh" . }}
