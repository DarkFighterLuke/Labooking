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
    let errMap=new Map();
    eraseErrorDivs();
    let tipologiaTest = document.getElementsByName("tipologia-test");
    let numTipologiaTest = 0;
    if(tipologiaTest==null){
        errMap.set("tipologia-test", false);
    }
    for (let i = 0; i < tipologiaTest.length; i++) {
        if (tipologiaTest[i].checked) {
            numTipologiaTest++;
        }
    }
    if (numTipologiaTest !== 1) {
        showErrorDiv(document.getElementById("div-tipologia-test").children[0], true, "È necessario selezionare una tipologia di test!");
        errMap.set("tipologia-test", false);
    }
    let tableOrariPrivati=document.getElementById("table-orari-privati");
    if (tableOrariPrivati != null && tableOrariPrivati.rows[1].cells.length === 5) {
        let numChecked=0;
        for (let i = 1; i < tableOrariPrivati.rows.length; i++) {
            if (tableOrariPrivati.rows[i].cells[0].children[0].checked) {
                numChecked++;
                let isPrivatoOk = /^[0-9]+$/.test(tableOrariPrivati.rows[i].cells[3].children[0].value);
                let isQuestionarioOk = false;
                if (tableOrariPrivati.rows[i].cells[4].children[0].files[0] != null) {
                    let file=tableOrariPrivati.rows[i].cells[4].children[0].files[0];
                    if (file!=null && file.name.substring(file.name.length - 4) === ".pdf") {
                        isQuestionarioOk = true;
                    }
                }
                if (!isPrivatoOk || !isQuestionarioOk) {
                    showErrorDiv(document.getElementById("div-orari-privati").children[0], true, "Controlla che i dati di prenotazione siano stati inseriti correttamente.");
                    errMap.set("prenotazioni", false);
                    break;
                }
            }
        }
        if(numChecked<1){
            showErrorDiv(document.getElementById("div-orari-privati").children[0], true, "È necessario selezionare almeno uno slot!");
            errMap.set("prenotazioni", false);
        }
    } else if (document.getElementById("table-orari-privati") != null && tableOrariPrivati.rows[0].cells.length===3) {
        let numChecked=0;
        for(let i=1; i<tableOrariPrivati.rows.length; i++){
            if(tableOrariPrivati.rows[i].cells[0].children[0].checked){
                numChecked++;
            }
        }
        if(numChecked!==1){
            showErrorDiv(document.getElementById("div-orari-privati").children[0], true, "È necessario prenotare 1 slot!");
            errMap.set("prenotazioni", false);
        }

        let isQuestionarioOk=false;
        let file = document.getElementById("questionario-anamnesi-upload").files[0];
        if (file!=null && file.name.substring(file.name.length - 4) === ".pdf") {
            isQuestionarioOk = true;
        }
        if(!isQuestionarioOk){
            showErrorDiv(document.getElementById("div-questionario").children[0], true, "È necessario caricare il questionario di anamnesi!");
            errMap.set("prenotazioni", false);
        }
    }
    else{
        errMap.set("generic", false);
    }

    let pagaOnline=document.getElementById("paga-online");
    let pagaPresenza=document.getElementById("paga-presenza");
    if(pagaOnline!=null && pagaOnline.checked){
        let numeroCarta=document.getElementById("numero-carta");
        let scadenza=document.getElementById("scadenza");
        let cvv=document.getElementById("cvv");
        if(numeroCarta==null || numeroCarta.value.length!==16 || !/^[0-9]+$/.test(numeroCarta.value) ||scadenza==null || scadenza.value=="" || cvv==null || cvv.value==""){
            showErrorDiv(document.getElementById("div-pagamenti").children[0], true, "Controllare i dati di pagamento inseriti e riprovare.");
            errMap.set("pagamenti", false);
        }
    }
    else if(pagaPresenza==null || !pagaPresenza.checked){
        showErrorDiv(document.getElementById("div-pagamenti").children[0], true, "Scegliere una modalità di pagamento.");
        errMap.set("pagamenti", false);
    }
    let isAllFine=true;
    errMap.forEach(function(value){
        if(value===false){
            isAllFine=false;
        }
    })
    return isAllFine;
}

function showErrorDiv(parentNode, isBefore, errMessage) {
    let div = document.createElement("div");
    div.className = "div-errore";
    let p = document.createElement("p");
    p.className = "p-errore";
    p.innerText = errMessage;
    div.appendChild(p);
    if (isBefore) {
        parentNode.before(div);
    } else {
        parentNode.after(div);
    }
}

function eraseErrorDivs() {
    let errorDivs = document.getElementsByClassName("div-errore");
    let errorDivsLength=errorDivs.length;
    for (let i = 0; i < errorDivsLength; i++) {
        errorDivs[0].parentNode.removeChild(errorDivs[0]);
    }
}
