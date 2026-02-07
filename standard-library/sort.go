//go:build ignore

package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age int
}

type UserSlice []User

// Old implementation
func (u UserSlice) Len() int {
	return len(u)
}

func (u UserSlice) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u UserSlice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	users := []User {
		{"Muhammad Iqbal", 20},
		{"Ummar Afandi", 22},
		{"Ahmad Zulkarnain", 19},
	}
	
	sortedUsers := UserSlice(users)
	sort.Sort(sortedUsers)
	fmt.Println(users)

	nums := []int{5,2,9,1}
	sort.Ints(nums)
	fmt.Println(nums)

	names := []string{"Iqbal", "Andi", "Budi"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.IntsAreSorted(nums))
	fmt.Println(sort.StringsAreSorted(names))

	sort.Slice(users, func(i, j int) bool {
		return users[i].Age > users[j].Age
	})
	fmt.Println(users)
}