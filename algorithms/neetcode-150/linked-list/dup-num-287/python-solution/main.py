class Solution:
    def findDuplicate0(self, nums: list[int]) -> int:
        counts = {}
        for num in nums:
            if counts.get(num):
                return num
            counts[num] = 1 
        return -1

    def findDuplicate(self, nums: list[int]) -> int:
        fast, slow = 0, 0
        while True:
            slow = nums[slow]
            fast = nums[nums[fast]]
            if fast == slow:
                break
        fast = 0
        while fast != slow:
            slow = nums[slow]
            fast = nums[fast]
        return fast

if __name__ == "__main__":
    nums = [ 1, 2, 3, 3, 4]
    nums2 = [1,3,4,2,2]
    s = Solution()
    print(s.findDuplicate(nums))
    print(s.findDuplicate(nums2))
