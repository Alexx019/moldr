# Moldr

Moldr is a tool to create, run and manage backed applications (ingots) based on precreated templates (molds).

## Dictionary
- **ingot**: A especific backed instance created from a mold.
- **mold**: A executable template that defines the structure of a backed application.

## Usage

```bash
# List molds and ingots
> moldr list (molds | ingots:default)

# Create a new ingot or mold
> moldr new ingot <ingot_name> (--mold=<mold_name>) (--port=<port>)
> moldr new mold <file_path>

# Actions
> moldr run <ingot_name>
> moldr stop <ingot_name>
> moldr delete <ingot_name>
> moldr logs <ingot_name>
```
