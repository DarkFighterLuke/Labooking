<h2 class="mt-4 content-tab-title">{{.Title}}</h2>

<div>
    <table id="table-attivita-recenti" class="table-responsive">
        <tbody>
            {{range .TestDiagnostici}}
                <tr>
                {{if .Referto}}
                    <td>
                        <div class="attivita">
                            <p> <strong> Referto test {{.TipologiaTest}} </strong> del {{.LastUpdate.Format "02/01/2006"}}</p>
                            <p> Il referto del tuo test diagnostico è pronto. Puoi consultarlo nella sezione
                                <a href="/dashboard/referti"> Referti</a>
                            </p>
                        </div>
                    <td>
                {{else}}
                    <td>
                        <div class="attivita">
                            <p> <strong> Prenotazione test {{.TipologiaTest}} </strong> del {{.LastUpdate.Format "02/01/2006"}}</p>
                        </div>
                    </td>
                {{end}}
                </tr>
            {{end}}
        </tbody>
    </table>
</div>

<div id="meno-recenti">
    <a href="/dashboard/home?all=true">Meno recenti</a>
</div>