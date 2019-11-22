{{.Test}}
</br>

{{range $idx, $item := .s}}
<div>
    {{$item.Id}} : <a href="http://big5.quanben.io{{$item.Url}}"> {{$item.Name}} </a>
</div>
{{end}}