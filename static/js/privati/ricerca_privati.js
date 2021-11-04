var cercaId=document.getElementById("cerca-id");
var cercaNome=document.getElementById("cerca-nome");
var cercaCognome=document.getElementById("cerca-cognome");
var cercaCodiceFiscale=document.getElementById("cerca-codice-fiscale");
var cercaIndirizzo=document.getElementById("cerca-indirizzo");
var cercaTelefono = document.getElementById("cerca-telefono");
var cercaEmail = document.getElementById("cerca-email");
var cercaDataNascita = document.getElementById("cerca-data-nascita");

var mapArray=[];
var foundArray=[];

function retrieveDataFromTable(startIndex) {
    let table = document.getElementById("table-privati");
    for (let i = startIndex; i < table.rows.length; i++) {
        let m = new Map();
        m.set("checkbox", table.rows[i].cells[0].innerHTML)
        m.set("id", table.rows[i].cells[1].innerText);
        m.set("nome", table.rows[i].cells[2].innerText);
        m.set("cognome", table.rows[i].cells[3].innerText);
        m.set("codiceFiscale", table.rows[i].cells[4].innerText);
        m.set("indirizzo", table.rows[i].cells[5].innerText);
        m.set("prefisso", table.rows[i].cells[6].innerText);
        m.set("telefono", table.rows[i].cells[7].innerText);
        m.set("email", table.rows[i].cells[8].innerText)
        m.set("dataNascita", table.rows[i].cells[9].innerText);
        mapArray.push(m);
    }
}

function cercaPrivati(){
    mapArray=[];
    foundArray = [];
    retrieveDataFromTable(2);

    let searchValues=new Map();
    searchValues.set("id", cercaId.value);
    searchValues.set("nome", cercaNome.value);
    searchValues.set("cognome", cercaCognome.value);
    searchValues.set("codiceFiscale", cercaCodiceFiscale.value);
    searchValues.set("indirizzo", cercaIndirizzo.value);
    searchValues.set("telefono", cercaTelefono.value);
    searchValues.set("email", cercaEmail.value);
    searchValues.set("dataNascita", cercaDataNascita.value);

    let parametersToSearch=[];
    if(searchValues.get("id")!==""){
        parametersToSearch.push("id");
    }
    if(searchValues.get("nome")!==""){
        parametersToSearch.push("nome");
    }
    if(searchValues.get("cognome")!==""){
        parametersToSearch.push("cognome");
    }
    if (searchValues.get("codiceFiscale") !== "") {
        parametersToSearch.push("codiceFiscale");
    }
    if (searchValues.get("indirizzo") !== "") {
        parametersToSearch.push("indirizzo");
    }
    if (searchValues.get("telefono") !== "") {
        parametersToSearch.push("telefono");
    }
    if (searchValues.get("email") !== "") {
        parametersToSearch.push("email");
    }
    if (searchValues.get("dataNascita") !== "") {
        parametersToSearch.push("dataNascita");
    }

    for (let i = 0; i < mapArray.length; i++) {
        let matchCounter = 0;
        for (let j = 0; j < parametersToSearch.length; j++) {
            if(parametersToSearch[j]==="id"){
                if(mapArray[i].get("id")===searchValues.get("id")){
                    matchCounter++;
                    continue;
                }
            }
            if(parametersToSearch[j]==="telefono"){
                if(mapArray[i].get("telefono")===searchValues.get("telefono")){
                    matchCounter++;
                    continue;
                }
            }
            if(parametersToSearch[j]==="dataNascita"){
                if(mapArray[i].get("dataNascita")===searchValues.get("dataNascita")){
                    matchCounter++;
                    continue;
                }
            }
            if(mapArray[i].get(parametersToSearch[j]).toLowerCase().includes(searchValues.get(parametersToSearch[j]).toLowerCase())){
                matchCounter++;
            }
        }
        if(matchCounter===parametersToSearch.length){
            foundArray.push(mapArray[i]);
        }
    }

    showSearchResults();
}

function showSearchResults(){
    document.getElementById("table-privati").remove();
    let resultsTable=document.createElement("table");
    resultsTable.className = "table-responsive";
    resultsTable.id = "table-privati";
    let parentDiv = document.getElementById("div-tabella-privati");
    resultsTable.insertRow(0);
    resultsTable.rows[0].insertCell(0).outerHTML = "<th></th>";
    resultsTable.rows[0].insertCell(1).outerHTML = "<th>Id Privato</th>";
    resultsTable.rows[0].insertCell(2).outerHTML = "<th>Nome</th>";
    resultsTable.rows[0].insertCell(3).outerHTML = "<th>Cognome</th>";
    resultsTable.rows[0].insertCell(4).outerHTML = "<th>Codice Fiscale</th>";
    resultsTable.rows[0].insertCell(5).outerHTML = "<th>Indirizzo</th>";
    resultsTable.rows[0].insertCell(6).outerHTML = "<th>Prefisso</th>";
    resultsTable.rows[0].insertCell(7).outerHTML = "<th>Telefono</th>";
    resultsTable.rows[0].insertCell(8).outerHTML = "<th>Email</th>";
    resultsTable.rows[0].insertCell(9).outerHTML = "<th>Data di nascita</th>";
    for (let i = 0; i < foundArray.length; i++) {
        resultsTable.insertRow(1);
        resultsTable.rows[1].insertCell(0);
        resultsTable.rows[1].insertCell(1);
        resultsTable.rows[1].insertCell(2);
        resultsTable.rows[1].insertCell(3);
        resultsTable.rows[1].insertCell(4);
        resultsTable.rows[1].insertCell(5);
        resultsTable.rows[1].insertCell(6);
        resultsTable.rows[1].insertCell(7);
        resultsTable.rows[1].insertCell(8);
        resultsTable.rows[1].insertCell(9);
        resultsTable.rows[1].cells[0].innerHTML = foundArray[i].get("checkbox");
        resultsTable.rows[1].cells[1].innerText = foundArray[i].get("id");
        resultsTable.rows[1].cells[2].innerText = foundArray[i].get("nome");
        resultsTable.rows[1].cells[3].innerText = foundArray[i].get("cognome");
        resultsTable.rows[1].cells[4].innerText = foundArray[i].get("codiceFiscale");
        resultsTable.rows[1].cells[5].innerText = foundArray[i].get("indirizzo");
        resultsTable.rows[1].cells[6].innerText = foundArray[i].get("prefisso");
        resultsTable.rows[1].cells[7].innerText = foundArray[i].get("telefono");
        resultsTable.rows[1].cells[8].innerText = foundArray[i].get("email");
        resultsTable.rows[1].cells[9].innerText = foundArray[i].get("dataNascita");
    }

    let eraseBtn = document.createElement("button");
    eraseBtn.onclick = eraseFilters;
    eraseBtn.innerText = "Cancella filtri";
    eraseBtn.className = "bg-lightblue";

    parentDiv.appendChild(eraseBtn);
    parentDiv.appendChild(resultsTable);
}

function eraseFilters(){
    location.reload();
}
