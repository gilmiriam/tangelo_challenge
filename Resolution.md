# Challenge Resolution
## Decisions
1. The first decision was what kind of API I use. This was an API REST because don't need to install additional libraries or software for make the petitions. It's easy to scale. You can choose in a easy what method to use and the endpoint for your application.

2. I choose to pass the slice via the request query, because thats a simple way to make them. This give and extra incovenient, the slice that the user give us only can be a string.

3. With a string value I have to change the way that I do the flat method. Instead of use a algorithm or something similar, I use the library strings for treat the string input and passing them into a flatten slice.

4. Other important decision was the way I do the test suite. I decided to make to type of test. One of the for test the input considering different scenarios like the user input a slice with words. The other test is for test the all implementation, calling the API and see the response.

5. Dockerize the project was another decision that makes more easy to develop, because you cam test it with only have docker and make installed.

## Pains
1. One of the difficulties was that I start developing a complex way to flat the slice, after receive the input from API. 