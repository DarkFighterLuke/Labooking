var nomeLaboratorio;
var partitaIvaLaboratorio;
var cittaLaboratorio;
var capLaboratorio;
var viaLaboratorio;
var civicoLaboratorio;
var telefonoLaboratorio;
var emailLaboratorio;
var passwordLaboratorio;
var tableOrariApertura;
var tableInfoTest;

function initElementsLaboratorio(){
    nomeLaboratorio=document.getElementById("nome-laboratorio");
    partitaIvaLaboratorio=document.getElementById("partita-iva-laboratorio");
    cittaLaboratorio=document.getElementById("citta-laboratorio");
    capLaboratorio=document.getElementById("cap-laboratorio");
    viaLaboratorio=document.getElementById("via-laboratorio");
    civicoLaboratorio=document.getElementById("civico-laboratorio");
    telefonoLaboratorio=document.getElementById("telefono-laboratorio");
    emailLaboratorio=document.getElementById("email-laboratorio");
    passwordLaboratorio=document.getElementById("password-laboratorio");
    confermaPasswordLaboratorio=document.getElementById("conferma-password-laboratorio");
    tableOrariApertura=document.getElementById("table-orari-apertura");
    tableInfoTest=document.getElementById("table-info-test");

    aggiungiSelectPrefissi("laboratorio");

    nomeLaboratorio.addEventListener("focusout", checkNomeLaboratorio);
    partitaIvaLaboratorio.addEventListener("focusout", checkPartitaIvaLaboratorio);
    cittaLaboratorio.addEventListener("focusout", checkCittaLaboratorio);
    capLaboratorio.addEventListener("focusout", checkCapLaboratorio);
    viaLaboratorio.addEventListener("focusout", checkViaLaboratorio);
    civicoLaboratorio.addEventListener("focusout", checkCivicoLaboratorio);
    telefonoLaboratorio.addEventListener("focusout", checkTelefonoLaboratorio);
    emailLaboratorio.addEventListener("focusout", checkEmailLaboratorio);
    passwordLaboratorio.addEventListener("focusout", checkPasswordLaboratorio);
    confermaPasswordLaboratorio.addEventListener("focusout", checkPasswordLaboratorio);
    codiceRegionale.addEventListener("change", checkCodiceRegionale);
}

function checkNomeLaboratorio(){
    if(nomeLaboratorio.value.length<1 || nomeLaboratorio.value.length>255 || nomeLaboratorio.value.match(/\d/)){
        nomeLaboratorio.style.backgroundColor="red";
        return false;
    }
    else{
        nomeLaboratorio.style.backgroundColor="white";
        return true;
    }
}

function checkPartitaIvaLaboratorio(){
    if(partitaIvaLaboratorio.value.length!==11){
        partitaIvaLaboratorio.style.backgroundColor="red";
        return false;
    }
    else{
        partitaIvaLaboratorio.style.backgroundColor="white";
        return true;
    }
}

function checkCittaLaboratorio(){
    if(cittaLaboratorio.value.length<1 || cittaLaboratorio.value.length>255 || cittaLaboratorio.value.match(/\d/)){
        cittaLaboratorio.style.backgroundColor="red";
        return false;
    }
    else{
        cittaLaboratorio.style.backgroundColor="white";
        return true;
    }
}

function checkCapLaboratorio(){
    if(capLaboratorio.value.length!==5 || !capLaboratorio.value.match(/^[0-9]+$/)){
        capLaboratorio.style.backgroundColor="red";
        return false;
    }
    else{
        capLaboratorio.style.backgroundColor="white";
        return true;
    }
}

function checkViaLaboratorio(){
    if(viaLaboratorio.value.length<1 || viaLaboratorio.value.length>255){
        viaLaboratorio.style.backgroundColor="red";
        return false;
    }
    else{
        viaLaboratorio.style.backgroundColor="white";
        return true;
    }
}

function checkCivicoLaboratorio(){
    if(civicoLaboratorio.value.length<1 || civicoLaboratorio.value.length>4 || !civicoLaboratorio.value.match(/^[0-9]+$/) || civicoLaboratorio.value==0){
        civicoLaboratorio.style.backgroundColor="red";
        return false;
    }
    else{
        civicoLaboratorio.style.backgroundColor="white";
        return true;
    }
}

function checkTelefonoLaboratorio(){
    if(telefonoLaboratorio.value.length!==10 || !telefonoLaboratorio.value.match(/^[0-9]+$/)){
        telefonoLaboratorio.style.backgroundColor="red";
        return false;
    }
    else{
        telefonoLaboratorio.style.backgroundColor="white";
        return true;
    }
}

