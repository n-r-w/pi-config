<id_guideline>
MUST use following ID prefixes for all findings, interpretations, hypotheses, open questions, recommendations, etc.:
    1. Context: CTX
    2. Problem Statement: PRB
    3. Trigger / Why Now: TRG
    4. Goal: GOL
    5. Success Metric: MET
    6. Non-Goal: NGL
    7. Actor: ATR
    8. Scenario: SCN
    9. In-Scope Item: ISP
    10. Out-of-Scope Item: OSP
    11. Functional Requirement: FRQ
    12. Non-Functional Requirement: NFQ
    13. Constraint: CNS
    14. Risk: RSK
    15. Assumption: ASM
    16. Open Question: QST
    17. Acceptance Criterion: ACC
    18. Decision: DEC
    19. Solution Overview Item: SOL
    20. Failure Mode: FLR
    21. Component: CMP
    22. Entity / Model: ENT
    23. API / Resource Contract: APC
    24. Event / Message Contract: EVC
    25. Trade-Off: TRD
    26. Phase: PHS
    27. Task / Work Item: TSK
    28. Deliverable: DLV
    29. Exit Criterion: EXC
    30. Dependency: DEP
    31. Blocker: BLK
    32. Release Step: RLS
    33. Rollback Condition: RBC
    34. Rollback Action: RBA
    35. Definition of Done Item: DOD
    36. Finding: FND
    37. Review Check Item: CHK
    38. Next Step Action: NXT
    39. Observation: OBS
    40. Interpretation: INT
    41. Hypothesis: HYP
    42. Option: OPT
    43. Recommendation: REC
    44. Action Plan Item: ACN
    45. Limitation: LIM
    46. Configuration Item: CFG
    47. Reference: REF
    48. Definition/Abbreviation: DEF
    49. Algorithm: ALG
    50. Algorithm/Workflow Step: STP
    51. Diagram: DGM

Rules:
    1. MUST use 2 digits for IDs, e.g. GOL-01, BLK-02.
    2. If there is no suitable prefix, create a custom one.
    3. IDs MUST be unique across document.
    4. If when editing document it is necessary to change order of items, you MUST keep their IDs. New items MUST receive IDs with an additional suffix so as not to break existing numbering. E.g.: there are `ID-1`, `ID-2`; you need to add new item between them -> new item receives an ID with an additional suffix `ID-1.1`
</id_guideline>