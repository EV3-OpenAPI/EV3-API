![ZHAW-logo](https://upload.wikimedia.org/wikipedia/commons/thumb/e/e6/ZHAW_Logo.svg/206px-ZHAW_Logo.svg.png)


# Lego Robot Mindstorm EV3-REST

This project deals with the goal to design the Lego Robot of the type mindstorm EV3 for the promotion of programming knowledge for non-informaticians, in such a way that the implemented methods in the server are callable over several programming technologies. The implemented methods should be callable via an application programming interface.
The server which is the Lego robot is implemented with [OpenAPI Specification](https://www.openapis.org/), which allows to automatically generate libraries for the client in several programming languages.

This Github project is the starting point for the Lego Robot Mindstorm EV3-RESt. Here you will find the information you need about the lego robot, simple examples of what it looks like, and some general information regarding the project.

# Current Version 1.0

TODO

# Tools and Libraries

TODO

# Participation

The current process for development of the Lego Robot is described in [Development Guidelines](https://github.com/PA-arslasel-machitic/EV3-REST/blob/main/DEVELOPMENT.md). Development of the next version of the Lego Robot is guided by the [ZHAW School of Engineering](https://www.zhaw.ch/en/engineering/). This group of committers bring their API expertise, incorporate feedback from the community, and expand the group of committers as appropriate. All development activity on the future specification will be performed as features and merged into this branch. Upon release of the future specification, this branch will be merged to <code>main</code>.

The ZHAW holds every 2 weeks web conferences to review to code and open pull requests and discuss open issues related to the evolving lego robot. 

The lego robot EV3 encourages participation from individuals and companies alike. If you want to participate in the evolution of the lego robot, consider taking the following actions:

* Review the current specification. The human-readable markdown file is the source of truth for the specification.
* Review the [development](https://github.com/PA-arslasel-machitic/EV3-REST/blob/main/DEVELOPMENT.md) process so you understand how the spec is evolving.
* Check the [issues](https://github.com/PA-arslasel-machitic/EV3-REST/issues) and [pull requests](https://github.com/PA-arslasel-machitic/EV3-REST/pulls) to see if someone has already documented your idea or feedback on the specification. You can follow an existing conversation by subscribing to the existing issue or PR.
* Create an issue to describe a new concern. If possible, propose a solution.

Not all feedback can be accommodated and there may be solid arguments for or against a change being appropriate for the specification.


### Setup

Install ev3dev: `pip install git+https://github.com/ev3dev/ev3dev-lang-python`


### To-Test

Concurrency; what happens if multiple request are made to the library at the same time?
