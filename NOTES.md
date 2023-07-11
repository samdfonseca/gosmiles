- Program only needs to handle alkanes (simple hydrocarbons) to greatly reduce the scope and complexity
  - Only single bonds. No need to model multiple types of bonds in the data structure.
  - Carbon and hydrogen only.
- Program only needs to handle straight-chain and branched alkanes, i.e. no cyclic
  - Need to be able to identify longest straight-chain (LSC).
  - Need to be able to identify where in LSC the branch occurs.


## Parsing
- Since we are only handling straight-chain and branched alkanes, we can use a tree data structure as a parse target

## Naming
- Needed for naming:
  - The longest continuous carbon chain (parent chain). Used to determine the parent name of the alkane.
    - If there are multiple choices of the same length, the parent chain is the longest chain with the greatest number of branches (substituents).
  - Number the chain beginning at the end closest to any substituents.
    - Need to be able to reverse the Tree data structure.
