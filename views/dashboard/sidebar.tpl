<!-- Sidebar-->
<div class="border-end bg-lightgray" id="sidebar-wrapper">
    <div class="sidebar-heading border-bottom bg-light">
        <strong>Menu</strong>
    </div>
    <div class="list-group list-group-flush bg-lightgray">
        <a class="list-group-item list-group-item-action list-group-item-light p-3" href="/dashboard/home">
            <img src="/img/icons/house-door-fill.svg" class="list-group-item-image"/>
            Home</a>
        {{if eq .Ruolo "medico"}}
        <a class="list-group-item list-group-item-action list-group-item-light p-3" href="/dashboard/pazienti">
            <img src="/img/icons/people-svgrepo-com.svg" class="list-group-item-image"/>
            Pazienti</a>
        {{end}}
        {{if eq .Ruolo "organizzazione"}}
        <a class="list-group-item list-group-item-action list-group-item-light p-3" href="/dashboard/dipendenti">
            <img src="/img/icons/people-svgrepo-com.svg" class="list-group-item-image"/>
            Dipendenti</a>
        {{end}}
        {{if or (eq .Ruolo "privato") (eq .Ruolo "medico") (eq .Ruolo "organizzazione")}}
        <a class="list-group-item list-group-item-action list-group-item-light p-3" href="/dashboard/prenota?action=ricerca">
            <img src="/img/icons/syringe-svgrepo-com.svg" class="list-group-item-image"/>
            Prenota test</a>
        {{else}}
        <a class="list-group-item list-group-item-action list-group-item-light p-3" href="/dashboard/prenotazioni">
            <img src="/img/icons/syringe-svgrepo-com.svg" class="list-group-item-image"/>
            Gestisci prenotazioni</a>
        {{end}}
        {{if or (eq .Ruolo "privato") (eq .Ruolo "medico") (eq .Ruolo "organizzazione")}}
        <a class="list-group-item list-group-item-action list-group-item-light p-3" href="/dashboard/referti">
            <img src="/img/icons/electrocardiogram-report-svgrepo-com.svg" class="list-group-item-image"/>
            Referti</a>
        {{end}}
        <a class="list-group-item list-group-item-action list-group-item-light p-3" href="/dashboard/calendario">
            <img src="/img/icons/calendar-svgrepo-com.svg" class="list-group-item-image"/>
            Calendario</a>
        {{if or (eq .Ruolo "privato") (eq .Ruolo "medico") (eq .Ruolo "organizzazione")}}
        <a class="list-group-item list-group-item-action list-group-item-light p-3" href="/dashboard/guida">
            <img src="/img/icons/help-svgrepo-com.svg" class="list-group-item-image"/>
            Guida ai test</a>
        {{end}}
    </div>
</div>
