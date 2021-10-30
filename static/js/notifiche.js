let isOpen = false;
let marked = false;

function onclickNotificationBell() {
    if (!marked) {
        marked = true;
        notifyServer();
    }
    if (!isOpen) {
        isOpen = true;
    } else {
        isOpen = false;
        document.getElementById("num-notifiche").style = "display: none";
    }
}

document.onclick = function () {
    if (isOpen) {
        isOpen = false;
        document.getElementById("num-notifiche").style = "display: none";
    }
}

function notifyServer() {
    let notifiche = document.getElementsByClassName("notifiche");
    let form = new FormData();
    for (let i = 0; i < notifiche.length; i++) {
        form.append(notifiche[i].name, notifiche[i].value);
    }

    let request = new Request("/api/notifiche", {
        method: "POST",
        body: form
    });
    fetch(request);
}
