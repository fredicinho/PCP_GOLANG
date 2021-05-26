package main

import "fmt"

const FEMALE = "Female"
const MALE = "Male"

type Person struct {
	name string
	sex  string
}

type Persons map[string]Person
type Parents map[string][]Person

/**
Übungen PCP Prolog SW02: Verwandtschaftsbeziehungen
*/
func main() {
	var females, males = createPersons()
	var parentRelationship = createParentRelationship(females, males)

	// a.) create predicate mother/2 and father/2
	fmt.Println("Father of tina is: ", father("tina", males, parentRelationship))
	//fmt.Println("Mother of tina is: ", mother("tina", females, parentRelationship))
	//fmt.Println("Father of tom is: ", father("tom", males, parentRelationship))
	//fmt.Println("Mother of tom is: ", mother("tom", females, parentRelationship))
	// a.) get childs of parent: I have already a slice of Parents. therefore just have to check the key of parentname to get the childs...

	// b.) Create predicate sibling/2
	fmt.Println("Are mia and fred siblings? ", areSiblings("mia", "fred", females, parentRelationship))
	fmt.Println("Siblings of mia: ", siblingsOf("mia", females, parentRelationship))
	//fmt.Println("Are mia and tom siblings? ", areSiblings("mia", "tom", females, parentRelationship))
	//fmt.Println("Siblings of tom: ", siblingsOf("tom", females, parentRelationship))

	// c.) Create predicate grandmother/2
	fmt.Println("Grandmother of ann is: ", grandmother("ann", females, males, parentRelationship))
	//fmt.Println("Grandchilds of liz are: ", grandchilds("liz", females, males, parentRelationship))

	// d.) Create predicate offspring/2
	//fmt.Println("Is ann offspring of mary? ", isOffspring("ann", "mary", females, males, parentRelationship))
	//fmt.Println("Are offsprings of mary: ", offspring("mary", females, males, parentRelationship))
}

func offspring(offspringName string, females, males Persons, parentRelationship Parents) []Person {
	var offsprings []Person
	var childs, ok = parentRelationship[offspringName]
	if !ok {
		return offsprings
	}

	offsprings = append(offsprings, childs...)
	for _, person := range offsprings {
		var childsOfChild, ok = parentRelationship[person.name]
		if ok {
			offsprings = append(offsprings, childsOfChild...)
		}
	}

	return offsprings

}

func isOffspring(offspringName string, affectedPersonName string, females, males Persons, parentRelationship Parents) bool {
	var female, okFemale = females[affectedPersonName]
	var male, _ = males[affectedPersonName]
	var affectedPerson Person
	if okFemale {
		affectedPerson = female
	} else {
		affectedPerson = male
	}

	var motherOfOffspring = mother(offspringName, females, parentRelationship)
	if (Person{} != motherOfOffspring && Person{} == affectedPerson) {
		return true
	}

	var fatherOfOffspring = father(offspringName, males, parentRelationship)
	if (Person{} != fatherOfOffspring && Person{} == affectedPerson) {
		return true
	}

	var grandmother = mother(motherOfOffspring.name, females, parentRelationship)
	if (Person{} != grandmother && Person{} == affectedPerson) {
		return true
	}

	var grandfather = father(fatherOfOffspring.name, males, parentRelationship)
	if (Person{} != grandfather && Person{} == affectedPerson) {
		return true
	}

	return false
}

func grandchilds(grandmotherName string, females Persons, males Persons, parentRelationship Parents) []Person {
	var grandchilds []Person
	for femaleName, female := range females {
		var grandmother = grandmother(femaleName, females, males, parentRelationship)
		if (Person{} != grandmother) {
			grandchilds = append(grandchilds, female)
		}
		var grandfather = grandfather(femaleName, females, males, parentRelationship)
		if (Person{} != grandfather) {
			grandchilds = append(grandchilds, female)
		}
	}
	for maleName, male := range males {
		var grandmother = grandmother(maleName, females, males, parentRelationship)
		if (Person{} != grandmother) {
			grandchilds = append(grandchilds, male)
		}
		var grandfather = grandfather(maleName, females, males, parentRelationship)
		if (Person{} != grandfather) {
			grandchilds = append(grandchilds, male)
		}
	}

	return grandchilds
}

func grandfather(childName string, females Persons, males Persons, parentRelationship Parents) Person {
	var grandmotherOfMotherside = father(mother(childName, females, parentRelationship).name, females, parentRelationship)
	if (Person{} != grandmotherOfMotherside) {
		return grandmotherOfMotherside
	}

	var grandMotherOfFatherside = father(father(childName, males, parentRelationship).name, females, parentRelationship)
	return grandMotherOfFatherside
}

func grandmother(childName string, females Persons, males Persons, parentRelationship Parents) Person {
	var grandmotherOfMotherside = mother(mother(childName, females, parentRelationship).name, females, parentRelationship)
	if (Person{} != grandmotherOfMotherside) {
		return grandmotherOfMotherside
	}

	var grandMotherOfFatherside = mother(father(childName, males, parentRelationship).name, females, parentRelationship)
	return grandMotherOfFatherside
}

func siblingsOf(childName string, females Persons, parentRelationship Parents) []Person {
	return parentRelationship[mother(childName, females, parentRelationship).name]
}

func areSiblings(childOne string, childTwo string, females Persons, parentRelationship Parents) bool {
	return mother(childOne, females, parentRelationship) == mother(childTwo, females, parentRelationship)
}

func father(nameOfChild string, males Persons, parents Parents) Person {
	for parentName, childs := range parents {
		for _, child := range childs {
			if child.name == nameOfChild {
				father, ok := males[parentName]
				if ok {
					return father
				}
			}
		}
	}
	return Person{}
}

func mother(nameOfChild string, females Persons, parents Parents) Person {
	for parentName, childs := range parents {
		for _, child := range childs {
			if child.name == nameOfChild {
				mother, ok := females[parentName]
				if ok {
					return mother
				}
			}
		}
	}
	return Person{}
}

func createPersons() (map[string]Person, map[string]Person) {
	femaleNames := []string{"mary", "liz", "mia", "tina", "ann", "sue"}
	females := make(Persons)
	for _, femaleName := range femaleNames {
		females[femaleName] = Person{femaleName, FEMALE}
	}

	maleNames := []string{"mike", "jack", "fred", "tom", "joe", "jim"}
	males := make(Persons)
	for _, maleName := range maleNames {
		males[maleName] = Person{maleName, MALE}
	}

	return females, males
}

func createParentRelationship(females, males map[string]Person) map[string][]Person {
	parentRelationship := make(Parents)
	parentRelationship["mary"] = []Person{females["mia"], males["fred"], females["tina"]}
	parentRelationship["mike"] = []Person{females["mia"], males["fred"], females["tina"]}
	parentRelationship["liz"] = []Person{males["tom"], males["joe"]}
	parentRelationship["jack"] = []Person{males["tom"], males["joe"]}
	parentRelationship["mia"] = []Person{females["ann"]}
	parentRelationship["tina"] = []Person{females["sue"], males["jim"]}
	parentRelationship["tom"] = []Person{females["sue"], males["jim"]}
	return parentRelationship
}
