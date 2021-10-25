var i=1;

function aggiungiRiga(){
    let table=document.getElementById("table-orari-apertura");
    i++;
    let newRow=table.insertRow(table.rows.length);
    let cellApertura=newRow.insertCell(0);
    let cellChiusura=newRow.insertCell(1);
    let cellGiorno=newRow.insertCell(2);
    let inputApertura=document.createElement("input");
    inputApertura.type="time";
    inputApertura.name="orario-apertura-".concat(i.toString(10));
    let inputChiusura=document.createElement("input");
    inputChiusura.type="time";
    inputChiusura.name="orario-chiusura-".concat(i.toString(10));
    let selectGiorno=document.createElement("select");
    selectGiorno.name="giorno-".concat(i.toString(10));
    let optLunedi=document.createElement("option");
    optLunedi.value="lunedi";
    optLunedi.text="Lunedì";
    let optMartedi=document.createElement("option");
    optMartedi.value="martedi";
    optMartedi.text="Martedì";
    let optMercoledi=document.createElement("option");
    optMercoledi.value="mercoledi";
    optMercoledi.text="Mercoledì";
    let optGiovedi=document.createElement("option");
    optGiovedi.value="giovedi";
    optGiovedi.text="Giovedì";
    let optVenerdi=document.createElement("option");
    optVenerdi.value="venerdi";
    optVenerdi.text="Venerdì";
    let optSabato=document.createElement("option");
    optSabato.value="sabato";
    optSabato.text="Sabato";
    let optDomenica=document.createElement("option");
    optDomenica.value="domenica";
    optDomenica.text="Domenica";
    selectGiorno.appendChild(optLunedi);
    selectGiorno.appendChild(optMartedi);
    selectGiorno.appendChild(optMercoledi);
    selectGiorno.appendChild(optGiovedi);
    selectGiorno.appendChild(optVenerdi);
    selectGiorno.appendChild(optSabato);
    selectGiorno.appendChild(optDomenica);
    cellApertura.appendChild(inputApertura);
    cellChiusura.appendChild(inputChiusura);
    cellGiorno.appendChild(selectGiorno);
    event.preventDefault();
    return false;
}