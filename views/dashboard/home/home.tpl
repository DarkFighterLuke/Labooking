<h2 class="mt-4 content-tab-title">{{.Title}}</h2>
<div>
    <table id="table-attivita-recenti" class="table-responsive">
        <tbody>
            {{range .TestDiagnostici}}
                <tr>
                {{if .Referto}}
                    <td>
                        <div class="attivita">
                            <p><strong> Referto test {{.TipologiaTest}} </strong> - {{.LastUpdate.Format "02/01/2006"}}
                            </p>
                            {{if eq $.Ruolo "laboratorio"}}
                            <p> Hai pubblicato il referto di un test diagnostico è pronto. Puoi consultarlo nella
                                sezione
                                <a href="/dashboard/referti"> Referti</a>
                            </p>
                            {{else}}
                            <p> Il referto del tuo test diagnostico è pronto. Puoi consultarlo nella sezione
                                <a href="/dashboard/referti"> Referti</a>
                            </p>
                            {{end}}
                        </div>
                    <td>
                {{else}}
                    <td>
                    <div class="attivita">
                        {{if eq $.Ruolo "laboratorio"}}
                        <p><strong> Hai ricevuto una prenotazione di un test {{.TipologiaTest}} </strong> -
                            {{.LastUpdate.Format "02/01/2006"}}</p>
                        {{else}}
                        <p><strong> Prenotazione test {{.TipologiaTest}} </strong> - {{.LastUpdate.Format "02/01/2006"}}
                        </p>
                        {{end}}
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