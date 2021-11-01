function calcolaStatistiche() {
    mapArray = [];
    retrieveDataFromTable(1);
    let numPositivi = 0;
    let numNegativi = 0;
    let numNulli = 0;
    for (let i = 0; i < mapArray.length; i++) {
        let esito = mapArray[i].get("esito");
        if (esito === "positivo") {
            numPositivi++;
        }
        if (esito === "negativo") {
            numNegativi++;
        }
        if (esito === "nullo") {
            numNulli++;
        }
    }
    mostraStatistiche(numPositivi, numNegativi, numNulli);
}

function mostraStatistiche(numPos, numNeg, numNul) {
    let divStatistiche = document.getElementById("div-statistiche");
    if (divStatistiche == null) {
        return;
    }
    let numChildren = divStatistiche.children.length;
    for (let i = 0; i < numChildren; i++) {
        divStatistiche.children[0].remove();
    }

    let pPositivi = document.createElement("p");
    pPositivi.id = "num-positivi";
    pPositivi.innerText = "Numero Positivi: ".concat(numPos);
    let pNegativi = document.createElement("p");
    pNegativi.id = "num-negativi";
    pNegativi.innerText = "Numero Negativi: ".concat(numNeg);
    let pNulli = document.createElement("p");
    pNulli.id = "num-nulli";
    pNulli.innerText = "Numero Nulli: ".concat(numNul);

    divStatistiche.appendChild(pPositivi);
    divStatistiche.appendChild(pNegativi);
    divStatistiche.appendChild(pNulli);
}

window.onload = calcolaStatistiche;
