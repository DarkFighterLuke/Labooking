<h2 class="mt-4 content-tab-title">{{.Title}}</h2>
<div id="div-tabella-privati" class="table-container">
    {{if eq .Ruolo "organizzazione"}}
    <a href="/dashboard/dipendenti/aggiunta"><button>Aggiungi dipendente</button></a>
    {{end}}
    {{if eq .Ruolo "medico"}}
    <a href="/dashboard/pazienti/aggiunta"><button>Aggiungi paziente</button></a>
    {{end}}
    <table id="table-privati" class="table-responsive">
        <tr>
            <th>ID Privato</th>
            <th>Nome</th>
            <th>Cognome</th>
            <th>Codice fiscale</th>
            <th>Indirizzo</th>
            <th>Prefisso</th>
            <th>Telefono</th>
            <th>Email</th>
            <th>Data di nascita</th>
        </tr>
        {{range .Privati}}
        <tr>
            <td>{{.IdPrivato}}</td>
            <td>{{.Nome}}</td>
            <td>{{.Cognome}}</td>
            <td>{{.CodiceFiscale}}</td>
            <td>{{.Indirizzo}}</td>
            <td>{{.Prefisso}}</td>
            <td>{{.Telefono}}</td>
            <td>{{.Email}}</td>
            <td>{{.DataNascita.Format "02/01/2006"}}</td>
        </tr>
        {{end}}
    </table>
</div>