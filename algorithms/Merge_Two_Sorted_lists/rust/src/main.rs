#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}

impl ListNode {
  #[inline]
  fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val
    }
  }
}
impl Solution {
    pub fn merge_two_lists(list1: Option<Box<ListNode>>, list2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if list1.is_none() {
            return list2;
        }
        if list2.is_none() {
            return list1;
        }

        let head = if &list1.as_ref().unwrap().val < &list2.as_ref().unwrap().val {
            list1
    } else {
            list2
        };

        let (mut l1, mut l2) = if head.as_ref().unwrap().val == list1.as_ref().unwrap().val {
            (list1, list2)
        } else {
            (list2, list1)
        };

        let mut current = head.as_ref();

        while l1.is_some() && l2.is_some() {
            if l1.as_ref().unwrap().val <= l2.as_ref().unwrap().val {
                let next = l1.as_mut().unwrap().next.take();
                current.unwrap().next = l1;
                l1 = next;
            } else {
                let next = l2.as_mut().unwrap().next.take();
                current.unwrap().next = l2;
                l2 = next;
            }
            current = current.unwrap().next.as_ref();
        }

        if l1.is_some() {
            current.unwrap().next = l1;
        } else {
            current.unwrap().next = l2;
        }

        head
    }
}