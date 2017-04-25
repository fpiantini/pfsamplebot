# PfSampleBot

Esempio di Telegram Bot scritto in Go che non fa assolutamente niente
se non rispondere staticamente a dei comandi.

Utilizza "telebot" reperibile su:
   https://github.com/cortinico/telebot

Di fatto non è altro che il programma di prova scritto nel README di
quel progetto.

Utilizzato per prendere confidenza con il linguaggio Go e con
la creazione di Bot Telegram

## Installazione

Per installare il programma utilizzare la variabile di environment GOPATH
come descritto in:
   https://golang.org/doc/code.html#GOPATH

 * Creare una directory base per i progetti go (p.e. $HOME/goprjs)
 * definire GOPATH=$HOME/goprjs
 * creare una sottodir _bin_ e inserirla nel PATH:

```bash
$ cd $GOPATH
$ mkdir bin
$ export PATH=$PATH:$(go env GOPATH)/bin
```

 * effettuare il clone del repository nella dir $GOPATH/src/github.com/fpiantini

```bash
$ mkdir -p src/github.com/fpiantini
$ cd src/github.com/fpiantini
$ git clone git@github.com:fpiantini/pfsamplebot.git
```

 * installare il programma:

```bash
$ go install github.com/fpiantini/pfsamplebot
```

# Esecuizione

Per lanciare il bot eseguire:

```bash
$ pfsamplebot
```

