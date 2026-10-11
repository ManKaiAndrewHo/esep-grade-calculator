package esepunittests

import (
	"testing"
)

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 70, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 70, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 65, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 62, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 40, Assignment)
	gradeCalculator.AddGrade("exam 1", 50, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 55, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestWeightsAreAppliedPerCategory(t *testing.T) {
	gradeCalculator := NewGradeCalculator()

	// Assignments 100, exams 70, essays 40 -> 50 + 24.5 + 6 = 80.5 -> "B"
	gradeCalculator.AddGrade("assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam", 70, Exam)
	gradeCalculator.AddGrade("essay", 40, Essay)

	if actual_value := gradeCalculator.GetFinalGrade(); actual_value != "B" {
		t.Errorf("Expected 'B'; got '%s' instead", actual_value)
	}
}

func TestMultipleGradesAreAveraged(t *testing.T) {
	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("assignment 1", 100, Assignment)
	gradeCalculator.AddGrade("assignment 2", 80, Assignment) // average 90
	gradeCalculator.AddGrade("exam 1", 90, Exam)
	gradeCalculator.AddGrade("essay 1", 90, Essay)

	if actual_value := gradeCalculator.GetFinalGrade(); actual_value != "A" {
		t.Errorf("Expected 'A'; got '%s' instead", actual_value)
	}
}

func TestGradeTypeString(t *testing.T) {
	cases := map[GradeType]string{
		Assignment: "assignment",
		Exam:       "exam",
		Essay:      "essay",
	}

	for gradeType, expected_value := range cases {
		if actual_value := gradeType.String(); actual_value != expected_value {
			t.Errorf("Expected String() to return '%s'; got '%s' instead", expected_value, actual_value)
		}
	}
}

func newPassFail(assignmentGrade, examGrade, essayGrade int) *GradeCalculator {
	GradeCalculator := NewGradeCalculatorWithMode(PassFail)

	GradeCalculator.AddGrade("assignment", assignmentGrade, Assignment)
	GradeCalculator.AddGrade("exam", examGrade, Exam)
	GradeCalculator.AddGrade("essay", essayGrade, Essay)

	return GradeCalculator
}

func TestPassFailWithA(t *testing.T) {
	if actual_value := newPassFail(100, 100, 100).GetFinalGrade(); actual_value != "Pass" {
		t.Errorf("Expected 'Pass'; got '%s' instead", actual_value)
	}
}

func TestPassFailWithC(t *testing.T) {
	if actual_value := newPassFail(70, 70, 70).GetFinalGrade(); actual_value != "Pass" {
		t.Errorf("Expected 'Pass'; got '%s' instead", actual_value)
	}
}

func TestPassFailWithD(t *testing.T) {
	if actual_value := newPassFail(69, 69, 69).GetFinalGrade(); actual_value != "Fail" {
		t.Errorf("Expected 'Fail'; got '%s' instead", actual_value)
	}
}

func TestPassFailWithF(t *testing.T) {
	if actual_value := newPassFail(40, 50, 55).GetFinalGrade(); actual_value != "Fail" {
		t.Errorf("Expected 'Fail'; got '%s' instead", actual_value)
	}
}

func TestDefaultReturnsLetterGrade(t *testing.T) {
	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam", 100, Exam)
	gradeCalculator.AddGrade("essay", 100, Essay)

	if actual_value := gradeCalculator.GetFinalGrade(); actual_value != "A" {
		t.Errorf("Expected 'A'; got '%s' instead", actual_value)
	}
}
