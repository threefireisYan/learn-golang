package main

type Player struct {
	Name        string
	HealthPoint int
	MagicPoint  int
}

func main() {
	tank := new(Player)
	//go经过语法糖 给我们进行简化的书写方式
	// 实际应该为 (*tank).Name = "xxx"
	tank.Name = "码神"
	tank.HealthPoint = 300
}