function checkEmailLaboratorio(){
    if(emailLaboratorio.value.length<1 || emailLaboratorio.value.length>255 || !emailLaboratorio.value.match("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])")){
        emailLaboratorio.style.backgroundColor="red";
        return false;
    }
    else{
        emailLaboratorio.style.backgroundColor="white";
        return true;
    }
}

function checkPasswordLaboratorio(){
    if(passwordLaboratorio.value.length<1 || passwordLaboratorio.value.length>255 || passwordLaboratorio.value!==confermaPasswordLaboratorio.value){
        passwordLaboratorio.style.backgroundColor="red";
        confermaPasswordLaboratorio.style.backgroundColor="red";
        return false;
    }
    else{
        passwordLaboratorio.style.backgroundColor="white";
        confermaPasswordLaboratorio.style.backgroundColor="white";
        return true;
    }
}

function checkCodiceRegionale(){
    if(codiceRegionale.value.length!==6){
        codiceRegionale.style.backgroundColor="red";
        return false;
    }
}

function checkOrariApertura(){
    eraseErrorDivs();
    let orariApertura=[];
    let orariChiusura=[];
    let giorni=[];
    for(let i=1; i<tableOrariApertura.rows.length; i++){
        let oa=tableOrariApertura.rows[i].cells[0].childNodes[1].valueAsDate;
        let oc=tableOrariApertura.rows[i].cells[1].childNodes[1].valueAsDate;
        let giorno=tableOrariApertura.rows[i].cells[2].childNodes[1].value;

        if(oa==null || oc==null || oa.getTime() >= oc.getTime()){
            mostraMessaggioErroreOrari();
            return false;
        }

        for(let j=0; j<orariApertura.length; j+=2){
            let oatemp=orariApertura[j];
            let octemp=orariChiusura[j+1];
            let gtemp=giorni[j];
            if(oa.getTime()>oatemp.getTime() && oa.getTime()<octemp.getTime() && gtemp===giorno || oc.getTime()>oatemp.getTime() && oc.getTime()>octemp.getTime() && gtemp===giorno){
                mostraMessaggioErroreOrari();
                return false;
            }
        }

        orariApertura.push(oa);
        orariChiusura.push(oc);
    }
    return true;
}

function mostraMessaggioErroreOrari(){
    console.log("bla");
    let div=document.createElement("div");
    div.className="div-errore";
    let p=document.createElement("p");
    p.className="p-errore";
    let messaggio="Sembra che qualcosa non vada con gli orari inseriti. Controlla che siano corretti e riprova.";
    p.innerText=messaggio;
    div.appendChild(p);
    tableOrariApertura.before(div);
}

function eraseErrorDivs(){
    let errorDivs=document.getElementsByClassName("div-errore");
    for(let i=0; i<errorDivs.length; i++){
        errorDivs[i].parentNode.removeChild(errorDivs[i]);
    }
}

function checkInfoTest(){
    eraseErrorDivs();
    let ore=[];
    let minuti=[];
    let costo=[];
    let effettua=[];

    for(let i=0; i<3; i++){
        effettua[i]=tableInfoTest.rows[i+1].cells[3].childNodes[1].checked;
        if(effettua[i]==true){
            ore[i]=tableInfoTest.rows[i+1].cells[1].childNodes[1].value;
            minuti[i]=tableInfoTest.rows[i+1].cells[1].childNodes[3].value;
            costo[i]=tableInfoTest.rows[i+1].cells[2].childNodes[1].value;
        }
    }

    for(let i=0; i<3; i++){
        if(ore[i]<0 || (minuti[i]!=0 && minuti[i]!=15 && minuti[i]!=30 && minuti[i]!=45) || costo[i]<0 || costo[i]>9999.99){
            mostraMessaggioErroreInfoTest();
            return false;
        }
    }
    return true;
}

function mostraMessaggioErroreInfoTest(){
    let div=document.createElement("div");
    div.className="div-errore";
    let p=document.createElement("p");
    p.className="p-errore";
    let messaggio="Sembra che qualcosa non vada con le info sui test diagnostici effettuati. Controlla che i dati siano corretti e riprova.";
    p.innerText=messaggio;
    div.appendChild(p);
    tableInfoTest.before(div);
}

function submitLaboratorio(){
    if(!(checkNomeLaboratorio() && checkPartitaIvaLaboratorio() && checkCittaLaboratorio() && checkCapLaboratorio() && checkViaLaboratorio() && checkCivicoLaboratorio() && checkTelefonoLaboratorio() && checkEmailLaboratorio() && checkPasswordLaboratorio() && checkOrariApertura() && checkInfoTest())){
        event.preventDefault();
        return false;
    }
}
