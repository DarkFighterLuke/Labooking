function mostraCampi() {
    if (document.getElementById("paga-online").checked) {
        document.getElementById("dati-pagamento-online").style = "display: block";
        document.getElementById("dati-pagamento-presenza").style = "display: none";
    } else {
        document.getElementById("dati-pagamento-online").style = "display: none";
        document.getElementById("dati-pagamento-presenza").style = "display: block";
    }
}


document.getElementById("dati-pagamento-presenza").style = "display: none";

function checkDatiInseriti() {
    let tipologiaTest = document.getElementsByName("tipologia-test");
    let numTipologiaTest = 0;
    for (let i = 0; i < tipologiaTest.length; i++) {
        if (tipologiaTest[i].checked) {
            numTipologiaTest++;
        }
    }
    if (numTipologiaTest !== 1) {
        // TODO: Stampare messaggio di errore
        return false;
    }
    if (document.getElementById("table-orari-privati").rows[0].cells === 4) {
        for (let i = 1; i < document.getElementById("table-orari-privati").rows.length; i++) {
            if (document.getElementById("table-orari-privati").rows[1].cells[0].children[0].checked) {
                let isPrivatoOk = /^[0-9]+$/.test(document.getElementById("table-orari-privati").rows[1].cells[3].children[0].value);
                let isQuestionarioOk = false;
                if (document.getElementById("table-orari-privati").rows[2].cells[4].children[0].files[0] != null) {
                    let fileName = document.getElementById("table-orari-privati").rows[2].cells[4].children[0].files[0].name;
                    if (fileName.substring(fileName.length - 4) === ".pdf") {
                        isQuestionarioOk = true;
                    }
                }
                if (!isPrivatoOk || !isQuestionarioOk) {
                    // TODO: Stampare messaggio di errore
                    return false;
                }
            }
        }
    } else if (document.getElementById("table-orari") != null) {

    }

}

function showErrorDiv(parentNode, isBefore, errMessage) {
    let div = document.createElement("div");
    div.className = "div-errore";
    let p = document.createElement("p");
    p.className = "p-errore";
    p.innerText = errMessage;
    div.appendChild(p);
    if (isBefore) {
        document.getElementById("div-datetime").before(div);
    } else {
        document.getElementById("div-datetime").after(div);
    }
}

function eraseErrorDivs() {
    let errorDivs = document.getElementsByClassName("div-errore");
    for (let i = 0; i < errorDivs.length; i++) {
        errorDivs[i].parentNode.removeChild(errorDivs[i]);
    }
}
