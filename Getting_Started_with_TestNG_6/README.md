* there is already a solution for many problems you can fast in testng
![1675436935620](image/README/1675436935620.png)

# Annotations and Assertions
* Test annotation makes a test a test
* always declare test methods public

*  Assert is where all assertion methods live
* testNG takes additional string message for all its methods
```java
assertEquals,

assertSame() 2 reference variables point to the same object
assertEqualsDeep() - checks for each key value pair within maps
assertEqualsNoOrder() takes only arrays and checks the conents are equal but doesnt check order
```

## Hard vs Soft Assertions
* test should have 1 reason to fail
* soft assertsions allow your tests to run till the end
```java
        // Arrange
        CloseableHttpClient client = HttpClientBuilder.create().build();
        SoftAssert sa = new SoftAssert();

        // Act
        CloseableHttpResponse response = client.execute(new HttpGet("https://api.github.com/"));


        // Assert
        System.out.println("First assert");
        sa.assertEquals(response.getStatusLine().getStatusCode(), 404);

        System.out.println("Second assert");
        sa.assertEquals(getOrDefault(response.getEntity()).getMimeType(), "application/xml");

        System.out.println("Third assert");
        sa.assertEquals(getOrDefault(response.getEntity()).getCharset().toString(), "UTF-8");


        client.close();
        response.close();

        sa.assertAll();
```
* make sure test start with clean sheet and clean state

##  Befroe and after tests
BeforeMethod,AfterMethod, BeforeClass,AfterClass

## Inheritied Setup
[unitetestbaseclass](courses\pluralsight\Getting_Started_with_TestNG_6\03\demos\m3-demo\src\test\java\baseclasses\UnitTestBaseClass.java)
* here you use BeforeSuite and AfterSuite
but there global before test
* global before method runs before local beforeMethod

## The SkipException
* you want to use Skip exception like a softAssert where for example, if status is not equal to 200 stop checking the rest of the headers
__FILE__ -WebServiceBaseClass
```java
    @BeforeClass
    public void setup() throws IOException {
      ...
        if(actualStatusCode != 200){
            throw new SkipException("Basic criteria failed," +
                    "was expecting code 200, but got: " + actualStatusCode);
        }
    }
```
so  now if status code is not equal to 200, skip all tests
* can use anywhere

# Getting More Functionality out of Annotations

## when you expcet an exception
* if this is supposed throw exception it will fail
* also has error mathching
expectedExceptions,expectedExceptionsMessageRegExp
```java
    @Test(expectedExceptions = DuplicateUserException.class,
            expectedExceptionsMessageRegExp = ".*already exists")
    public void addDuplicateThrowsException() throws DuplicateUserException {
        // Act
        um.addUser("same@email.com");
        um.addUser("same@email.com");
    }

```

## timeout
* a test cannot take over .5 seconds
* not for performance test
[TestTimeout](courses\pluralsight\Getting_Started_with_TestNG_6\04\demos\m4-demo\src\test\java\api\TestTimeout.java)
```java
    @BeforeMethod(timeOut = 800) // enough time - will pass
    public void setup(){
        // Arrange
        client = HttpClientBuilder.create().build();

    }
    @Test(timeOut = 500) // not enough time - will fail
    public void testIsTooSlow() throws IOException {

        // Act
        CloseableHttpResponse response = client.execute(new HttpGet("https://api.github.com/"));

        // Assert
        Assert.assertEquals(response.getStatusLine().getStatusCode(), 200);

    }
```

## Disabled
* for flaky tests
* for classes use @Ignore - global runs
* for classes @Test(enabled= false) - nothing runs
```java
@Ignore
public class DisabledTest extends UnitTestBaseClass {



    @Test
    public void unstableTest1(){
        System.out.println("Test 1");
    }

    @Test(enabled=false)
    public void unstableTest2(){
        System.out.println("Test 2");
    }
}

```

## Always Run
* when test fail and browser windows dont close use this
[e2e test](courses\pluralsight\Getting_Started_with_TestNG_6\04\demos\m4-demo\src\test\java\ui\UiPluralsightTest.java)
```java
    @AfterMethod(alwaysRun = true)
    public void closeBrowser(){
        System.out.println("Closing down the browser");
        driver.close();
    }
```

