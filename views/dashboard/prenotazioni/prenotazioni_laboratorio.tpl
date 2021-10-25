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
                    {{.DataEsecuzione}}
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
                   <a href="{{.Referto.IdReferto}}"><button></button></a> <!--TODO: inserire immagine-->
                    {{end}}
                </td>
                <td>
                    <a href="/dashboard/questionari/{{.Questionario.Nome}}.pdf" download><button></button></a>
                </td>
            </tr>
            {{end}}
            </tbody>
        </table>
    </div>
</div>