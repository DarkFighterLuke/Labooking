<div>
    <h2 id="lab-title" class="mt-4 content-tab-title">{{.Title}}</h2>
</div>
<div>
    <h3>Orari</h3>
    <table id="table-orari-apertura" class="table w-75">
        <tbody class="w-100">
        <tr>
            <th class="w-25 text-center">Orario apertura</th>
            <th class="w-25 text-center">Orario chiusura</th>
            <th class="w-50 text-center">Giorno</th>
        </tr>
        {{range .Orari}}
        <tr>
            <td class="w-25 text-center">
                {{.OrarioAperturaStr}}
            </td>
            <td class="w-25 text-center">
                {{.OrarioChiusuraStr}}
            </td>
            <td class="w-50 text-center">
                {{.Giorno}}
            </td>
        </tr>
        {{end}}
        </tbody>
    </table>
</div>
<div>
    <h3>Tempi e Costi</h3>
    <table id="table-tamponi" class="table w-75">
        <tbody class="w-100">
        <tr>
            <th class="w-25 text-center"> Tipologia test</th>
            <th class="w-25 text-center"> Costo</th>
            <th class="w-50 text-center"> Tempo necessario all'analisi</th>
        </tr>
        {{range .InfoTest}}
        <tr>
            <td class="w-25 text-center">{{.TipologiaTest}}</td>
            <td class="w-25 text-center">
                {{.Costo}}
            </td>
            <td class="w-50 text-center">
                {{.TempiStr}}
            </td>
        </tr>
         {{end}}
        </tbody>
    </table>
</div>
<div>
    <button class="bg-lightblue" onclick="return prenota()">Prenota</button>
</div>
<script>
    function prenota(){
        let urlPrenotazione="/prenota".concat(window.location.search, "&action=prenotazione");
        window.location.href=urlPrenotazione;
    }
</script>
