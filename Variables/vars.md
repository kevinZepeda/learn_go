# Summary
1. Variable Declaration
  - var foo int
  - var foo int = 43
  - foo := 42
2. Variables can't redeclare variables, but can shadow then.
3. All variables must be used.
4. Visibility
  - Lower case fisrt letter for package scope
  - Upper case firts letter to export
  - No private exports
5. Naming conventions
  - Pascal or camelCase
    - Capitalize acronyms (HTTP, URL)
  - As short as reasonable
    - Longer names for Longer lives
6. Type conversion
  - destinationType(variable)
  - use strconv package for strings
