API Workshop - Endpoint Design

## Sections:

* [Bread Crumb Navigation](#bread-crumb-navigation)

#### Restful URL

Use nouns not verbs for your resources

Let us make student resources

* `GET /students` - Retrieves a list of students

* `GET /students/:id` - Retrieves a specific student

* `POST /students` - Creates a new student

* `PUT /students/:id` - Updates student #5

* `PATCH /students/:id` - Partially updates student #5

* `DELETE /students/:id` - Deletes student #5

The `/students` endpoint is for a collection

The `/students/:id` endpoint is for a specific element in the collection like for example `/students/5`

Doing this will help you get HTTP verbs/actions out of your base URLs.

##### Why are verbs bad in URLs?

Let us try to think about all the particular types of actions we can do with students:

* `/getAllStudents`
* `/checkStudentGrades`
* `/retrieveStudent`
* `/updateStudents`
* `/updateSpecificStudent`
* `/deleteAllStudents`
* `/deleteSpecificStudent`

Mapping urls to actions can quickly get out of hand and is not intuitive for app developers.

**Keep verbs out of your base URLs**

**Use HTTP verbs to operate on the collections and elements.**

##### CRUD

For our student resources, we have two base URLs that use nouns as labels, and we can operate on them with HTTP verbs. 

Our HTTP verbs are `POST`, `GET`, `PUT`, `PATCH`, and `DELETE`. 

We can map them to the acronym CRUD: `CREATE READ UPDATE DELETE`

##### Resource Table

| Resource | POST | GET | PUT | PATCH | DELETE |
| --- | --- | --- | --- | --- | --- |
| `/students` | Create a New Student | List all students | Bulk Update of all Students | Partial Update | Delete all Students |
| `/students/:id` | Error | Show Rocky Balboa | If Rocky exists then update him else Error | Delete Rocky Balboa |

As you can see With the two resources (`/students` and `/students/:id`) and the 5 HTTP verbs, we have a rich set of capabilities that is intuitive to use.

**Tip:**

* Use two base URLs per resource.

* Keep verbs out of your base URLs.  

* Use HTTP verbs to operate on the collections and elements.

#### Plural and Concrete Names

Try to favor plural names for a resource but above all else be consistent, don't use a plural name for one resource and then a singular name for another, this can quickly get confusing for an API consumer.

For example we used `/students` over `/student` for our student resources

Choose concrete names over abstract names

For example what does an `/items` resource convey to you, not much right

But what does `/toys` resource convey to you

The level of abstraction can depend on scenarios but concrete names are more intuitive for API consumers.

#### Resource Associations

Resources tend to have relationships with other resources.

Let us go back to our students resource.

Now let us say that each student will have an associated teachers

If we wanted to get all the students belonging to a specific teacher we could do the following:

`GET /teachers/:id/students` so for example an http request like this `GET /teachers/754/students` would get all students belonging to teacher number 754

If we wanted to add a new student to a particular teacher we could do the following:

`POST /teachers/:id/students` so for example an http request like this `POST /teachers/754/student` would create a new student under teacher number 754

Relationships can get very complex however and it is not uncommon to see APIs nest 5 to 6 levels deep.

Once you have the primary key of a relationship then you shouldn't need to include relationship levels above the specific noun/object.

Meaning there should not be many cases where you have to make particular URL  more than 3 levels deep `/resource/identifier/resource`.

#### Limit complexity by using query strings

Many APIs have many different states that a resource can be in.

Simplify complexities by utilizing query string parameters.

*Any optional state that a particular resource can be in, can be conveyed by a particular query string.*

To get all students living in dorms perhaps you could use the following query:

`GET /students?dorms=true`

Keep the API intuitive by using query strings where appropriate.

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [Query String](./query-string.md) | [API Testing](./api-testing.md) →
