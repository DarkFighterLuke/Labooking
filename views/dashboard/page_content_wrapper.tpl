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
    {{template "dashboard/sidebar.html"}}
    <!-- Page content wrapper-->
    <div id="page-content-wrapper">
        <!-- Page content-->
        <div class="container-fluid">
            <h2 class="mt-4 content-tab-title">{{.title}}</h2>
            <!--Renderizzare contenuto specifico-->
        </div>
    </div>
    {{template "dashboard/scripts.tpl"}}
</body>
