// Example 2: Variables and Mutability
// Demonstrates immutable and mutable variables

fn main() {
    // Immutable variable (default in Rust)
    let x = 5;
    println!("The value of x is: {}", x);
    
    // This would cause an error:
    // x = 6;  // Cannot assign twice to immutable variable
    
    // Mutable variable (explicitly marked with 'mut')
    let mut y = 10;
    println!("The value of y is: {}", y);
    
    y = 15;  // This works because y is mutable
    println!("The new value of y is: {}", y);
    
    // Shadowing - creating a new variable with the same name
    let z = 5;
    let z = z + 1;  // New variable, shadows the first
    let z = z * 2;  // Another new variable
    println!("The value of z is: {}", z);  // Prints 12
    
    // Shadowing allows changing types
    let spaces = "   ";  // String
    let spaces = spaces.len();  // Number
    println!("Number of spaces: {}", spaces);
}

/*
 * Key Concepts:
 * - Variables are immutable by default (Rust's safety feature)
 * - Use 'mut' to make a variable mutable
 * - Shadowing creates new variables with the same name
 * - Shadowing allows type changes, mutation doesn't
 */
