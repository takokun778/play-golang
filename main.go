package main

import "play/logger"

func main() {
	// src := []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(src)

	// 配列から配列を取り除く
	// list1 := []int{1, 2, 3, 4, 5, 6}
	// list2 := []int{1, 3, 5}
	// fmt.Println(array.Take(list1, list2))

	// 配列を逆順に並び替える
	// fmt.Println(src)
	// dst := array.Reverse(src)
	// fmt.Println(dst)
	// fmt.Println(src)

	// 配列を指定したサイズで分割する
	// dst := array.Split(src, 2)
	// fmt.Println(dst)
	// fmt.Println(src)

	// 配列を条件でフィルタリングする
	// condisions := func(n int) bool {
	// 	return n%2 == 0
	// }
	// dst := array.Filter(src, condisions)
	// fmt.Println(dst)
	// fmt.Println(src)

	// timer.Sleep(10)

	// type Result struct {
	// 	Number int
	// 	Error  error
	// }

	// results := make([]Result, 0)

	// wg := new(sync.WaitGroup)
	// mu := new(sync.Mutex)

	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func(wg *sync.WaitGroup, mu *sync.Mutex) {
	// 		defer wg.Done()
	// 		defer mu.Unlock()
	// 		result, err := parallel.AsyncFunc()
	// 		mu.Lock()
	// 		results = append(results, Result{
	// 			Number: result,
	// 			Error:  err,
	// 		})
	// 	}(wg, mu)
	// }

	// wg.Wait()

	// for _, result := range results {
	// 	fmt.Println(result)
	// }

	// fmt.Println(regexp.CheckEmail("m7107208636@dea-21olympic.com"))
	// regexp.CheckPassword("m7107206A36@dea-21olympic.com")

	secret := Secret{
		Name: "rei ozono",
	}

	logger.Logger().Info(secret)
}

type Secret struct {
	Name Name
}

type Name string

func (n Name) String() string {
	return "*****"
}
