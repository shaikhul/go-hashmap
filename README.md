Problem Statement:
------------------
Implement a hash based map of X number of (key, value) entries with.
Provide basic functionality like Get, Put, Del methods.
Use a loading factor to double its size when (no of pairs/total pairs) exceeds the loading factor.
Address hash collision with open addressing.

Assumption:
----------
1. Start with map size 3
2. Key are string but value can be anything
3. Loading factor 70%
4. Should have Get, Put and Del method
5. Add tests

Test Result:
-----------
```
➜  hashmap go test -v     
=== RUN   TestHashMapGet
--- PASS: TestHashMapGet (0.00s)
=== RUN   TestHashMapPutABool
--- PASS: TestHashMapPutABool (0.00s)
=== RUN   TestHashMapPutAString
--- PASS: TestHashMapPutAString (0.00s)
=== RUN   TestHashMapDel
--- PASS: TestHashMapDel (0.00s)
=== RUN   TestHashMapRehashing
--- PASS: TestHashMapRehashing (0.00s)
PASS
ok      
```