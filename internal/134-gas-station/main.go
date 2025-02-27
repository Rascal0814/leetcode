package _34_gas_station

const _ = `
134. 加油站
中等
相关标签
相关企业
在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。

你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。

给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。
`

func canCompleteCircuit(gas []int, cost []int) int {
	for start, gasNum := 0, len(gas); start < gasNum; {
		cur, totalGas := 0, 0 // 当前车站相对，油箱总油量
		for cur < gasNum {
			j := (cur + start) % gasNum //当前车站绝对坐标值
			totalGas += gas[j] - cost[j]
			if totalGas < 0 {
				break
			}
			cur++
		}
		if cur == len(gas) {
			return start
		} else {
			start += cur + 1 // 如果从起始位置到达cur加油站无法继续前行，最差情况就是一路的油箱剩余都为0，所以直接从cur下一个站点前进
		}
	}

	return -1
}
