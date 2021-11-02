<h2 class="mt-4 content-tab-title">{{.Title}}</h2>
<div>
    {{.ErrMsg}}
</div>
<div id="privato">
    <!-- Form privato -->
    <form class="form-labooking" id="form-privato" action="/signup?idForm=privato&goToLogin=false" method="POST"
          autocomplete="off">
        <input type="text" name="privato" hidden>
        {{.FormPrivato | renderform}}
        <br>Data di nascita: <input name="DataNascita" type="date" value="" id="data-nascita-privato">
        <br><br><input id="submit-privato" type="submit" onclick="return submitPrivato()">
    </form>
</div>
<script src="/js/utils/autocompleteAddress.js"></script>
<script src="/js/registrazione/prefissi_telefonici.js"></script>
<script src="/js/dipendenti/aggiunta.js"></script>
