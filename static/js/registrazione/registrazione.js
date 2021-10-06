initElementsPrivato();
initElementsMedico();
initElementsLaboratorio();

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

function aggiungiSelectPrefissi(userType){
    let pref=JSON.parse(prefissi);
    let refnode=document.getElementById("civico-".concat(userType));
    let select=document.createElement("select");
    select.name="Prefisso";
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
