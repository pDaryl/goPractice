class Student {
    constructor(studentName, ID, grade) {
        this.studentName = studentName;
        this.ID = ID;
        this.grades = grade || [];
    }
}

function listStudents(students){
    for (const s of students){
        console.log(`Student: ${s.studentName}, ID: ${s.ID}, Grades: ${s.grades}`);
    }
}

function averageGrades(grades){
    var sum = 0;

    for (const g of grades){
        sum += g;
    }
    var GPA = sum / grades.length;

    return GPA;
}

function topStudent(students){
    var topStudent;
    var highestGPA = 0;

    for(const s of students){
       var currentGPA = averageGrades(s.grades);
       if(currentGPA > highestGPA){
        topStudent = s;
        highestGPA = currentGPA;
       }
    }
    return topStudent;
}

function printStudentDetails(students){
    console.log(`Student: ${students.studentName}, ID: ${students.ID}, Grades: ${students.grades}`);
}

const students = [
    new Student("Daryl", "55", [80, 78, 92, 88]), 
    new Student("Rae", "69", [82, 79, 90, 89]), 
    new Student("Boo", "1", [99, 90, 80, 70]), 
];

console.log(`All Students:`);
for(const s of students){
    printStudentDetails(s);
}

console.log(`Top Student Details:`);
const theTopStudent = topStudent(students);
printStudentDetails(theTopStudent);