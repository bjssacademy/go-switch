# Go - Switch

In Go (and many other programming languages) you can replace complex conditional statements with a switch statement.

Whilst nested ```if...else if``` blocks are fine in some instances, generally for ease of readability and maintainability, if you find yourself with a long list of ```if...else if``` statements, try replacing with a ```switch```.

## Syntax

```
<variable> := <whatever>
switch {
case <value1>:
    // Code to execute if <variable> equals <value1>
case <value2>:
    // Code to execute if <variable> equals <value2>
... //any other cases you need
default:
    // Code to execute if <variable> does not match any case
}
```

You can use any operators (=, !=, <>, >, < and so on to compare your variable to your value).

This is known as an expressionless switch because we are not checking for an exact match, we are checking if the time is less than something.

Here's an example of hoow you can use a switch statement when checking for exact matches using switch:

```
switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }


//using an if/else

day := time.Now().Weekday() 
if day == time.Saturday || day == time.Sunday {
    fmt.Println("It's the weekend")
} else {
        fmt.Println("It's a weekday")
}
```

You'll also note you can have *more than one* value in the same *case* line, which is an implicit OR operator (the || or "double pipe" operator in the if/else code), separated by a single comma.

There's a lot more information on switch with examples [here](https://www.programiz.com/golang/switch).