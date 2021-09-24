<!-- Top navigation mobile-->
<nav class="navbar navbar-expand-lg bg-lightblue border-bottom" id="mobile-nav">
    <div class="container-fluid d-flex justify-content-between">
        <div><img src="https://icons.getbootstrap.com/assets/icons/list.svg" class="btn" width=140% id="sidebarToggle"></div>
        <div><h1 class="title">Labooking</h1></div>
        <a href="#" class="notification">
            <img src="img/icons/bell-svgrepo-com.svg" width=120%>
            <span class="badge">{{.NumNotifiche}}</span>
        </a>
        <!--<button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation"><span class="navbar-toggler-icon"></span></button>-->
        <button class="navbar-toggler" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle user menu" onclick="openUserMenu();"><img src="img/icons/user-avatar-filled-alt-svgrepo-com.svg" alt="Nome Utente"></button>
        <!--suppress XmlDuplicatedId -->
    </div>
    <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav ms-auto mt-2 mt-lg-0">
            <li class="nav-item dropdown">
                <a class="nav-link" id="navbarDropdown" href="#" role="button" data-bs-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                    <div class="d-flex flex-row">
                        <label id="username" class="nav-link dropdown-toggle">{{.NomeUtente}}</label>
                        <img id="user-avatar" src="img/icons/user-avatar-filled-alt-svgrepo-com.svg" alt="{{.NomeUtente}}">
                    </div>
                </a>

                <!--suppress XmlDuplicatedId -->
                <div class="dropdown-menu dropdown-menu-end" aria-labelledby="navbarDropdown" id="navbarSupportedContent">
                    <a class="dropdown-item" href="#!">Il mio profilo</a>
                    <a class="dropdown-item" href="#!">Messaggi<span class="badge badge-danger">{{.NumNotifiche}}</span></a>
                    <a class="dropdown-item" href="#!">I miei pagamenti</a>
                    <div class="dropdown-divider"></div>
                    <a class="dropdown-item" href="#!">Logout</a>
                </div>
            </li>
        </ul>
    </div>
</nav>

<!-- Top navigation-->
<nav class="navbar navbar-expand-lg bg-lightblue border-bottom" id="normal-nav">
    <div class="container-fluid d-flex justify-content-between">
        <div><img src="https://icons.getbootstrap.com/assets/icons/list.svg" class="btn" width=140% id="sidebarToggle"></div>
        <div><h1 class="title">Labooking</h1></div>
        <!--<button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation"><span class="navbar-toggler-icon"></span></button>-->
        <div class="d-flex flex-row">
            <a href="#" class="notification">
                <img src="img/icons/bell-svgrepo-com.svg" width=120%>
                <span class="badge">{{.NumNotifiche}}</span>
            </a>
            <button class="navbar-toggler" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle user menu" onclick="openUserMenu();"><img src="img/icons/user-avatar-filled-alt-svgrepo-com.svg" alt="Nome Utente"></button>
            <!--suppress XmlDuplicatedId -->
            <div class="collapse navbar-collapse" id="navbarSupportedContent">
                <ul class="navbar-nav ms-auto mt-2 mt-lg-0">
                    <li class="nav-item dropdown">
                        <a class="nav-link" id="navbarDropdown" href="#" role="button" data-bs-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                            <div class="d-flex flex-row">
                                <label id="username" class="nav-link dropdown-toggle">{{.NomeUtente}}</label>
                                <img id="user-avatar" src="img/icons/user-avatar-filled-alt-svgrepo-com.svg" alt="{{.NomeUtente}}">
                            </div>
                        </a>

                        <!--suppress XmlDuplicatedId -->
                        <div class="dropdown-menu dropdown-menu-end" aria-labelledby="navbarDropdown" id="navbarSupportedContent">
                            <a class="dropdown-item" href="#!">Il mio profilo</a>
                            <a class="dropdown-item" href="#!">Messaggi<span class="badge badge-danger">{{.NumNotifiche}}</span></a>
                            <a class="dropdown-item" href="#!">I miei pagamenti</a>
                            <div class="dropdown-divider"></div>
                            <a class="dropdown-item" href="#!">Logout</a>
                        </div>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</nav>
