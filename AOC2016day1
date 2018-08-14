package main

import (	"fmt"
		  	"io/ioutil"
		  	"strings"		
			"strconv"
		)

//The "turner" function here uses numbers 0-3 to represent directions.
//N = 0. E = 1. S = 2. W = 3.
//The function has a built-in "reset" if the numbers get outside of the range.
//-1 becomes 3. 4 becomes 0. And the cycle continues.
func turner(a string,b int) int {
	if a[0] == 'R' {
		b += 1
		} else {
		b -= 1
		}
	if b == -1 {
		b = 3}
	if b == 4 {
		b = 0}
	return b
}

//If our heading is 0 (North) or 1 (East), we add to the distance value.
//If our heading is 2 (South) or 3 (West), we subtract from the distance value.
//This is essentially like a coordinate grid.
func moveme(distance int, facing int, instruction string) int {
	int_instruction, err := strconv.Atoi(instruction[1:])
	if err != nil {
        fmt.Print(err)
    }	
	if facing < 2 {
		distance += int(int_instruction)
		} else {
			distance -= int(int_instruction)
			}
	return distance
}

//Note that func moveme was deprecated by 1B.
//stepmove operates under a similar principle.
//It moves one step at a time and adds every position to a slice of slices.
func stepmove(position []int, facing int, instruction string, pos_list [][]int) [][]int {
	int_instruction, err := strconv.Atoi(instruction[1:])
	if err != nil {
        fmt.Print(err)
    }
    if facing == 0 {
	 	for j := 1; j <= int_instruction; j++ {
	 		pos_list = append(pos_list,nil)	 		
	 		pos_list[len(pos_list)-1] = append(pos_list[len(pos_list)-1], position[0],position[1]+j)
	 		}
	 	} else if facing == 1 {
		for j := 1; j <= int_instruction; j++ {
	 		pos_list = append(pos_list,nil)	 		
	 		pos_list[len(pos_list)-1] = append(pos_list[len(pos_list)-1], position[0]+j,position[1])
	 		}
	 	} else if facing == 2 {
	 	for j := 1; j <= int_instruction; j++ {
	 		pos_list = append(pos_list,nil)	 		
	 		pos_list[len(pos_list)-1] = append(pos_list[len(pos_list)-1], position[0],position[1]-j)
	 		}
	 	} else {
	 	for j := 1; j <= int_instruction; j++ {
	 		pos_list = append(pos_list,nil)	 		
	 		pos_list[len(pos_list)-1] = append(pos_list[len(pos_list)-1], position[0]-j,position[1])
	 		}    
    }
    return pos_list

}


//Simple function: is string a in map b?
func check_map (a string, b map[string]int) bool {
_, found := b[a]
	return found
}

//Absolute value. Apparently this doesn't exist in Golang.
func Abs(x int) int {
	if x < 0 {
		return -x
	} else {
	return x
	}
}

//Making a concatenated string for our map.
func stringmaker(a []int) string {
	values := []string{"x", strconv.Itoa(a[0]), "y", strconv.Itoa(a[1])}
	return strings.Join(values,"")
}

func main() {	
		//Read the input, and error handle.		
		inputtext, err := ioutil.ReadFile("day1input.txt")
		
		if err != nil {
        fmt.Print(err)
    }		
		
		//Convert the input text into a slice.
		instructions := strings.Split(string(inputtext), ", ")
		
    //Initialize some variables.
		vdistance := 0
		hdistance := 0
		facing_var := 0
		locations := make([][]int,1)
		pos_slice := []int{hdistance,vdistance}	
    answer_1b := 0
		loc_map := make(map[string]int)

    
    //1A for loop.
    //Set the direction, then keep adding positions to the locations slice.
    //Then move to the next position via pos_slice.
    //This for loop will iterate through the whole list of instructions.
    for i := 0; i < len(instructions); i++ {		
			facing_var = turner(instructions[i],facing_var)
			locations = stepmove(pos_slice, facing_var, instructions[i], locations)		
			pos_slice[0] = locations[len(locations) - 1][0]
			pos_slice[1] = locations[len(locations) - 1][1]
			}
      
     //Calculate 1A solution: what is our final position?
		answer_1a := 0
		answer_1a = Abs(pos_slice[0]) + Abs(pos_slice[1])
		
		//Print solution for 1A.
		fmt.Printf("1a: ")
		fmt.Printf("%v\n",answer_1a)		
		
    //1B is a little tougher.
    //We need to see when we get our first "repeat."
    //We'll use a FOR loop that requires a repeat to break.
    //If the coordinate is not in the map, then we iterate to the next coordinate.
    //Once we find a coordinate that IS in the map, we've found our repeat.
    //At that point, we do the math to calculate the distance.
		i := 1
		for {
			checker := stringmaker(locations[i])
		
			if check_map(checker, loc_map) == false {
				loc_map[checker] = 1
				i += 1
			} else {			
				loc_map[checker] += 1				
				answer_1b = Abs(locations[i][0]) + Abs(locations[i][1])
				break
			}
		}
		
		//Print our 1B solution.
		fmt.Printf("1b: ")
		fmt.Printf("%v\n",answer_1b)
		}
		
