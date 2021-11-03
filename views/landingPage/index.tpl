<!DOCTYPE html>
<html lang="en">
  {{template "landingPage/head.html"}}

  <body>
    {{template "landingPage/top_bar.html"}}
    <div class="container w-100">
      <div id="carousel-landing-page" class="carousel slide h-100 w-100" data-ride="carousel">

        <!-- The slideshow -->
        <div class="carousel-inner h-100">

          <div class="carousel-item active h-100 w-100">
            <div class="card h-100 border border-3 border-info bg-light text-center">
                <div class="card-body">
                    <h5 class="card-title">Servizio Prenotazione Test Diagnostico per Privati</h5>
                    <p class="card-text">Sei un privato e cerchi un test diagnostico COVID-19?</p>
                    <p class="card-text">Se la risposta è sì, sei nel posto giusto. Labooking è la piattaforma di riferimento in Italia
                        per la prenotazione di test diagnostici COVID-19 per privati.</p>
                    <p class="card-text">Usando il nostro servizio, potrai confrontare prezzi e disponibilità di tutti i laboratori nella tua zona
                        e prenotare il tuo test diagnostico in pochi semplici click, senza code e lunghe chiamate.
                    </p>
                    <a href="/signup" class="btn btn-primary">Prenota un test ora</a>
                </div>
                <div class="card-footer text-muted"> Labooking SRL </div>
            </div>
          </div>

          <div class="carousel-item h-100 w-100">
            <div class="card h-100 border border-3 border-info bg-light text-center">
                <div class="card-body">
                    <h5 class="card-title">Servizi Esclusivi per le Aziende</h5>
                    <p class="card-text">Semplifica la gestione dei test diagnostici per i tuoi dipendenti e mantieni la tua azienda
                        al sicuro dal COVID-19.</p>
                    <p class="card-text">Grazie ai servizi dedicati di Labooking potrai creare un account aziendale,
                        registrare i tuoi dipendenti e prenotare un test diagnostico per il personale della tua azienda
                        in pochi semplici passaggi, risparmiando tempo e denaro.</p>
                    <p class="card-text">Inoltre, avrai accesso alle statistiche e ai referti di tutti i tuoi dipendenti
                    in unica piattaforma.</p>
                    <a href="/signup" class="btn btn-primary">Registra la tua azienda</a>
                </div>
                <div class="card-footer text-muted"> Labooking SRL </div>
            </div>
          </div>

          <div class="carousel-item h-100 w-100">
            <div class="card h-100 border border-3 border-info bg-light text-center">
                <div class="card-body">
                    <h5 class="card-title">Gestione Studio di Medicina Generale</h5>
                    <p class="card-text">Labooking aiuta anche i medici di base nel semplificare la gestione del proprio studio.</p>
                    <p class="card-text">Se sei un medico di medicina generale e cerchi una soluzione per fornire ai tuoi pazienti il servizio di
                    prenotazione dei test diagnostici COVID-19, Labooking può esserti di aiuto.</p>
                    <p class="card-text">Effettua la registrazione come medico, aggiungi i tuoi pazienti e, in pochi secondi potrai
                    prenotare per loro un test diagnostico scegliendo tra i molti laboratori disponibili.
                    </p>
                    <a href="/signup" class="btn btn-primary">Registrati come medico</a>
                </div>
                <div class="card-footer text-muted"> Labooking SRL </div>
            </div>
          </div>

          <a class="carousel-control-prev" href="#carousel-landing-page" role="button" data-slide="prev">
            <span class="carousel-control-prev-icon" aria-hidden="true"></span>
            <span class="sr-only">Previous</span>
          </a>
          <a class="carousel-control-next" href="#carousel-landing-page" role="button" data-slide="next">
            <span class="carousel-control-next-icon" aria-hidden="true"></span>
            <span class="sr-only">Next</span>
          </a>

        </div>
      </div>
    </div>
  </body>
</html>
