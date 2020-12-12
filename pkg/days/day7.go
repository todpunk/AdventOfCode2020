package days

import (
	"strconv"
	"strings"
)

type bag_rule struct {
	color string
	count int64
}

func Day7a(input []string) int {
	var my_color = "shiny gold"
	var bags = map[string][]bag_rule{}
	var contained_in = map[string][]bag_rule{}

	for _, rule := range input {
		var splits = strings.Split(rule, " bags contain ")
		bags[splits[0]] = []bag_rule{}
		for _, rule_part := range strings.Split(splits[1], ", ") {
			var bag = rule_part[:strings.Index(rule_part, " bag")]
			if bag != "no other" {
				var bag_num, _ = strconv.ParseInt(bag[:strings.Index(bag, " ")], 10, 64)
				var bag_name = bag[strings.Index(bag, " ")+1:]
				bags[splits[0]] = append(bags[splits[0]], bag_rule{bag_name, bag_num})
				if _, ok := contained_in[bag_name]; !ok {
					contained_in[bag_name] = []bag_rule{}
				}
				contained_in[bag_name] = append(contained_in[bag_name], bag_rule{splits[0], bag_num})
			}
		}
	}

	var my_bags = contained_in[my_color]
	var seen = map[string]bool{}
	for i := 0; i < len(my_bags); i++ {
		for _, next_bag := range contained_in[my_bags[i].color] {
			if _, ok := seen[next_bag.color]; !ok {
				my_bags = append(my_bags, next_bag)
				seen[next_bag.color] = true
			}
		}
	}
	return len(my_bags)
}

func Day7b(input []string) int {
	var my_color = "shiny gold"
	var bags = map[string][]bag_rule{}

	for _, rule := range input {
		var splits = strings.Split(rule, " bags contain ")
		bags[splits[0]] = []bag_rule{}
		for _, rule_part := range strings.Split(splits[1], ", ") {
			var bag = rule_part[:strings.Index(rule_part, " bag")]
			if bag != "no other" {
				var bag_num, _ = strconv.ParseInt(bag[:strings.Index(bag, " ")], 10, 64)
				var bag_name = bag[strings.Index(bag, " ")+1:]
				bags[splits[0]] = append(bags[splits[0]], bag_rule{bag_name, bag_num})
			}
		}
	}

	var my_bags = []bag_rule{bag_rule{my_color, 0}}
	var n int64 = 0
	for i := 0; i < len(my_bags); i++ {
		for _, next_bag := range bags[my_bags[i].color] {
			for n = 0; n < next_bag.count; n++ {
				my_bags = append(my_bags, next_bag)
			}
		}
	}
	return len(my_bags) - 1
}
