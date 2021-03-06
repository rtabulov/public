## matrix_display

### Instructions

Use the Matrix struct given in the [expected struct](#expected-functions-and-struct) and implement the `std::fmt::Display` trait so it prints the matrix like in the [usage](#usage).

You will also have to implement the associated function `new` that creates a matrix from a slice of slices.

### Expected Functions and Struct

```rust
pub struct Matrix(pub Vec<Vec<i32>>);

pub fn new(slice: &[&[i32]]) -> Self {
}

use std::fmt;

impl fmt::Display for Matrix {
}
```

### Usage

Here is a possible program to test your function

```rust
use matrix_display::*;

fn main() {
	let matrix = Matrix::new(&[&[1, 2, 3], &[4, 5, 6], &[7, 8, 9]]);
	println!("{}", matrix);
}
```

And it's output:

```console
student@ubuntu:~/[[ROOT]]/test$ cargo run
(1 2 3)
(4 5 6)
(7 8 9)
student@ubuntu:~/[[ROOT]]/test$
```
