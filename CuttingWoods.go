package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	r := bufio.NewReader(os.Stdin) // 표준 입력에서 한 줄씩 읽기 위한 Reader 생성
	w := bufio.NewWriter(os.Stdout) // 표준 출력에 쓰기 위한 Writer 생성
	defer w.Flush() // main 함수가 종료될 때 Writer 비우기

	var n, m int
	fmt.Fscan(r, &n, &m) // 표준 입력에서 두 개의 정수를 읽어와 변수 n과 m에 저장

	woods := make([]int, n) // 길이가 n인 정수형 배열 생성
	max := 0 // 나무의 최대 높이를 저장할 변수 초기화

	for i := 0; i < n; i++{ // n번 반복하여 나무의 높이를 읽어와서 배열에 저장
		fmt.Fscan(r, &woods[i]) // 표준 입력에서 한 개의 정수를 읽어와 woods 배열에 저장

		if max < woods[i]{ // 현재 나무의 높이가 최대 높이보다 크면
			max = woods[i] // 최대 높이를 갱신
		}
	}

	// 이진 탐색 함수를 호출하여 최적의 높이를 찾고 결과를 출력
	fmt.Fprintln(w, binSearch(woods, 1, max, m))
}

// 이진 탐색 함수
func binSearch(arr []int, l, r, target int) (result int){
	mid := (l + r) / 2 // 중간 높이 계산
	cut := cut(arr, mid) // 중간 높이로 나무를 잘랐을 때 얻게 되는 나무의 길이

	// 범위가 역전된 경우 종료
	if l > r{
		if l == target{ // l이 목표 길이와 같으면 l 반환
			return l
		}
		return r // 그렇지 않으면 r 반환
	}

	// 이진 탐색 수행
	if cut > target{ // 잘린 나무의 길이가 목표 길이보다 크면
		result = binSearch(arr, mid + 1, r, target) // 오른쪽 반을 탐색
	} else if cut < target{ // 잘린 나무의 길이가 목표 길이보다 작으면
		result = binSearch(arr, l, mid - 1, target) // 왼쪽 반을 탐색
	} else { // 잘린 나무의 길이가 목표 길이와 같으면
		result = mid // 현재 높이를 결과로 반환
	}

	return result
}

// 주어진 높이로 나무를 잘랐을 때 얻게 되는 나무의 길이를 계산하는 함수
func cut(arr []int, h int) (result int){
	result = 0 // 결과를 저장할 변수 초기화

	// 모든 나무에 대해 반복하여
	for _, v := range arr{
		if v > h{ // 나무의 높이가 주어진 높이보다 크면
			result += v - h // 잘린 나무의 길이를 누적
		}
	}

	return // 잘린 나무의 총 길이 반환
}
