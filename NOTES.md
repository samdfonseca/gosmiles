## Thoughts
- Program only needs to handle alkanes (simple hydrocarbons) to greatly reduce the scope and complexity
  - Only single bonds. No need to model multiple types of bonds in the data structure.
  - Carbon and hydrogen only.
- Program only needs to handle straight-chain and branched alkanes, i.e. no cyclic
  - Need to be able to identify longest straight-chain (LSC).
  - Need to be able to identify where in LSC the branch occurs.


### Parsing
- Since we are only handling straight-chain and branched alkanes, we can use a tree data structure as a parse target
  - Data structure is directional, not cyclical.

### Naming
- Needed for naming:
  - The longest continuous carbon chain (parent chain). Used to determine the parent name of the alkane.
    - If there are multiple choices of the same length, the parent chain is the longest chain with the greatest number of branches (substituents).
  - Location of branches off of parent chain.
    - Start at root with counter=1, then advance to next node in parent chain (root node can't have branches) and increment counter.
  - Location of first branch off of parent chain.
    - Chains are reversable.
    - Number the chain beginning at the end closest to any branches to ensure the lowest possible numbers for the position of branches.
    - If parent chain length is 8, and first branch is at node 6, renumber node 6 as node 3
    - f(n) = pcl - n + 1, where pcl is parent chain length and n is node number
  - Number the chain beginning at the end closest to any substituents.
    - Need to be able to reverse the Tree data structure.
  - Root node can have branches.
    - First node of a branch, i.e. 1,1-dimethylethyl


## Solution Notes
### Organization
- This solution is organized into 3 packages.
  - A lexer + parser for reading the SMILES string into a data structure.
  - The node tree data structure used to represent the molecule.
  - A namer package to produce IUPAC nomenclature from the node tree data structure.

### Limitations
- Molecules with complex branches (branches that have branches) are not supported.
  - Only supports molecules with straight branches.
- Input sanitizing is minimal.
  - Depending on the use case for such a tool, additional input sanitizing may be required.
  - Currently, the assumption is that the user controls the input and that input will only be well formed SMILES strings.
- Some valid SMILES strings are not supported.
  - SMILES has several optional features.
  - Currently, only unique (canonicalized) SMILES strings are supported.
  - Brackets and explicit hydrogen bonds are not supported.
  - For example, "\[CH3][CH2][CH3]" must be written as "CCC".

### Further Work
- Adding support for molecules with complex branches requires rethinking some shortcuts used in naming.
  - Molecule naming is recursive. Naming is based on chains of carbon atoms. Branches are their own chains. Branches off
    of branches are their own chains. Etc.
  - From a high level, naming a molecule is naming chains recursively until all chains have been named.
  - Complexity around certain rules, such as alphabetization of branches, makes this a non-trivial matter.
    - Straight branches are ordered based off their base, i.e. 4-methyl before 3-propyl. Multiplier prefixes are ignored.
    - Complex branches are ordered based off their full name, i.e. 1-dimethylethyl before 2-trimethylpropyl.
  - A lack of standardization leaves ambiguity around handling certain scenarios involving complex branches.
    - For example, "CCC(C)CC(C)C(C(C)(CC)C)C(CC)C(C(C)(CC)C)CCC" has two 1,1-dimethylethyl branches. Where the multiplier
      prefix, if one at all, should be places is unclear.
