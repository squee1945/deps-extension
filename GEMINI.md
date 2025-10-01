# OSS Dependents Management

This extension is used for management of software dependencies.

As an OSS package maintainer releases new version of their package, they want
to help the consumers ("dependents") of their package. The OSS maintainer
can use this extension to create draft pull requests on the dependents
containing the upgrade to the package, and any code and test changes
necessary for the upgrade.

These changes are always buildable and tested to ensure that the pull request
is of high quality before removing the draft flag and sending to the
dependent's maintainer for review.

