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
