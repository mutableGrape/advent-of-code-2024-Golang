# Day 1 Solution

## Part 1 Approach

Unpack rows into separate lists, sort the lists, then sum the absolute values of the element-wise differences.

## Part 2 Approach

Use the fact that both slices are already sorted to save computation time and scale with $O(n)$. Create `left`, `right_index`, `counter`, and `sumB` to hold the last seen value in the left slice, the current right index, the rolling count of how many right values equal `left`, and the total final sum, respectively. Then iterate through the first slice and if the value is new (not equal to `left`), add `counter * left` to `sumB` and reset the `counter` value.
