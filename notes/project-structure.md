# Project Structure
 
Golang gives you the opportunity to structure your project however you want, but there are some things to take note of:


- The `internal/`directory
This is a special Go convention - code in internal/ can only be imported by code in the same project. It's like making things "private" to your project. Good for your business logic that shouldn't be used by other projects.
- The `pkg/` directory
This is for code that could be reused by other projects.

Next: [Environment Variables](/notes/env-vars.md).