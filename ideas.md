# Corsair Text
An open world (universe), single player, single ship space adventure, in text.
## Goals
* Single player
* Single ship/fleet (single location)
* No internet needed
* Economic/Transportation/conflict simulation
* Abstracted conflict
* Trader/bounty hunter/thief/courier/miner/smuggler
* Economies evolve over time
* No instant communication (all remote data is historical)
* No boring flying/cut scenes
* Cool new ships to buy
* Cool toys to put on ships
## UI
* Text only
* Terminal scrolling
* Curses style?
* Text commands (with single letter shortcuts)
* Journal?
## Setting
* Generated universe
* Systems/Planets/Ports
* Other discoverable spots in Systems
* Asteroids for ore
* gas giants for hydrogen
* unofficial ‘ports’
* distances are time and fuel
* no full map
* (“look left” gets you a list of closest systems to your left?)
* Inside systems everything is just in a bag (no real distances)
* Arrival spot? For bad guys to camp and good guys to protect? Or as you get closer to the port does the danger increase?
### Ports
* market
* shipyard
  * ships
  * toys
* job board
* police
### Ports have attributes:
* Corporate
* Outlaw
* Religious
* Open/Closed
## Data
* All data is stored in json
* All aspects are data driven
## Commands
### Universal Commands
* (H)elp
* (L)ook
* (M)ap
* (I)nventory
* (G)o (destination)
### Area specific commands
#### System
* (F)ight
* (R)un
* (E)xplore
#### Market
* (B)uy # (product)
* (S)ell # (product)
#### Shipyard
* (B)uy (ship)
* (B)uy (toy)
* (S)ell (toy)
#### Asteroid
* (M)ine
#### Job Board
* (A)ccept (job)
## Economics
* prices will be very close (or identical?) within a system
* System prices will vary over time but not wildly
* distance allows greater differences (but does not require them)
* As the player travels to a system, the economy can be run since the last time visited
* The rest of the economies are frozen until the player gets back
* jobs/news affects economy?

