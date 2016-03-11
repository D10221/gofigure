package inutil

type Anything interface {

}

func FirstOrDefault(anythings []Anything, isFound func(t Anything) bool) Anything {
	for i := 0; i < len(anythings); i++ {
		if isFound(anythings[i]) {
			return anythings[i]
		}
	}
	return nil
}
