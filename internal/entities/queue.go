package entities

// implementation 1
type QueueData struct {
	Move map[int]string
}

type Queue struct {
	Steps []QueueData
}

// this is easiear to access and ideomatic in terms of name. But order of the map is not guarantied.

// implementation 2

type QueueData2 struct {
	AntName  int
	NextRoom string
}

type Queue2 struct {
	Step []QueueData2
}

// and then return will be []Queue2
// but it might be better to rename it to something else but queue
