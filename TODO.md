# TODO list

## Application

- use cobra for commands
- introduce viper for configuration
  - create modular configuration system that can be passed to server, sandbox, checker, etc.
- introduce servers
  - common interfaces exposed by all servers including following functionality
    - reading stdout / stderr
    - sending signals
    - handling config templates
  - the implementation will be following
    - fpm
    - nginx
    - apache
    - caddy
- introduce sandbox
  - create `Sandbox` interface
  - create local sandbox to execute
- introduce more generic checker to replace instance
  - it should be modular and server / sandbox independent
    - it should just use some common interface exposed by servers / sandboxes
  - expectations should be configurable and not defined in the code
  - output matching should support patterns and there should be some operators to count number of matches
  
## CI

- Set up GitHub actions for running tests and linting 