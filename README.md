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
1 Transaction =>  dacfabc4806537828b243eb9aee20a467018b35ce4a0a2d83bd6c7c1cf940f23  count =>  282
2 Transaction =>  ef020113d2161d25643acf9fa83d5c102f2cf7aa2a66bafc4de33dd45d3ab1e5  count =>  155
3 Transaction =>  6412035b2bcc0065cc4b9c5c4c7a68bf0ff6d843086cb0b2fcb8045b45fe8e67  count =>  125
4 Transaction =>  4f2ab4a877a0519e834384805885f0d20a41511b3cb00a7132cdfde639f6f790  count =>  117
5 Transaction =>  ea9c6e11aa46baf568be3921b3017220deff4b3920bdee5e88b8b0ca196d952d  count =>  116
6 Transaction =>  87f344d0c95f2c86a21568acc885b752e9d1bf8a9e9d339e0f1e552f230007df  count =>  111
7 Transaction =>  5c245535776fc28cdc9933cf2923d14dcd1a298590041c4175879d02ef22a4b0  count =>  100
8 Transaction =>  6ac33f7ef0c1f19ef90f82d88d664a9a97e5c7f7b65a5fbc9c0c48667eb6512d  count =>  100
9 Transaction =>  4ff8b7315c83c718926290fb636f2f112143439192640ebda0fd23dbfe108aed  count =>  100
10 Transaction =>  847578ab1bfbca68ca3d37898765362f5d22060a18c98c4ad6fc3f946e244e6f  count =>  100


-> find_ancestors test
1 Transaction =>  tx1  count =>  9
2 Transaction =>  tx2  count =>  5
3 Transaction =>  tx7  count =>  3
4 Transaction =>  tx11  count =>  1
```
### For smaller input (test.json file)


![Screenshot](https://github.com/balampbv/largest-ancestry-sets/blob/main/Screenshot%202022-04-13%20180537.jpg)
