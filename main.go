package main

import (
	//"fmt"
	"lem-in/graphs"
	"lem-in/rooms"
)



func main(){
	myGraph := graphs.Graph{}
	 myGraph.NewGraph([]string{"start","r1","r2","r3","r4","r5","r6","r7","r8","r9","r10","r11","end"})
	//room_start := rooms.NewRoom("s",1,2)
	room_start := rooms.NewRoom("start",1,2)
	room_r1 := rooms.NewRoom("r1",1,2)
	room_r2 := rooms.NewRoom("r2",1,2)
	room_r3:= rooms.NewRoom("r3",1,2)
	room_r4 := rooms.NewRoom("r4",1,2)
	room_r5 := rooms.NewRoom("r5",1,2)
	room_r6 := rooms.NewRoom("r6",1,2)
	room_r7 := rooms.NewRoom("r7",1,2)
	room_r8 := rooms.NewRoom("r8",1,2)
	room_r9 := rooms.NewRoom("r9",1,2)
	room_r10 := rooms.NewRoom("r10",1,2)
	room_r11 := rooms.NewRoom("r11",1,2)
	room_end := rooms.NewRoom("end",1,2)

	 myGraph.AddRoom(room_start,room_r1)
	 myGraph.AddRoom(room_start,room_r2)
	 myGraph.AddRoom(room_start,room_r3)
	 myGraph.AddRoom(room_r1,room_r4)
	 myGraph.AddRoom(room_r2,room_r5) 
	 myGraph.AddRoom(room_r3,room_r6)
	 myGraph.AddRoom(room_r4,room_r7)
	 myGraph.AddRoom(room_r5,room_r8)
	 myGraph.AddRoom(room_r6,room_r9)
	 myGraph.AddRoom(room_r6,room_r5)
	 myGraph.AddRoom(room_r6,room_r10)
	 myGraph.AddRoom(room_r10,room_r11)
	 myGraph.AddRoom(room_r7,room_end)
	 myGraph.AddRoom(room_r8,room_end)
	 myGraph.AddRoom(room_r9,room_end)
	 myGraph.DFSExplore(room_start)
	//myGraph.ShowAdjList()
	/*myGraph := NewGraph(11)
	myGraph.AddNode(0,4)
	myGraph.AddNode(0,6)
	myGraph.AddNode(1,3)
	myGraph.AddNode(4,3)
	myGraph.AddNode(5,2)
	myGraph.AddNode(3,5)
	myGraph.AddNode(4,2)
	myGraph.AddNode(2,1)
	myGraph.AddNode(7,6)
	myGraph.AddNode(7,2)
	myGraph.AddNode(7,4)
	myGraph.AddNode(6,5)
	myGraph.DFSsearch(0,5)*/
}