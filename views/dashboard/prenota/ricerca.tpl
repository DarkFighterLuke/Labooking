<h2 class="mt-4 content-tab-title">{{.Title}}</h2>
<div>
    <input type="text" id="luogo" name="luogo" placeholder="Cerca nei dintorni di...">
</div>

<div>
    <label for="costo">Costo: </label>
    <input id="costo" name="costo" type="range" min="0" max="100">
    <br>
    <label for="tempo">Tempo: </label>
    <input id="tempo" name="tempo" type="range" min="0" max="100">
</div>

<div>
    <label for="molecolare">Molecolare</label>
    <input type="checkbox" id="molecolare" name="molecolare"><br>
    <label for="antigenico">Antigenico</label>
    <input type="checkbox" id="antigenico" name="antigenico"><br>
    <label for="sierologico">Sierologico</label>
    <input type="checkbox" id="sierologico" name="sierologico"><br>
</div>

<div>
    <label for="inizio-intervallo">Dalle </label>
    <input type="time" id="inizio-intervallo" name="inizio-intervallo">
    <label for="fine-intervallo">alle </label>
    <input type="time" id="fine-intervallo" name="fine-intervallo">
    <label for="giorno">il </label>
    <select id="giorno" name="giorno">
        <option selected value>Tutti</option>
        <option value="lunedi">Lunedì</option>
        <option value="martedi">Martedì</option>
        <option value="mercoledi">Mercoledì</option>
        <option value="giovedi">Giovedì</option>
        <option value="venerdi">Venerdì</option>
        <option value="sabato">Sabato</option>
        <option value="domenica">Domenica</option>
    </select>
    <br>
    <input type="button" value="Cerca" onclick="return sendFilters()">
    <br>
</div>

<div id="map" style="height: 650px"></div>

<script src="/js/prenota/ricerca/ricerca.js"></script>