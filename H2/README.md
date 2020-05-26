## Go Course exercises
This repository is to store exercises done for [Tero Karvinen's Go Programming](http://terokarvinen.com/2020/go-programming-course-2020-w22/#laksyt) course.

### H2
>a) LinWinMac! Käännä ohjelma kolmelle alustalle: Windows, Linux, Mac. Tee staattinen käännös niin, että ohjelma toimii ilman mitään riippuvuuksia tai kirjastoja. Testaa ohjelman toiminta ainakin joillain näistä alustoista ja ota ruutukaappaukset. 'CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build tero.go' (GOOS darwin, windows, linux)

Käännetään staattinen 64-bittinen binääri kolmelle alustalle.

![img1](img/img1)

Ajetaan ohjelma Linux ja Windows -käyttöjärjestelmillä.

![img2](img/img2)

<hr>

>b) Kirjastoja kohti. Kirjoita ohjelma, joka käyttää kahta uutta ominaisuutta tai kirjastoa Go by Example-kirjasta. Voit valita mitä vain, helppoja vaihtoehtoja ovat esimerkiksi satunnaisluvut, merkkijonojen käsittely ja aika.
>c) Vapaaehtoinen: Patterit mukana. Vapaaehtoinen: kirjoita ohjelma, joka käyttää jotain uutta ominaisuutta Go:n standardikirjastoista.

[LetterCount](LetterCount/program.go)<br>
Ohjelma joka laskee vokaalit ja konsonantit syötetystä tekstistä.
-f parametrilla voidaan antaa tiedosto.
~~~~
$ ./LetterCount-linux-64bit 
LetterCount. v.0.01
Enter string: Ääkkösiä ja ruotsalainen Å
Vowels: 13 Consonants: 10 Total: 23
$ cat example.txt 
Vesihiisi sihisi hississä # 11, 12, 23

!?:))##!?")=!¤/!=163584213--__$$ # 0
$ ./LetterCount-linux-64bit -f example.txt 
LetterCount. v.0.01
Vowels: 11 Consonants: 12 Total: 23
~~~~

<hr>

>e) Vapaaehtoinen: Kirjoita yksikkötesti (unit test).
