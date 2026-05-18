package main

func main() {
	worker := NewPool(4)
	defer worker.Close()
	for i := 0; i < 100; i++ {
		worker.AddTask(func() {
			println(i)
		})
	}

}
