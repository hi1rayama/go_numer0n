package main
import (
   "fmt"
   "os"
   "bufio"
   "strings"
   "strconv"
   "math/rand"
   "time"
)


/** ユーザー情報の構造体*/
type userInfo struct{
	name string
	numbers [3]int
}

/* 入力を行う関数 */
func getInputValue() (string){
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	return stdin.Text()

}

func hanteiNumber(value string) (bool,[3]int){
	/*
	説明 : 3桁の数字かつ各桁の数字がそれぞれ異なることを判定する関数
	引数 : value [string] : ユーザが入力した値
	返り値 : bool(条件を満たしていればTrue), [3]int(各桁の数字)
	*/
	numbers := [3]int{-1,-1,-1}
	if len(value) != 3{
		return false,numbers
	}else{
		slice := strings.Split(value,"")
		for i,s := range slice{
			n,err := strconv.Atoi(s)
			if err != nil{
				return false,numbers
			}else if arrayContains(numbers,n){
				return false,numbers
			}else{
				numbers[i] = n
			}
		}
	}
	return true, numbers
}


func arrayContains(arr [3]int, number int) bool{
	/*
	説明 : 配列の重複を確認する関数
	引数 : arr [3]int : 重複があるか確認する3桁の数字が入った配列, number : 比較する数値
	返り値 : bool(重複していたら, True)
	*/

	for _, v := range arr{
	  if v == number{
		return true
	  }
	}
	return false
  }


func getCPUNumbers(cpu *userInfo) {
	/*
	説明 : コンピュータの数字を決定する関数
	引数 : cpu *userInfo : cpuの情報を格納する構造体
	返り値 : なし
	*/
	for  i := 0; i < 3; i++{
		rand.Seed(time.Now().UnixNano())
		 cpu.numbers[i] = rand.Intn(10)
	}
} 

func cpuGuessNumber(enemyNumber [3]int) (bool){
	/*
	説明 : コンピュータが相手の数字を当てる処理をする関数
	引数 : enemyNumber [3]int: 相手の3桁の数値
	返り値 : bool(予測した数値が当たっていればtrue)
	*/

	fmt.Println("相手の3桁の数字を予想してください")
	var selectNumber [3]int
	// 何らかの手法で相手の数を求める(最初はランダムで...)
	for  i := 0; i < 3; i++{
		rand.Seed(time.Now().UnixNano())
		selectNumber[i] = rand.Intn(10)
		fmt.Printf("%d",selectNumber[i])
	}
	fmt.Println()
	// 対戦相手の数とユーザーが入力した数を比較
	eat,bite := compNumber(selectNumber,enemyNumber)

	// 結果を出力, 対戦した値が一致したらtrue, 一致しなかったら, eat,byteを表示し, falseを返す
	if eat == 3{
		return true
	}else{
		fmt.Printf("eat : %d, bite : %d\n",eat,bite)
		return false
	}
}

func compNumber(forecast [3]int,correct [3]int) (int,int){
	/*
	説明 : 2つの配列を比較し, eat,bite判定を行う関数
	引数 : forecast [3]int : 予想した数値, correct [3]int : 相手の3桁の数値
	返り値 : int(eatの数),int(biteの数)
	*/
	var eat int = 0
	var bite int = 0

	for i:=0;i<3;i++{
		if(forecast[i] == correct[i]){
			eat++
		}else{
			for j:=0;j<3;j++{
				if(i != j && forecast[i] == correct[j]){
					bite++
					break
				}
			}
		}
	}
	return eat,bite

}

func userGuessNumber(enemyNumbers [3]int) (bool){
	/*
	説明 : ユーザが相手の数字を当てる処理をする関数
	引数 : enemyNumber [3]int: 相手の3桁の数値
	返り値 : bool(予測した数値が当たっていればtrue)
	*/
	var selectNumber [3]int
	fmt.Println("相手の3桁の数字を予想してください")
	for {
		value := getInputValue()
		flag,numbers := hanteiNumber(value)
		if flag{
			selectNumber = numbers
			break
		}else{
			fmt.Println("半角で3桁の数字を入力してください!")
		}
	}
	eat,bite := compNumber(selectNumber,enemyNumbers)
	if eat == 3{
		return true
	} else{
		fmt.Printf("eat : %d, bite : %d\n",eat,bite)
		return false
	}
}


func main() {
	var user userInfo // ユーザの情報
	var computer userInfo // CPUの情報
	computer.name = "CPU"
	getCPUNumbers(&computer)

	fmt.Println("CPUと対戦するヌメロンです.")
	fmt.Printf("名前を入力してください >>")
	user.name = getInputValue()
	fmt.Printf("%sさんこんちわ!\n次に自分の数字を3桁で入力してください(半角数字) >>",user.name)

	for {
		inputValue := getInputValue()
		flag,numbers := hanteiNumber(inputValue)
		if flag{
			user.numbers = numbers
			break
		}else{
			fmt.Println("半角で3桁の数字を入力してください!")
		}
	}

	fmt.Println("あなたが選んだ数字")
	fmt.Println("ーーーーーーーーーーーーーーー")
	for i,s := range user.numbers{
		fmt.Printf("%d桁目 : %d \n",i + 1, s)
	}

	fmt.Println("ーーーーーーーーーーーーーーー")
	fmt.Println("game Start!!!")
	fmt.Println("先攻はあなたです")
	fmt.Println(computer.numbers)

	// CPUとユーザーが交互に数値を当て合う処理
	i := 1
	for{
		if (i % 2 == 0){
			fmt.Println("----------------------CPUのターン----------------------")
			if cpuGuessNumber(user.numbers){
				fmt.Println("CPUの勝ちです... , あなたの負けぇぇぇ!!!!!")
				break
			}
		}else{
			fmt.Println("--------------------あなたのターン--------------------")
			if userGuessNumber(computer.numbers){
				fmt.Println("おめでとうございます! あなたの勝ちです!!!")
				break
			}
		}
		i = i + 1
	}

}

