package LCR037

/*
给定一个整数数组 asteroids，表示在同一行的小行星。
对于数组中的每一个元素，其绝对值表示小行星的大小，正负表示小行星的移动方向（正表示向右移动，负表示向左移动）。每一颗小行星以相同的速度移动。
找出碰撞后剩下的所有小行星。碰撞规则：两个行星相互碰撞，较小的行星会爆炸。如果两颗行星大小相同，则两颗行星都会爆炸。两颗移动方向相同的行星，永远不会发生碰撞。
*/
func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	// 使用压栈的方式来模拟碰撞
	for i := 0; i < len(asteroids); i++ {
		num := asteroids[i]

		if len(stack) == 0 {
			stack = append(stack, num)
			continue
		}

		if num < 0 {
			// 获得栈顶元素
			top := stack[len(stack)-1]
			if top > 0 {
				stack = stack[:len(stack)-1]
				for len(stack) >= 0 {
					sum := top + num

					if sum > 0 {
						stack = append(stack, top)
						break
					}

					if sum == 0 {
						break
					}

					if len(stack) == 0 {
						stack = append(stack, num)
						break
					}

					top = stack[len(stack)-1]
					if top < 0 {
						stack = append(stack, num)
						break
					}

					stack = stack[:len(stack)-1]
				}
			} else {
				stack = append(stack, num)
			}
		} else {
			stack = append(stack, num)
		}
	}

	return stack
}
