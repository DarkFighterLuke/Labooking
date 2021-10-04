var nomePrivato=document.getElementById("nome-privato");
var cognomePrivato=document.getElementById("cognome-privato");
var codiceFiscalePrivato=document.getElementById("codice-fiscale-privato");
var numeroTesseraSanitariaPrivato=document.getElementById("numero-tessera-sanitaria-privato");
var cittaPrivato=document.getElementById("citta-privato");
var capPrivato=document.getElementById("cap-privato");
var viaPrivato=document.getElementById("via-privato");
var civicoPrivato=document.getElementById("civico-privato");
var telefonoPrivato=document.getElementById("telefono-privato");
var emailPrivato=document.getElementById("email-privato");
var passwordPrivato=document.getElementById("password-privato");
var confermaPasswordPrivato=document.getElementById("conferma-password-privato");
var dataNascitaPrivato=document.getElementById("data-nascita-privato");

aggiungiSelectPrefissi("privato");
aggiungiSelectPrefissi("medico");

nomePrivato.addEventListener("focusout", checkNome);
cognomePrivato.addEventListener("focusout", checkCognome);
codiceFiscalePrivato.addEventListener("focusout", checkCodiceFiscale);
numeroTesseraSanitariaPrivato.addEventListener("focusout", checkNumeroTesseraSanitaria);
cittaPrivato.addEventListener("focusout", checkCitta);
capPrivato.addEventListener("focusout", checkCap);
viaPrivato.addEventListener("focusout", checkVia);
civicoPrivato.addEventListener("focusout", checkCivico);
telefonoPrivato.addEventListener("focusout", checkTelefono);
emailPrivato.addEventListener("focusout", checkEmail);
passwordPrivato.addEventListener("focusout", checkPassword);
confermaPasswordPrivato.addEventListener("focusout", checkPassword);
dataNascitaPrivato.addEventListener("change", checkDataNascita);

function checkNome(){
    if(nomePrivato.value.length<1 || nomePrivato.value.length>255 || nomePrivato.value.match(/\d/)){
        nomePrivato.style.backgroundColor="red";
        return false;
    }
    else{
        nomePrivato.style.backgroundColor="white";
        return true;
    }
}

function checkCognome(){
    if(cognomePrivato.value.length<1 || cognomePrivato.value.length>255 || cognomePrivato.value.match(/\d/)){
        cognomePrivato.style.backgroundColor="red";
        return false;
    }
    else{
        cognomePrivato.style.backgroundColor="white";
        return true;
    }
}

function checkCodiceFiscale(){
    if(codiceFiscalePrivato.value.length!==16){
        codiceFiscalePrivato.style.backgroundColor="red";
        return false;
    }
    else{
        codiceFiscalePrivato.style.backgroundColor="white";
        return true;
    }
}

function checkNumeroTesseraSanitaria(){
    if(numeroTesseraSanitariaPrivato.value.length!==20 || !numeroTesseraSanitariaPrivato.value.match(/^[0-9]+$/)){
        numeroTesseraSanitariaPrivato.style.backgroundColor="red";
        return false;
    }
    else{
        numeroTesseraSanitariaPrivato.style.backgroundColor="white";
        return true;
    }
}

function checkCitta(){
    if(cittaPrivato.value.length<1 || cittaPrivato.value.length>255 || cittaPrivato.value.match(/\d/)){
        cittaPrivato.style.backgroundColor="red";
        return false;
    }
    else{
        cittaPrivato.style.backgroundColor="white";
        return true;
    }
}

function checkCap(){
    if(capPrivato.value.length!==5 || !capPrivato.value.match(/^[0-9]+$/)){
        capPrivato.style.backgroundColor="red";
        return false;
    }
    else{
        capPrivato.style.backgroundColor="white";
        return true;
    }
}

function checkVia(){
    if(viaPrivato.value.length<1 || viaPrivato.value.length>255){
        nomePrivato.style.backgroundColor="red";
        return false;
    }
    else{
        cognomePrivato.style.backgroundColor="white";
        return true;
    }
}

function checkCivico(){
    if(civicoPrivato.value.length<1 || civicoPrivato.value.length>4 || !civicoPrivato.value.match(/^[0-9]+$/) || civicoPrivato.value==0){
        civicoPrivato.style.backgroundColor="red";
        return false;
    }
    else{
        civicoPrivato.style.backgroundColor="white";
        return true;
    }
}

function checkPrefisso(){
    // TODO: controllare se il prefisso è presente nella lista di prefissi
}

function checkTelefono(){
    if(telefonoPrivato.value.length!==10 || !telefonoPrivato.value.match(/^[0-9]+$/)){
        telefonoPrivato.style.backgroundColor="red";
        return false;
    }
    else{
        telefonoPrivato.style.backgroundColor="white";
        return true;
    }
}

function checkEmail(){
    if(emailPrivato.value.length<1 || emailPrivato.value.length>255 || !emailPrivato.value.match("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])")){
        emailPrivato.style.backgroundColor="red";
        return false;
    }
    else{
        emailPrivato.style.backgroundColor="white";
        return true;
    }
}

function checkPassword(){
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

function checkDataNascita(){
    if(dataNascitaPrivato.value==="" || !underAgeValidate(dataNascitaPrivato.value)){
        dataNascitaPrivato.style.backgroundColor="red";
        return false;
    }
    else{
        dataNascitaPrivato.style.backgroundColor="white";
        return true;
    }
}

function underAgeValidate(birthday){
    // it will accept two types of format yyyy-mm-dd and yyyy/mm/dd
    let optimizedBirthday = birthday.replace(/-/g, "/");

    //set date based on birthday at 01:00:00 hours GMT+0100 (CET)
    let myBirthday = new Date(optimizedBirthday);

    // set current day on 01:00:00 hours GMT+0100 (CET)
    let currentDate = new Date().toJSON().slice(0,10)+' 01:00:00';

    // calculate age comparing current date and borthday
    let myAge = ~~((Date.now(currentDate) - myBirthday) / (31557600000));

    return myAge >= 18;

}

function submitPrivato(){
    if(!(checkNome() && checkCognome() && checkCodiceFiscale() && checkNumeroTesseraSanitaria() && checkCitta() && checkCap() && checkVia() && checkCivico() && checkTelefono() && checkEmail() && checkPassword() && checkDataNascita())){
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
