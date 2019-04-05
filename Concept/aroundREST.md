## REST
----
##GET & POST.

### Dos

 - Use Get if the interaction is more like a question (i.e., it is a safe operation such as a query, read     operation, or lookup)
 
 - Use POST if 
    The interaction is more like an order, or
    The interaction changes the state of the resource in a way that the user would perceive (e.g., a subscription to a service), or
    The user be held accountable for the results of the interaction.

### Don'ts

 - Use Get if the interaction produces any side effect like altering data.
 
 - Use POST if it is only to retrieve data.