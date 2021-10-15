var accessToken='pk.eyJ1IjoiZGFya2ZpZ2h0ZXJsdWtlIiwiYSI6ImNrdWd6eWtkZTBlazEycW15bWd3dmRpMDUifQ.z-QRPdZnwbHgFsnAbUjVjw';
var filtersEndpoint=window.location.host.concat(":", window.location.port);

var mymap = L.map('map', {center: [41.951, 13.887], zoom: 6});
L.tileLayer('https://api.mapbox.com/styles/v1/{id}/tiles/{z}/{x}/{y}?access_token='.concat(accessToken), {
    attribution: 'Map data &copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors, Imagery © <a href="https://www.mapbox.com/">Mapbox</a>',
    maxZoom: 18,
    id: 'mapbox/streets-v11',
    tileSize: 512,
    zoomOffset: -1,
    accessToken: accessToken
}).addTo(mymap);

async function sendFilters(){
    let luogo=document.getElementById("luogo").value;
    let costo=document.getElementById("costo").value;
    let tempo=document.getElementById("tempo").value;

    let filters=new FormData();
    filters.append("costo", costo);
    filters.append("tempo", tempo);

    let request=new Request(filtersEndpoint, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: filters
    });
    fetch(request).then(response => setLabMap(response));
}

function setLabMap(response){

}
