def __merge(l1: list[int], l2: list[int]):
    """
        returns sorted list if l1 & l2 are sorted
    """
    out: list[int] = []
    i1, i2 = 0, 0
    while i1 < len(l1) and i2 < len(l2):
        if l1[i1] < l2[i2]:
            out.append(l1[i1])
            i1 += 1
        else:
            out.append(l2[i2])
            i2 += 1
    while i1 < len(l1):
        out.append(l1[i1])
        i1 += 1
    while i2 < len(l2):
        out.append(l2[i2])
        i2 += 1
    return out


def merge_sort(nums: list[int]) -> list[int]:
    # break up the list into separate pieces & zip them back together correctly
    if len(nums) <= 1: return nums
    mid: int = len(nums)//2
    return __merge(merge_sort(nums[:mid]),merge_sort(nums[mid:])) 

if __name__ == "__main__":
    print(merge_sort([3,4,2,3]))
    print(merge_sort([]))
    print(merge_sort([50,40,30,20,10]))
    
