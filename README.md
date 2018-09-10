PaaS Internal CLI
=================

A CLI for operators of The Government PaaS.

Written in Golang, initially as a learning experience.

Planned interface
-----------------

Generally:

```
paas $noun $verb ...flags
```

Examples:

```
paas bosh-cli open dev # Opens a bosh-cli pointing at the dev environment
paas docs open
paas manual open
paas status open
paas concourse open dev towers
```

Ultimately this should take over all of the functionality of the Makefiles in
https://github.com/alphagov/paas-cf/blob/master/Makefile and
https://github.com/alphagov/paas-bootstrap/blob/master/Makefile
