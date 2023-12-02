# Storing and accessing data

## objectives

**Storing and accessing data using arrays**
- Create arrays
- Adding and updating array data
- Enumerating array data

**Measuring algorithmic complexity**
- Asymptotic analysis
- Big-O notation


</br>

## Why arrays?

- Arrays are a data structure that allow us to store and access data sequentially.
  They are a fundamental building block of computer science and are used in almost every program.

</br>

## Properties of arrays

- Can contain multiple instances of a single data type
- Numeric indexing
- fixed-sized or dynamic in size

==example (static size):==
```
## if we make a reading every 5 seconds for an hour then we know we will have 720 readings
## and we can create an array of size 720 to store the data.

# Creating a static array (fixed size) of 5 readings
Readings[] myStaticArray = new Reading[5];
myStaticArray[0] = Guage.Reading();
myStaticArray[1] = Guage.Reading();
myStaticArray[2] = Guage.Reading();
myStaticArray[3] = Guage.Reading();
myStaticArray[4] = Guage.Reading();

# Reading from thee static array
Reading firstReading = myStaticArray[0];
Reading fifthReading = myStaticArray[4];

# Looping through the static array
for (int i = 0; i < 5; i++)
{
    Reading r = myStaticArray[i];
}

# updating the value of a static array
myStaticArray[2] = Guage.Reading();
```

The List<T> class is a generic collection in C# that internally uses an array to store elements. It automatically handles resizing the underlying array as needed, providing a dynamic aspect. It allows you to add, remove, and access elements with flexibility
==example (dynamic size):==
```
## if we are allowing the user to change the frequency of the readings then we don't know
## how many readings we will have before runtime and therefore we need a dynamic array.

List<int> dynamicArray = new List<int>(); // Creating a dynamic array (List) with no fixed size
dynamicArray.Add(1);
dynamicArray.Add(2);
dynamicArray.Add(3);
```

- fixed size arrays are faster to access than dynamic arrays but cannot be changed in size once created
- a fixes-size array will be allocated on the stack whereas a dynamic array will be allocated on the heap
  meaning the the static array is more performant than the dynamic array.

</br>


## Asymptotic Analysis of Algorithms

When talkin about asymptotic analysis we are talking about resources. Resources are things like:
- Operations (the number of times we need to perform an operation)
- Memory (the amount of memory we need to perform an operation)
- Network (the amount of network traffic we need to perform an operation)

so anything we can measure that might effect the performance of our algorithm.


==example (of operations):==
```
char[] letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
int index = 0;

while(letters[index] != 'G')
  index++;

# in this example we are loopin through the letters in the array until we find the letter 'G'
# each tim we loop and check a letter it will have a cost of 1 operation
# in this example it will have a cost of 7 operations
# The maximum number of operations is 26 and this number will be named as 'n'
```

</br>

## Big-O Notation

This is a way of defining the cost of an algorithm in terms of the number of operations it takes to complete.
this can also be called the classification of an algorithm.

Big-O notation represents the upper cost of an algorithm. This means that it represents the worst case scenario
and will be represented as 'O(n)' where 'n' is the number of operations.


==Asymptote==
```
The asymptote of a curve is a line where the distance between the curve and the line
approaches zero as as they tend towards infinity.

=> in other words

asymtotic analysis is the measurement of how the inputs of an algorithm
affect the behaviour as the inputs approach some limit.
```

## Big-O example
**1. O(1) - fixes cost (the cost is always 1)**

|input size|cost|
|-|-|
|1|1|
|100|1|
|1000|1|
|1000000|1|

fixed != fast. Here the amount of work is just not dependant of the size of the input.


**2. O(n) - linear cost (the cost increases linearly with the size of the input)**

in the example below we are looping through the array and printing each letter
this is a liniar algorithm and can be identified by looking for a single loop over a collection of data
```
char[] letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";

for(int i = 0; i < letters.length; i++) {
  Console.WriteLine(letters[i]);
}
```

**3. O(log n) - logarithmic cost**

as input grows the cost of the algorithm does not increase at the same rate as the input

works by dividing a large problem into smaller and smaller chunks.
A common example of an O(log n) algorithm is binary search, which finds an
element in a sorted array by comparing it with the middle element and then
repeating the process on the relevant half of the array 1. Therefore,
binary search is an example of an algorithm that has a time complexity of O(log n).

|input size|cost|
|-|-|
|1|1|
|100|1|
|1000|3|
|1000000|6|

the cost of the algorithm is not increasing at the same rate as the input size

**4. O(n^2) - quadratic cost**

resource usage is the square of the input size

exmaple:
```
void quad(char[] input, int count) {
  for (int i = 0; i < count; i++) {
    for (int x = 0; x < count; x++) {
      process(input, i, x);
    }
  }
}
```

in the example aboce lets say that our big-o will be O(n^2) where n is 1000
in this example count == n
the outer forloop will execute 1000 times and everytime it executes the inner forloop will execute 1000 times
so the total number of operations will be 1000 * 1000 = 1000000, and procees will be called 1000000 times

so these types of algorithms are very expensive

|input size|cost|
|-|-|
|1|1|
|10|100|
|1000|10000000|
|1000000|1e+12|


**5. O(nm) - polynomial cost**

in these types of algorithms we have multiple inputs which will be represented as 'n' and 'm'

example:
```
void nm(char[] n, int nc, char[] m, int mc){
  for (int i = 0; i < nc; i++) {
    for (int x = 0; x < mc; x++) {
      process(n, i, m, x);
    }
  }
}
```

a nested loop which itereates over two distinct collections of data might indicate an O(nm) algorithm
