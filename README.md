# largest-ancestry-sets
CLI to fetch the largest ancestry sets for the block using blockstream. 

### extract_block [Block-Hash] 
  Will fetch the transactions from the blockstream and persistent in a json file
### find_ancestors [Block-Hash]
  Will find the ancestors for the given block hash. 


#### Dry run
```
---------------------
-> find_ancestors 000000000000000000076c036ff5119e5a5a74df77abf64203473364509f7732
1 Transaction =>  20ab5295e1aa20d76145a478b0471b41e2915df433c802e43d822d0ce904e933  count =>  646
2 Transaction =>  6422defd90f656b222f7f0ec1d92d6a48c9e1d3d8990499d550d3b7524d32379  count =>  583
3 Transaction =>  dacfabc4806537828b243eb9aee20a467018b35ce4a0a2d83bd6c7c1cf940f23  count =>  282
4 Transaction =>  ef020113d2161d25643acf9fa83d5c102f2cf7aa2a66bafc4de33dd45d3ab1e5  count =>  155
5 Transaction =>  6412035b2bcc0065cc4b9c5c4c7a68bf0ff6d843086cb0b2fcb8045b45fe8e67  count =>  125
6 Transaction =>  4f2ab4a877a0519e834384805885f0d20a41511b3cb00a7132cdfde639f6f790  count =>  117
7 Transaction =>  ea9c6e11aa46baf568be3921b3017220deff4b3920bdee5e88b8b0ca196d952d  count =>  116
8 Transaction =>  87f344d0c95f2c86a21568acc885b752e9d1bf8a9e9d339e0f1e552f230007df  count =>  111
9 Transaction =>  d48bf3f362c443929b198d5eaa6422fbd6b9dc9a519565fcc5f1aff6c77568ef  count =>  100
10 Transaction =>  2a9d32bd628cb2b0ad26d0d66138940ab973f7828f14d7334c41fe89d26daaa9  count =>  100
```
