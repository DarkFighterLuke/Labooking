<!DOCTYPE html>
<html lang="en">
   <head>
      <meta charset="utf-8">
      <meta name="viewport" content="width=device-width, initial-scale=1">
      <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css">
      <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap.min.css">
      <link href="https://fonts.googleapis.com/css2?family=Oleo+Script+Swash+Caps:wght@400;700&display=swap" rel="stylesheet">
      <link href="https://fonts.googleapis.com/css?family=Open+Sans" rel="stylesheet">
      <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
      <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/js/bootstrap.min.js"></script>
      <link rel="stylesheet" href="css/stylesRegistrazione.css">
   </head>
   <body class="bg-light">
      <div>
         <nav class="navbar navbar-dark bg-primary h-100">
            <a class="navbar-brand" href="/">Labooking</a>
         </nav>
      </div>
      <div class="container">
         <div class="row justify-content-center d-flex align-items-end h-100">
            <h1>Registrazione</h1>
         </div>
      </div>
      <div class="container bg-transparent">
         <div class="card text-center">
            <div class="card-header">
               <ul class="nav nav-pills">
                  <li class="nav-item active"><a data-toggle="pill" href="#privato">Privato</a></li>
                  <li class="nav-item"><a data-toggle="pill" href="#organizzazione">Organizzazione</a></li>
                  <li class="nav-item"><a data-toggle="pill" href="#medico">Medico</a></li>
                  <li class="nav-item"><a data-toggle="pill" href="#laboratorio">Laboratorio</a></li>
               </ul>
            </div>
            <div class="tab-content">
               <div id="privato" class="tab-pane fade in active">
                  <!-- Form privato -->
                  <form class="form-labooking" id="form-privato" action="/signup?idForm=privato" method="POST" arseonsubmit="return submitPrivato()">
                     <input type="text" name="privato" hidden>
                     {{.FormPrivato | renderform}}
                     <br>Data di nascita: <input name="DataNascita" type="date" value="" id="data-nascita-privato">
                     <br><br><input id="submit-privato" type="submit">
                  </form>
               </div>
               <div id="organizzazione" class="tab-pane fade">
                  <p>ok</p>
               </div>
               <div id="medico" class="tab-pane fade">
                  <!-- Form medico -->
                  <form class="form-labooking" id="form-medico" action="/signup?idForm=medico" method="POST">
                     <input type="text" name="medico" hidden>
                     {{.FormMedico | renderform}}
                     <br><br><input id="submit-medico" type="submit">
                  </form>
               </div>
               <div id="laboratorio" class="tab-pane fade">
                  <!-- Form laboratorio -->
                  <form class="form-labooking" id="form-laboratorio" action="/signup?idForm=laboratorio" method="POST">
                     <input type="text" name="laboratorio" hidden>
                     {{.FormLaboratorio | renderform}}
                     <br><br><input id="submit-laboratorio" type="submit">
                  </form>
               </div>
            </div>
         </div>
      </div>
      <div class="container already-registred">
         <div class="row justify-content-center">
            <a href="#" class="align-self-center">Sei già registrato? Effettua il login</a>
         </div>
      </div>
      <script src="js/registrazione/prefissi_telefonici.js"></script>
      <script src="js/registrazione/registrazione.js"></script>
   </body>
</html>