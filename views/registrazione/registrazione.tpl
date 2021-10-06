<!DOCTYPE html>
<html lang="en">
   {{template "registrazione/head.html"}}
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
                  <form class="form-labooking" id="form-privato" action="/signup?idForm=privato" method="POST" onsubmit="return submitPrivato()">
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
                  <form class="form-labooking" id="form-medico" action="/signup?idForm=medico" method="POST" onsubmit="return submitMedico()">
                     <input type="text" name="medico" hidden>
                     {{.FormMedico | renderform}}
                     <br><br><input id="submit-medico" type="submit">
                  </form>
               </div>
               <div id="laboratorio" class="tab-pane fade">
                  <!-- Form laboratorio -->
                  <form class="form-labooking" id="form-laboratorio" action="/signup?idForm=laboratorio" method="POST" onsubmit="return submitLaboratorio()">
                     <input type="text" name="laboratorio" hidden>
                     {{.FormLaboratorio | renderform}}
                     <table id="table-orari-apertura">
                        <tr>
                           <th>Orario apertura</th>
                           <th>Orario chiusura</th>
                           <th>Giorno</th>
                        </tr>
                        <tr>
                           <td>
                              <input name="orario-apertura-1" type="time">
                           </td>
                           <td>
                              <input name="orario-chiusura-1" type="time">
                           </td>
                           <td>
                              <select name="giorno-1">
                                 <option value="lunedi">Lunedì</option>
                                 <option value="martedi">Martedì</option>
                                 <option value="mercoledi">Mercoledì</option>
                                 <option value="giovedi">Giovedì</option>
                                 <option value="venerdi">Venerdì</option>
                                 <option value="sabato">Sabato</option>
                                 <option value="domenica">Domenica</option>
                              </select>
                           </td>
                        </tr>
                     </table>
                     <button onclick="return aggiungiRiga()">Aggiungi</button>
                     <script src="js/registrazione/orari_apertura.js"></script>
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
      <script src="js/registrazione/registrazione_privato.js"></script>
      <script src="js/registrazione/registrazione_medico.js"></script>
      <script src="js/registrazione/registrazione_laboratorio.js"></script>
      <script src="js/registrazione/registrazione.js"></script>
   </body>
</html>