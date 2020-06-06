Tiny, opinionated authentication library for Go. Work in progress and not usable right now.

### Concepts
- An to add/edit/delete users.
- Users are identified by an identifier (a string: email, username, phone, anything).
- Supports hashed password and password authentication.
- Password authentication can be turned off to enable external OAuth.
- Permissions are represented as a list of strings per user. Eg: `post.edit, post.create` etc.
- User account statuses: `pending, enabled, disabled` etc.
- Pluggable functions for password reset, recovery etc.
- Multiple "store" backends (Postgres, Redis etc.) via a Store interface.
