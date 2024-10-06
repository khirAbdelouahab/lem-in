package readata

import (
	"regexp"
	"strings"
)

type Start struct {
	start      string
	startCheck bool
	x          string
	y          string
}

type End struct {
	end      string
	endCheck bool
	x        string
	y        string
}

func ExtractStartAndEnd(data []string) {
	var s Start
	var e End
	for i, v := range data {
		if v == "##start" && i != len(data)-1 && !s.startCheck {
			hold := strings.Split(data[i+1], " ")
			s.start = hold[0]
			s.startCheck = true
		}
		if v == "##end" && i != len(data)-1 && !e.endCheck {
			hold := strings.Split(data[i+1], " ")
			e.end = hold[0]
			e.endCheck = true
		}
	}
}

func ExtractEdgeList(data []string) [][]string {
	var EdgeList [][]string
	for _, v := range data {
		pattern := regexp.MustCompile(`^[a-zA-Z0-9]+-[a-zA-Z0-9$]+`)
		edge := pattern.FindAllString(v, -1)
		if len(edge) > 0 {
			EdgeList = append(EdgeList, edge)
		}
	}

	return EdgeList
}

func EdgeListToAdjList(List [][]string) map[string][]string {
	var hold [][]string
	Map := make(map[string][]string)
	for _, Node := range List {
		for _, v := range Node {

			pattern := regexp.MustCompile(`^[a-zA-Z0-9]+-[a-zA-Z0-9]+$`)
			found := pattern.FindAllString(v, -1)
			if len(found) > 0 {
				h := strings.Split(found[0], "-")
				hold = append(hold, h)
			}

		}
	}

	for _, v := range hold {

		Map[v[0]] = append(Map[v[0]], v[1])

		Map[v[1]] = append(Map[v[1]], v[0])

	}
	//fmt.Println(m)

	return Map
}
