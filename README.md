# causal
Server cause and effect

## Currently in development.

Causal sits on a node and runs watchers every x seconds. These can be anything - disk space used, size of files, folders, memory free
and on and on. When these watchers fail they fire an alert. 

These alerts can be anything as well! Email, hipchat, slack, graphite and on and on.

Causal is written in Go!

## TODO  
In no order :)

* Sort out the config 
* Create a scheduler that controls when all the watchers run
* Increase test coverage
* Write more watchers and alerters!
* Implements a RPC watcher and alerter for chaining nodes
* Create a nice commander
* Examples!