package algorithm

import "fmt"

/**
所有排序默认升序
 */

func swap(arr []int, index1, index2 int) {
	arr[index1] ,arr[index2] = arr[index2], arr[index1]
}

// 冒泡
// 时间复杂度：O(N2)，稳定性：稳定
// 依次比较相邻两元素，若前一元素大于后一元素则交换之，直至最后一个元素即为最大；然后重新从首元素开始重复同样的操作，直至倒数第二个元素即为次大元素；依次类推。如同水中的气泡，依次将最大或最小元素气泡浮出水面
func BubbleSort(arr []int)  {
	length := len(arr)
	for i :=0; i < length -1 ; i++  {
		for j := 1; j < length - i; j++  {
			if arr[j] < arr[j - 1] {
				swap(arr, j, j-1)
			}
		}
	}
}

// 鸡尾酒排序(冒泡优化版本)：1、增加标志位，若已经完成则退出，2、双向进行排序，分别把最大值与最小值进行移动
// 鸡尾酒排序优点在于比冒泡需要更少的遍历，性能上在乱序数组时，这两者均表现较差
func CocktailSort(arr []int) {
	end := len(arr) -1
	beg := 0
	for beg < end {
		top, bot := end, beg
		for i := beg; i < end; i++ {
			if arr[i] > arr[i+1] {
				top = i
				swap(arr, i, i+1)
			}
		}

		if top == end {
			break
		}
		end = top

		for j := end; j > beg ; j--  {
			if arr[j] < arr[j-1] {
				bot = j
				swap(arr, j, j-1)
			}
		}

		if bot == beg {
			break
		}
		bot = beg
	}
}

// 选择排序
// 时间复杂度：O(N2)，稳定性：不稳定
// 首先初始化最小元素索引值为首元素，依次遍历待排序数列，若遇到小于该最小索引位置处的元素则刷新最小索引为该较小元素的位置，直至遇到尾元素，结束一次遍历，并将最小索引处元素与首元素交换；然后，初始化最小索引值为第二个待排序数列元素位置，同样的操作，可得到数列第二个元素即为次小元素；以此类推。
func SelectSor()  {

}

// 快速排序
// 时间复杂度：O(NlogN)，稳定性：不稳定
// 类似于选择排序的定位思想）选一基准元素，依次将剩余元素中小于该基准元素的值放置其左侧，大于等于该基准元素的值放置其右侧；然后，取基准元素的前半部分和后半部分分别进行同样的处理；以此类推，直至各子序列剩余一个元素时，即排序完成（类比二叉树的思想，from up to down
func QuickSort(arr []int, beg ,end int) {
	if beg < end {
		p := partition(arr, beg, end)
		QuickSort(arr,beg, p-1)
		QuickSort(arr,p+1, end)
	}
}

func partition(arr []int, beg ,end int) int {
	p := arr[beg]
	i, j := beg + 1, end
	for i <= j {
		if arr[i] < p {
			i++
		}else if arr[j] > p {
			j--
		}else if arr[i] >= p && arr[j] <= p {
			swap(arr,i,j)
			i++
			j--
		}
	}

	fmt.Println(arr,"i,j:",i,j)
	swap(arr, j, beg)
	return j
}

// 插入排序
// 时间复杂度：O(N2)，稳定性：稳定
// 数列前面部分看为有序，依次将后面的无序数列元素插入到前面的有序数列中，初始状态有序数列仅有一个元素，即首元素。在将无序数列元素插入有序数列的过程中，采用了逆序遍历有序数列，相较于顺序遍历会稍显繁琐，但当数列本身已近排序状态效率会更高
func InsertSort(arr []int)  {
	
}

// 希尔排序
// 时间复杂度：通常认为是O(N3/2) ，未验证，稳定性：不稳定
// 插入排序的改进版。为了减少数据的移动次数，在初始序列较大时取较大的步长，通常取序列长度的一半，此时只有两个元素比较，交换一次；之后步长依次减半直至步长为1，即为插入排序，由于此时序列已接近有序，故插入元素时数据移动的次数会相对较少，效率得到了提高
func ShellSort()  {

}

// 基数排序
// 时间复杂度：O(x*N)，稳定性：稳定
// 桶排序的改进版，桶的大小固定为10，减少了内存空间的开销。首先，找出待排序列中得最大元素max，并依次按max的低位到高位对所有元素排序；桶元素10个元素的大小即为待排序数列元素对应数值为相等元素的个数，即每次遍历待排序数列，桶将其按对应数值位大小分为了10个层级，桶内元素值得和为待排序数列元素个数
func CountSort()  {

}

// 归并排序
// 时间复杂度：O(NlogN)，稳定性：稳定
// 采用了分治和递归的思想，递归&分治-排序整个数列如同排序两个有序数列，依次执行这个过程直至排序末端的两个元素，再依次向上层输送排序好的两个子列进行排序直至整个数列有序（类比二叉树的思想，from down to up）
func MergeSortInOrder()  {

}
