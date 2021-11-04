<!-- Top navigation mobile-->
<!--suppress XmlDuplicatedId -->
<nav class="navbar navbar-expand-lg bg-lightblue border-bottom" id="mobile-nav">
    <div class="container-fluid d-flex justify-content-between">
        <div><img src="https://icons.getbootstrap.com/assets/icons/list.svg" class="btn" width=140%
                  id="sidebarToggleMobile"></div>
        <div><h1 class="title"><a id="brand-name" href=/dashboard/home>Labooking</a></h1></div>
        {{if eq .Ruolo "privato"}}
        <div>
            <a href="#" class="notification nav-link" role="button" data-bs-toggle="dropdown" aria-haspopup="true"
               aria-expanded="false" {{if ne .NumNotifiche "0"}}onclick="onclickNotificationBell()"{{end}}>
            <img src="/img/icons/bell-svgrepo-com.svg" width=120%>
            {{if ne .NumNotifiche "0"}}
            <span id="num-notifiche-mobile" class="badge">{{.NumNotifiche}}</span>
            {{end}}
            </a>

            {{if ne .NumNotifiche "0"}}
            <div class="dropdown-menu dropdown-menu-end me-5 mt-n1" aria-labelledby="navbarDropdown">
                {{range $i, $v := .Notifiche}}
                <div class="bg-lightblue notifica">
                    <input type="hidden" class="notifiche" name="notifica-{{$i}}" value="{{.IdTestDiagnostico}}">
                    <a class="a-notifica" href="/dashboard/referti">Referto del {{.DataEsecuzione.Format "01/02/2006"}} pronto!</a>
                </div>
                {{end}}
            </div>
            {{end}}

            <button id="button-profilo-mobile" class="navbar-toggler" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent"
                    aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle user menu"
                    onclick="openUserMenu();">
                    <img src="/img/icons/user-avatar-filled-alt-svgrepo-com.svg" alt="Nome Utente">
            </button>
        </div>
        {{end}}

        <!--suppress XmlDuplicatedId -->
    </div>
    <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav ms-auto mt-2 mt-lg-0">
            <li class="nav-item dropdown">
                <a class="nav-link" id="navbarDropdown" href="#" role="button" data-bs-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                    <div class="d-flex flex-row">
                        <label id="username" class="nav-link dropdown-toggle">{{.NomeUtente}}</label>
                        <img id="user-avatar" src="/img/icons/user-avatar-filled-alt-svgrepo-com.svg"
                             alt="{{.NomeUtente}}">
                    </div>
                </a>

                <!--suppress XmlDuplicatedId -->
                <div class="dropdown-menu dropdown-menu-end" aria-labelledby="navbarDropdown" id="navbarSupportedContent">
                    <a class="dropdown-item" href="/logout">Logout</a>
                </div>
            </li>
        </ul>
    </div>
</nav>

<!-- Top navigation-->
<nav class="navbar navbar-expand-lg bg-lightblue border-bottom" id="normal-nav">
    <div class="container-fluid d-flex justify-content-between">
        <div><img src="https://icons.getbootstrap.com/assets/icons/list.svg" class="btn" width=140% id="sidebarToggle"></div>
        <div><h1 class="title"><a id="brand-name" href=/dashboard/home>Labooking</a></h1></div>
        <!--<button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation"><span class="navbar-toggler-icon"></span></button>-->
        <div class="d-flex flex-row">
            {{if eq .Ruolo "privato"}}
            <div>
                <a href="#" class="notification nav-link" role="button" data-bs-toggle="dropdown" aria-haspopup="true"
                   aria-expanded="false" {{if ne .NumNotifiche "0"}}onclick="onclickNotificationBell()"{{end}}>
                <img src="/img/icons/bell-svgrepo-com.svg" width=120%>
                {{if ne .NumNotifiche "0"}}
                <span id="num-notifiche" class="badge">{{.NumNotifiche}}</span>
                {{end}}
                </a>

                {{if ne .NumNotifiche "0"}}
                <div class="dropdown-menu dropdown-menu-end me-5 mt-n1" aria-labelledby="navbarDropdown">
                    {{range $i, $v := .Notifiche}}
                    <div class="bg-lightblue notifica">
                        <input type="hidden" class="notifiche" name="notifica-{{$i}}" value="{{.IdTestDiagnostico}}">
                        <a class="a-notifica" href="/dashboard/referti">Referto del {{.DataEsecuzione.Format "01/02/2006"}} pronto!</a>
                    </div>
                    {{end}}
                </div>
                {{end}}
            </div>
            <script src="/js/notifiche.js"></script>
            {{end}}
            <button class="navbar-toggler" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent"
                    aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle user menu"
                    onclick="openUserMenu();">
                        <img src="/img/icons/user-avatar-filled-alt-svgrepo-com.svg" alt="Nome Utente">
            </button>
            <!--suppress XmlDuplicatedId -->
            <div class="collapse navbar-collapse" id="navbarSupportedContent">
                <ul class="navbar-nav ms-auto mt-2 mt-lg-0">
                    <li class="nav-item dropdown">
                        <a class="nav-link" id="navbarDropdown" href="#" role="button" data-bs-toggle="dropdown"
                           aria-haspopup="true" aria-expanded="false">
                            <div class="d-flex flex-row">
                                <label id="username" class="nav-link dropdown-toggle">{{.NomeUtente}}</label>
                                <img id="user-avatar" src="/img/icons/user-avatar-filled-alt-svgrepo-com.svg"
                                     alt="{{.NomeUtente}}">
                            </div>
                        </a>

                        <!--suppress XmlDuplicatedId -->
                        <div class="dropdown-menu dropdown-menu-end" aria-labelledby="navbarDropdown"
                             id="navbarSupportedContent">
                            <a class="dropdown-item" href="/logout">Logout</a>
                        </div>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</nav>
