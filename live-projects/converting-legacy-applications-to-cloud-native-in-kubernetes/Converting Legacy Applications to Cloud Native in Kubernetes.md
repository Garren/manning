# Converting Legacy Applications to Cloud Native in Kubernetes

## 1.0 Initial Setup

* Java - installed with homebrew 14.0.1

* Maven - installed with homebrew `mvn`

* Docker...previously installed

* MySQL --- using docker`https://towardsdatascience.com/connect-to-mysql-running-in-docker-container-from-a-local-machine-6d996c574e55`

  ```shell
  docker volume create mysql-volume
  
  docker run --name=mysql-clacnk \
  -p3306:3306 \
  -v mysql-volume:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=s3cr3t2020 \
  -d mysql/mysql-server:8.0.20 
  
  ```

* Minikube...installed with docker

* Tomcat - installed with homebrew

  ```shell
  To have launchd start tomcat now and restart at login:
    brew services start tomcat
  Or, if you don't want/need a background service you can just run:
    catalina run
  ```

  

## 1.1 Running the legacy application

1. install apache tomcat - done
2. install mysql, create database and a user that can read/write for thet db...

```shell
$ docker exec -it mysql-clacnk bash
bash-4.2# mysql -u root -p
Enter password:
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 12
Server version: 8.0.20 MySQL Community Server - GPL

Copyright (c) 2000, 2020, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> create database profiles;
Query OK, 1 row affected (0.00 sec)

mysql> CREATE USER empuser@'localhost' IDENTIFIED BY 'password';
Query OK, 0 rows affected (0.01 sec)

mysql> CREATE USER empuser@'%' IDENTIFIED BY 'password';
Query OK, 0 rows affected (0.00 sec)

mysql> ALTER USER 'empuser'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';
Query OK, 0 rows affected (0.00 sec)

mysql> ALTER USER 'empuser'@'%' IDENTIFIED WITH mysql_native_password BY 'password';
Query OK, 0 rows affected (0.01 sec)

mysql> GRANT ALL PRIVILEGES ON profiles.* to empuser@'localhost';
Query OK, 0 rows affected, 1 warning (0.00 sec)

mysql> GRANT ALL PRIVILEGES ON profiles.* to empuser@'%';
Query OK, 0 rows affected (0.00 sec)

```

3. reconfigure the local filesystem....? 
   * apparently points to /tmp which should work for OS X
4. Reconfigure the mysql url and username in the webapp configuration
   * uh....looks good?
5. Deploy the application to tomcat
   1. `https://github.com/grafpoo/MigratingToK8s/tree/master/starter-webapp`
   2. `mvn compile && mvn package`
   3. Open local tomcat, access management Gui, add rolename and username as described in 401 error ("manager-gui" and "tomcat"). Restart tomcat `catalina stop && catalina configtest && catalina start`
   4. Navigate to manager gui, hit deploy war, select profiles.war 
6. http://localhost:8080/profiles is accessible
7. http://localhost:8080/profiles/profiles/garren created



## 1.2 "'Bootify" the application using Spring Boot

1. We're not using the "default package", we have an explicit "liveproject.m2k8s" package
2. ... fucking 2-3 hours of dicking with pom.xml file ...
3. Profit.

This required a little bit to get up and running. Quite a few of the dependencies specified in pom.xml just go away or get replaced by a `org.springframework.boot` equivalent. However, not being totally familiar with maven, I spent a lot of time getting a pom that finally worked.

