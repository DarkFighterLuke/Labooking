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
    <br>
    <input type="button" value="Cerca" onclick="return sendFilters()">
    <br>
</div>

<div id="map" style="height: 650px"></div>

<script src="/js/prenota/ricerca/ricerca.js"></script>
