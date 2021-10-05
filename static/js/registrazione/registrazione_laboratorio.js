var nomeLaboratorio;
var partitaIvaLaboratorio;
var cittaLaboratorio;
var capLaboratorio;
var viaLaboratorio;
var civicoLaboratorio;
var telefonoLaboratorio;
var emailLaboratorio;
var passwordLaboratorio;

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

function submitLaboratorio(){
    if(!(checkNomeLaboratorio() && checkPartitaIvaLaboratorio() && checkCittaLaboratorio() && checkCapLaboratorio() && checkViaLaboratorio() && checkCivicoLaboratorio() && checkTelefonoLaboratorio() && checkEmailLaboratorio() && checkPasswordLaboratorio())){
        event.preventDefault();
        return false;
    }
}
