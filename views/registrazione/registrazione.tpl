<!DOCTYPE html>
<html lang="en">
    {{template "landingPage/head.html"}}

    <body class="bg-light">
        {{template "landingPage/top_bar.html"}}

        <!-- Form privato -->
        <form id="privato" action="" method="POST">
            <!--Aggiungere campo tipo form-->
            {{.FormPrivato | renderform}}
            <!--Aggiungere input per la data di nascita-->
        </form>

        <!-- Form medico -->
        <form id="medico" action="" method="POST">
            <!--Aggiungere campo tipo form-->
            {{.FormMedico | renderform}}
            <!--Aggiungere input per la data di nascita-->
        </form>

        <!-- Form laboratorio -->
        <form id="laboratorio" action="" method="POST">
            <!--Aggiungere campo tipo form-->
            {{.FormLaboratorio | renderform}}
        </form>
    </body>
</html>
