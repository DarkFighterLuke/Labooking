var nomeMedico;
var cognomeMedico;
var codiceFiscaleMedico;
var cittaMedico;
var capMedico;
var viaMedico;
var civicoMedico;
var telefonoMedico;
var emailMedico;
var passwordMedico;
var confermaPasswordMedico;

function initElementsMedico(){
    nomeMedico=document.getElementById("nome-medico");
    cognomeMedico=document.getElementById("cognome-medico");
    codiceFiscaleMedico=document.getElementById("codice-fiscale-medico");
    cittaMedico=document.getElementById("citta-medico");
    capMedico=document.getElementById("cap-medico");
    viaMedico=document.getElementById("via-medico");
    civicoMedico=document.getElementById("civico-medico");
    telefonoMedico=document.getElementById("telefono-medico");
    emailMedico=document.getElementById("email-medico");
    passwordMedico=document.getElementById("password-medico");
    confermaPasswordMedico=document.getElementById("conferma-password-medico");
    codiceRegionale=document.getElementById("codice-regionale-medico");

    aggiungiSelectPrefissi("medico");

    nomeMedico.addEventListener("focusout", checkNomeMedico);
    cognomeMedico.addEventListener("focusout", checkCognomeMedico);
    codiceFiscaleMedico.addEventListener("focusout", checkCodiceFiscaleMedico);
    cittaMedico.addEventListener("focusout", checkCittaMedico);
    capMedico.addEventListener("focusout", checkCapMedico);
    viaMedico.addEventListener("focusout", checkViaMedico);
    civicoMedico.addEventListener("focusout", checkCivicoMedico);
    telefonoMedico.addEventListener("focusout", checkTelefonoMedico);
    emailMedico.addEventListener("focusout", checkEmailMedico);
    passwordMedico.addEventListener("focusout", checkPasswordMedico);
    confermaPasswordMedico.addEventListener("focusout", checkPasswordMedico);
    codiceRegionale.addEventListener("change", checkCodiceRegionale);
}

function checkNomeMedico(){
    if(nomeMedico.value.length<1 || nomeMedico.value.length>255 || nomeMedico.value.match(/\d/)){
        nomeMedico.style.backgroundColor="red";
        return false;
    }
    else{
        nomeMedico.style.backgroundColor="white";
        return true;
    }
}

function checkCognomeMedico(){
    if(cognomeMedico.value.length<1 || cognomeMedico.value.length>255 || cognomeMedico.value.match(/\d/)){
        cognomeMedico.style.backgroundColor="red";
        return false;
    }
    else{
        cognomeMedico.style.backgroundColor="white";
        return true;
    }
}

function checkCodiceFiscaleMedico(){
    if(codiceFiscaleMedico.value.length!==16){
        codiceFiscaleMedico.style.backgroundColor="red";
        return false;
    }
    else{
        codiceFiscaleMedico.style.backgroundColor="white";
        return true;
    }
}

function checkCittaMedico(){
    if(cittaMedico.value.length<1 || cittaMedico.value.length>255 || cittaMedico.value.match(/\d/)){
        cittaMedico.style.backgroundColor="red";
        return false;
    }
    else{
        cittaMedico.style.backgroundColor="white";
        return true;
    }
}

function checkCapMedico(){
    if(capMedico.value.length!==5 || !capMedico.value.match(/^[0-9]+$/)){
        capMedico.style.backgroundColor="red";
        return false;
    }
    else{
        capMedico.style.backgroundColor="white";
        return true;
    }
}

function checkViaMedico(){
    if(viaMedico.value.length<1 || viaMedico.value.length>255){
        nomeMedico.style.backgroundColor="red";
        return false;
    }
    else{
        cognomeMedico.style.backgroundColor="white";
        return true;
    }
}

function checkCivicoMedico(){
    if(civicoMedico.value.length<1 || civicoMedico.value.length>4 || !civicoMedico.value.match(/^[0-9]+$/) || civicoMedico.value==0){
        civicoMedico.style.backgroundColor="red";
        return false;
    }
    else{
        civicoMedico.style.backgroundColor="white";
        return true;
    }
}

function checkTelefonoMedico(){
    if(telefonoMedico.value.length!==10 || !telefonoMedico.value.match(/^[0-9]+$/)){
        telefonoMedico.style.backgroundColor="red";
        return false;
    }
    else{
        telefonoMedico.style.backgroundColor="white";
        return true;
    }
}

function checkEmailMedico(){
    if(emailMedico.value.length<1 || emailMedico.value.length>255 || !emailMedico.value.match("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])")){
        emailMedico.style.backgroundColor="red";
        return false;
    }
    else{
        emailMedico.style.backgroundColor="white";
        return true;
    }
}

function checkPasswordMedico(){
    if(passwordMedico.value.length<1 || passwordMedico.value.length>255 || passwordMedico.value!==confermaPasswordMedico.value){
        passwordMedico.style.backgroundColor="red";
        confermaPasswordMedico.style.backgroundColor="red";
        return false;
    }
    else{
        passwordMedico.style.backgroundColor="white";
        confermaPasswordMedico.style.backgroundColor="white";
        return true;
    }
}

function checkCodiceRegionale(){
    if(codiceRegionale.value.length!==6){
        codiceRegionale.style.backgroundColor="red";
        return false;
    }
}

function submitMedico(){
    if(!(checkNomeMedico() && checkCognomeMedico() && checkCodiceFiscaleMedico() && checkCittaMedico() && checkCapMedico() && checkViaMedico() && checkCivicoMedico() && checkTelefonoMedico() && checkEmailMedico() && checkPasswordMedico())){
        event.preventDefault();
        return false;
    }
}
