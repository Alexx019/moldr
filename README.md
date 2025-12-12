# Moldr

Moldr is a tool to create, run and manage backed applications.

## Dictionary
- **ingot**: A especific backed instance created from a mold.
- **mold**: A executable template that defines how an ingot is created, run and managed.

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
```
