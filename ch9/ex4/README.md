On a i5-6200U CPU @ 2.30GHz and 16gb of ram I've was able to reach 7.3M goroutines.
However traversing took way too long (I didn't wait).
With 5.5M goroutines in the chain, the traversal took 2.5s. 
Anything above 5.5M took ages to run probably due to heavy swapping.