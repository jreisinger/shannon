Shannon calculates the entropy of a string. The higher the entropy the more
random a string is.

```
$ for i in $(seq 1 10); do echo $RANDOM; done | shannon
0.970951	24424
1.000000	67
1.370951	22124
1.500000	9794
1.921928	23528
2.321928	23819
2.321928	15390
2.321928	25614
2.321928	24569
2.321928	28906
```