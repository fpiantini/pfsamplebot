// From https://github.com/cortinico/telebot

package main

import "github.com/cortinico/telebot"

func main() {
  conf := telebot.Configuration{
    BotName: "PfSampleBot",
    ApiKey:  "351486743:AAEJHdZfw1VgrktLClJi6RCbaNT3D5nUCdE",
  }

  var bot telebot.Bot

  bot.Start(conf, func(mess string) (string, error) {
    var answer string
    switch mess {
    case "/start":
      answer = pfsbotStartCommand()
    case "/test":
      answer = "Il comando test Funziona :)"
    case "/prova":
      answer = "Il comando prova Funziona :)"
    case "/energia":
      answer = "E = mc^2"
    case "/materia":
      answer = "Entità provvista di una propria consistenza fisica, dotata di peso e di inerzia, capace di adeguarsi a una forma"
    case "/antimateria":
      answer = "materia costituita da antiparticelle, corrispondenti per massa alle particelle della materia ordinaria, ma aventi alcuni numeri quantici, come ad esempio la carica elettrica, di segno opposto"
    case "/alienroy":
      answer = "https://soundcloud.com/alien-roy"
    case "/serj":
      answer = "http://www.videospazio.it/ (versione vecchia...)\nquella nuova: http://www.videospazio.norstrilia.net/"
    case "/videospazio":
      answer = "vedi alla voce: serj"
    case "/pf":
      answer = "http://www.fpiantini.it/chi-sono/"

    default:
      answer = "Comando sconosciuto: " + mess
    }
    return answer, nil
  })
}

func pfsbotStartCommand() string {
  var answer string
  answer = "Ciao, sono un bot di test scritto da PF. Non faccio niente di particolare"

  return answer
}


