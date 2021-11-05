var nomeLaboratorio;
var ibanLaboratorio;
var partitaIvaLaboratorio;
var indirizzoLaboratorio;
var telefonoLaboratorio;
var emailLaboratorio;
var passwordLaboratorio;
var confermaPasswordLaboratorio;
var testPerOraLaboratorio;
var tableOrariApertura;
var tableInfoTest;

function initElementsLaboratorio(){
    nomeLaboratorio = document.getElementById("nome-laboratorio");
    ibanLaboratorio = document.getElementById("iban-laboratorio");
    partitaIvaLaboratorio = document.getElementById("partita-iva-laboratorio");
    indirizzoLaboratorio=document.getElementById("indirizzo-laboratorio");
    telefonoLaboratorio=document.getElementById("telefono-laboratorio");
    emailLaboratorio=document.getElementById("email-laboratorio");
    passwordLaboratorio=document.getElementById("password-laboratorio");
    confermaPasswordLaboratorio = document.getElementById("conferma-password-laboratorio");
    testPerOraLaboratorio = document.getElementById("test-per-ora-laboratorio");
    tableOrariApertura = document.getElementById("table-orari-apertura");
    tableInfoTest=document.getElementById("table-info-test");

    autocomplete(indirizzoLaboratorio);
    aggiungiSelectPrefissi("laboratorio");

    nomeLaboratorio.addEventListener("focusout", checkNomeLaboratorio);
    ibanLaboratorio.addEventListener("focusout", checkIbanLaboratorio);
    partitaIvaLaboratorio.addEventListener("focusout", checkPartitaIvaLaboratorio);
    indirizzoLaboratorio.addEventListener("focusout", checkIndirizzoLaboratorio);
    telefonoLaboratorio.addEventListener("focusout", checkTelefonoLaboratorio);
    emailLaboratorio.addEventListener("focusout", checkEmailLaboratorio);
    passwordLaboratorio.addEventListener("focusout", checkPasswordLaboratorio);
    confermaPasswordLaboratorio.addEventListener("focusout", checkPasswordLaboratorio);
    testPerOraLaboratorio.addEventListener("focusout", checkTestPerOraLaboratorio);
}

function checkNomeLaboratorio() {
    if (nomeLaboratorio.value.length < 1 || nomeLaboratorio.value.length > 255 || nomeLaboratorio.value.match(/\d/)) {
        nomeLaboratorio.style.backgroundColor = "#ff7b5a";
        return false;
    } else {
        nomeLaboratorio.style.backgroundColor = "white";
        return true;
    }
}

function checkIbanLaboratorio() {
    if (ibanLaboratorio.value.length < 15) {
        ibanLaboratorio.style.backgroundColor = "#ff7b5a";
        return false;
    } else {
        ibanLaboratorio.style.backgroundColor = "white";
        return true;
    }
}

function checkPartitaIvaLaboratorio() {
    if (partitaIvaLaboratorio.value.length !== 11) {
        partitaIvaLaboratorio.style.backgroundColor = "#ff7b5a";
        return false;
    } else {
        partitaIvaLaboratorio.style.backgroundColor = "white";
        return true;
    }
}

function checkIndirizzoLaboratorio(){
    if(indirizzoLaboratorio.value.length<1 || indirizzoLaboratorio.value.length>255){
        indirizzoLaboratorio.style.backgroundColor="#ff7b5a";
        return false;
    }
    else{
        indirizzoLaboratorio.style.backgroundColor="white";
        return true;
    }
}

function checkTelefonoLaboratorio(){
    if(telefonoLaboratorio.value.length!==10 || !telefonoLaboratorio.value.match(/^[0-9]+$/)){
        telefonoLaboratorio.style.backgroundColor="#ff7b5a";
        return false;
    }
    else{
        telefonoLaboratorio.style.backgroundColor="white";
        return true;
    }
}

function checkEmailLaboratorio(){
    if(emailLaboratorio.value.length<1 || emailLaboratorio.value.length>255 || !emailLaboratorio.value.match("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])")){
        emailLaboratorio.style.backgroundColor="#ff7b5a";
        return false;
    }
    else{
        emailLaboratorio.style.backgroundColor="white";
        return true;
    }
}

