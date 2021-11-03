<h2 class="mt-4 content-tab-title">{{.Title}}</h2>

<div>
    {{range .TestDiagnostici}}
    <div>
        {{if .Referto}}
        <p>Referto Test {{.TipologiaTest}} - {{.LastUpdate.Format "02/01/2006"}}</p>
        <p>Accedi alla pagina referti per visualizzarlo</p>
        {{else}}
        <p>Prenotazione Test {{.TipologiaTest}} - {{.LastUpdate.Format "02/01/2006"}}</p>
        {{end}}
    </div>
    {{end}}
</div>
<div><a href="/dashboard/home?all=true">Meno recenti...</a></div>