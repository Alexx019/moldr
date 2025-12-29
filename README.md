# Moldr

Moldr is a tool to create, run and manage backed applications (ingots) based on precreated templates (molds).

## Dictionary
- **ingot**: A especific backed instance created from a mold.
- **mold**: A executable template that defines the structure of a backed application.

## Usage

```bash
# List ingots
> moldr list

# Create a new ingot
> moldr new <ingot_name> (--mold=<mold_name>) (--port=<port>)

# Actions
> moldr run <ingot_name>
> moldr stop <ingot_name>
> moldr delete <ingot_name>
> moldr logs <ingot_name>
```