function checkPasswordLaboratorio(){
    if(passwordLaboratorio.value.length<1 || passwordLaboratorio.value.length>255 || passwordLaboratorio.value!==confermaPasswordLaboratorio.value){
        passwordLaboratorio.style.backgroundColor = "#ff7b5a";
        confermaPasswordLaboratorio.style.backgroundColor = "#ff7b5a";
        return false;
    } else {
        passwordLaboratorio.style.backgroundColor = "white";
        confermaPasswordLaboratorio.style.backgroundColor = "white";
        return true;
    }
}

function checkTestPerOraLaboratorio() {
    if (testPerOraLaboratorio.value < 1) {
        testPerOraLaboratorio.style.backgroundColor = "#ff7b5a";
        return false;
    } else {
        testPerOraLaboratorio.style.backgroundColor = "white";
        return true;
    }
}

function checkOrariApertura() {
    eraseErrorDivs();
    let orariApertura = [];
    let orariChiusura = [];
    let giorni = [];
    for (let i = 1; i < tableOrariApertura.rows.length; i++) {
        let oa = tableOrariApertura.rows[i].cells[0].children[0].valueAsDate;
        let oc = tableOrariApertura.rows[i].cells[1].children[0].valueAsDate;
        let giorno = tableOrariApertura.rows[i].cells[2].children[0].value;

        if (oa == null || oc == null || oa.getTime() >= oc.getTime()) {
            let messaggio = "Sembra che qualcosa non vada con gli orari inseriti. Controlla che siano corretti e riprova.";
            showErrorDiv(tableOrariApertura, true, messaggio);
            return false;
        }

        for (let j = 0; j < orariApertura.length; j += 2) {
            let oatemp = orariApertura[j];
            let octemp = orariChiusura[j];
            let gtemp = giorni[j];
            if (oa.getTime() > oatemp.getTime() && oa.getTime() < octemp.getTime() && gtemp === giorno || oc.getTime() > oatemp.getTime() && oc.getTime() < octemp.getTime() && gtemp === giorno) {
                let messaggio = "Sembra che qualcosa non vada con gli orari inseriti. Controlla che siano corretti e riprova.";
                showErrorDiv(tableOrariApertura, true, messaggio);
                return false;
            }
        }

        orariApertura.push(oa);
        orariChiusura.push(oc);
        giorni.push(giorno);
    }
    return true;
}

function checkInfoTest(){
    eraseErrorDivs();
    let ore=[];
    let minuti=[];
    let costo=[];
    let effettua=[];

    let checked = 0;
    for (let i = 0; i < 3; i++) {
        effettua[i] = tableInfoTest.rows[i + 1].cells[3].children[0].checked;
        if (effettua[i] == true) {
            ore[i] = tableInfoTest.rows[i + 1].cells[1].children[0].value;
            minuti[i] = tableInfoTest.rows[i + 1].cells[1].children[1].value;
            costo[i] = tableInfoTest.rows[i + 1].cells[2].children[0].value;
            checked++;
        }
    }
    if (checked < 1) {
        let messaggio = "Seleziona almeno un tipo di test diagnostico e riprova.";
        showErrorDiv(tableInfoTest, true, messaggio);
        return false;
    }

    for (let i = 0; i < 3; i++) {
        if (effettua[i] == true) {
            if (ore[i] === "" || ore[i] < 0 || (minuti[i] != 0 && minuti[i] != 15 && minuti[i] != 30 && minuti[i] != 45) || costo[i] === "" || costo[i] < 0 || costo[i] > 9999.99 || (ore[i] == 0 && minuti[i] == 0)) {
                let messaggio = "Sembra che qualcosa non vada con le info sui test diagnostici effettuati. Controlla che i dati siano corretti e riprova.";
                showErrorDiv(tableInfoTest, true, messaggio);
                return false;
            }
        }
    }
    return true;
}

function submitLaboratorio(){
    if (!(checkNomeLaboratorio() && checkIbanLaboratorio() && checkPartitaIvaLaboratorio() && checkIndirizzoLaboratorio() && checkTelefonoLaboratorio() && checkEmailLaboratorio() && checkPasswordLaboratorio() && checkIbanLaboratorio() && checkOrariApertura() && checkInfoTest())) {
        event.preventDefault();
        return false;
    }
}
