# Application model

The application is developed based on the Onion architecture and dependency injection technique.

According to Onion architecture pattern it's divided to application core, which includes domain, domain services and app services levels, and infrastructure level which contains services implementations.

This is described with UML class diagrams.

## Domain
![Domain class diagram](domain.png)

## Domain and domain services level

![Domain and domain services class diagram](domain-services.png)

## App services level

![App and domain services class diagram](app-domain_services.png)

## Infrastructure level

![Infrastructure level class diagram](infrastructure.png)

### Domain implementation

![Infrastructure domain implementation class diagram](infrastructure-domain.png)

### App services implementation

![Infrastructure app services implementation class diagram](infrastructure-app_services.png)