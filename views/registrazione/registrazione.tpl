<!DOCTYPE html>
<html lang="en">
    {{template "landingPage/head.html"}}

    <body class="bg-light">
        {{template "landingPage/top_bar.html"}}

        <!-- Form privato -->
        <form action="" method="POST">
            {{.FormPrivato | renderform}}
        </form>

        <!-- Form medico -->
        <form action="" method="POST">
            {{.FormMedico | renderform}}
        </form>
    </body>
</html>
