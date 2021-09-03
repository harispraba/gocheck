# Gocheck

Gocheck adalah sebuah simple tools yang digunakan untuk mengecek expired dari ssl sebuah domain. Jika tanggal expirednya kurang dari 30 hari maka akan mengirimkan notifikasi via webhook discord ( yang bisa dicustom sesuai dengan webhook kalian ).

Installasi :
```bash
git clone https://github.com/mrofisr/gocheck
cd gocheck
go build
```

ada dua metode penggunaan gocheck, yang pertama hanya mengecek satu domain, dan yang kedua bisa menggunakan list.txt yang berisikan banyak domain.

```bash
gocheck -l list.txt -webhook <url>
gocheck -d domain.com -webhook <url>
```