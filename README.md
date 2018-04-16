# Matomat - ReWrite
The old matomat code needs to be rewritten.

## Goals & Implementation strategy
- Strive for "MVP" and add more functionality later.

## Architectural requirements
- Split the former monolithic concept into a service and a (g)ui layer.
- Code should be modular enough to swap out and / or add new components (e.g. different auth methods, different ways to "consume" a drink)
- Split authentication from authorization to allow various authentication mechanisms.

## Overall functional requirements
The new system should at least provide the functionality listed below.
(At the current point in time no inventorization of drinks is planned. So no need to  e.g. check if enough drinks are available when performing an "consumption" of a drink)

### authentication and authorization
- login
    - username / password
    - rfid
- logout
- have at least two user roles: admin and regular user

### user management
- user create (by admin)
- user delete (by admin or the user)
- user edit (by user)
- password change (by user)
- password reset (by admin)

### credit management
- add credit (for and by user)
- use credit (by user)
- show credit (for user)
- [transfer credit to other user (by user)]

### drink inventory management
- add drink(s)
- remove drink(s)
- show available drinks

### (drink) consumption
- consume drink

### stats
- track consumed drinks
- show consumed drinks
- [send consumption / consumption highlights to mqtt?]
- show current total credits in system

## General requirements of GUI MVP
- Needs to run in "text only" mode on a terminal