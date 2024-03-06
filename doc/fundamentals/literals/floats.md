# Floats

Go has two floating-point types, `float32` and `float64`. Go uses the [IEEE 754 specification](https://ieeexplore.ieee.org/document/30711), giving a large range and limited precision. Unless you have to be compatible with an existing format, `use float64` since float literals defaul value is float64.

## Precision Problems

Similar to many programming languages, Go's floating-point numbers have an extensive range, yet `they cannot precisely represent every value` within that spectrum; rather, **`they store the closest approximation`**. Given their inherent imprecision, floating-point numbers are suitable only for scenarios where approximate values suffice or where the nuances of floating-point arithmetic are well understood.

Although Go permits the use of `==` and `!=` operators for comparing floats, it's advisable to avoid doing so. **`Due to the inherent imprecision of floats, two floating-point values may not be considered equal even when they appear identical`**. Instead, consider defining a maximum permissible variance and comparing whether the difference between two floats falls below this threshold. This threshold value (often referred to as **epsilon**), should be chosen based on the precision requirements of your application.

!!! tip "Help to define variance value"
    Checkout: [Floating-Point Guide](https://floating-point-gui.de/errors/comparison/)
