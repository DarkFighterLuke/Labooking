<h2 class="mt-4 content-tab-title">{{.Title}}</h2>
<div id="div-tabella-privati" class="table-container">
    {{if eq .Ruolo "organizzazione"}}
    <a href="/dashboard/dipendenti/aggiunta"><button class="bg-lightblue button-visualizzazione-privati">Aggiungi dipendente</button></a>
    <button class="bg-lightblue button-visualizzazione-privati" onclick="return eliminaSelezionati()">Elimina dipendente</button>
    {{end}}
    {{if eq .Ruolo "medico"}}
    <a href="/dashboard/pazienti/aggiunta"><button class="bg-lightblue button-visualizzazione-privati">Aggiungi paziente</button></a>
    <button class="bg-lightblue button-visualizzazione-privati" onclick="return eliminaSelezionati()">Elimina paziente</button>
    {{end}}
    <table id="table-privati" class="table-responsive">
        <tbody>
        <tr>
            <td></td>
            <td class="td-id-test">
                <input type="number" min="1" id="cerca-id">
            </td>
            <td>
                <input type="text" id="cerca-nome">
            </td>
            <td>
                <input type="text" id="cerca-cognome">
            </td>
            <td>
                <input type="text" max="16" id="cerca-codice-fiscale">
            </td>
            <td>
                <input type="text" id="cerca-indirizzo">
            </td>
            <td></td>
            <td>
                <input type="text" id="cerca-telefono">
            </td>
            <td>
                <input type="text" id="cerca-email">
            </td>
            <td>
                <input type="date" id="cerca-data-nascita">
            </td>
            <td id="td-button" colspan="2">
                <button class="bg-lightblue w-100" id="cerca" onclick="return cercaPrivati()">Cerca
            </button>
        </td>

        </tr>
        <tr>
            <th></th>
            <th>ID Privato</th>
            <th>Nome</th>
            <th>Cognome</th>
            <th>Codice fiscale</th>
            <th>Indirizzo</th>
            <th>Prefisso</th>
            <th>Telefono</th>
            <th>Email</th>
            <th>Data di nascita</th>
            <th></th>
        </tr>
        {{range .Privati}}
        <tr>
            <td><input type="checkbox" id="{{.IdPrivato}}"></td>
            <td>{{.IdPrivato}}</td>
            <td>{{.Nome}}</td>
            <td>{{.Cognome}}</td>
            <td>{{.CodiceFiscale}}</td>
            <td>{{.Indirizzo}}</td>
            <td>{{.Prefisso}}</td>
            <td>{{.Telefono}}</td>
            <td>{{.Email}}</td>
            <td>{{.DataNascita.Format "2006-01-02"}}</td>
            <td></td>
        </tr>
        {{end}}
        </tbody>
    </table>
</div>
<script src="/js/privati/ricerca_privati.js"></script>
<script src="/js/privati/eliminazione_privati.js"></script>