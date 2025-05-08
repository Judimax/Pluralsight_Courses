# Getting Started with Web and API testing

![1675453366168](image/README/1675453366168.png)

## Writing Basic GET Tests for Response Headers

* helping files go in  main
[API test examples](courses\pluralsight\Getting_started_with_Web_API_Test_Automation_in_Java\module4)
* you need to close and open connect, http client allows by default only 2 concurrent connection
* Response Header Example
*courses\pluralsight\Getting_started_with_Web_API_Test_Automation_in_Java\module5\src\main\java\ResponseUtils.java
* courses\pluralsight\Getting_started_with_Web_API_Test_Automation_in_Java\module5\src\test\java\ResponseHeaders.java
* write headers for all the headers you need

# Writing Advanced GET Tests for Response Payload
* convert json to a structure
[JSON to Map](courses\pluralsight\Getting_started_with_Web_API_Test_Automation_in_Java\module5\src\test\java\BodyTestWithSimpleMap.java)
```java
        HttpGet get = new HttpGet(BASE_ENDPOINT + "/users/andrejss88");

        response = client.execute(get);

        String jsonBody = EntityUtils.toString(response.getEntity());

        JSONObject jsonObject = new JSONObject(jsonBody);

        String loginValue = (String) getValueFor(jsonObject, LOGIN);

        Assert.assertEquals(loginValue, "andrejss88");
```

## (Un)Marshalling
*
![1675455909776](image/README/1675455909776.png)
* you dont want to write out all the properties on an object so use this to avoid failure
```java
                .configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false)
```
```java

    public static <T> T unmarshallGeneric(CloseableHttpResponse response, Class<T> clazz) throws IOException {

        String jsonBody = EntityUtils.toString(response.getEntity());

        return new ObjectMapper()
                .configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false)
                .readValue(jsonBody, clazz);
    }
```
## (Un)Marshalling Nested JSON
* @JsonProperty("resources") means check for top level resources object
```java
package entities;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.Map;

public class RateLimit {

    // private int coreLimit;

    // private String searchLimit;

    // public int getCoreLimit() {
    //     return coreLimit;
    // }

    // public String getSearchLimit() {
    //     return searchLimit;
    // }
    @SuppressWarnings("unchecked")
    @JsonProperty("resources")
    private void unmarshallNested(Map<String, Object> resources){
        Map<String, Integer> core = (Map<String, Integer>) resources.get("core");
        coreLimit = core.get("limit");

        Map<String, String> search = (Map<String, String>) resources.get("search");
        searchLimit = String.valueOf(search.get("limit"));
    }
}
```
## JSON to schema to POJO
use jsonschema2pojo.org
or use the example in the [pom.xml](courses\pluralsight\Getting_started_with_Web_API_Test_Automation_in_Java\module5\pom.xml) and have  maven build it for you

## Writing Test for Options
```java
    @Test
    public void optionsReturnsCorrectMethodsList() throws IOException {

        String header = "Access-Control-Allow-Methods";
        String expectedReply = "GET, POST, PATCH, PUT, DELETE";

        HttpOptions request = new HttpOptions(BASE_ENDPOINT);
        response = client.execute(request);

        String actualValue = ResponseUtils.getHeader(response, header);

        Assert.assertEquals(actualValue, expectedReply);
    }
```

* adding request headers (delete)
```java
    @Test(description = "This test will work only if you set a valid Token in the Credentials Class")
    public void deleteIsSuccessful() throws IOException {

        HttpDelete request = new HttpDelete("https://api.github.com/repos/andrejss88/deleteme");

        request.setHeader(HttpHeaders.AUTHORIZATION, "token " + Credentials.TOKEN);
        response = client.execute(request);

        int actualStatusCode = response.getStatusLine().getStatusCode();

        Assert.assertEquals(actualStatusCode, 204);
    }

```

## Post Request
* [other http methods test](courses\pluralsight\Getting_started_with_Web_API_Test_Automation_in_Java\module6\src\test\java\DeleteAndPost.java)
```java

    // Instead, use this POST test using the Web Token
    @Test
    public void createRepoReturns201WebToken() throws IOException {

        // Create an HttpPost with a valid Endpoint
        HttpPost request = new HttpPost("https://api.github.com/user/repos");

        request.setHeader(HttpHeaders.AUTHORIZATION, "token " + Credentials.TOKEN);

        // Define Json to Post and set as Entity
        String json = "{\"name\": \"deleteme\"}";
        request.setEntity(new StringEntity(json, ContentType.APPLICATION_JSON));

        // Send
        response = client.execute(request);

        int actualStatusCode = response.getStatusLine().getStatusCode();
        Assert.assertEquals(actualStatusCode, 201);
    }
```
* always encode over https
```java
        String auth = Credentials.EMAIL + ":" + Credentials.PASSWORD;
        byte[] encodedAuth = Base64.encodeBase64(auth.getBytes(StandardCharsets.ISO_8859_1));
```
