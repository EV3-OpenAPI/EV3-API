# Development Guidelines

This document intends to establish guidelines which build a transparent, open mechanism for deciding how to evolve the LEGO® MINDSTORMS® EV3-REST.
[ZHAW](https://www.zhaw.ch/en/engineering/) will initially follow these processes when merging changes from external contributors or from the
ZHAW itself. This guideline document will be adjusted as practicality dictates.

## Specification Change Criteria

The specification will *evolve over time*. Changes may be made when any of the following criteria are met:

- Clarity. The current "way" something is done doesn't make sense, is complicated, or not clear.

- Consistency. A portion of the specification is not consistent with the rest, or with the industry standard terminology.

- Necessary functionality. We are missing functionality because of a certain design of the specification.

- Forward-looking designs. We should always consider what the next important functionality should be.

## Tracking Process

- GitHub is the medium of record for all spec designs, use cases, and so on.

- The **human readable** document is the source of truth. If using a JSON Schema again to document the spec, it is secondary to the human documentation. The documentation should live in a *.md file, in parallel to the latest document (versions/1.0.0.md for example).

- At any given time, there would be at most 4 work branches. The branches would exist if work has started on them. Assuming a current version of 1.0.0:

- **main** - Current stable version. No PRs would be accepted directly to modify the specification. PRs against supporting files can be accepted.

- v1.0.1-dev - The next PATCH version of the specification. This would include non-breaking changes such as typo fixes, document fixes, wording clarifications.

- v1.1.0 - The next MINOR version.

- v2.0.0 - The next MAJOR version.

- The main branch shall remain the current, released LEGO® MINDSTORMS® EV3-REST. We will describe and link the work branch(es) on the default README.md on main.

- Examples of how something is described *currently* vs. the proposed solution should accompany any change proposal.

- New features should be done in *feature/name* which, upon approval, are merged into the proper work branch.

- Use labels for the workflow of specification changes.

- An issue will be opened for each feature change. Embedded in the issue, or ideally linked in a file via pull-request (PR), a document about use cases should be supplied with the change.

- A PR will be used to describe the proposed solution and linked to the original issue.

- When the work branch is ready and approved, the branch will be merged to main.

## Participation

The evolution of the specification happens through the participation of members of the developer community at large. Any person willing to contribute to the effort is welcome, and contributions may include filing or participating in issues, creating pull requests, or helping others with such activities.
