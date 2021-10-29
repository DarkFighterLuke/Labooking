function showErrorDiv(parentNode, isBefore, errMessage) {
    let div = document.createElement("div");
    div.className = "div-errore";
    let p = document.createElement("p");
    p.className = "p-errore";
    p.innerText = errMessage;
    div.appendChild(p);
    if (isBefore) {
        parentNode.before(div);
    } else {
        parentNode.after(div);
    }
}

function eraseErrorDivs() {
    let errorDivs = document.getElementsByClassName("div-errore");
    let errorDivsLength = errorDivs.length;
    for (let i = 0; i < errorDivsLength; i++) {
        errorDivs[0].parentNode.removeChild(errorDivs[0]);
    }
}
