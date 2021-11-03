<h2 class="mt-4 content-tab-title">{{.Title}}</h2>

<div>
    {{range .TestDiagnostici}}
    <div>
        <p>Referto Test {{.TipologiaTest}} - {{.LastUpdate}}</p>
        <p>Accedi alla pagina referti per visualizzarlo</p>
    </div>
    {{end}}
</div>