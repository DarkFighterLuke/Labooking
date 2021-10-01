<!-- Page content-->
<div class="container-fluid">
    <h2 class="mt-4 content-tab-title">{{.Title}}</h2>
    <!--Contenuto specifico della sezione-->
    {{if eq .TabName "home"}}<!--Renderizza template attività recenti-->
    {{else}}{{if eq .TabName "prenota"}}<!--Renderizza template prenotazione test-->
    {{else}}{{if eq .TabName "referti"}}<!--Renderizza template referti-->
    {{else}}{{if eq .TabName "calendario"}}<!--Renderizza template calendario-->
    {{else}}{{if eq .TabName "guida"}}<!--Renderizza template guida ai test-->
    {{end}}{{end}}{{end}}{{end}}{{end}}
</div>
