<h2 class="mt-4 content-tab-title">{{.Title}}</h2>
<div class="parent">
    <form action="/dashboard/prenota" method="POST" enctype="multipart/form-data">
        <input hidden name="id-laboratorio" value="{{.IdLaboratorio}}">
        <input hidden name="data" value="{{.DataPrenotazione}}">
        <div id="div-tipologia-test">
            <h3>Tipologie Test</h3>
            {{range .InfoTest}}
            <label for="{{.TipologiaTest}}">{{.TipologiaTest}}</label>
            <input type="radio" id="{{.TipologiaTest}}" value="{{.TipologiaTest}}" name="tipologia-test">
            {{end}}
        </div>
        {{if ne .Ruolo "privato"}}
        <div>
            <h4>Modello questionario di anamnesi</h4>
            <label for="questionario-anamnesi-download">Scarica qui il questionario di anamnesi: </label>
            <a href="/pdf/questionario-anamnesi.pdf" id="questionario-anamnesi-download" download>Questionario</a>
        </div>
        {{end}}

        <div id="div-orari-privati">
            <h3>Orari prenotabili</h3>
            <table id="table-orari-privati" class="table-responsive">
                <tbody>
                {{$ruolo := .Ruolo}}
                {{$privati := .Privati}}
                <tr>
                    <th></th>
                    <th>Orario</th>
                    <th>Disponibilità</th>
                    {{if ne $ruolo "privato"}}
                    {{if eq $ruolo "organizzazione"}}
                    <th>Dipendente</th>
                        {{else}}
                        <th>Paziente</th>
                        {{end}}
                        <th>Questionario</th>
                    {{end}}
                </tr>
                {{range .Slots}}
                <tr {{if not .Disponibile}}style="color: darkred" {{end}}>
                    <td>
                        {{if eq $ruolo "privato"}}
                        <input type="radio" name="slot" value="{{.Orario}}" {{if not .Disponibile}}disabled{{end}}>
                        {{else}}
                        <input type="checkbox" name="slot" value="{{.Orario}}" {{if not .Disponibile}}disabled{{end}}>
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
                    {{if ne $ruolo "privato"}}
                    <td>
                        <select name="privato-{{.Orario}}">
                            <option>---</option>
                            {{range $privati}}
                            <option value="{{.IdPrivato}}">{{.Nome}} {{.Cognome}} --- {{.CodiceFiscale}}</option>
                            {{end}}
                        </select>
                    </td>
                    <td><input type="file" id="questionario-anamnesi-upload" name="questionario-anamnesi-upload-{{.Orario}}"></td>
                    {{end}}
                </tr>
                {{end}}
                </tbody>
            </table>
        </div>
        <br>

        {{if eq .Ruolo "privato"}}
        <div id="div-questionario">
            <h3>Questionario di anamnesi</h3>
            <div>
                <label for="questionario-anamnesi-download">Scarica qui il questionario di anamnesi: </label>
                <a href="/pdf/questionario-anamnesi.pdf" id="questionario-anamnesi-download" download>Questionario</a>
            </div>

            <div>
                <label for="questionario-anamnesi-upload">Caricare qui il questionario di anamnesi: </label>
                <input type="file" id="questionario-anamnesi-upload" name="questionario-anamnesi-upload">
            </div>
        </div>
        <br>
        {{end}}

        <div id="div-pagamenti">
            <h3>Pagamento</h3>
            {{if ne .Ruolo "medico"}}
            <div id="div-paga-online">
                <p>Scegli metodo pagamento</p>
                <input type="radio" id="paga-online" name="metodo-pagamento" value="true" checked
                       onchange="return mostraCampi()">
                <label for="paga-online">Pagamento online</label>
            </div>
            {{end}}

            <div id="div-paga-presenza">
                <input type="radio" id="paga-presenza" name="metodo-pagamento" value="false"
                       onchange="return mostraCampi()" {{if eq .Ruolo "medico"}}checked{{end}}>
                <label for="paga-presenza">Pagamento in presenza</label>
            </div>
        </div>

        <div id="dati-pagamento-online">
            <label for="numero-carta">Numero di carta: </label>
            <input id="numero-carta" name="numero-carta" type="text" maxlength="16">
            <label for="scadenza">Scadenza: </label>
            <input id="scadenza" name="scadenza" type="month">
            <label for="cvv">CVV: </label>
            <input id="cvv" name="cvv" type="text" maxlength="3">
        </div>

        <div id="dati-pagamento-presenza">
            Iban: {{.Iban}}
        </div>

        <input id="conferma-prenotazione" type="submit" class="bg-lightblue" value="Conferma prenotazione" onclick="return checkDatiInseriti()">
    </form>
</div>
<script src="/js/prenota/prenota.js"></script>