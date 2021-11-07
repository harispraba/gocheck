# Gocheck

Gocheck adalah sebuah simple tools yang digunakan untuk mengecek expired dari ssl sebuah domain. Jika tanggal expirednya kurang dari 30 hari maka akan mengirimkan notifikasi via webhook discord ( yang bisa dicustom sesuai dengan webhook kalian ).

Installasi :
```bash
git clone https://github.com/mrofisr/gocheck
cd gocheck
go build
```

ada dua metode penggunaan gocheck, yang pertama hanya mengecek satu domain, dan yang kedua bisa menggunakan list.txt yang berisikan banyak domain.

Format list domain
```text
domain1.com
domain2.com
domain3.com
domain4.com
```

Cara menggunakan :

```bash
gocheck -l list.txt -webhook <url>
gocheck -d domain.com -webhook <url>
```


Menggunakan Kubernetes :

```yaml
apiVersion: v1
kind: ConfigMap
data:
  list.txt: |-
    mrofisr.dev
    google.com
    youtube.com
    bing.com
    mtarget.co
    akusukajeruk.com
metadata:
  name: gocheck-ssl-conf
  namespace: default

---
apiVersion: batch/v1
kind: CronJob
metadata:
  name: gocheck-ssl
  labels:
    name: gocheck-ssl
spec:
  schedule: "0 0 1 * *"
  jobTemplate:
    spec:
      template:
        metadata:
          name: gocheck-ssl
          labels:
            name: gocheck-ssl
        spec:
          containers:
            - name: gocheck-ssl
              image: ghcr.io/mrofisr/gocheck:latest
              command:
                [
                  "/bin/gocheck",
                  "-L=/data/gocheck/list.txt",
                  "-webhook=https://<your discord webhook url>",
                ]
              volumeMounts:
                - name: domain-list
                  mountPath: /data/gocheck
          volumes:
            - name: domain-list
              configMap:
                name: gocheck-ssl-conf
          restartPolicy: OnFailure
```