var ragioneSocialeOrganizzazione;
var partitaIvaOrganizzazione;
var indirizzoOrganizzazione;
var telefonoOrganizzazione;
var emailOrganizzazione;
var passwordOrganizzazione;
var confermaPasswordOrganizzazione;

function initElementsOrganizzazione(){
    ragioneSocialeOrganizzazione=document.getElementById("ragione-sociale-organizzazione");
    partitaIvaOrganizzazione=document.getElementById("partita-iva-organizzazione");
    indirizzoOrganizzazione=document.getElementById("indirizzo-organizzazione");
    telefonoOrganizzazione=document.getElementById("telefono-organizzazione");
    emailOrganizzazione=document.getElementById("email-organizzazione");
    passwordOrganizzazione=document.getElementById("password-organizzazione");
    confermaPasswordOrganizzazione=document.getElementById("conferma-password-organizzazione");

    autocomplete(indirizzoOrganizzazione);
    aggiungiSelectPrefissi("organizzazione");

    ragioneSocialeOrganizzazione.addEventListener("focusout", checkRagioneSocialeOrganizzazione);
    partitaIvaOrganizzazione.addEventListener("focusout", checkPartitaIvaOrganizzazione);
    indirizzoOrganizzazione.addEventListener("focusout", checkIndirizzoOrganizzazione());
    telefonoOrganizzazione.addEventListener("focusout", checkTelefonoOrganizzazione);
    emailOrganizzazione.addEventListener("focusout", checkEmailOrganizzazione);
    passwordOrganizzazione.addEventListener("focusout", checkPasswordOrganizzazione);
    confermaPasswordOrganizzazione.addEventListener("focusout", checkPasswordOrganizzazione);
}

function checkRagioneSocialeOrganizzazione(){
    if(ragioneSocialeOrganizzazione.value.length<1 || ragioneSocialeOrganizzazione.value.length>255 || ragioneSocialeOrganizzazione.value.match(/\d/)){
        ragioneSocialeOrganizzazione.style.backgroundColor="red";
        return false;
    }
    else{
        ragioneSocialeOrganizzazione.style.backgroundColor="white";
        return true;
    }
}

function checkPartitaIvaOrganizzazione(){
    if(partitaIvaOrganizzazione.value.length!==11){
        partitaIvaOrganizzazione.style.backgroundColor="red";
        return false;
    }
    else{
        partitaIvaOrganizzazione.style.backgroundColor="white";
        return true;
    }
}

function checkIndirizzoOrganizzazione(){
    if(indirizzoOrganizzazione.value.length < 0 || indirizzoOrganizzazione.value.length > 255){
        indirizzoOrganizzazione.style.backgroundColor="red";
        return false;
    }
    else{
        indirizzoOrganizzazione.style.backgroundColor="white";
    }
}

function checkTelefonoOrganizzazione(){
    if(telefonoOrganizzazione.value.length!==10 || !telefonoOrganizzazione.value.match(/^[0-9]+$/)){
        telefonoOrganizzazione.style.backgroundColor="red";
        return false;
    }
    else{
        telefonoOrganizzazione.style.backgroundColor="white";
        return true;
    }
}

function checkEmailOrganizzazione(){
    if(emailOrganizzazione.value.length<1 || emailOrganizzazione.value.length>255 || !emailOrganizzazione.value.match("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])")){
        emailOrganizzazione.style.backgroundColor="red";
        return false;
    }
    else{
        emailOrganizzazione.style.backgroundColor="white";
        return true;
    }
}

function checkPasswordOrganizzazione(){
    if(passwordOrganizzazione.value.length<1 || passwordOrganizzazione.value.length>255 || passwordOrganizzazione.value!==confermaPasswordOrganizzazione.value){
        passwordOrganizzazione.style.backgroundColor="red";
        confermaPasswordOrganizzazione.style.backgroundColor="red";
        return false;
    }
    else{
        passwordOrganizzazione.style.backgroundColor="white";
        confermaPasswordOrganizzazione.style.backgroundColor="white";
        return true;
    }
}

function submitOrganizzazione(){
    if(!(checkRagioneSocialeOrganizzazione() && checkPartitaIvaOrganizzazione() && checkIndirizzoOrganizzazione() && checkTelefonoOrganizzazione() && checkEmailOrganizzazione() && checkPasswordOrganizzazione())){
        event.preventDefault();
        return false;
    }
}

