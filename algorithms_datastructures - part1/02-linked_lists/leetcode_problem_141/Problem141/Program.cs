
  // Definition for singly-linked list.
public class ListNode {
    public int val;
    public ListNode next;
    public ListNode(int x) {
        val = x;
        next = null;
    }
}


// Solution
public class Solution {
    public bool HasCycle(ListNode head) {
        var slow = head;
        var fast = head;

        while (fast != null && fast.next != null)
        {
            slow = slow.next;
            fast = fast.next.next;
            if (slow == fast)
            {
                return true; // fast caught up with slow meaning there is a loop
            }
        }

        return false; // fast reached a null meaning the end of the linked list
    }
}