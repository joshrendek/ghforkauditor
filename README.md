## What is it?

Given an org and a oauth token, this will print a table showing the repository name, URL, # of forks, and users who have forked the repository.

Flags:

```
  -oauth string
        oauth token to use
  -org string
        org name
```

Example output:

```
+------------------------------------+-------------------------------------------------------------------------+-------+
| Repository                | URL                                                   | Forks | Users                    |                                                                                        |
+------------------------------------+-------------------------------------------------------------------------+-------+
| SomeProject               | https://github.com/your_org/SomeProject               | 2     | UserA, UserB             |
+------------------------------------+-------------------------------------------------------------------------+-------+
```
