class Supplier {
    constructor(supplierName, country) {
        this.supplierName = supplierName;
        this.country = country;
    }
}

class Product {
    constructor(productName, price, supplier) {
        this.productName = productName;
        this.price = price;
        this.supplier = supplier;
    }
}

function createProduct(productName, price, supplierName, country){
    const supplier = new Supplier(supplierName, country);
    const product = new Product(productName, price, supplier);
    return product;
}

function printProduct(product){
    console.log(`Product(Name: ${product.productName}, 
        Price: ${parseFloat(product.price)},
         Supplier: [Name: ${product.supplier.supplierName}, 
         Country: ${product.supplier.country}])`);
}

const p = createProduct("Laptop", 99.99, "Ameratech", "USA");

printProduct(p);

