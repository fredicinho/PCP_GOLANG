package MemoryModel

var p *int

func main() {
	done := make(chan bool)
	// "done"-Variable wird in der Main-Methode
	// sowie in der Goroutine verwendet, somit
	// landet dieser im Heap

	go func() {
		x, y, z := 123, 456, 789
		_ = z  // z kann auf dem Stack abgelegt werden
		p = &x // x und y landen auf dem Heap, weil sie
		p = &y // über das globale p referenziert werden

		// x wird nun nicht mehr referenziert
		// und kann durch den Garbage Collector
		// aufgesammelt werden

		p = nil
		// Y wird nun auch nicht mehr referenziert
		// und kann durch den Garbage Collector
		// gesammelt werden

		done <- true
	}()

	<-done
}
