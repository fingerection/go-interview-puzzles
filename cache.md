# Cache comment

Cache is commmon used class and seldom people think about the internal structure for cache. So the first point, design a basic interface for a cache. It always depends on whether you have used a cache library or not.

Then you’ll face up to second point: data structure. In cache system, you can easily use a data structure called `map` to store data. But when cache is full, one always use LRU strategy to handle. That is delete last visited key in map. You can easily use timestamp stored as part of map value, but you have to iterate map values to decide which one to delete. 

You use link list for lru information, there’s a queue for it. If you use one-direction link, when you delete last visited node, you should iterate it to last element, which also O(N), so better solution is use bi-direction link list which you can iterate from last element. (source code is written for one-direction link list)  
