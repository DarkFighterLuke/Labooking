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
                  <form class="form-labooking" id="form-privato" action="/signup?idForm=privato" method="POST">
                     <input type="text" name="privato" hidden>
                     {{.FormPrivato | renderform}}
                     <br>Data di nascita: <input name="DataNascita" type="date" value="" id="data-nascita-privato">
                     <br><br><input id="submit-privato" type="submit" onclick="return submitPrivato()">
                  </form>
               </div>
               <div id="organizzazione" class="tab-pane fade">
                  <!-- Form organizzazione -->
                  <form class="form-labooking" id="form-organizzazione" action="/signup?idForm=organizzazione"
                        method="POST">
                     <input type="text" name="organizzazione" hidden>
                     {{.FormOrganizzazione | renderform}}
                     <br><br><input id="submit-organizzazione" type="submit" onclick="return submitOrganizzazione()">
                  </form>
               </div>
               <div id="medico" class="tab-pane fade">
                  <!-- Form medico -->
                  <form class="form-labooking" id="form-medico" action="/signup?idForm=medico" method="POST">
                     <input type="text" name="medico" hidden>
                     {{.FormMedico | renderform}}
                     <br><br><input id="submit-medico" type="submit" onclick="return submitMedico()">
                  </form>
               </div>
               <div id="laboratorio" class="tab-pane fade">
                  <!-- Form laboratorio -->
                  <form class="form-labooking" id="form-laboratorio" action="/signup?idForm=laboratorio" method="POST">
                     <input type="text" name="laboratorio" hidden>
                     {{.FormLaboratorio | renderform}}
                     <div class="d-flex justify-content-center">
                        <table id="table-orari-apertura" class="table w-50">
                           <tbody class="w-100">
                           <tr>
                              <th class="w-33 text-center">Orario apertura</th>
                              <th class="w-33 text-center">Orario chiusura</th>
                              <th class="w-33 text-center">Giorno</th>
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
                            </tbody>
                        </table>
                     </div>
                     <div>
                         <button onclick="return aggiungiRiga()"> + </button>
                         <script src="js/registrazione/orari_apertura.js"></script>
                     </div>
                     <div class="d-flex justify-content-center">
                        <table id="table-tamponi" class="table w-75">
                            <tbody class="w-100">
                               <tr class="row w-100 d-flex align-items-center">
                                  <th class="w-25 text-center"> Tipologia test</th>
                                  <th class="w-25 text-center"> Tempo necessario all'analisi</th>
                                  <th class="w-25 text-center"> Costo</th>
                                  <th class="w-25 text-center"> Effettua</th>
                               </tr>
                               <tr class="row w-100 d-flex align-items-center">
                                  <td class="w-25 text-center">Molecolare</td>
                                  <td class="w-25">
                                     <input class="w-25" type="number" name="molecolare-ore" min="0">
                                     <select class="w-50" name="molecolare-minuti">
                                        <option value="0">0</option>
                                        <option value="15">15</option>
                                        <option value="30">30</option>
                                        <option value="45">45</option>
                                     </select>
                                  </td>
                                  <td class="w-25">
                                     <input class= "w-75" type="number" name="molecolare-costo" min="0" step="0.5">
                                  </td>
                                  <td class="w-25">
                                     <input type="checkbox" name="molecolare-effettua">
                                  </td>
                               </tr>
                               <tr class="row w-100 d-flex align-items-center">
                                  <td class="w-25 text-center">Antigenico</td>
                                  <td class="w-25">
                                     <input class="w-25" type="number" name="antigenico-ore" min="0">
                                     <select class="w-50" name="antigenico-minuti">
                                        <option value="0">0</option>
                                        <option value="15">15</option>
                                        <option value="30">30</option>
                                        <option value="45">45</option>
                                     </select>
                                  </td>
                                  <td class="w-25">
                                     <input class= "w-75" type="number" name="antigenico-costo" min="0" step="0.5">
                                  </td>
                                  <td class="w-25">
                                     <input type="checkbox" name="antigenico-effettua">
                                  </td>
                               </tr>
                               <tr class="row w-100 d-flex align-items-center">
                                  <td class="w-25 text-center">Sierologico</td>
                                  <td class="w-25">
                                     <input class="w-25" type="number" name="sierologico-ore" min="0">
                                     <select class="w-50" name="sierologico-minuti">
                                        <option value="0">0</option>
                                        <option value="15">15</option>
                                        <option value="30">30</option>
                                        <option value="45">45</option>
                                     </select>
                                  </td>
                                  <td class="w-25">
                                     <input class="w-75" type="number" name="sierologico-costo" min="0" step="0.5">
                                  </td>
                                  <td class="w-25">
                                     <input type="checkbox" name="sierologico-effettua">
                                  </td>
                               </tr>
                            </tbody>
                        </table>
                     </div>
                     <br><br><input id="submit-laboratorio" type="submit" onclick="return submitLaboratorio()">
                  </form>
               </div>
            </div>
         </div>
      </div>
      <div class="container already-registred">
         <div class="row justify-content-center">
            <a href="/login" class="align-self-center">Sei già registrato? Effettua il login</a>
         </div>
      </div>
      <script src="js/utils/autocompleteAddress.js"></script>
      <script src="js/registrazione/prefissi_telefonici.js"></script>
      <script src="js/registrazione/registrazione_privato.js"></script>
      <script src="js/registrazione/registrazione_medico.js"></script>
      <script src="js/registrazione/registrazione_laboratorio.js"></script>
      <script src="js/registrazione/registrazione_organizzazione.js"></script>
      <script src="js/registrazione/registrazione.js"></script>
   </body>
</html>