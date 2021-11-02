var nomePrivato;
var cognomePrivato;
var codiceFiscalePrivato;
var numeroTesseraSanitariaPrivato;
var indirizzoPrivato;
var telefonoPrivato;
var emailPrivato;
var passwordPrivato;
var confermaPasswordPrivato;
var dataNascitaPrivato;

function initElementsPrivato(){
    nomePrivato=document.getElementById("nome-privato");
    cognomePrivato=document.getElementById("cognome-privato");
    codiceFiscalePrivato=document.getElementById("codice-fiscale-privato");
    numeroTesseraSanitariaPrivato=document.getElementById("numero-tessera-sanitaria-privato");
    indirizzoPrivato=document.getElementById("indirizzo-privato");
    telefonoPrivato=document.getElementById("telefono-privato");
    emailPrivato=document.getElementById("email-privato");
    passwordPrivato=document.getElementById("password-privato");
    confermaPasswordPrivato=document.getElementById("conferma-password-privato");
    dataNascitaPrivato=document.getElementById("data-nascita-privato");

    autocomplete(indirizzoPrivato);
    aggiungiSelectPrefissi("privato");

    nomePrivato.addEventListener("focusout", checkNomePrivato);
    cognomePrivato.addEventListener("focusout", checkCognomePrivato);
    codiceFiscalePrivato.addEventListener("focusout", checkCodiceFiscalePrivato);
    numeroTesseraSanitariaPrivato.addEventListener("focusout", checkNumeroTesseraSanitariaPrivato);
    indirizzoPrivato.addEventListener("focusout", checkIndirizzoPrivato);
    telefonoPrivato.addEventListener("focusout", checkTelefonoPrivato);
    emailPrivato.addEventListener("focusout", checkEmailPrivato);
    passwordPrivato.addEventListener("focusout", checkPasswordPrivato);
    confermaPasswordPrivato.addEventListener("focusout", checkPasswordPrivato);
    dataNascitaPrivato.addEventListener("change", checkDataNascitaPrivato);
}

function checkNomePrivato(){
    if(nomePrivato.value.length<1 || nomePrivato.value.length>255 || nomePrivato.value.match(/\d/)){
        nomePrivato.style.backgroundColor="#ff7b5a";
        return false;
    }
    else{
        nomePrivato.style.backgroundColor="white";
        return true;
    }
}

function checkCognomePrivato(){
    if(cognomePrivato.value.length<1 || cognomePrivato.value.length>255 || cognomePrivato.value.match(/\d/)){
        cognomePrivato.style.backgroundColor="#ff7b5a";
        return false;
    }
    else{
        cognomePrivato.style.backgroundColor="white";
        return true;
    }
}

function checkCodiceFiscalePrivato(){
    if(codiceFiscalePrivato.value.length!==16){
        codiceFiscalePrivato.style.backgroundColor="#ff7b5a";
        return false;
    }
    else{
        codiceFiscalePrivato.style.backgroundColor="white";
        return true;
    }
}

function checkNumeroTesseraSanitariaPrivato(){
    if(numeroTesseraSanitariaPrivato.value.length!==20 || !numeroTesseraSanitariaPrivato.value.match(/^[0-9]+$/)){
        numeroTesseraSanitariaPrivato.style.backgroundColor="#ff7b5a";
        return false;
    }
    else{
        numeroTesseraSanitariaPrivato.style.backgroundColor="white";
        return true;
    }
}

function checkIndirizzoPrivato(){
    if(indirizzoPrivato.value.length<1 || indirizzoPrivato.value.length>255) {
        indirizzoPrivato.style.backgroundColor="#ff7b5a";
        return false;
    }
    else{
        indirizzoPrivato.style.backgroundColor="white";
        return true;
    }
}

function checkTelefonoPrivato(){
    if(telefonoPrivato.value.length!==10 || !telefonoPrivato.value.match(/^[0-9]+$/)){
        telefonoPrivato.style.backgroundColor="#ff7b5a";
        return false;
    }
    else{
        telefonoPrivato.style.backgroundColor="white";
        return true;
    }
}

function checkEmailPrivato(){
    if(emailPrivato.value.length<1 || emailPrivato.value.length>255 || !emailPrivato.value.match("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])")){
        emailPrivato.style.backgroundColor="#ff7b5a";
        return false;
    }
    else{
        emailPrivato.style.backgroundColor="white";
        return true;
    }
}

function checkPasswordPrivato(){
    if(passwordPrivato.value.length<1 || passwordPrivato.value.length>255 || passwordPrivato.value!==confermaPasswordPrivato.value){
        passwordPrivato.style.backgroundColor="#ff7b5a";
        confermaPasswordPrivato.style.backgroundColor="#ff7b5a";
        return false;
    }
    else{
        passwordPrivato.style.backgroundColor="white";
        confermaPasswordPrivato.style.backgroundColor="white";
        return true;
    }
}

function checkDataNascitaPrivato(){
    if(dataNascitaPrivato.value==="" || !underAgeValidate(dataNascitaPrivato.value)){
        dataNascitaPrivato.style.backgroundColor="#ff7b5a";
        return false;
    }
    else{
        dataNascitaPrivato.style.backgroundColor="white";
        return true;
    }
}

function submitPrivato(){
    if (!(checkNomePrivato() && checkCognomePrivato() && checkCodiceFiscalePrivato() && checkNumeroTesseraSanitariaPrivato() && checkIndirizzoPrivato() && checkTelefonoPrivato() && checkEmailPrivato() && checkPasswordPrivato() && checkDataNascitaPrivato())) {
        event.preventDefault();
        return false;
    }
}
