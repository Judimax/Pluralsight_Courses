* you can do record and playback and selenuium will generate a script for you
* python c# java, are the most used
* there is selenium ide, browser extension
* web driver for common test cases
* selenium grid, sort of a kubernetes for unit testing
  * test on multiple browsers and operating systems
  * grid run tests in parallel
  * all of its is free
* no warranty or liablity(dont expect support)

* there is applitools,katalon studio
* use a cloud provider to run tests in parallel

![1675351914370](image/README/1675351914370.png)
* you need many unit test and few e2e tests
* JMeter is a performance tool

![1675352358880](image/README/1675352358880.png)

## Demo Writing A Basic Script
[Home Page](Selenium_4_the_big_picture\exercise_files\Selenium-The-Big-Picture-demo-code\src\web\home.html)
* need selennium webdriver lib,testNG and download chromedriver
* spend time writing elements on UI scripts
* Assert.assertTrue is from TestNG


## Finding Elements
![1675384985225](image/README/1675384985225.png)
* put id on everything makes automation soo much simpler

* all things on a page will be access using web element
* you have to know which elements have certain methods
![1675385663561](image/README/1675385663561.png)
*and properties
![1675385713105](image/README/1675385713105.png)

* radio buttons are hard to select
[radiobtn test](courses\pluralsight\Automated_Web_Testing_with_Selenium_and_WebDriver_4_Using_Java\05\demos\m5\after\src\main\java\com\pluralsight\WebDriverRadioButtons.java)
* really means search by red
```java
By.name
```

* [checkboxes example](courses\pluralsight\Automated_Web_Testing_with_Selenium_and_WebDriver_4_Using_Java\05\demos\m5\after\src\main\java\com\pluralsight\WebDriverCheckboxes.java)


* [dropdown example](courses\pluralsight\Automated_Web_Testing_with_Selenium_and_WebDriver_4_Using_Java\05\demos\m5\after\src\main\java\com\pluralsight\WebDriverSelectItem.java)

* select jar makes things easier
(wrapper pattern  )
```java
import org.openqa.selenium.support.ui.Select;

        WebElement selectElement = driver.findElement(By.id("select1"));
        Select select = new Select(selectElement);
        select.selectByVisibleText("Maybe");
```

* [table example](courses\pluralsight\Automated_Web_Testing_with_Selenium_and_WebDriver_4_Using_Java\05\demos\m5\after\src\main\java\com\pluralsight\WebDriverTables.java)
* xpath may be the best way to access a row on the table so they say
```java
WebElement outerTable = driver.findElement(By.tagName("table"));
WebElement innerTable = outerTable.findElement(By.tagName("table"));
WebElement row = innerTable.findElements(By.tagName("td")).get(1);
```
or
```java
WebElement row = driver.findElement(By.xpath("/html/body/table/tbody/tr/td[2]/table/tbody/tr[2]/td"));
System.out.println(row.getText());
driver.close();
```
*  wait pings the DOM every 500 ms to see if the item is available

* explicit wait is asynchronoucs

```java
WebDriverWait wait = new WebDriverWait(driver, Duration.ofSeconds(10));
wait.until(ExpectedConditions.presenceOfElementLocated(By.linkText("Images")));
```

## Selenium Server
* remotely execute tests
* setup selenium server
* selenuum server grid mode has a hub orchersating everyhting, like kubernetes


* download [selenium grid](https://www.selenium.dev/downloads/) then run
```ps1
 java -jar .\selenium-server-4.8.0.jar standalone
```

* now to run remotely
```java
        ChromeOptions chromeOptions = new ChromeOptions();
        chromeOptions.addArguments("start-maximized");
        WebDriver driver = new RemoteWebDriver(new URL("http://localhost:4444/wd/hub"), chromeOptions);
```

* look at the console output make sure there is nothing weird like a stacktrace
* selenium is using hub to run in parallel
to do this

use [url to hub]/grid/console to see the ui


```ps1
# for a new hub
 java -jar .\selenium-server-4.8.0.jar  hub

#  to  add additional nodes
  java -jar .\selenium-server-4.8.0.jar  node --hub [url to hub]/grid/register --port [Subeequent port]


```

head to the [grid dashboard](http://192.168.1.154:4444/grid/console)
* can run in detached,headless mode

[test example](courses\pluralsight\Automated_Web_Testing_with_Selenium_and_WebDriver_4_Using_Java\06\demos\m6\after\src\main\java\com\pluralsight\WebDriverDemo.java)

## Parallel Processing
* grid makes it possible you have to initiate


# Building A Framework


## Page Object Model
object shold reflect your page
several pages could have shopping cart but there needs to have common methods it can use
![1675395123846](image/README/1675395123846.png)
