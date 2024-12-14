# dmngo

find short, available domain names.

no support and no warranty under no circumstance. see
[LICENSE](LICENSE), have fun.

## usage

```bash
go build .
./dmngo -help
```

## example

```bash
./dmngo -length 3 -tld org -vowel
...yfo.org is likely taken.
...xuu.org is likely taken.
...cuo.org is likely taken.
...ote.org is likely taken.
...xoi.org is likely taken

./dmngo -length 4 -tld com
...daxn.com is likely taken.
...fspo.com is likely taken.
...xtea.com is likely taken.
...wnkx.com is likely taken.

./dmngo -length 4 -tld dk -vowel
ztxo.dk MAY be free. adding to domains.txt
zepu.dk MAY be free. adding to domains.txt
ijru.dk MAY be free. adding to domains.txt

./dmngo -input examples/cities.txt -tld se
...stockholm.se is likely taken.
...goteborg.se is likely taken.
...malmo.se is likely taken.
```

domains that _may_ be free are saved to `domains.txt`
