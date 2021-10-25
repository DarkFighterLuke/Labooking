<h2 class="mt-4 content-tab-title">{{.Title}}</h2>
<div class="parent">
    <div id="div-tabella">
        <table id="table-orari-privati" class="table-responsive">
            <tbody>
            <tr>
                <td>
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
                <td><button id="cerca" onclick="return cercaPrenotazioni()">Cerca</button></td>
            </tr>
            <tr>
                <th>ID test</th>
                <th>Data esecuzione</th>
                <th>Pagato</th>
                <th>Tipologia test</th>
                <th>Stato</th>
                <th>Privato</th>
                <th>Referto</th>
                <th>Questionario anamnesi</th>
            </tr>
            {{range .TestDiagnostici}}
            <tr>
                <td>
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
                    <a href="{{.Referto.IdReferto}}">
                        <img src="/img/icons/electrocardiogram-report-svgrepo-com.svg" class="list-group-item-image"/>
                        Referto
                    </a>
                    {{end}}
                </td>
                <td>
                    <a href="/dashboard/questionari/{{.Questionario.Nome}}.pdf" download><img src="/img/icons/contact-form-svgrepo-com.svg" width="60" class="list-group-item-image"/></a>
                </td>
            </tr>
            {{end}}
            </tbody>
        </table>
    </div>
</div>
<script src="/js/prenotazioni/ricerca_prenotazioni.js"></script>