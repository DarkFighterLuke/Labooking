<!DOCTYPE html>
<html lang="en">
   {{template "login/head.html"}}
   <body class="bg-light">
      <div>
         <nav class="navbar navbar-dark bg-primary h-100">
            <a class="navbar-brand" href="/">Labooking</a>
         </nav>
      </div>
      <div class="container">
         <div class="row justify-content-center d-flex align-items-end h-100">
            <h1>Login</h1>
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
                  <form class="form-labooking" id="form-privato" action="/login?idForm=privato" method="POST" onsubmit="return loginPrivato()">
                      Email:
                      <input id="email-privato" name="email-privato" type="text" value="">
                      <br>
                      Password:
                      <input id="password-privato" name="password-privato" type="password" value="">
                      <br>
                      <a href="#">Hai dimenticato la password?</a>
                      <br>
                      <br>
                      <input id="submit-privato" type="submit">
                  </form>
               </div>
               <div id="organizzazione" class="tab-pane fade">
                  <p>ok</p>
               </div>
               <div id="medico" class="tab-pane fade">
                  <!-- Form medico -->
                  <form class="form-labooking" id="form-medico" action="/login?idForm=medico" method="POST" onsubmit="return loginMedico()">
                      Email:
                      <input id="email-medico" name="email-medico" type="text" value="">
                      <br>
                      Password:
                      <input id="password-medico" name="password-medico" type="password" value="">
                      <br>
                      <a href="#">Hai dimenticato la password?</a>
                      <br>
                      <br>
                      <input id="submit-medico" type="submit">
                  </form>
               </div>
               <div id="laboratorio" class="tab-pane fade">
                  <!-- Form laboratorio -->
                  <form class="form-labooking" id="form-laboratorio" action="/login?idForm=laboratorio" method="POST" onsubmit="return loginLaboratorio()">
                      Email:
                      <input id="email-laboratorio" name="email-laboratorio" type="text" value="">
                      <br>
                      Password:
                      <input id="password-laboratorio" name="password-laboratorio" type="password" value="">
                      <br>
                      <a href="#">Hai dimenticato la password?</a>
                      <br>
                      <br>
                      <input id="submit-laboratorio" type="submit">
                  </form>
               </div>
            </div>
         </div>
      </div>
      <div class="container not-registred">
          <div class="row justify-content-center">
              <p>{{.errmsg}}</p>
          </div>
         <div class="row justify-content-center">
            <a href="#" class="align-self-center">Non sei registrato? Registrati ora</a>
         </div>
      </div>
   </body>
</html>