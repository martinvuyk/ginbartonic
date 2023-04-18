# [GinBarTonic MVC REST Framework](https://github.com/martinvuyk/ginbartonic) - The Bartender that only serves Gin Tonic

This will be a Gin REST MVC Framework with connection to a DB, automated deployment, and auto docs<br />
The idea of this Framework is to allow any developer to quickly setup an API that is scalable but also quick to develop with, and flexible enough for future readjustments that would otherwise be painful

## TODOs:

<!-- - [ ] Stuff -->

- [ ] Get our own cute Go Gopher bartender with a gintonic in a plate
- [x] Make the Data Abstraction layer work
- [ ] Build a CLI tool for installation of the FW
- [x] Check, update, and fork any dependencies in tonic that have fewer than 1000 dependents
- [ ] Setup CI/CD, smoketest for develop branch, regression tests for alfa

#### Database handling

The idea is to design an abstract enough interaction schema where the developer can change the database handling system of their choice. Be it an ORM like gorm, or an SQL library like the standard library, or SQLX, or SQLC with only a couple of files and functions to be changed in the models directory.

To achieve this, a Data Abstraction Layer is implemented between the database Model definition and it's "DataAspect", which is what the controllers wll know and interact with. This DataAspect will have other data representation types (e.g. "JsonResponse1", "DataReprEndpointX") where different fields use different properties of the DataAspect according to what it represents.

As such, the data definitions remain in one file, the database Model definition in another, the getters and setters that interact with the database each in a file, and the serializers which transform data from one representation to the other remain orderly and extendable.

The current attempt tries to generalize the basic functions in order for the Framework user to be able to recycle them and avoiding code redundancy but mantaining type safety

#### Error handling

Error handling is managed with the help of another project called tonic which was forked in (github.com/martinvuyk/gadgeto/tree/master/tonic), which facilitates using hooks such as validators, and error handlers, as well as the routes themselves.

The communication of errors between the controller and the endpoint must be through the juju errors package (github.com/juju/errors).

Every endpoint has to return an established ApiResponse struct, adapted to the data each endpoint's output establishes through use of Type Generics

## Contributing

#### Branches:

- main: only releases and readme updates through pull requests

- alfa: regression tested deployable abstract framework, edits through pull requests

- develop: freely alterable example use of the framework, changes are welcome and encouraged, always trying to keep everything functional through basic smoketests. Override and edit as you see fit.

#### Guidelines:

- When making a pull request to alfa, keep the changes as small as possible for each request
- Comment your intentions when altering something that was already functional
- Type safety is a must, no use putting interface{} in every return type, if we wanted that we wouldn't use Go in the first place
- Don't download a library to add 2+2 but also don't reinvent the wheel
- Keep flexibility in mind when adding any functionality or dependency
- Start small: build a use case for your idea, then try to generalize from there. If you can't, leave the example there and commit the code with comments explaining what you were trying to do
- When you have interesting ideas, leave them in a "// TODO:" comment!!!
