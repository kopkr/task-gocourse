# Go Course exercises
This repository is to store exercises done for [Tero Karvinen's Go Programming](http://terokarvinen.com/2020/go-programming-course-2020-w22/#laksyt) course.

## [H1](H1/)
>a) Kirjoita itse keksimäsi ohjelma, joka lukee käyttäjältä syötteen (flag) ja tulostaa käyttäjälle tekstiä.

[Primes](Primes/)<br>
Yksinkertainen ohjelma, joka tulostaa kaikki alkuluvut kahden ja annetun luvun väliltä.
Esim.
~~~~
$ go build
$ ./Primes
Tulosta alkuluvut väliltä 2 ja syötetty arvo.
Määritä yläraja: 30
2
3
5
7
11
13
17
19
23
29
$ ./Primes -max 30
Tulosta alkuluvut väliltä 2 ja syötetty arvo.
Annettu yläraja 30 (flag)
2
3
5
7
11
13
17
19
23
29
~~~~
<hr>
>e) Vapaaehtoinen: Tee rot13-salakirjoitusohjelma. Se siirtää kirjaimia a-zA-Z 13 askelta eteenpäin. Ameriikkalaisia aakkosia on 26, tästä puolet on 13, joten sama ohjelma salaa ja purkaa. Esimerkiksi "Tero" -> "Greb".

[Rot13](Rot13/)<br>
Ohjelma, joka kääntää stringin rot13-salakirjoitusmenetelmällä.
Esim.
~~~~
$ build go
$ ./Rot13 
Rot13-kääntäjä
Syötä tekstiä käännettäväksi: Hello, world!
Uryyb, jbeyq!
$ ./Rot13 -s "Uryyb, jbeyq!"
Rot13-kääntäjä
Teksti käännettäväksi: Uryyb, jbeyq!
Hello, world!
~~~~
