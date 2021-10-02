<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
            <meta name="viewport" content="width=device-width, initial-scale=1">
            <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css">
            <link href="https://fonts.googleapis.com/css2?family=Oleo+Script+Swash+Caps:wght@400;700&display=swap" rel="stylesheet">
            <link href="https://fonts.googleapis.com/css?family=Open+Sans" rel="stylesheet">

            <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap.min.css">
            <link rel="stylesheet" href="css/stylesRegistrazione.css">
            <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
            <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/js/bootstrap.min.js"></script>
    </head>

    <body class="bg-light">
        <div>
            <nav class="navbar navbar-dark bg-primary">
                <a class="navbar-brand" href="/">Labooking</a>
            </nav>
        </div>

        <div class="container h-10">
            <div class="row justify-content-center d-flex align-items-end h-100">
                <h1>Registrazione</h1>
            </div>
        </div>

        <div class="container bg-transparent">
                <div class="card text-center">
                    <div class="card-header">
                        <ul class="nav nav-pills">
                            <li class="nav-item active"><a data-toggle="pill" href="#utente">Utente</a></li>
                            <li class="nav-item"><a data-toggle="pill" href="#organizzazione">Organizzazione</a></li>
                            <li class="nav-item"><a data-toggle="pill" href="#medico">Medico</a></li>
                            <li class="nav-item"><a data-toggle="pill" href="#laboratorio">Laboratorio</a></li>
                        </ul>
                    </div>
                    <div class="tab-content">

                        <div id="utente" class="tab-pane fade in active">

                            <!-- Form privato -->
                            <form id="privato" action="" method="POST">
                                <input type="text" name="privato" hidden="true">
                                {{.FormPrivato | renderform}}
                                <!--Aggiungere input per la data di nascita-->
                                <br><input id="submit-privato" type="submit">
                            </form>
                        </div>

                        <div id="organizzazione" class="tab-pane fade">
                            <p>Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat</p>
                        </div>

                        <div id="medico" class="tab-pane fade">
                            <p>Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam</p>
                        </div>

                        <div id="laboratorio" class="tab-pane fade">
                            <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua</p>
                        </div>

                    </div>
                </div>
            </div>

            <div class="container">
                <div class="row justify-content-center">
                    <a href="#" class="align-self-center">Sei già registrato? Effettua il login</a>
                </div>
            </div>
    </body>
</html>
