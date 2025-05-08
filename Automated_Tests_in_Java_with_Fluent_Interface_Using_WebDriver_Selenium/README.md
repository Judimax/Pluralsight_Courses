# Applying the Page Object Pattern
* if you have webdriver api in test methods you are doing it wrong
s
[use this](courses\pluralsight\Automated_Tests_in_Java_with_Fluent_Interface_Using_WebDriver_Selenium\automated-tests-in-java-with-fluent-interface-using-webdriver-selenium-exercise-files)
to see the page object model in the main

![1675459676980](image/README/1675459676980.png)


## Advanced  Fluent Interface
*
there is ActController and Veryify that what we do in UI testing
* use static factory classes for encapsulation

```java
        return new SearchPage(new SearchActController(),
                              new SearchVerifyController(),
                              new SearchGetController());
```

* so now if you want to act you have to go throgh act if you want to verify youhave to go through verify
```java
        search.act()
                .filterBySkillLevel(SkillLevel.BEGINNER)
                .filterByRole(Role.SOFTWARE_DEVELOPMENT)
                .selectTab(Tab.COURSES)
                .selectCourse("Java Fundamentals: The Java Language");


        course.verify()
               .coursePreviewIsDisplayed()
               .freeTrialIsDisplayed();
```

## Using AssertJ
* complex verifications we can delegate to a 3rd party verify
* with getcontroller you can do complex get

```java

public class SearchGetController {

    WebDriver driver = getChromeDriver();


    public List<String> courses(){
        List<WebElement> courses = driver.findElements(
                By.xpath("//div[@id='search-results-category-target']//div[@class='search-result__title']"));

        assertTrue(courses.size() != 0, "List is empty, failed to collect courses");

        return courses.stream()
                      .map(WebElement::getText)
                      .collect(toList());

    }


}

```

```java
        assertThat(search.get().courses())
                .hasSize(25)
                .containsOnlyOnce("Java Fundamentals: The Java Language")
                .doesNotContain("Python");
```
