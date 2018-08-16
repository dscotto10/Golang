package main

import 	(	
		"fmt"
		"io/ioutil"
		"strings"
		"reflect"		
		"strconv"
		)

func fileimport(a string) []uint8 {

b, err := ioutil.ReadFile(a)
		
	if err != nil {
        	fmt.Print(err)
    	}
    	return b
}

func listmaker(a []uint8) []string {
	b := strings.Split(string(a), ", ")
	return b
}

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

//stepmove is a complex function. We'll go through it piece by piece.
//"position []int is the "coordinate" at which the moving piece rests at the start of an instruction.
//"facing int" is the direction that the piece is resting, as determined by "turner."
//"instruction" is the input from the instructions file. (R2, L5, etc etc etc).
//"pos_list" is a slice of slices, with every item in the slice being a coordinate.
func stepmove(position []int, facing int, instruction string, pos_list [][]int) [][]int {
	//First, we take just the "distance" from the instruction, and convert it into an integer using Atoi.
	int_instruction, err := strconv.Atoi(instruction[1:])
	//Handle errors.
	if err != nil {
        fmt.Print(err)
    	}
	//There are four possible FOR loops that we could run through, but they are all roughly the same.
	//The FOR loop is designed to append every position we pass through to the "pos_list."
	//We adjust the number of the position based on the instruction.
	//If we're facing North, we add j to the y coordinate.
	//If we're facing East, we add j to the x coordinate.
	//If we're facing South, we subtract a number from the y coordinate.
	//If we're facing West, we subtract a number from the x coordinate.
	//Each requires two steps: first, we append a "nil" item to the pos_list slice.
	//Then we replace the "nil" with our new coordinate.
	//This function returns the updated pos_list.
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

//We need an absolute value function; Golang doesn't have one built in.
//Very simple math.
func abs(x int) int {
	if x < 0 {
		return -x
	} else {
	return x
	}
}

//Simple function: is string a in map b?
//Will return "true" if yes, "false" if no.
func check_map (a string, b map[string]int) bool {
_, found := b[a]
	return found
}

//Making a concatenated string for our map.
//Two steps here: we create a list of strings: [x, x-coordinate, y, y-coordinate].
//Then we concatenate them into a single string.
func stringmaker(a []int) string {
	values := []string{"x", strconv.Itoa(a[0]), "y", strconv.Itoa(a[1])}
	return strings.Join(values,"")
}

//Our main program.
func main() {	

	inputtext := fileimport("day1input.txt")
	//fmt.Println("\n",inputtext)

	instructions := listmaker(inputtext)
	//fmt.Println("\n",instructions)

	//fmt.Println(reflect.TypeOf(instructions))
   
   //Initialize some variables to solve 1A.
	vdistance := 0
	hdistance := 0
	facing_var := 0
	//Locations is our slice of slices.
	locations := make([][]int,1)
	//pos_slice handles the individual coordinates. Our first is 0,0.
	pos_slice := []int{hdistance,vdistance}	
    
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
		answer_1a = abs(pos_slice[0]) + abs(pos_slice[1])
		
		//Print solution for 1A.
		fmt.Printf("1a: ")
		fmt.Printf("%v\n",answer_1a)		
		
    //1B is a little tougher.
    //We need to see when we get our first "repeat."
    //We'll use a FOR loop that requires a repeat to break.
    //If the coordinate is not in the map, then we iterate to the next coordinate.
    //Once we find a coordinate that IS in the map, we've found our repeat.
    //At that point, we do the math to calculate the distance.

	//Initialize two more variables.
	answer_1b := 0
	//loc_map is the map that will track coordinates, for the purposes of our check.
	loc_map := make(map[string]int)
	
	//Straightforward FOR loop here; we'll break when we get a hit.
	i := 1
	for {
		//Make our coordinate into a string, format "x0y0."
		checker := stringmaker(locations[i])
		//Check our map for that string.
		if check_map(checker, loc_map) == false {
			//If it's not there, add the string as a key to our map, with value 1.
			loc_map[checker] = 1
			//Iterate the i to go to the next location.
			i += 1
		//Otherwise, break out of this loop.
		} else {			
			break
			}
		}
		
	//Calculate 1B answer.
	answer_1b = abs(locations[i][0]) + abs(locations[i][1])

	//Print 1B solution.
	fmt.Printf("1b: ")
	fmt.Printf("%v\n",answer_1b)
}
