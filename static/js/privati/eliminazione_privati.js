function eliminaSelezionati(){
    let table=document.getElementById("table-privati");
    let form=new FormData();
    let numChecked=0;
    for(let i=2; i<table.rows.length; i++){
        if(table.rows[i].cells[0].children[0].checked){
            form.append("id-privato", table.rows[i].cells[0].children[0].id);
            numChecked++;
        }
    }
    if(numChecked<1){
        return;
    }

    let request=new Request(location.href.concat("/eliminazione"), {
        method: "POST",
        body: form
    });
    fetch(request).then(eraseFilters);
}