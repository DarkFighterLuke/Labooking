<h2 class="mt-4 content-tab-title">{{.Title}}</h2>
<div class="parent">
    <div>
        <table id="table-orari-privati" class="table-responsive">
            <tbody>
            <tr>
                <th>ID Test</th>
                <th>Data Esecuzione</th>
                <th>Pagato</th>
                <th>Tipologia Test</th>
                <th>Stato</th>
                <th>Privato</th>
                <th>Referto</th>
                <th>Questionario Anamnesi</th>
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