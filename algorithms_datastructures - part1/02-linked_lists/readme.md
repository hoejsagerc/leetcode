# Linked Lists

## Objectives
- What are linked lists?
  - Nodes
  - Node chains
- Single linked lists
- doubly linked lists
- sorted linked lists

</br>

## Introducing Linked Lists

A linked list is a container where nodes of data are linked together in to a list.

a concept of linking simple nodes into a complex data structure is the basis of many data structures.

can be used as the underlying implementation of other very important data structures such as stacks, queues, and hash tables.

think of it as a chain. each link in the chain is a node containing data and has a link to the next node in the chain.

you cannot jump from the first link to the last link without traversing the chain. since each link only knows about the next link in the chain.

</br>

## Nodes and Node chaining

each link in th chain is a node, and each node contains:
- value (single value)
- reference (link to the next node in the chain)

==example of a node==
```
class Node
{
    public Node(int value)
    {
        this.Value = value;
        this.Next = null;
    }

    public Node Next;
    public int Value;
}
```

connecting nodes into a list:
```
# Creating a new node with a value of 1 and the Next property set to null
Node head = new Node(1);

# We are now linking the head node to a new node with a value of 2 and the Next property of head as the new Node
head.Next = new Node(2);

# Creating a third node will look like:
head.Next.Next = new Node(3);
```
This kind of list can extend indeffenitely. The only limitation is the memory available to the program.


</br>


## Single and Doubly Linked Lists

**Single Linked List**
This is a list where we can start at the first node and follow the chain to the last node.
iterating through the list is done by starting at the head node and following the chain until we reach the end of the chain.

example: generic class containing the data and reference to the next node:
```
class LinkedListNode<TNode> {
    public LinkedListNode(TNode value, LinkedListNode<TNode> next = null)
    {
        this.Value = value;
        this.Next = next;
    }

    public LinkedListNode<TNode> Next;
    public TNode Value;
}
```


**Doubly Linked List**
This builds on the Single linked list by adding a reference to the previous node in the chain.
enabling to traverse the list in both directions.

```
class Node
{
    public Node(int value)
    {
        this.Value = value;
        this.Previous = null;
        this.Next = null;
    }

    public Node Previous;
    public Node Next;
    public int Value;
}
```

Lets build a doubly linked list
```
Node node1 = new Node(1);
Node node2 = new Node(2);
Node node3 = new Node(3);

node1.Next = node2;
node2.Previous = node1;

node2.Next = node3;
node3.Previous = node2;
```

note that node3.Next and node1.Previous are both null. This indicates the start and end of the list.


example: generic class containing the data and reference to the next node:
```
class DoublyLinkedListNode<TNode> {
    public DoublyLinkedListNode(TNode value,
        Node<TNode> prev = null,
        Node<TNode> next = null)
    {
        this.Value = value;
        this.Previous = prev;
        this.Next = next;
    }

    public Node<TNode> Next;
    public Node<TNode> Previous;
    public TNode Value;
}
```

</br>

## Adding items to Linked Lists

| Function | Behaviour | Complexity |
| --- | --- | --- |
| AddHead | Adds a value to the beginning of the list | O(1) |
| AddTail | Adds a value to the end of the list | O(1) |

the cost will be O(1) because we are only adding a single node to the list.
which is one of the benefits of linked lists.

example: adding a node to the head of a linked list:
```
public void AddHead(T value)
{
    DoublyLinkedListNode<T> adding = new DoublyLinkedListNode<T>(value, null, head);

    if (head != null)
    {
        head.Previous = adding;
    }

    head = adding;

    if (tail == null)
    {
        tail = head;
    }

    Count++;
}


DoublyLinkedList<int> ints = new DoublyLinkedList<int>();

for (int i = 0; i < 5; i++)
{
    ints.AddHead(i);
}
```

the exact same logic can be used to add a node to the tail of the list.

</br>

## Finding items in Linked Lists


| Function | Behaviour | Complexity |
| --- | --- | --- |
| Find | Find the first node whose value equals the provided argument | O(n) |
| Contains | Returns true if the specified value exists in the list, false otherwise | O(n) |

these functions needs to iterate through the list to find the node with the specified value.
therfore the cost is O(n).


example: finding a node in a linked list:
```
private LinkedListNode<T> Find(T value)
{
    LinkedListNode<T> current = head;

    while (current != null)
    {
        if (current.Value.Equals(value))
        {
            return current;
        }

        current = current.Next;
    }

    return null;
}
```

iterate through the list until we find the node with the specified value.
we then return the node if found or otherwise we return null if it is not found.



example: if list contains a value:
```
public bool Contains(T value)
{
    return Find(value) != null;
}
```

here we simply just use the Find method from before and then return a bool based on if the node is found or not.


</br>


## Removing items from Linked Lists

| Function | Behaviour | Complexity |
| --- | --- | --- |
| Remove | Removes the first node whose value is equal to the argument | O(n) |

Here we need to iterate through the list to find the node with the specified value.