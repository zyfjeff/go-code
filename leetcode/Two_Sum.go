package  main
func twoSum(nums []int, target int) []int {
  ret := make([]int,2)
  nums_map := make(map[int]int)

  for index,num := range nums {
    peer_index,ok := nums_map[target - num]

    if(!ok) {
      nums_map[num] = index
    } else {
      ret[0] = peer_index
      ret[1] = index
    }
  }

  return ret

}
