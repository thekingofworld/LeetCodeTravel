package main

import "fmt"

func main() {
	nums1 := []int{2}
	nums2 := []int{1, 3, 4, 5}
	find := findMedianSortedArrays(nums1, nums2)
	fmt.Println(find)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	median := (len1 + len2) / 2
	if len1 == 0 && len2 == 0 {
		return 0
	}
	if len1 == 0 {
		if (len1 + len2) % 2 == 0 {
			if len2 == 1 {
				return float64(nums2[0])
			}
			return (float64(nums2[median-1]) + float64(nums2[median]))/2
		} else {
			return float64(nums2[median])
		}
	}
	if len2 == 0 {
		if (len1 + len2) % 2 == 0 {
			if len1 == 1 {
				return float64(nums1[0])
			}
			return (float64(nums1[median-1]) + float64(nums1[median]))/2
		} else {
			return float64(nums1[median])
		}
	}
	if nums2[0] >= nums1[len1-1] {
		if (len1 + len2) % 2 == 0 {
			var one,other float64
			if median <= len1 - 1 {
				one = float64(nums1[median])
			} else {
				one = float64(nums2[median - len1])
			}

			if (median - 1) <= len1 - 1 {
				other = float64(nums1[(median - 1)])
			} else {
				other = float64(nums2[(median - 1) - len1])
			}

			return (one + other) / 2
		} else {
			if median <= len1 - 1 {
				return float64(nums1[median])
			}
			return float64(nums2[median - len1])
		}
	} else if nums2[len2-1] <= nums1[0] {
		if (len1 + len2) % 2 == 0 {
			var one,other float64
			if median <= len2 - 1 {
				one = float64(nums2[median])
			} else {
				one = float64(nums1[median - len2])
			}

			if (median - 1) <= len2 - 1 {
				other = float64(nums2[(median - 1)])
			} else {
				other = float64(nums1[(median - 1) - len2])
			}

			return (one + other) / 2
		} else {
			if median <= len2 - 1 {
				return float64(nums2[median])
			}
			return float64(nums1[median - len2])
		}
	} else {
		if (len1 + len2) % 2 == 0 {
			var x,y,find int
			var one,other float64
			var othermedian = median - 1
			for x <= len1 - 1 || y <= len2 - 1 {
				if x <= len1 - 1 && y <= len2 - 1 {
					if nums1[x] > nums2[y] {
						find = nums2[y]
						y++
					} else {
						find = nums1[x]
						x++
					}
				} else if x <= len1 -1 {
					find = nums1[x]
					x++
				} else if y <= len2 -1 {
					find = nums2[y]
					y++
				}
				median--
				othermedian--
				if median == -1 {
					one = float64(find)
				}
				if othermedian == -1 {
					other = float64(find)
				}
				if median == -1 && othermedian == -2 {
					return (one + other) / 2
				}
			}
		} else {
			var x,y,find int
			for x <= len1 - 1 || y <= len2 - 1 {
				if x <= len1 -1 && y <= len2 -1 {
					if nums1[x] > nums2[y] {
						find = nums2[y]
						y++
					} else {
						find = nums1[x]
						x++
					}
				} else if x <= len1-1 {
					find = nums1[x]
					x++
				} else if y <= len2-1 {
					find = nums2[y]
					y++
				}
				median--
				if median == -1 {
					return float64(find)
				}
			}
		}
	}
	return 0
}