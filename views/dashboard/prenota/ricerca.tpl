<h2 class="mt-4 content-tab-title">{{.Title}}</h2>
<div class="parent">

    <div id="luogo-autocomplete" class="autocomplete">
        <input type="text" id="luogo" name="luogo" placeholder="Cerca nei dintorni di...">
    </div>

    <div class="area cover">
        {{if ne .Ruolo "privato"}}
        <div>
            <label for="numero-persone">Numero persone: </label>
            <input type="number" id="numero-persone" name="numero-persone" min="1" placeholder="1">
        </div>
        {{end}}

        <div>
            <label for="costo">Costo: </label>
            <input id="costo" name="costo" value="50" type="range" min="{{.MinCosto}}" max="{{.MaxCosto}}" oninput="this.nextElementSibling.value = this.value">
            <output>50</output>€
            <br>
            <label for="tempo">Tempo: </label>
            <input id="tempo" name="tempo" value="25" type="range" min="{{.MinTempi}}" max="{{.MaxTempi}}" oninput="this.nextElementSibling.value = this.value">
            <output>25</output>h
        </div>

        <div>
            <label for="molecolare">Molecolare</label>
            <input type="checkbox" id="molecolare" name="molecolare"><br>
            <label for="antigenico">Antigenico</label>
            <input type="checkbox" id="antigenico" name="antigenico"><br>
            <label for="sierologico">Sierologico</label>
            <input type="checkbox" id="sierologico" name="sierologico"><br>
        </div>

        <div id="div-datetime">
            <label for="inizio-intervallo">Dalle </label>
            <input type="time" id="inizio-intervallo" name="inizio-intervallo" value="{{.OraInizio}}">
            <label for="fine-intervallo">alle </label>
            <input type="time" id="fine-intervallo" name="fine-intervallo" value="{{.OraFine}}">
            <label for="data">del </label>
            <input type="date" id="data" name="data" value="{{.Data}}">
            <br>
            <input id="button-cerca" class="bg-lightblue" type="button" value="Cerca" onclick="return sendFilters()">
            <br>
        </div>

        <div id="map"></div>
    </div>
</div>
<script src="/js/utils/autocompleteAddress.js"></script>
<script src="/js/prenota/ricerca/ricerca.js"></script>
