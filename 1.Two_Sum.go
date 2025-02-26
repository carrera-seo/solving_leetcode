package main

func twoSum(nums []int, target int) []int {

	hashmap := make(map[int]int) // 해시맵 생성
	for i := 0; i < len(nums); i++ {
		hashmap[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		if j, exists := hashmap[target-nums[i]]; j != i && exists {
			return []int{i, j} // 두 인덱스 반환
		}

	}

	return []int{} // 솔루션이 없으면 빈 슬라이스 반환
}
