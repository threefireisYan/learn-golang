package main

import (
	"fmt"
	"math"
)

type Character struct {
	TargetX, TargetY   float64 // 目标坐标
	CurrentX, CurrentY float64 // 当前坐标
	Speed              float64 // 移动速度
}

func (c *Character) Move() {
	dx := c.TargetX - c.CurrentX
	dy := c.TargetY - c.CurrentY

	distance := math.Sqrt(dx*dx + dy*dy)
	if distance <= c.Speed {
		c.CurrentX = c.TargetX
		c.CurrentY = c.TargetY
	} else {
		ratio := c.Speed / distance
		c.CurrentX += dx * ratio
		c.CurrentY += dy * ratio
	}
}

func main() {
	// 创建一个人物角色
	character := Character{
		TargetX:  100,
		TargetY:  -201,
		CurrentX: 0,
		CurrentY: 0,
		Speed:    5,
	}

	// 模拟移动过程
	for character.CurrentX != character.TargetX || character.CurrentY != character.TargetY {
		character.Move()
		fmt.Printf("Current position: (%.2f, %.2f)\n", character.CurrentX, character.CurrentY)
	}

	// 打印最终位置
	fmt.Printf("Final position: (%.2f, %.2f)\n", character.CurrentX, character.CurrentY)
}
