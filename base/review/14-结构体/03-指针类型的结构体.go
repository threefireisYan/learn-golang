package main

type Player struct {
	Name        string
	HealthPoint int
	MagicPoint  int
}

func main() {
	tank := new(Player)
	tank.Name = "码神"
	tank.HealthPoint = 300
}
