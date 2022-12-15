package main

import (
	circle "main/lib"
)

func main() {

	ll := make([]circle.Circle, 0)
	for i := 0; i <= 20; i++ {
		ll = append(ll, *circle.New("asd", i))
	}

	for _, i := range ll {
		println(i.Get_integger())
	}
}
