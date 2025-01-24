class Person {
    name;
    age;
    constructor(name, age) {
        this.name = name;
        this.age = age;
    }
}

function createPerson(name, age){
    var person = new Person(name, age);
    return person;
}

const person = createPerson("Daryl", 28);
console.log(`Name: ${person.name}, Age: ${person.age}`);

