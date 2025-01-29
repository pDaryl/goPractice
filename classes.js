class Employee {
    constructor(employeeName, ID, tasks) {
        this.employeeName = employeeName;
        this.ID = ID;
        this.tasks = tasks || [];
    }

    addTask(task){
     this.tasks.push(task);
    }

}

function listEmployees(employeeName, ID, tasks){
    const employees = new Employee(employeeName, ID, tasks);
    return employees;
}

function findMostTasks(employees){
    var mostTasks = 0;
    var topEmployee = null;

    for(const employee of employees){
        if(employee.tasks.length > mostTasks){
            mostTasks = employee.tasks.length;
            topEmployee = employee;
        }
    }
    return topEmployee;
}

const employees = [
    new Employee(
		"daryl", 
		"69", 
		["clean computer", "take a bath"],
    ), 
    new Employee(
		"Rae", 
		"1", 
		["eat a bagel", "read a book", "pick a movie"],
    ), 
	new Employee(
		"Boo", 
		"55", 
		["meow a lot", "cuddle with daryl"],
    ), 
];

employees[0].addTask("go to gym")
employees[0].addTask("clean the house")
employees[2].addTask("use the litter box")


console.log("All Employees:");
console.log(employees);

var topEmployee = findMostTasks(employees);
console.log(`Top Employee:`);
console.log(topEmployee);


