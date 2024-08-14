def __recurse(l1: list[int], l2: list[int]):
    if len(l1) == len(l2) == 1:
        return [l1[0], l2[0]] if l1[0] < l2[0] else [l2[0], l1[0]]
    l1_fh = l1[:len(l1)//2]
    l1_sh = l1[len(l1)//2:]
    l2_fh = l1[:len(l1)//2]
    l2_sh = l1[len(l1)//2:]
    return __recurse(l1_fs, l2_fs)

def merge_sort(nums: list[int]) -> list[int]:
    # break up the list into separate pieces & zip them back together correctly
    return [nums[0]]

if __name__ == "__main__":
    print(__recurse([2],[1]))
    
