# PfSampleBot

Esempio di Telegram Bot scritto in Go che non fa assolutamente niente
se non rispondere staticamente ai comandi /start e /test

Utilizza "telebot" reperibile su:
   https://github.com/cortinico/telebot

Di fatto non è altro che il programma di prova scritto nel README di
quel progetto

Utilizzato per prendere confidenza con il linguaggio Go e con
la creazione di Bot Telegram

## Installazione

Per installarlo utilizzare la variabile di environment GOPATH
come descritto in:
   https://golang.org/doc/code.html#GOPATH

 * Creare una directory base per i progetti go (p.e. $HOME/goproj)
 * definire GOPATH=$HOME/goproj
 * creare una sottodir _bin_ e inserirla nel PATH:

```bash
   $ export PATH=$PATH:$(go env GOPATH)/bin
```


