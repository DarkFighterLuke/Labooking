<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />
    <meta name="description" content="" />
    <meta name="author" content="" />
    <title>Labooking</title>
    <!-- Favicon-->
    <link rel="icon" type="image/x-icon" href="static/favicon.ico" />
    <!-- Core theme CSS (includes Bootstrap)-->
    <link href="css/styles.css" rel="stylesheet" />
</head>
<body>
    <div class="d-flex" id="wrapper">
        {{template "dashboard/sidebar.html"}}
        <!-- Page content wrapper-->
        <div id="page-content-wrapper">
            {{template "dashboard/navbar.tpl"}}
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
        </div>
        {{template "dashboard/scripts.tpl"}}
    </div>
</body>
