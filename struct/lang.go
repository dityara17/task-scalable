package _struct

type Lang struct {
	Language       string    `json:"language"`
	Appeared       int       `json:"appeared"`
	Created        []string  `json:"created"`
	Functional     bool      `json:"functional"`
	ObjectOriented bool      `json:"object-oriented"`
	Relation       Relations `json:"relation"`
}

type Relations struct {
	InfluencedBy []string `json:"influenced-by"`
	Influences []string `json:"influences"`
}

type ResponseRelation struct {
	Status int
	Desc   interface{}
}

//{
//   "language": "C",
//   "appeared": 1972,
//   "created": [
//       "Dennis Ritchie"
//   ],
//   "functional": true,
//   "object-oriented": false,
//   "relation": {
//       "influenced-by": [
//           "B",
//           "ALGOL 68",
//           "Assembly",
//           "FORTRAN"
//       ],
//       "influences": [
//           "C++",
//           "Objective-C",
//           "C#",
//           "Java",
//           "JavaScript",
//           "PHP",
//           "Go"
//       ]
//   }
//}
