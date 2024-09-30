# E-Mail-Verifier
This Project is a small version of a Email Verification Tool.

In  this Project-

-At First we take the Domain,

-then we look for a MX Record(Mail Exchange Records) using the package in GO (net.LookupMX)

-after that by using another package net.LookupTXT we look for a spfRecords which is type of DNS record that helps identify which mail servers are permitted to send emails on behalf of a domain.

-then using the same we look for dmarc records (Domain-based Message Authentication, Reporting, and Conformance).They define a domain’s email authentication policy, specifying how to handle emails that fail authentication checks.
