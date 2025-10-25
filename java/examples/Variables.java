// Example 2: Variables and Data Types
// Demonstrates Java's variable declarations and primitive types

public class Variables {
    public static void main(String[] args) {
        // Integer types
        byte smallNumber = 127;           // 8-bit: -128 to 127
        short mediumNumber = 32000;       // 16-bit: -32,768 to 32,767
        int regularNumber = 2000000000;   // 32-bit: most common
        long bigNumber = 9000000000L;     // 64-bit: note the L suffix
        
        System.out.println("Numbers: " + smallNumber + ", " + mediumNumber + 
                          ", " + regularNumber + ", " + bigNumber);
        
        // Floating point types
        float decimalFloat = 3.14f;       // 32-bit: note the f suffix
        double decimalDouble = 3.14159;   // 64-bit: default for decimals
        
        System.out.println("Decimals: " + decimalFloat + ", " + decimalDouble);
        
        // Character and Boolean
        char letter = 'A';                // 16-bit Unicode character
        boolean isJavaFun = true;         // true or false
        
        System.out.println("Char: " + letter + ", Boolean: " + isJavaFun);
        
        // String (not a primitive, but commonly used)
        String message = "Hello, Java!";
        System.out.println("String: " + message);
        
        // Variable reassignment
        regularNumber = 42;
        message = "Variables can change!";
        System.out.println("Updated: " + regularNumber + ", " + message);
        
        // Constants (final keyword)
        final double PI = 3.14159;
        final String GREETING = "Welcome";
        System.out.println(GREETING + "! PI is " + PI);
        
        // This would cause an error:
        // PI = 3.14;  // Cannot assign a value to final variable
        
        // Type inference with var (Java 10+)
        var age = 25;              // Inferred as int
        var name = "Alice";        // Inferred as String
        var price = 19.99;         // Inferred as double
        
        System.out.println("Inferred: " + name + " is " + age + 
                          " years old, price: $" + price);
    }
}

/*
 * Key Concepts:
 * - Java has 8 primitive types: byte, short, int, long, float, double, char, boolean
 * - String is a class (not primitive) but used like a primitive
 * - Variables must be declared with a type (or use 'var' for inference)
 * - 'final' keyword makes variables constant
 * - Type inference with 'var' available in Java 10+
 * 
 * To compile: javac Variables.java
 * To run: java Variables
 */
