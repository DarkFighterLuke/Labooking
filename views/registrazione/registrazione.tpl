<!DOCTYPE html>
<html lang="en">
    {{template "landingPage/head.html"}}

    <body class="bg-light">
        {{template "landingPage/top_bar.html"}}
        <form action="" method="POST">
            {{.Form | renderform}}
        </form>
    </body>
</html>
