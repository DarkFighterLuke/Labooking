function rimuoviCampiPassword(){
    document.getElementById("password-privato").remove();
    document.getElementById("conferma-password-privato").remove();
    document.getElementById("form-privato").innerHTML=document.getElementById("form-privato").innerHTML.replace("<br>Password: ", "")
    document.getElementById("form-privato").innerHTML=document.getElementById("form-privato").innerHTML.replace("<br>Conferma password: ", "")
}

window.onload=rimuoviCampiPassword;
