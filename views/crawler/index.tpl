{{.Test}}
</br>

{{range $idx, $item := .s}}
<div>
    {{$idx}} : {{$item.Name}}
</div>
{{end}}