var accessToken = 'pk.eyJ1IjoiZGFya2ZpZ2h0ZXJsdWtlIiwiYSI6ImNrdWd6eWtkZTBlazEycW15bWd3dmRpMDUifQ.z-QRPdZnwbHgFsnAbUjVjw';
var filtersEndpoint = "/api/ricerca";
var visualizzaLaboratorioEndpoint = "/dashboard/laboratorio";
var nominatimApi = "https://nominatim.openstreetmap.org/search?format=json&q="

var markers = [];

var parametriGet = "";
var firstLaunch = true;

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
    let luogo = document.getElementById("luogo").value;
    let numeroPersoneElement = document.getElementById("numero-persone");
    let numeroPersone = "";
    if (numeroPersoneElement != null) {
        numeroPersone = numeroPersoneElement.value;
    }
    let costo = document.getElementById("costo").value;
    let tempo = document.getElementById("tempo").value;
    let molecolare = document.getElementById("molecolare").checked;
    let antigenico = document.getElementById("antigenico").checked;
    let sierologico = document.getElementById("sierologico").checked;
    let inizio = document.getElementById("inizio-intervallo").value;
    let fine = document.getElementById("fine-intervallo").value;
    let data = document.getElementById("data").value;
    if (!checkDateTimeFields(inizio, fine, data)) {
        return false;
    }

    parametriGet = "&data=".concat(data, "&inizio=", inizio, "&fine=", fine, "&persone=", numeroPersone);

    let filters = new FormData();
    filters.append("numero-persone", numeroPersone);
    filters.append("costo", costo);
    filters.append("tempo", tempo);
    filters.append("molecolare", molecolare);
    filters.append("antigenico", antigenico);
    filters.append("sierologico", sierologico);
    filters.append("inizio-intervallo", inizio);
    filters.append("fine-intervallo", fine);
    filters.append("data", data);

    let request=new Request(filtersEndpoint, {
        method: "POST",
        body: filters
    });
    fetch(request).then(response => setLabMap(response));
    if (firstLaunch === true) firstLaunch = false;
    setMapView(luogo);
}

async function setLabMap(response) {
    removeAllMarkers();
    let dati = JSON.parse(await response.text());
    if (dati !== null) {
        for (let i = 0; i < dati.length; i++) {
            let marker = L.marker([dati[i].lat, dati[i].long]).addTo(mymap);
            let labLink = visualizzaLaboratorioEndpoint.concat("?idLab=", dati[i].id_laboratorio, parametriGet);
            let popupContent = "<b>".concat(dati[i].nome, "</b></br>");
            if (!firstLaunch) {
                popupContent = popupContent.concat("<a href='", labLink, "'>Vedi</a>");
            }
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

function checkDateTimeFields(inizio, fine, data) {
    eraseErrorDivs();
    let flag = true;
    let messaggio = "";
    if (inizio === "" || fine === "" || data === "") {
        flag = false;
        messaggio = "Gli orari di inizio e fine e la data non possono essere vuoti!";
    } else {
        let inizioObj = new Date(1, 1, 1, inizio.slice(0, 2), inizio.slice(3));
        let fineObj = new Date(1, 1, 1, fine.slice(0, 2), fine.slice(3));
        if (inizioObj.getTime() >= fineObj.getTime()) {
            flag = false;
            messaggio = "L'orario di fine non può essere inferiore a quello di inizio!";
        }
    }
    let dataObj = new Date(data);
    let today = new Date();
    dataObj.setHours(0, 0, 0, 0);
    today.setHours(0, 0, 0, 0);
    if (dataObj.getUnixTime() < today.getUnixTime()) {
        flag = false;
        messaggio = "La data non può essere precedente a quella odierna!"
    }

    if (!flag) {
        let div = document.createElement("div");
        div.className = "div-errore";
        let p = document.createElement("p");
        p.className = "p-errore";
        p.innerText = messaggio;
        div.appendChild(p);
        document.getElementById("div-datetime").after(div);
    }

    return flag;
}

function eraseErrorDivs() {
    let errorDivs = document.getElementsByClassName("div-errore");
    for (let i = 0; i < errorDivs.length; i++) {
        errorDivs[i].parentNode.removeChild(errorDivs[i]);
    }
}

Date.prototype.getUnixTime = function () {
    return this.getTime() / 1000 | 0
};
if (!Date.now) Date.now = function () {
    return new Date();
}
Date.time = function () {
    return Date.now().getUnixTime();
}
