package main

import "fmt"

type node struct {
	data int
	next *node
}

func main() {

	var first, current, new, ctr, prev *node
	var num, choice int

	for {
		fmt.Println("\n=============================\n")
		fmt.Println("1. Insert a number in list.")
		fmt.Println("2. Print the list.")
		fmt.Println("3. Delete a number from list.")
		fmt.Println("4. Quit.")
		fmt.Printf("\nEnter your choice(1-4): ")
		fmt.Scanf("%d", &choice)

		if choice == 4 {
			break
		}
		switch choice {
		case 1:
			fmt.Printf("Enter the number to insert: ")
			fmt.Scanf("%d", &num)
			if first == nil {
				first = &(node{num, nil})
				current = first
			} else {
				new = &(node{num, nil})
				current.next = new
				current = new
			}
			fmt.Printf("Number has been inserted successfully.\n")
		case 2:
			if first == nil {
				fmt.Println("List is empty.")
				break
			}
			ctr = first
			for {
				fmt.Printf("%d", ctr.data)
				if ctr.next == nil {
					break
				}
				ctr = ctr.next
				fmt.Printf(" -> ")
			}
		case 3:
			fmt.Printf("Enter the number to delete: ")
			fmt.Scanf("%d", &num)
			ctr = first
			for {
				if ctr.data == num {
					if ctr == first {
						ctr = ctr.next
						first = ctr
					} else {
						prev.next = ctr.next
					}
					fmt.Println("Number has been deleted successfully.")
					break
				}
				if ctr.next == nil {
					fmt.Println("Number is not present in the list.")
					break
				}
				prev = ctr
				ctr = ctr.next
			}
		default:
			fmt.Println("\nInvalid choice, Please enter again. \n")
		}
	}

}
