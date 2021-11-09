# Labooking
La piattaforma n. 1 per la prenotazione dei tuoi test diagnostici COVID19.<br>
Demo: https://labooking.tk

## Istruzioni per riprodurre la build
- Scaricare l'eseguibile per il proprio sistema
- Inserire il file app.conf (disponibile nella sezione Documents del progetto su Redmine) nella cartella `conf/`
- Lanciare l'eseguibile

### Linux
```
cd Labooking
mkdir conf
cp /path/to/your/app.conf conf/
./Labooking
```

### Windows
```
cd Labooking
mkdir conf
cp /path/to/your/app.conf conf/
Labooking.exe
```

## Istruzioni per compilare la build
È possibile compilare il codice sorgente da soli, ma ulteriori tool saranno necessari.
Per semplicità qui si assume che si possieda già il Go SDK >= 1.16 e Git.

```
go get -d github.com/beego/bee/v2
cd $GOPATH/src/github.com/beego/bee
go install .
```
È necessario aggiungere la cartella `$GOPATH/bin` alla variabile d'ambiente `PATH` del proprio sistema.<br>
In un nuovo terminale:

```
git clone https://github.com/DarkFighterLuke/Labooking.git
cd Labooking
go mod tidy
bee pack
```
Sarà ora disponibile un archivio contenente tutti i file necessari per il funzionamento, basterà estrarne il contenuto e lanciare l'eseguibile come mostrato sopra.
