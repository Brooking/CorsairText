# Architecture
## Universe
The base component
* Holds all state
* Narrow top interface
* Static data
  * Universe Map
    * (Not fully generated apriori, but const once discovered)
  * Object library
    * Commodities
    * Ships
    * Toys
  * Name generation roots
  * Politics
* Dynamic data
  * Player data
  * Economics
  * News & Events
  * npc data
* All objects use dependency injection
## UI
### Text
The initial UI is console text only
* may contain translations of all Univers interfaces
### Curses basic
* Operates like Text, but with separate view and input areas
### Curses advanced
* Close to full graphics using ascii art
### Local HTTP
* Actually two pieces: the RESTful json server and the stateless viewer
