package main

func arraySign(nums []int) int {
    var result int
    for i:=0;i<len(nums);i++{
        if  nums[i]==0{
            return 0
        }
        if nums[i]<0{
            result++
        }
    }
    if result%2!=0{
        return -1
    }    
    return 1
}

// RuneTime:=0ms Beats 100.00%
// Memory:= 3.16MB Beats 88.99%