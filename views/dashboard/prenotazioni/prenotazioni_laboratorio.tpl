<h2 class="mt-4 content-tab-title">{{.Title}}</h2>

{{if eq .Ruolo "organizzazione"}}
<script src="https://cdn.syncfusion.com/ej2/dist/ej2.min.js" type="text/javascript"></script>
<script src="bar.js" type="text/javascript"></script>
<div id="container-statistiche"></div>
{{end}}

<div class="parent">
    <div id="div-tabella" class="table-container table-responsive">
        <form method="POST" action="/dashboard/referti" enctype="multipart/form-data">
            {{if eq .Ruolo "laboratorio"}}
            <input id="button-salva-modifiche" type="submit" class="bg-lightblue" value="Salva modifiche">
            {{end}}
            <table id="table-orari-privati" class="table-responsive">
                <tbody>
                    <tr id="tr-prenotazioni">
                        <td id="td-id-test" class="td-id-test">
                            <input type="number" min="1" id="cerca-id">
                        </td>
                        <td>
                            <input type="date" id="cerca-data-esecuzione">
                        </td>
                        <td>
                            <select id="cerca-pagato">
                                <option value=""></option>
                                <option value="No">No</option>
                                <option value="Sì">Sì</option>
                            </select>
                        </td>
                        <td>
                            <select id="cerca-tipologia-test">
                                <option value=""></option>
                                <option value="molecolare">Molecolare</option>
                                <option value="antigenico">Antigenico</option>
                                <option value="sierologico">Sierologico</option>
                            </select>
                        </td>
                        <td>
                            <select id="cerca-stato">
                                <option value=""></option>
                                <option value="prenotato">Prenotato</option>
                                <option value="eseguito">Eseguito</option>
                                <option value="notificato">Notificato</option>
                            </select>
                        </td>
                        <td>
                            <input type="text" id="cerca-privato">
                        </td>
                        <td>
                            <select id="cerca-esito">
                                <option value=""></option>
                                <option value="nullo">Nullo</option>
                                <option value="negativo">Negativo</option>
                                <option value="positivo">Positivo</option>
                            </select>
                        </td>
                        <td id="td-button" colspan="2">
                            <button class="bg-lightblue w-75" id="cerca" onclick="return cercaPrenotazioni()">Cerca
                            </button>
                        </td>
                    </tr>
                    <tr>
                        <th>ID test</th>
                        <th>Data esecuzione</th>
                        <th>Pagato online</th>
                        <th>Tipologia test</th>
                        <th>Stato</th>
                        <th>Privato</th>
                        <th>Esito</th>
                        <th>Referto</th>
                        <th>Questionario anamnesi</th>
                    </tr>
                    {{range $i, $v := .TestDiagnostici}}
                    <tr>
                        <td>
                            <input type="hidden" name="test-diagnostico-{{$i}}" value="{{.IdTestDiagnostico}}">
                            {{.IdTestDiagnostico}}
                        </td>
                        <td>
                            {{.DataEsecuzione.Format "2006-01-02 15:04"}}
                        </td>
                        <td>
                            {{if .Pagato}}
                            Sì
                            {{else}}
                            No
                        {{end}}
                        </td>
                        <td>
                            {{.TipologiaTest}}
                        </td>
                        <td>
                            {{.Stato}}
                        </td>
                        <td>
                            {{.Privato.Nome}} {{.Privato.Cognome}} -- {{.Privato.CodiceFiscale}}
                        </td>
                        <td>
                            {{if .Referto}}
                                {{.Referto.Risultato}}
                            {{else}}
                            --
                            {{end}}
                        </td>
                        <td id="td-referto">
                            {{if not .Referto}}
                                {{if eq $.Ruolo "laboratorio"}}
                                <input type="file" name="referto-upload-{{$i}}" accept="application/pdf">
                                <br>
                                <label for="esito-{{$i}}">Esito:</label>
                                <select id="esito-{{$i}}" name="esito-{{$i}}">
                                    <option value="nullo">Nullo</option>
                                    <option value="negativo">Negativo</option>
                                    <option value="positivo">Positivo</option>
                                </select>
                                {{else}}
                                <p>Il referto non è ancora disponibile</p>
                                {{end}}
                            {{else}}
                            <a href="/dashboard/referti/download?idReferto={{.Referto.IdReferto}}">
                                <img id="img-referto" src="/img/icons/electrocardiogram-report-svgrepo-com.svg"
                                    class="list-group-item-image"/>
                            </a>
                        {{end}}
                        </td>
                        <td id="td-questionario">
                            <a href="/questionari/{{.Questionario.Nome}}.pdf">
                                <img id="img-questionario" src="/img/icons/contact-form-svgrepo-com.svg"
                                    class="list-group-item-image"/>
                            </a>
                        </td>
                    </tr>
                    {{end}}
                </tbody>
            </table>
    </div>
</div>
<script src="/js/prenotazioni/ricerca_prenotazioni.js"></script>
<script src="/js/referto/statistiche_referti.js"></script>
