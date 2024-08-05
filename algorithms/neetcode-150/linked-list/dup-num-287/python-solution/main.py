class Solution:
    def findDuplicate0(self, nums: list[int]) -> int:
        counts = {}
        for num in nums:
            if counts.get(num):
                return num
            counts[num] = 1 
        return -1

    def findDuplicate(self, nums: list[int]) -> int:
        return -1
