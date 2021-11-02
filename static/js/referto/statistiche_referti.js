function calcolaStatistiche() {
    mapArray = [];
    retrieveDataFromTable(1);
    let numPositivi = 0;
    let numNegativi = 0;
    let numNulli = 0;
    for (let i = 0; i < mapArray.length; i++) {
        let esito = mapArray[i].get("esito");
        if (esito === "positivo") {
            numPositivi++;
        }
        if (esito === "negativo") {
            numNegativi++;
        }
        if (esito === "nullo") {
            numNulli++;
        }
    }
    return [numPositivi, numNegativi, numNulli]
}

var values = calcolaStatistiche()

var chart = new ej.charts.Chart({

    primaryXAxis: {
        valueType: "Category",
        majorGridLines : {
            color : 'white',
            width : 0
        },
        minorGridLines : {
            color : 'white',
            width : 0
        }
    },

    primaryYAxis: {
        title: "Numero Referti",
        majorGridLines : {
            color : 'white',
            width : 0
        },
        minorGridLines : {
            color : 'white',
            width : 0
        }
    },

    //Initializing Chart Series
    series: [
        {
            type: "Bar",
            dataSource: [
                { month: "Positivi", visitors: values[0] },
                { month: "Negativi", visitors: values[1] },
                { month: "Nulli", visitors: values[2] },

            ],
            xName: "month",
            yName: "visitors",
            marker: {
                // Data label for chart series
                dataLabel: {
                    visible: true
                }
            }
        }
    ],
    height: '300'

});

chart.appendTo("#container-statistiche");

