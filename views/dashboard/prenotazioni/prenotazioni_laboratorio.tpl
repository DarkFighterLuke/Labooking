<h2 class="mt-4 content-tab-title">{{.Title}}</h2>
<div class="parent">
    <div>
        <table id="table-orari-privati" class="table-responsive">
            <tbody>
            {{range .TestDiagnostici}}
            <tr>
                <th>ID Referto</th>
                <th>Data Esecuzione</th>
                <th>Pagato</th>
                <th>Tipologia Test</th>
                <th>Stato</th>
                <th>Privato</th>
                <th>Referto</th>
                <th>Qestionario Anamnesi</th>
            </tr>
            <tr>
                <td>
                    {{.IdReferto}}
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
                   <a href="#"><button></button></a> <!--TODO: inserire immagine-->
                </td>
                <td>
                    <a href="#"><button></button></a> <!-- TODO: creare link di download-->
                </td>
            </tr>
            {{end}}
            </tbody>
        </table>
    </div>
</div>