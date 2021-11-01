<h2 class="mt-4 content-tab-title">{{.Title}}</h2>
<div>
    <form class="form-labooking" id="form-privato" action="/signup?idForm=privato" method="POST" onsubmit="return submitPrivato()">
        <input type="text" name="privato" hidden>
        {{.FormPrivato | renderform}}
        <br>Data di nascita: <input name="DataNascita" type="date" value="" id="data-nascita-privato">
        <br><br><input id="submit-privato" type="submit" value="Aggiungi">
    </form>
</div>
<script src="/js/dipendenti/aggiunta.js"></script>