## Other fields
[Documentation Reference](testng.org/doc/documentation-main.html)

# Levarigng the Data Providers
* 3 values in data provider means test run 3 times

* commons-validator is a mvn package that gives us email validator


here to make our dataProviders generic we use dataProviderClass
```java
    @Test(dataProvider = "endpointsRequiringAuthentication", dataProviderClass = CommonApiDataProviders.class)
    public void userEndpointReturns401(String endpoint) throws IOException {
        // Act
        response = client.execute(new HttpGet("https://api.github.com/" + endpoint));
        int actualStatusCode = response.getStatusLine().getStatusCode();

        // Assert
        Assert.assertEquals(actualStatusCode, 401);
    }
```

[apitest](courses\pluralsight\Getting_Started_with_TestNG_6\05\demos\m5-demo\src\test\java\api\ApiTestWithDataProviders.java)


* multiple values with a single data provider
```java

    @DataProvider
    private Object[][] invalidCredentials(){
        return new Object[][] {
                // 1st param , 2nd param
                {"let_me_in", ""                             },
                {    ""     , "safest_password_in_the_world" },
        };
    }

    @Test(dataProvider = "invalidCredentials")
    public void emptyPasswordFailsLogin(String login, String password){

    }
```

## Different Input Output

* you can have many different output fo rinput as well
![1675447435315](image/README/1675447435315.png)

![1675447552222](image/README/1675447552222.png)
* be carefull
```java
@DataProvider
Object[][] twoInputProvider() {
  return new Object[][] {
    {"", "email must not be empty"}
    {"joe", "email must have an @ sign"}
    {"joe@email", "email must have one . after @ sign"}
  };
}
```

## Refactoring DataProviders
* use composition
![1675448082065](image/README/1675448082065.png)
* Kiran question, if we have imported data providers do we keep them under test, or remove scope ( if gradle has such a thing)
* testNg uses java reflection so some errors are not soo obvious
* you get data provider mismatch that means you data providers,and method params dont match up

# Organizing your tests
* test should run in any order any # of times and produce same result

* sometimes test have business importance
![1675448511788](image/README/1675448511788.png)
![1675448557946](image/README/1675448557946.png)

## XML Suite

[like a test package.json](courses\pluralsight\Getting_Started_with_TestNG_6\06\demos\m6-demo\src\test\resources\tests.xml)
* use packages over classes
```xml
  <packages>
      <package name="unit.*"/>
  </packages>

  <classes>
      <class name="unit.FirstTestNGTest"></class>
  </classes>
```
* use sepearte files if you have ci
* order tests like this
![1675449553134](image/README/1675449553134.png)
[tests.xml](courses\pluralsight\Getting_Started_with_TestNG_6\06\demos\m6-demo\src\test\resources\tests.xml)

## Group and Depends On
* Sanity tests

```java

@Test(groups="")
public void sanityTest1(){}


@Test(dependsOnGroups="·.")
public void dependentTest(){}
```

* but if you move sainty test to sanity class testNG throws error because it cant figure out it moved to another class
* run the test from tests.xml
* if you move sanity out of the unitpackage, testNg will fail move it back to get tests to pass

[Sanity class](courses\pluralsight\Getting_Started_with_TestNG_6\06\demos\m6-demo\src\test\java\unit\sanity\UnitSanityTests.java)

[Class that depends on Sanity](courses\pluralsight\Getting_Started_with_TestNG_6\06\demos\m6-demo\src\test\java\unit\DisabledTest.java)

## Test Setup vs Sanity Test
![1675450860876](image/README/1675450860876.png)

## Priority
it goes by numnbers you make want to make this an enum

## Questions
* are we using magic numbers with priority
* are we using fluent language

## Features to Explore
* there is sucessPercentage field,
* there is @BeforeTest (diff from BeforeMethod), @Parameters @Factory(like dataProviders)
* Reports ( need to integrate with reporting tool)
* ITestListener, allows you to really take control of your test
