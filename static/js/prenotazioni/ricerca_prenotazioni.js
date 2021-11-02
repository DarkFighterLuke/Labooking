var cercaId=document.getElementById("cerca-id");
var cercaDataEsecuzione=document.getElementById("cerca-data-esecuzione");
var cercaPagato=document.getElementById("cerca-pagato");
var cercaTipologiaTest=document.getElementById("cerca-tipologia-test");
var cercaStato=document.getElementById("cerca-stato");
var cercaPrivato = document.getElementById("cerca-privato");
var cercaEsito = document.getElementById("cerca-esito");

var mapArray=[];
var foundArray=[];

function retrieveDataFromTable(startIndex) {
    let table = document.getElementById("table-orari-privati");
    for (let i = startIndex; i < table.rows.length; i++) {
        let m = new Map();
        m.set("id", table.rows[i].cells[0].innerText);
        m.set("dataEsecuzione", table.rows[i].cells[1].innerText);
        m.set("pagato", table.rows[i].cells[2].innerText);
        m.set("tipologiaTest", table.rows[i].cells[3].innerText);
        m.set("stato", table.rows[i].cells[4].innerText);
        m.set("privato", table.rows[i].cells[5].innerText);
        m.set("esito", table.rows[i].cells[6].innerText)
        m.set("referto", table.rows[i].cells[7].innerHTML);
        m.set("questionario", table.rows[i].cells[8].innerHTML);
        mapArray.push(m);
    }
}

function cercaPrenotazioni(){
    mapArray=[];
    foundArray = [];
    retrieveDataFromTable(2);

    let searchValues=new Map();
    searchValues.set("id", cercaId.value);
    searchValues.set("dataEsecuzione", cercaDataEsecuzione.value);
    searchValues.set("pagato", cercaPagato.value);
    searchValues.set("tipologiaTest", cercaTipologiaTest.value);
    searchValues.set("stato", cercaStato.value);
    searchValues.set("privato", cercaPrivato.value);
    searchValues.set("esito", cercaEsito.value);

    let parametersToSearch=[];
    if(searchValues.get("id")!==""){
        parametersToSearch.push("id");
    }
    if(searchValues.get("dataEsecuzione")!==""){
        parametersToSearch.push("dataEsecuzione");
    }
    if(searchValues.get("pagato")!==""){
        parametersToSearch.push("pagato");
    }
    if (searchValues.get("tipologiaTest") !== "") {
        parametersToSearch.push("tipologiaTest");
    }
    if (searchValues.get("stato") !== "") {
        parametersToSearch.push("stato");
    }
    if (searchValues.get("privato") !== "") {
        parametersToSearch.push("privato");
    }
    if (searchValues.get("esito") !== "") {
        parametersToSearch.push("esito");
    }

    for (let i = 0; i < mapArray.length; i++) {
        let matchCounter = 0;
        for (let j = 0; j < parametersToSearch.length; j++) {
            if (parametersToSearch[j] === "dataEsecuzione") {
                if (mapArray[i].get("dataEsecuzione").substring(0, 10) === searchValues.get("dataEsecuzione")) {
                    matchCounter++;
                    continue;
                }
            }
            if(parametersToSearch[j]==="privato"){
                if(mapArray[i].get("privato").toLowerCase().includes(searchValues.get("privato").toLowerCase())){
                    matchCounter++;
                    continue;
                }
            }
            if(mapArray[i].get(parametersToSearch[j])===searchValues.get(parametersToSearch[j])){
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
    document.getElementById("table-orari-privati").remove();
    let resultsTable=document.createElement("table");
    resultsTable.className = "table-responsive";
    resultsTable.id = "table-orari-privati";
    let parentDiv = document.getElementById("div-tabella");
    resultsTable.insertRow(0);
    resultsTable.rows[0].insertCell(0).outerHTML = "<th>Id test</th>";
    resultsTable.rows[0].insertCell(1).outerHTML = "<th>Data esecuzione</th>";
    resultsTable.rows[0].insertCell(2).outerHTML = "<th>Pagato online</th>";
    resultsTable.rows[0].insertCell(3).outerHTML = "<th>Tipologia test</th>";
    resultsTable.rows[0].insertCell(4).outerHTML = "<th>Stato</th>";
    resultsTable.rows[0].insertCell(5).outerHTML = "<th>Privato</th>";
    resultsTable.rows[0].insertCell(6).outerHTML = "<th>Esito</th>";
    resultsTable.rows[0].insertCell(7).outerHTML = "<th>Referto</th>";
    resultsTable.rows[0].insertCell(8).outerHTML = "<th>Questionario anamnesi</th>";
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
        resultsTable.rows[1].cells[0].innerText = foundArray[i].get("id");
        resultsTable.rows[1].cells[1].innerText = foundArray[i].get("dataEsecuzione");
        resultsTable.rows[1].cells[2].innerText = foundArray[i].get("pagato");
        resultsTable.rows[1].cells[3].innerText = foundArray[i].get("tipologiaTest");
        resultsTable.rows[1].cells[4].innerText = foundArray[i].get("stato");
        resultsTable.rows[1].cells[5].innerText = foundArray[i].get("privato");
        resultsTable.rows[1].cells[6].innerHTML = foundArray[i].get("esito");
        resultsTable.rows[1].cells[7].innerHTML = foundArray[i].get("referto");
        resultsTable.rows[1].cells[8].innerHTML = foundArray[i].get("questionario");
    }

    let eraseBtn = document.createElement("button");
    eraseBtn.onclick = eraseFilters;
    eraseBtn.innerText = "Cancella filtri";
    eraseBtn.className = "bg-lightblue";

    parentDiv.appendChild(eraseBtn);
    parentDiv.appendChild(resultsTable);

    calcolaStatistiche();
}

function eraseFilters(){
    location.reload();
}
