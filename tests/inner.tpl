Prior Experience:
{{ range .Experiences -}} 
  * since {{ .From }} until {{ .To }} as {{ padleft .Description ">>>" -}}
{{ end }}