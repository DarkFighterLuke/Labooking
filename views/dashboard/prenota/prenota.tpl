<h2 class="mt-4 content-tab-title">{{.Title}}</h2>
<div class="parent">
    <form action="/dashboard/prenota" method="POST"  enctype="multipart/form-data">
        {{if ne .Ruolo "privato"}}
        <div>
            {{if eq .Ruolo "medico"}}
            <h3><label for="privati">Pazienti: </label></h3>
            {{else}}
            <h3><label for="privati">Dipendenti: </label></h3>
            {{end}}
            <table id="privati" class="table-responsive">
                <tbody>
                <tr>
                    <th></th>
                    <th>Codice Fiscale</th>
                    <th>Nome</th>
                    <th>Cognome</th>
                </tr>
                {{range .Privati}}
                <tr>
                    <td><input type="checkbox" name="id-privato" value="{{.IdPrivato}}"></td>
                    <td>{{.CodiceFiscale}}</td>
                    <td>{{.Nome}}</td>
                    <td>{{.Cognome}}</td>
                </tr>
                {{end}}
                </tbody>
            </table>
        </div>
        {{end}}

        <div>
            <table>
                <tbody>
                <tr>
                    <th></th>
                    <th>Orario</th>
                    <th>Disponibilità</th>
                </tr>
                {{range .Slots}}
                <tr {{if not .Disponibile}}style="color: darkred" {{end}}>
                    <td>
                        {{if ne .Ruolo "privato"}}
                        <input type="checkbox" name="slot" value="{{.Orario}}">
                        {{else}}
                        <input type="radio" name="slot" value="{{.Orario}}">
                        {{end}}

                    </td>
                    <td>{{.Orario}}</td>
                    <td>
                        {{if .Disponibile}}
                        Disponibile
                        {{else}}
                        Non disponibile
                        {{end}}
                    </td>
                </tr>
                {{end}}
                </tbody>
            </table>
        </div>

        <div>
            <div>
                <label for="questionario-anamnesi-download">Scarica questionario di anamnesi: </label>
                <a href="/pdf/questionario-anamnesi.pdf" id="questionario-anamnesi-download" download>Questionario</a>
            </div>
            <div>
                <label for="questionario-anamnesi-upload">Caricare questionario di anamnesi: </label>
                <input type="file" id="questionario-anamnesi-upload" name="questionario-anamnesi-upload">
            </div>
        </div>

        <div>
            {{if ne .Ruolo "medico"}}
            <div id="div-paga-online">
            <p>Scegli metodo pagamento</p>
                <input type="radio" id="paga-online" name="metodo-pagamento" value="true" checked onchange="return mostraCampi()">
                <label for="paga-online">Pagamento online</label>
            </div>
            {{end}}

            <div id="div-paga-presenza">
                <input type="radio" id="paga-presenza" name="metodo-pagamento" value="false" onchange="return mostraCampi()">
                <label for="paga-presenza">Pagamento in presenza</label>
            </div>
        </div>

        <div id="dati-pagamento-online">
            <label for="numero-carta">Numero di carta: </label>
            <input id="numero-carta" name="numero-carta" type="text" maxlength="16" >
            <label for="scadenza">Scadenza: </label>
            <input id="scadenza" name="scadenza" type="month">
            <label for="cvv">CVV: </label>
            <input id="cvv" name="cvv" type="text" maxlength="4">
        </div>

        <div id="dati-pagamento-presenza">
            Iban: {{.Iban}}
        </div>

        <input type="submit" class="bg-lightblue" value="Conferma prenotazione">
    </form>
</div>
<script src="/js/prenota/prenota.js"></script>