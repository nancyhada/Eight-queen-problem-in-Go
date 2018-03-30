package main

import "fmt"

func display(file [8]int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if file[i] == j {
				fmt.Print("Q\t")
			} else {
				fmt.Print("-\t")
			}
		}
		fmt.Println("\v")
	}
	var x int
	fmt.Println("Press any key")
	fmt.Scanf("%d", &x)
}

func check(file [8]int) bool {
	var x, y int
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 8; j++ {

			x = file[i] - file[j]
			y = i - j
			if x == y || x == -y {
				return false
			}
		}
	}
	return true
}
func main() {
	var file [8]int
	var i1, i2, i3, i4, i5, i6, i7, i8 int
	var res bool
	for i1 = 0; i1 < 8; i1++ {
		file[0] = i1
		for i2 = 0; i2 < 8; i2++ {
			if i2 == i1 {
				continue
			}
			file[1] = i2
			for i3 = 0; i3 < 8; i3++ {
				if i3 == i1 || i3 == i2 {
					continue
				}
				file[2] = i3
				for i4 = 0; i4 < 8; i4++ {
					if i4 == i1 || i4 == i2 || i4 == i3 {
						continue
					}
					file[3] = i4
					for i5 = 0; i5 < 8; i5++ {
						if i5 == i1 || i5 == i2 || i5 == i3 || i5 == i4 {
							continue
						}
						file[4] = i5
						for i6 = 0; i6 < 8; i6++ {
							if i6 == i1 || i6 == i2 || i6 == i3 || i6 == i4 || i6 == i5 {
								continue
							}
							file[5] = i6
							for i7 = 0; i7 < 8; i7++ {
								if i7 == i1 || i7 == i2 || i7 == i3 || i7 == i4 || i7 == i5 || i7 == i6 {
									continue
								}
								file[6] = i7
								for i8 = 0; i8 < 8; i8++ {
									if i8 == i1 || i8 == i2 || i8 == i3 || i8 == i4 || i8 == i5 || i8 == i6 || i8 == i7 {
										continue
									}
									file[7] = i8
									res = check(file)
									if res == true {
										display(file)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
