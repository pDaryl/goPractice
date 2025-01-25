class Employee {
    name;
    ID;
    constructor(name, ID) {
        this.name = name;
        this.ID = ID;
    }
}

class Address {
    street
    city
    postalCode
    constructor(street, city, postalCode) {
        this.street = street;
        this.city = city;
        this.postalCode = postalCode;
    }
}

function createEmployee(name, ID, street, city, postalCode){
    var employee = new Employee(name, ID);
    var address = new Address(street, city, postalCode);
    return {employee, address};
}

function printEmployee({employee, address}){
    console.log(`Employee(Name:${employee.name}, 
        ID:${employee.ID}, 
        Address: [Street:${address.street}, 
        City:${address.city}, 
        PostalCode:${address.postalCode}])`);
}

const {employee, address} = createEmployee("Daryl", 100, "Zeta", "Belle Chasse", "70037");
printEmployee({employee, address});


