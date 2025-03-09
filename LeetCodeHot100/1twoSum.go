package LeetCodeHot100

import (
	"sort"
)

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

type node struct {
	value int
	index int
}

func fasttwoSum(nums []int, target int) []int {
	var nodes []node
	for i, v := range nums {
		nodes = append(nodes, node{v, i})
	}
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].value == nodes[j].value {
			return nodes[i].index < nodes[j].index
		}
		return nodes[i].value < nodes[j].value
	})
	i, j := 0, len(nodes)-1
	for i < j {
		if nodes[i].value+nodes[j].value == target {
			return []int{nodes[i].index, nodes[j].index}
		} else if nodes[i].value+nodes[j].value < target {
			i++
		} else {
			j--
		}
	}
	return nil
}
