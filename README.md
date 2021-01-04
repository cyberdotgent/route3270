# Route/3270

A simple TN3270 router that can be used to route access to several applications or machines behind a single gateway

## Features

* Crude access control system for proxied services
* Authentication

### Screenshots

![](doc/login.png)
_login screen_

![](doc/selection.png)
_service selection screen_

## Usage

    route3270 -c example.toml
    
Review the `example.toml` file for usage instructions.

## Current limitations

* Each user can only have up to 14 services to chose from.
* Does not support SSL (both server and client)