# Tests that Bring Little Value
![1675396842182](image/README/1675396842182.png)

![1675397006322](image/README/1675397006322.png)

# Test should be first
![1675397356418](image/README/1675397356418.png)

![1675398437788](image/README/1675398437788.png)

# BICEP Princples
*if you ask what if that is a boundary conditions
* check inverse relationships
  * verify something did not happen
![1675427633763](image/README/1675427633763.png)
* you should cross check web with api
[cross check with many api examples](courses\pluralsight\Fundamentals_of_Test_Automation_in_Java\test-automation-principles-in-java\src\test\java\test\principles\bicep\CrossCheckTestRepoCount.java)

* write test that force errors to occur
* dont catch with  java top level catch, shutdown gracefully and the integration test should see it does get a 500 response
* heartbeats are a way for the client to see if the server is running
* this can result in number overflow
something in the main class
```java
        int passengerNum = Integer.parseInt(args[0]);
```
* you can end up taking java course
* dont try to optimize something that is working fine now
* dont try to guess where the problem is

# Making Tests CORRECT
![1675428984916](image/README/1675428984916.png)
* stings, numbers, format should conform
* but peoples names dontt conform
![1675429728038](image/README/1675429728038.png)
![1675429828418](image/README/1675429828418.png)
![1675429869308](image/README/1675429869308.png)
* backend must verify data
* can I alter the senquence of events
* orde rhappens in time (async) and space

![1675430072236](image/README/1675430072236.png)
* what if the thing I need is not there
if there is 0 check against diviision
if there is division check against 0
![1675430346843](image/README/1675430346843.png)
* watch out for loops where the client repeadly tries to connect without doing something
![1675430489793](image/README/1675430489793.png)

# Avoiding Common Test Anti-Patterns

## Poor Test Name
*

## Clueless test
* descriptive name, and do and or or
* better to have small failed test
* such test overlap

## Complex Test
* Unstructured, too DRY,Unreadable, more difficult to read than necessary

* use libraries that use fluent interface
* if you see if else do flat test

[refer to refacotred test](courses\pluralsight\Fundamentals_of_Test_Automation_in_Java\test-automation-principles-in-java\src\test\java\test\principles\antipatterns)
* you want to in this case, test invalid, date, then invalid nanme, then invalid numbers


repeated data with DAMP
```java
    @DataProvider
    public static Object[][] invalidDateInput() {
        return new Object[][]{
                {"15/10-2019"},
                {"15-10/2019"},
                {"15102019"},
                {"15-1-2019"},
                {"15-10-20"}
        };
    }

    @Test(dataProvider = "invalidDateInput")
    void searchRejectsInvalidDateFormat(String date) {
        Assert.assertThrows(InvalidDateFormatException.class, () -> searchService.search("New York", "Atlanta", date, 1));
    }
```

