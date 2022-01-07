package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "Averil",
			LastName:  "Simen",
			Grades: []Grade{
				{
					Title: "quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "quiz 2",
					Type:  GradeHomework,
					Score: 90,
				},
				{
					Title: "quiz 3",
					Type:  GradeTest,
					Score: 40,
				},
			},
		},
		{
			ID:        2,
			FirstName: "Defil",
			LastName:  "Simen",
			Grades: []Grade{
				{
					Title: "quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "quiz 2",
					Type:  GradeHomework,
					Score: 90,
				},
				{
					Title: "quiz 3",
					Type:  GradeTest,
					Score: 40,
				},
			},
		},
		{
			ID:        3,
			FirstName: "Feril",
			LastName:  "Simen",
			Grades: []Grade{
				{
					Title: "quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "quiz 2",
					Type:  GradeHomework,
					Score: 90,
				},
				{
					Title: "quiz 3",
					Type:  GradeTest,
					Score: 40,
				},
			},
		},
		{
			ID:        4,
			FirstName: "Weril",
			LastName:  "Simen",
			Grades: []Grade{
				{
					Title: "quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "quiz 2",
					Type:  GradeHomework,
					Score: 90,
				},
				{
					Title: "quiz 3",
					Type:  GradeTest,
					Score: 40,
				},
			},
		},
	}
}
