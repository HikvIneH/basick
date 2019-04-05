## Code Smell
----
##Refused Bequest.

Quoting Martin Fowler's book Refactoring: improving the design of existing code: 

> Subclasses get to inherit the methods 
> and data of their parents. But what if 
> they don't want or need what they are 
> given? They are given all these great
> gifts and pick just a few to play with.

For example, the SubClass below, it ignores the functionalities of BaseClass and overrides it. 
Taking a bad decision, we decided that our SubClass would inherit from BaseClass.
Furthermore the inherited class would not be able to replace the code without affecting its 
funcionality.

```java
public class BaseClass{
 
    public virtual void MethodA(){
        Console.WriteLine("Do Something Usefull");
    } 
}
 
public class SubClass:BaseClass{
    public override void MethodA(){
        Console.WriteLine("Do Something More Usefull");
    }
}
```
In this case, like the example above, there is a workaround called Push Down method, 
where the code is refactored in a way that the method is moved completely to the SubClass, 
as none of the SubClass uses the functionality from the BaseClass. The example of the workaround
would be something like the code below. 


```java
public class BaseClass{
// Other used methods
}
 
public class SubClass:BaseClass{
    public override void Method(){
        Console.WriteLine("Do Something More Usefull");
    }
}
```
----

## Dependency Injection
----
Well, as brief as possible:

> Dependency Injection is a realization  
> of the inversion control. It worked to  
> create a dependency outside of the class  
> that uses it. Inject it from the outside
> and gain control over their creation away
> from inside of our current class.







