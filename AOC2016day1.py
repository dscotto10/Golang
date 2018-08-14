
def fileimport(input_file1):
	input1 = open(input_file1, 'r')
	input1 = input1.read()
	return input1
	
def listmaker(a):
	b = a.split(", ")
	return b
	
def turner (a,b):
	if a[0] == 'R':
		b += 1
	else:
		b -= 1
	if b == -1:
		b = 3
	if b == 4:
		b = 0
	return b

def stepmove(position, facing, instruction, pos_list):
	int_instruction = instruction[1:]
	if facing == 0:
		for j in range(1,int(int_instruction)+1):
			pos_list.append([position[0],position[1]+j])
	elif facing == 1:
		for j in range(1,int(int_instruction)+1):
			pos_list.append([position[0]+j,position[1]])
	elif facing == 2:
		for j in range(1,int(int_instruction)+1):
			pos_list.append([position[0],position[1]-j])
	else:
		for j in range(1,int(int_instruction)+1):
			pos_list.append([position[0]-j,position[1]])
	return pos_list
		

inputtext = fileimport("day1input.txt")
instructions = listmaker(inputtext)
vdistance = 0
hdistance = 0
facing_var = 0
locations = []
pos_slice = [hdistance, vdistance]

for i in instructions:
	facing_var = turner(i,facing_var)
	locations = stepmove(pos_slice, facing_var, i, locations) 
	pos_slice[0] = locations[len(locations) - 1][0]
	pos_slice[1] = locations[len(locations) - 1][1]

answer_1a = abs(pos_slice[0]) + abs(pos_slice[1])

print "1a: " + str(answer_1a)

loc_map = []

for i in range(0,len(locations)):
	if locations[i] not in loc_map:
		loc_map.append(locations[i])
	else:
		break
answer_1b = abs(locations[i][0]) + abs(locations[i][1])

print "1b: " + str(answer_1b)
