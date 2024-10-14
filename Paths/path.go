package paths

import (
	"fmt"
	"sort"
	"strconv"
)

var (
	PathsList = make([]path, 0)
)

type path struct {
	length       int
	Ants         []string
	arrayOfRooms []string
}

func GetAllPaths(arrayOfPaths [][]string) {
	arrayOfPaths = sortPaths(arrayOfPaths)
	for _, p := range arrayOfPaths {
		mypath := path{length: len(p), arrayOfRooms: make([]string, 0)}
		for _, room := range p {
			mypath.arrayOfRooms = append(mypath.arrayOfRooms, room)
		}
		PathsList = append(PathsList, mypath)
	}
}

func ShowPathList() {
	for _, path := range PathsList {
		fmt.Println("====================================")
		fmt.Printf("length : %d \n", path.length)
		fmt.Printf("Ants : %v \n", path.Ants)
		for _, room := range path.arrayOfRooms {
			fmt.Printf("-> %v", room)
		}
		fmt.Println()
	}
}

/*func sortPaths(allPaths []path) {

}*/

func MakeAntsInPlaces(number_Ants int) {
	paths_size := len(PathsList)
	ant_Id := 1
	index := 0
	for ant_Id <= number_Ants {
		if index < paths_size-1 {
			room_size := PathsList[index].length
			ants_Size := len(PathsList[index].Ants)
			if room_size+ants_Size > PathsList[index+1].length {
				index++
				PathsList[index].Ants = append(PathsList[index].Ants, "L"+strconv.Itoa(ant_Id))
			} else {
				index = 0
				PathsList[index].Ants = append(PathsList[index].Ants, "L"+strconv.Itoa(ant_Id))
			}
		} else {
			index = 0
			PathsList[index].Ants = append(PathsList[index].Ants, "L"+strconv.Itoa(ant_Id))
		}
		ant_Id++
	}
}

//abdelilah code ==============================================================

func sortPaths(mypaths [][]string) [][]string {
	sort.Slice(mypaths, func(i, j int) bool {
		if len(mypaths[i]) != len(mypaths[j]) {
			return len(mypaths[i]) < len(mypaths[j])
		}
		return false
	})
	return mypaths
}

func arePathsEqual(path1, path2 []string) bool {
	fmt.Println(path1[1:len(path1)-1], " , ", path2[1:len(path2)-1])
	for i := 0; i < len(path1[1:len(path1)-1]); i++ {
		for j := 0; j < len(path2[1:len(path2)-1]); j++ {
			if path1[1 : len(path1)-1][i] == path2[1 : len(path2)-1][j] {
				return true
			}
		}
	}
	return false
}

func GroupUniquePaths(paths [][]string) [][][]string {
	paths = sortPaths(paths)
	var groups [][][]string
	used := make(map[int]bool)
	for i, path := range paths {
		if used[i] {
			continue
		}
		group := [][]string{path}
		for j := i + 1; j < len(paths); j++ {
			if used[j] {
				continue
			}
			unique := true
			for _, groupPath := range group {
				if arePathsEqual(paths[j], groupPath) {
					unique = false
					break
				}
			}
			if unique {
				group = append(group, paths[j])
			}
		}
		groups = append(groups, group)
	}

	return groups
}

func PrintData(paths [][]string) {
	groupedPaths := GroupUniquePaths(paths)

	for i, group := range groupedPaths {
		fmt.Printf("Group %d:\n", i+1)
		for _, path := range group {
			fmt.Printf("  %v\n", path)
		}
		fmt.Println()
	}
}
