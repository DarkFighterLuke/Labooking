<!DOCTYPE html>
<html lang="en">
    {{template "landingPage/head.html"}}

    <body class="bg-light">
        {{template "landingPage/top_bar.html"}}

        <!-- Form privato -->
        <form id="privato" action="" method="POST">
            <input type="text" name="privato" hidden="true">
            {{.FormPrivato | renderform}}
            <!--Aggiungere input per la data di nascita-->
        </form>

        <!-- Form medico -->
        <form id="medico" action="" method="POST">
            <input type="text" name="medico" hidden="true">
            {{.FormMedico | renderform}}
            <!--Aggiungere input per la data di nascita-->
        </form>

        <!-- Form laboratorio -->
        <form id="laboratorio" action="" method="POST">
            <input type="text" name="laboratorio" hidden="true">
            {{.FormLaboratorio | renderform}}
        </form>
    </body>
</html>
