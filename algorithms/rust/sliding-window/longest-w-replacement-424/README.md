# Problem
You are given a string `s` and an integer `k`. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most `k` times.

Return _the length of the longest substring containing the same letter you can get after performing the above operations_.

## Example 1:

Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.

## Example 2:

Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
There may exists other ways to achieve this answer too.

## Constraints:

- `1 <= s.length <= 105`
- `s` consists of only uppercase English letters.
- `0 <= k <= s.length`
# Solution
This solution implements a sliding window with the help of a hash map, `counts`, that keeps track of the items in the array. The important information from this hash map is extracted into `maxf`, telling us what the largest character count is in the current sliding window. 
We compare the window size minus `maxf` to `k` to determine whether we overshot the size of the window. If so, we need to increment the left pointer to shrink the current window size.