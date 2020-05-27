# Go Course exercises
This repository exists to store exercises done for [Tero Karvinen's Go Programming](http://terokarvinen.com/2020/go-programming-course-2020-w22/) course.

[Assignments at TeroKarvinen.com(Finnish)](http://terokarvinen.com/2020/go-programming-course-2020-w22/#laksyt)

## [H1](H1/)

[Primes](H1/Primes/program.go), [Rot13](H1/Rot13/program.go)

>a) Kirjoita itse keksimäsi ohjelma, joka lukee käyttäjältä syötteen (flag) ja tulostaa käyttäjälle tekstiä.

>b) Tee Tour of Go: Welcome. Tätä b-alakohtaa ei tarvitse raportoida. Ei tarvitse tehdä lisäkohtaa "Go offline (optional)".

>c) Vapaaehtoinen, helppo: Tee Tour of Go: "Basics": "Packages, variables and functions" Tätä c-alakohtaa ei tarvitse raportoida.

>d) Vapaaehtoinen: Tee Tour of Go: "Flow control statements: for, if else, switch and defer". Tätä d-kohtaa ei tarvitse raportoida. Tuo "Exercise: Loops and Functions" -harjoitus voi olla haastavampi, voit halutessasi jättää sen väliin. Muut kohdat ovat suoraviivaisia.

>e) Vapaaehtoinen: Tee rot13-salakirjoitusohjelma. Se siirtää kirjaimia a-zA-Z 13 askelta eteenpäin. Ameriikkalaisia aakkosia on 26, tästä puolet on 13, joten sama ohjelma salaa ja purkaa. Esimerkiksi "Tero" -> "Greb".

## [H2](H2/)

[LetterCount](H2/LetterCount/program.go), [AgeCalc](H2/AgeCalc/program.go)

>a) LinWinMac! Käännä ohjelma kolmelle alustalle: Windows, Linux, Mac. Tee staattinen käännös niin, että ohjelma toimii ilman mitään riippuvuuksia tai kirjastoja. Testaa ohjelman toiminta ainakin joillain näistä alustoista ja ota ruutukaappaukset. 'CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build tero.go' (GOOS darwin, windows, linux)

>b) Kirjastoja kohti. Kirjoita ohjelma, joka käyttää kahta uutta ominaisuutta tai kirjastoa Go by Example-kirjasta. Voit valita mitä vain, helppoja vaihtoehtoja ovat esimerkiksi satunnaisluvut, merkkijonojen käsittely ja aika.

>c) Vapaaehtoinen: Patterit mukana. Vapaaehtoinen: kirjoita ohjelma, joka käyttää jotain uutta ominaisuutta Go:n standardikirjastoista.

>d) Vapaaehtoinen: Tee Tour of Go alusta luvun "More types: structs, slices, and maps" loppuun asti. Jos "Exercise"-kohdat ovat hankalia, voit hypätä niistä yli tai jättää ne viimeisiksi. Tätä d-kohtaa ei tarvitse raportoida.

>e) Vapaaehtoinen: Kirjoita yksikkötesti (unit test).

## [H3](H3/)

[FileBackup](H3/FileBackup/program.go)

>a) Kirjastosta päivää. Tee ohjelma, joka käyttää ainakin kahta uutta kirjastoa. Jos ohjelmallasi on jokin käyttötarkoitus, se on mainiota. Muista ensin kokeilla kutakin kirjastoa erikseen. Kirjastoja löytyy GoByExample ja Go:n virallisesta dokumentaatiosta. Käytä jotain uutta kirjastoa, eli ei pelkästään rand, fmt, time eikä string.

>b) Lähteet. Katso, että olet viitannut kaikissa tehtävävastauksissasi kaikkiin lähteisiin, joita olet käyttänyt. Kurssiin, tehtävänantoihin, käyttämääsi materiaaliin (GoByExample tms), muiden koodeihin, StackOverflown vastauksiin ja kaikkiin muihinkin lähteisiin.

>c) Kaikki tai jotain. Valitse aihe tai pari vaihtoehtoa kurssin loppuprojektille. Palastele: mieti, miten voit ensin toteuttaa pienen, toimivan kokonaisuuden. Ja sen päälle vaihtoehtoja laajentaa.

>d) Vapaaehtoinen, helppo ja suositeltava: Asenna Go omaan isäntäjärjestelmääsi. Senhän saa Linuxille, Windowsille ja Macille. https://golang.org/ Jos Windowsin virustorjuntaohjelmasi poistelee mielivaltaisesti tiedostoja, kerro, kuinka annoit itsellesi virustorjunnan asetuksista luvan laittaa binäärejä johonkin hakemistoon.

>e) Vapaaehtoinen: Grep. Tee ohjelma, joka näyttää tekstitiedostosta vain rivit, joilla mainitaan hakusana. Kun ohjelmasi toimii ja haluat lisähaastetta, voit tehdä lisäominaisuuksia: case insensitive, regular expressions, rekursio. Jos kaipaat vielä lisähaastetta, voit kokeilla, toimiiko ohjelmasi isolla (yli 50 GB) tiedostoilla, esim. tietokantadumpeilla?

>f) Vapaaehtoinen: Tee Tour of Go loppuun. Jos "Exercise"-kohdat ovat hankalia, voit hypätä niistä yli tai jättää ne viimeisiksi. Tätä d-kohtaa ei tarvitse raportoida.