```xml
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/maven-v4_0_0.xsd">
    <parent>
      <groupId>org.springframework.boot</groupId>
      <artifactId>spring-boot-starter-parent</artifactId>
      <version>2.3.1.RELEASE</version>       
    </parent>
    <modelVersion>4.0.0</modelVersion>
    <groupId>com.manning.pl</groupId>
    <artifactId>profile-mvc</artifactId>
    <version>1.0</version>
    <name>Profile Management</name>
    <description>Profile management</description>
    <packaging>jar</packaging>
    <properties>
        <springVersion>4.0.7.RELEASE</springVersion>
        <hibernateVersion>4.1.6.Final</hibernateVersion>
        <hibernateEntityManagerVersion>4.0.1.Final</hibernateEntityManagerVersion>
        <thymeleafVersion>2.1.3.RELEASE</thymeleafVersion>
        <h2-version>1.4.182</h2-version>
        <hibernateValidatorVersion>5.0.1.Final</hibernateValidatorVersion>
        <commonsLangVersion>3.1</commonsLangVersion>
        <junitVersion>4.11</junitVersion>
        <mockitoVersion>1.9.5</mockitoVersion>
        <hamcrestVersion>1.3</hamcrestVersion>
    </properties>

    <repositories>
        <repository>
            <id>spring-release</id>
            <url>http://maven.springframework.org/release</url>
        </repository>
        <repository>
            <id>spring-milestone</id>
            <url>http://maven.springframework.org/milestone</url>
        </repository>
        <repository>
            <id>spring-snnapshot</id>
            <url>http://maven.springframework.org/snapshot</url>
        </repository>
        <repository>
            <id>maven-2</id>
            <url>http://download.java.net/maven/</url>
        </repository>
    </repositories>
    <dependencies>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
            <optional>true</optional>
        </dependency>

        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
        </dependency>

        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-thymeleaf</artifactId>
        </dependency>

        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-data-jpa</artifactId>
        </dependency>

        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-jdbc</artifactId>
        </dependency>

        <dependency>
            <groupId>commons-fileupload</groupId>
            <artifactId>commons-fileupload</artifactId>
            <version>1.2</version>
        </dependency>
        <dependency>
            <groupId>commons-io</groupId>
            <artifactId>commons-io</artifactId>
            <version>2.5</version>
        </dependency>

        <dependency>
            <groupId>org.projectlombok</groupId>
            <artifactId>lombok</artifactId>
            <version>1.18.10</version>
            <optional>true</optional>
        </dependency>
        <dependency>
            <groupId>org.hibernate</groupId>
            <artifactId>hibernate-validator</artifactId>
            <version>${hibernateValidatorVersion}</version>
        </dependency>
        <dependency>
            <groupId>org.apache.commons</groupId>
            <artifactId>commons-lang3</artifactId>
            <version>3.1</version>
        </dependency>

        <!-- test scope -->
        <dependency>
            <groupId>junit</groupId>
            <artifactId>junit</artifactId>
            <version>${junitVersion}</version>
            <scope>test</scope>
        </dependency>
        <dependency>
            <groupId>org.springframework</groupId>
            <artifactId>spring-test</artifactId>
            <version>${springVersion}</version>
        </dependency>
        <dependency>
            <groupId>org.mockito</groupId>
            <artifactId>mockito-core</artifactId>
            <version>${mockitoVersion}</version>
        </dependency>
        <dependency>
            <groupId>org.hamcrest</groupId>
            <artifactId>hamcrest-library</artifactId>
            <version>${hamcrestVersion}</version>
        </dependency>

        <dependency>
            <groupId>mysql</groupId>
            <artifactId>mysql-connector-java</artifactId>
            <version>8.0.11</version>
        </dependency>
    </dependencies>
    <build>
        <finalName>profiles</finalName>
        <plugins>
          <plugin>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-maven-plugin</artifactId>
          </plugin>
        </plugins>
    </build>
</project>

```

`Application.java` changed to:

```java
package liveproject.m2k8s;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class Application {
  public static void main(String[] args) {
    SpringApplication.run(Application.class, args);
  }
}

```

`WebConfig.java` changed to:

```java
package liveproject.m2k8s.web;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.ViewResolver;
import org.springframework.web.servlet.config.annotation.DefaultServletHandlerConfigurer;
import org.springframework.web.servlet.config.annotation.EnableWebMvc;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurerAdapter;
import org.springframework.web.servlet.config.annotation.ResourceHandlerRegistry;
import org.springframework.context.ApplicationContext;
import org.springframework.context.ApplicationContextAware;
import org.thymeleaf.spring5.SpringTemplateEngine;
import org.thymeleaf.spring5.view.ThymeleafViewResolver;
import org.thymeleaf.spring5.templateresolver.SpringResourceTemplateResolver;

@Configuration
@EnableWebMvc
@ComponentScan("liveproject.m2k8s")
public class WebConfig extends WebMvcConfigurerAdapter implements ApplicationContextAware {

  private ApplicationContext applicationContext;

  @Override
  public void setApplicationContext(ApplicationContext applicationContext) {
    this.applicationContext = applicationContext;
  }

  @Override
  public void addResourceHandlers(ResourceHandlerRegistry registry) {
    registry
      .addResourceHandler("/resources/**")
      .addResourceLocations("/resources/");
  }

  @Bean
  public ViewResolver viewResolver() {
    ThymeleafViewResolver viewResolver = new ThymeleafViewResolver();
    viewResolver.setTemplateEngine(templateEngine());
    return viewResolver;
  }
  @Bean
  public SpringTemplateEngine templateEngine() {
    SpringTemplateEngine templateEngine = new SpringTemplateEngine();
    templateEngine.setTemplateResolver(templateResolver());
    return templateEngine;
  }

  @Bean
  public SpringResourceTemplateResolver templateResolver() {
    SpringResourceTemplateResolver templateResolver = new SpringResourceTemplateResolver();
    templateResolver.setApplicationContext(this.applicationContext);
    templateResolver.setPrefix("/WEB-INF/views/");
    templateResolver.setSuffix(".html");
    templateResolver.setTemplateMode("HTML5");
    return templateResolver;
  }
}

```

The `SpringResoufceTemplateResolver` and the `addResourceHandlers` override where tricky. Again, I know nothing about Spring or Spring Boot. The later is necessary to serve static content (i.e., style.css)

Lastly, I had to change the `@GeneratedValue` annotation in `Profile.java` to `@GeneratedValue(strategy = GenerationType.IDENTITY)` to get around a 'missing hibernate_sequence' table error when registering new users. This wasn't an issue /before/ spring boot, so I'm at a loss as to why it's necessary now.

With that, running `mvn clean && mvn dependency:tree && mvn compile && mvn package && mvn spring-boot:run` loads the app. I can register new users, access previous users, and upload new images.