// From https://github.com/cortinico/telebot

package main

import (
  //"fmt";
  "syscall";
  "github.com/cortinico/telebot"
)

func main() {
  conf := telebot.Configuration{
    BotName: "PfSampleBot",
    ApiKey:  "351486743:AAEJHdZfw1VgrktLClJi6RCbaNT3D5nUCdE",
  }

  var version = "1.0.0"
  var bot telebot.Bot

  bot.Start(conf, func(mess string) (string, error) {
    var answer string
    switch mess {
    case "/start":
      answer = pfsbotStartCommand()
    case "/ver":
      answer = "Version: " + version
    case "/uname":
      answer = pfsbotUnameCommand()
    case "/test":
      answer = "Il comando test Funziona :)"
    case "/prova":
      answer = "Il comando prova Funziona :)"
    case "/energia":
      answer = "E = mc^2"
    case "/materia":
      answer = "Entità provvista di una propria consistenza fisica, " +
      "dotata di peso e di inerzia, capace di adeguarsi a una forma"
    case "/antimateria":
      answer = "materia costituita da antiparticelle, " +
      "corrispondenti per massa alle particelle della materia ordinaria, " +
      "ma aventi alcuni numeri quantici, come ad esempio la carica elettrica, " +
      "di segno opposto"
    case "/alienroy":
      answer = "https://soundcloud.com/alien-roy"
    case "/serj":
      answer = "http://www.videospazio.it/ (versione vecchia...)\n" +
      "quella nuova: http://www.videospazio.norstrilia.net/"
    case "/videospazio":
      answer = "vedi alla voce: serj"
    case "/pf":
      answer = "http://www.fpiantini.it/chi-sono/"
    case "/code":
      answer = "https://github.com/fpiantini/pfsamplebot/"

    default:
      answer = "Comando sconosciuto: " + mess
    }
    return answer, nil
  })
}

func pfsbotStartCommand() string {
  var answer string
  answer = "Ciao, sono un bot di test scritto da PF. " +
  "Non faccio niente di particolare"

  return answer
}

func pfsbotUnameCommand() string {
  var utsname syscall.Utsname

  _ = syscall.Uname(&utsname)

  // fmt.Println(utsname)
  str := "Sysname: " + B2S(utsname.Sysname[:]) + "\n" +
  "Nodename: " + B2S(utsname.Nodename[:]) + "\n" +
  "Release: " + B2S(utsname.Release[:]) + "\n" +
  "Version: " + B2S(utsname.Version[:]) + "\n" +
  "Machine: " + B2S(utsname.Machine[:]) + "\n" +
  "Domainname: " + B2S(utsname.Domainname[:])

  return str
}

func B2S(bs []int8) string {
  s := make([]byte, len(bs))
  var lens int
  for ; lens < len(bs); lens++ {
    if bs[lens] == 0 {
      break
    }
    s[lens] = uint8(bs[lens])
  }
  return string(s[0:lens])
}

