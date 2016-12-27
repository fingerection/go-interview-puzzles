# quicksort comment

Quicksort problem is not an easy problem. It includes several steps to write a good one.

First, you should construct the basic function declearation for it. If you use `func quicksort(l []int)`, you can't do it for recursion. 

Second, you know you can use `func quicksort(l []int, start int, end int)` to do it. What is quicksort indeed? Its core concept is you put element to the right place, which has left side smaller than it and right side larger than it. And you only need to iterate array once. Then just do it for the left part and right part.

That is the key part for the puzzle. You treat the first element as the element to be put in the right place, and now you have an empty *SLOT* for putting all element less than first element to that slot, in coding, we use j for slot index. It is the largest index which is less than your picked value. So you maintain such  array:

'' [values less than picked  SLOT] [values larger than picked]

So now you can treat new element well.

Last tip is when you initial a list, slot index is the first element. That make things begin to run.
