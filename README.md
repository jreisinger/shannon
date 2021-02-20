Shannon calculates the entropy of a string. The higher the entropy the more
random a string is.

```
$ for i in $(seq 1 10); do echo $RANDOM; done | shannon
11191                                                                 0.721928
30008                                                                 1.370951
22825                                                                 1.370951
3616                                                                  1.500000
3385                                                                  1.500000
24747                                                                 1.521928
23208                                                                 1.921928
4597                                                                  2.000000
21095                                                                 2.321928
23416                                                                 2.321928
```