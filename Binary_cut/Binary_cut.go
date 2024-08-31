package main
import "fmt"

func main () {
	var n, m int
	fmt.Scanf("%d %d", &n, %m)
	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}

	for ; m > 0; m-- {
		var x int
		fmt.Scanf("%d", &x)
		//逼近找左端，起始点
		l, r := 0, n-1
		for l < r {
			mid := (l + r) / 2
			if q[mid] >= x {
				r = mid
			} else {
				l = mid + 1
			}
		}
		//上面这段的mid为什么不是l+r+1，因为如果l和r只差1位，也就是n-1 n
		//并且这是在找起始点，也就是n-1=l指向的是小于x的数，n=r指向x
		//那mid=(l+r+1)/2=n=r 这时候还是指向x，mid的值赋予r相当于r=r，
		//并且l仍小于r，循环继续，陷入死循环

		//那为什么要l=mid+1，仍然还是，n-1 n
		//这时候mid向下取整，也就是mid=l
		//那mid指向小于x的数，会执行l=mid+1，加了一位后，就是l=r，会跳出for，反之则不会

		//那如果mid=l+r+1和l=mid结合呢，仍然还是， n-1 n
		//这时候mid=n=r，并且mid指向的数等于x，执行r=mid，实际上根本没有执行l=mid的机会
		//然后陷入到上面的死循环

		if q[l] != x {
			fmt.Println("-1 -1")	//这是找到了l=r跳出了for之后还是没有找到x
									//判断x不存在
		} else {
			fmt.Print(l, " ")
			l, r := 0, n-1
			//逼近找右端点，结束点
			for l < r {
				mid := (l + r + 1) / 2
				if q[mid] <= x {
					l = mid
				} else {
					r = mid - 1
				}
			}
			//这里其实是一样的，mid=l+r+1，向上取整，使得mid=r
			//这时候差一位的l和r分别是l指向x，r指向大于x的数，还是n-1 n
			//如果执行了mid，q[mid]=q[r]>x ，那么就会是执行r = mid-1 = l
			//就会跳出循环，输出l

			//那如果是mid=l+r，mid将会等于l
			//也就是q[mid]=x，执行l=mid=l，这时候l仍然小于r，陷入死循环

			fmt.Println(l)
			//Print和Println共同输出(l l(处理后))
		
		}
	}

}