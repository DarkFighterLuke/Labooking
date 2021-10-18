var accessToken = 'pk.eyJ1IjoiZGFya2ZpZ2h0ZXJsdWtlIiwiYSI6ImNrdWd6eWtkZTBlazEycW15bWd3dmRpMDUifQ.z-QRPdZnwbHgFsnAbUjVjw';
var filtersEndpoint = "/api/ricerca";
var visualizzaLaboratorioEndpoint = "/dashboard/laboratorio";
var nominatimApi = "https://nominatim.openstreetmap.org/search?format=json&q="

var markers = [];

var parametriGet="";

var luogo = document.getElementById("luogo");
autocomplete(luogo);

var mymap = L.map('map', {center: [41.951, 13.887], zoom: 6});
L.tileLayer('https://api.mapbox.com/styles/v1/{id}/tiles/{z}/{x}/{y}?access_token='.concat(accessToken), {
    attribution: 'Map data &copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors, Imagery © <a href="https://www.mapbox.com/">Mapbox</a>',
    maxZoom: 18,
    id: 'mapbox/streets-v11',
    tileSize: 512,
    zoomOffset: -1,
    accessToken: accessToken
}).addTo(mymap);

retrieveAllLab();

async function sendFilters(){
    let luogo=document.getElementById("luogo").value;
    let costo=document.getElementById("costo").value;
    let tempo=document.getElementById("tempo").value;
    let molecolare=document.getElementById("molecolare").checked;
    let antigenico=document.getElementById("antigenico").checked;
    let sierologico=document.getElementById("sierologico").checked;
    let inizio=document.getElementById("inizio-intervallo").value;
    let fine=document.getElementById("fine-intervallo").value;
    let data=document.getElementById("data").value;
    let dataObj=new Date(data)
    let giorno=parseDayOfWeek(dataObj.getDay())
    parametriGet="&data=".concat(data, "&ora_inizio=", inizio, "&ora_fine=", fine,
        "&molecolare=", molecolare, "&antigenico=", antigenico, "&sierologico=", sierologico);

    let filters=new FormData();
    filters.append("costo", costo);
    filters.append("tempo", tempo);
    filters.append("molecolare", molecolare);
    filters.append("antigenico", antigenico);
    filters.append("sierologico", sierologico);
    filters.append("inizio-intervallo", inizio);
    filters.append("fine-intervallo", fine);
    // TODO: Inviare data per filtrare anche in base a disponibilità laboratorio
    filters.append("giorno", giorno);

    let request=new Request(filtersEndpoint, {
        method: "POST",
        body: filters
    });
    fetch(request).then(response => setLabMap(response));
    setMapView(luogo);
}

async function setLabMap(response) {
    removeAllMarkers();
    let dati = JSON.parse(await response.text());
    if (dati !== null) {
        for (let i = 0; i < dati.length; i++) {
            let marker = L.marker([dati[i].lat, dati[i].long]).addTo(mymap);
            let labLink = visualizzaLaboratorioEndpoint.concat("?idLab=", dati[i].id_laboratorio, parametriGet);
            let popupContent = "<b>".concat(dati[i].nome, "</b></br><a href='", labLink, "'>Vedi</a>");
            marker.bindPopup(popupContent);
            markers.push(marker);
        }
    }
}

function removeAllMarkers() {
    for (let i = 0; i < markers.length; i++) {
        markers[i].remove();
    }
    markers = [];
}

async function setMapView(luogo) {
    if(luogo!==""){
        let request = new Request(nominatimApi.concat(encodeURIComponent(luogo)), {
            method: "GET"
        });
        fetch(request).then(async function (response) {
            let jsonData = JSON.parse(await response.text());
            mymap.setView([jsonData[0].lat, jsonData[0].lon], 15);
        });
    }
}

function retrieveAllLab() {
    let request = new Request(filtersEndpoint, {
        method: "GET",
        headers: {
            "Access-Control-Allow-Origin": "*"
        }
    });
    fetch(request).then(response => setLabMap(response));
}

function parseDayOfWeek(number){
    switch (number){
        case 1:
            return 'lunedi'
        case 2:
            return 'martedi'
        case 3:
            return 'mercoledi'
        case 4:
            return 'giovedi'
        case 5:
            return 'venerdi'
        case 6:
            return 'sabato'
        case 0:
            return 'domenica'
    }
}
