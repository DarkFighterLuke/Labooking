var nomePrivato;
var cognomePrivato;
var codiceFiscalePrivato;
var numeroTesseraSanitariaPrivato;
var cittaPrivato;
var capPrivato;
var viaPrivato;
var civicoPrivato;
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
    cittaPrivato=document.getElementById("citta-privato");
    capPrivato=document.getElementById("cap-privato");
    viaPrivato=document.getElementById("via-privato");
    civicoPrivato=document.getElementById("civico-privato");
    telefonoPrivato=document.getElementById("telefono-privato");
    emailPrivato=document.getElementById("email-privato");
    passwordPrivato=document.getElementById("password-privato");
    confermaPasswordPrivato=document.getElementById("conferma-password-privato");
    dataNascitaPrivato=document.getElementById("data-nascita-privato");

    aggiungiSelectPrefissi("privato");

    nomePrivato.addEventListener("focusout", checkNomePrivato);
    cognomePrivato.addEventListener("focusout", checkCognomePrivato);
    codiceFiscalePrivato.addEventListener("focusout", checkCodiceFiscalePrivato);
    numeroTesseraSanitariaPrivato.addEventListener("focusout", checkNumeroTesseraSanitariaPrivato);
    cittaPrivato.addEventListener("focusout", checkCittaPrivato);
    capPrivato.addEventListener("focusout", checkCapPrivato);
    viaPrivato.addEventListener("focusout", checkViaPrivato);
    civicoPrivato.addEventListener("focusout", checkCivicoPrivato);
    telefonoPrivato.addEventListener("focusout", checkTelefonoPrivato);
    emailPrivato.addEventListener("focusout", checkEmailPrivato);
    passwordPrivato.addEventListener("focusout", checkPasswordPrivato);
    confermaPasswordPrivato.addEventListener("focusout", checkPasswordPrivato);
    dataNascitaPrivato.addEventListener("change", checkDataNascitaPrivato);
}

function checkNomePrivato(){
    if(nomePrivato.value.length<1 || nomePrivato.value.length>255 || nomePrivato.value.match(/\d/)){
        nomePrivato.style.backgroundColor="red";
        return false;
    }
    else{
        nomePrivato.style.backgroundColor="white";
        return true;
    }
}

function checkCognomePrivato(){
    if(cognomePrivato.value.length<1 || cognomePrivato.value.length>255 || cognomePrivato.value.match(/\d/)){
        cognomePrivato.style.backgroundColor="red";
        return false;
    }
    else{
        cognomePrivato.style.backgroundColor="white";
        return true;
    }
}

function checkCodiceFiscalePrivato(){
    if(codiceFiscalePrivato.value.length!==16){
        codiceFiscalePrivato.style.backgroundColor="red";
        return false;
    }
    else{
        codiceFiscalePrivato.style.backgroundColor="white";
        return true;
    }
}

function checkNumeroTesseraSanitariaPrivato(){
    if(numeroTesseraSanitariaPrivato.value.length!==20 || !numeroTesseraSanitariaPrivato.value.match(/^[0-9]+$/)){
        numeroTesseraSanitariaPrivato.style.backgroundColor="red";
        return false;
    }
    else{
        numeroTesseraSanitariaPrivato.style.backgroundColor="white";
        return true;
    }
}

function checkCittaPrivato(){
    if(cittaPrivato.value.length<1 || cittaPrivato.value.length>255 || cittaPrivato.value.match(/\d/)){
        cittaPrivato.style.backgroundColor="red";
        return false;
    }
    else{
        cittaPrivato.style.backgroundColor="white";
        return true;
    }
}

function checkCapPrivato(){
    if(capPrivato.value.length!==5 || !capPrivato.value.match(/^[0-9]+$/)){
        capPrivato.style.backgroundColor="red";
        return false;
    }
    else{
        capPrivato.style.backgroundColor="white";
        return true;
    }
}

function checkViaPrivato(){
    if(viaPrivato.value.length<1 || viaPrivato.value.length>255){
        viaPrivato.style.backgroundColor="red";
        return false;
    }
    else{
        viaPrivato.style.backgroundColor="white";
        return true;
    }
}

function checkCivicoPrivato(){
    if(civicoPrivato.value.length<1 || civicoPrivato.value.length>4 || !civicoPrivato.value.match(/^[0-9]+$/) || civicoPrivato.value==0){
        civicoPrivato.style.backgroundColor="red";
        return false;
    }
    else{
        civicoPrivato.style.backgroundColor="white";
        return true;
    }
}

function checkTelefonoPrivato(){
    if(telefonoPrivato.value.length!==10 || !telefonoPrivato.value.match(/^[0-9]+$/)){
        telefonoPrivato.style.backgroundColor="red";
        return false;
    }
    else{
        telefonoPrivato.style.backgroundColor="white";
        return true;
    }
}

function checkEmailPrivato(){
    if(emailPrivato.value.length<1 || emailPrivato.value.length>255 || !emailPrivato.value.match("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])")){
        emailPrivato.style.backgroundColor="red";
        return false;
    }
    else{
        emailPrivato.style.backgroundColor="white";
        return true;
    }
}

function checkPasswordPrivato(){
    if(passwordPrivato.value.length<1 || passwordPrivato.value.length>255 || passwordPrivato.value!==confermaPasswordPrivato.value){
        passwordPrivato.style.backgroundColor="red";
        confermaPasswordPrivato.style.backgroundColor="red";
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
        dataNascitaPrivato.style.backgroundColor="red";
        return false;
    }
    else{
        dataNascitaPrivato.style.backgroundColor="white";
        return true;
    }
}

function submitPrivato(){
    if(!(checkNomePrivato() && checkCognomePrivato() && checkCodiceFiscalePrivato() && checkNumeroTesseraSanitariaPrivato() && checkCittaPrivato() && checkCapPrivato() && checkViaPrivato() && checkCivicoPrivato() && checkTelefonoPrivato() && checkEmailPrivato() && checkPasswordPrivato() && checkDataNascitaPrivato())){
        event.preventDefault();
        return false;
    }
}

function aggiungiSelectPrefissi(userType){
    let pref=JSON.parse(prefissi);
    let refnode=document.getElementById("civico-".concat(userType));
    let select=document.createElement("select");
    select.id="prefisso-".concat(userType);
    for(let i=0; i<pref.countries.length; i++){
        let option=document.createElement("option");
        option.value=pref.countries[i].code;
        option.text=pref.countries[i].code.concat(" -- ", pref.countries[i].name);
        select.appendChild(option);
    }
    let br=document.createElement("br");
    let label=document.createElement("label")
    label.textContent="Prefisso: ";
    label.style.display="inline";
    label.style.fontWeight="normal";
    refnode.after(br, label, select);
}
