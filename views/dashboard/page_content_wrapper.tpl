<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />
    <meta name="description" content="" />
    <meta name="author" content="" />
    <title>Labooking</title>
    <!-- Favicon-->
    <link rel="icon" type="image/x-icon" href="/img/favicon.ico"/>
    <!-- Core theme CSS (includes Bootstrap)-->
    <link href="/css/styles.css" rel="stylesheet"/>
    {{.Head}}
</head>
<body>
<div class="d-flex" id="wrapper">
    {{.Sidebar}}
    <!-- Page content wrapper-->
    <div id="page-content-wrapper">
        {{.Navbar}}
        <div class="container-fluid">
            {{.LayoutContent}}
        </div>
    </div>
    <!-- Bootstrap core JS-->
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js"></script>
    <!-- Core theme JS-->
    <script src="/js/scripts.js"></script>
</div>
</body>
