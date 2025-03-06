# cognizant

-----------------------
# ASSIGNMENT checklist
1. Use persistent storage (e.g., Postgres database or physical file). (done)
2. Implement logging for API requests and responses.
(done)
3. Add validation for each API (e.g., missing or invalid input).
(done)
4. Write tests for coverage and regression.
(done)
5. Return appropriate HTTP status codes for success and error scenarios.
(done)
6. Write at least one unit test for each endpoint using Go’s net/http/httptest package.
(done)
7. The Library should have at least the following objects. You may add in any new objects or fields as you deemed necessary to complete the program
(done)

BookDetail
1. Title (string): Unique identifier for the book.
2. AvailableCopies (int): No of available copies of the book that can be loaned. 
(done)

LoanDetail
1. NameOfBorrower (string): Name of borrower.
2. LoanDate (date): Date where the book was borrowed.
3. ReturnDate (date): Date where the book should be returned.
(TBC)

8. Expose the following RESTful endpoints:
o GET /Book to retrieve the detail and available copies of a book title.
(done)

-----------------------
# IMPORTANT ----

function extensively and well tested, DB initialized and working

connection via localhost:3000, return and log errors extensively

unit tests were completed extensively, 100% coverage.

dao, service, transport layers were tested extensively

assessment for cognizant
using net http for routing, simple and easy to use, best suited for this assessment

-- later on added gorillamux for easier routing to specific HTTP methods (GET, POST, etc.),
e.g POST book, GET book, able to use the same routes with different specifications

created and tested first API, 
server runs on localhost:3000

using datadog mocksql to mock several behaviours throughout the layers
for easier testing

TO WORK ON:
db connection pooling in order to handle larger userbase(which my project wont have XD)

authentication for connection 

other crud requests
