## Go Course exercises
This repository is to store exercises done for [Tero Karvinen's Go Programming](http://terokarvinen.com/2020/go-programming-course-2020-w22/#laksyt) course.

### H3
>a) Kirjastosta päivää. Tee ohjelma, joka käyttää ainakin kahta uutta kirjastoa. Jos ohjelmallasi on jokin käyttötarkoitus, se on mainiota. Muista ensin kokeilla kutakin kirjastoa erikseen. Kirjastoja löytyy GoByExample ja Go:n virallisesta dokumentaatiosta. Käytä jotain uutta kirjastoa, eli ei pelkästään rand, fmt, time eikä string.
>b) Lähteet. Katso, että olet viitannut kaikissa tehtävävastauksissasi kaikkiin lähteisiin, joita olet käyttänyt. Kurssiin, tehtävänantoihin, käyttämääsi materiaaliin (GoByExample tms), muiden koodeihin, StackOverflown vastauksiin ja kaikkiin muihinkin lähteisiin.

### [FileBackup](FileBackup/program.go)<br>
Ohjelma, joka varmuuskopioi tiedoston. Harjoittelua tiedostojen käsittelyyn os, io ja path/filepath kirjastoilla.<br>
`-from` (pakollinen) - Ottaa vastaan tiedoston, joka varmuuskopioidaan.<br>
`-to` - Ottaa vastaan tallennussijainnin ja vapaavalintaisesti myös tiedostonimen. Oletuksena tallentaa samaan hakemistoon syötetyn tiedoston nimellä + .bak pääte. (foo.txt > foo.txt.bak)

Esim. ilman `-to` parametria
~~~~
$ tree
.
├── hello.txt
├── program.go
└── test

1 directory, 3 files
$ cat hello.txt 
Hello, world
$ go run program.go -from hello.txt 
Saved backup file to: hello.txt.bak
$ cat hello.txt.bak 
Hello, world
~~~~

Esim. `-to` parametrin kanssa ja eri kohdekansioilla.
~~~~
$ go run program.go -from ~/heimaailma.txt -to test/
Saved backup file to: test/heimaailma.txt.bak
$ cat test/heimaailma.txt.bak 
Hei, maailma!
~~~~

Esim. varmuuskopiointi eri tiedostonimellä. (Lisää silti .bak päätteen)
~~~~
$ go run program.go -from hello.txt -to olleh.txt
Saved backup file to: olleh.txt.bak
$ cat olleh.txt.bak 
Hello, world
~~~~

Materiaalit<br>
[https://golangbot.com/write-files/]https://golangbot.com/write-files/
[https://gobyexample.com/directories](https://gobyexample.com/directories)
[https://gobyexample.com/file-paths](https://gobyexample.com/file-paths)
[https://golang.org/pkg/os/](https://golang.org/pkg/os/)
[https://golang.org/pkg/path/filepath/](https://golang.org/pkg/path/filepath/)
[https://golang.org/pkg/io/](https://golang.org/pkg/io/)
[https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang](https://stackoverflow.com/questions/21060945/simple-way-to-copy-a-file-in-golang)
[https://golangbyexample.com/check-if-file-or-directory-exists-go/](https://golangbyexample.com/check-if-file-or-directory-exists-go/)

<hr>
