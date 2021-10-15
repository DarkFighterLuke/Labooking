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
            <h1>Cambio Password</h1>
         </div>
      </div>
      <div class="container bg-transparent">
         <div class="card text-center">
            <div class="tab-content">
               <div class="tab-pane fade in active">
                  <form class="form-labooking" id="form-cambio-password" action="" method="POST" onsubmit="return cambioPassword()">
                     Password:
                     <input name="password" type="password" value="">
                     <br>
                     Conferma Passowrd:
                     <input name="conferma-password" type="password" value="">
                     <br>
                     <br>
                     <input type="submit">
                  </form>
               </div>
            </div>
         </div>
      </div>
   </body>
</html>