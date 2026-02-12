---
updated_at: 2026-02-07T17:48:20.312+10:00
tags:
  - Naval
  - Deepnight
  - Mercenary
aliases:
  - CEI
---
### **Analysis of CEI Rules from Three Sources**

**Common Elements (All Three Sources):**
*   **Core Concept:** CEI (Crew/Company Effectiveness Index) is a 0-15 rating representing the overall training, competence, cohesion, and morale of a group (ship's crew or military unit). It acts as a Die Modifier (DM) for tasks undertaken by the group as a whole.
*   **CEI Table:** All three sources use an identical table linking CEI scores (0-15) to Training Level descriptions, expected Specialist/General Skill Levels for typical members, and a Task DM.
*   **Dynamic Nature:** CEI can change during a mission/campaign due to events. It can improve from training, successes, or good supply; it can degrade from casualties, leadership crises, or major setbacks.
*   **Crisis & Improvement Mechanics:** Specific events (5% casualties, loss of CO, major system/position loss, leadership crisis) can trigger a potential CEI loss, avoidable with a Difficult (10+) Leadership check. Improvements can be triggered by specific positive events (training, success, provisioning, new personnel) and require an Average (8+) Leadership check.
*   **Periodic Checks:** Every 2D weeks, a Leadership check determines if the group's effectiveness (and morale) improves, declines, or stays the same, using an identical results table.
*   **Heroic/Extreme Measures:** A leader can inspire a special effort by expending points of CEI/ECEI. Each point spent grants DM+3 on related task rolls or allows the attempt of a near-impossible task. This expenditure is permanent, representing exhaustion and sacrifice.
*   **Detachments (DEI):** All sources introduce the Detachment Efficiency Index (DEI) to model the effectiveness of a sub-group (ship's department, mercenary platoon, mission team).
*   **Forming Detachments:** The process for forming a detachment and calculating its DEI is identical: a Leadership check, modified by the detachment's size relative to the parent group, and a roll on the "Total DEI" table. Options exist to form an elite "A-team" (DM+3) or a weak detachment (DM-3).
*   **Weakening:** Rules for how forming detachments, taking casualties, or adding inexperienced personnel can weaken the parent group's CEI/DEI are functionally identical, using the same modifiers and results table.

**Key Differences & Source-Specific Details:**

1.  **Naval Sourcebook:**
    *   **Scope:** Standard rules for naval and merchant vessels.
    *   **Short-Term Effectiveness:** Uses **Effective CEI (ECEI)**, a variable rating starting at base CEI at mission start and fluctuating based on events.
    *   **Permanent Change:** Permanent CEI change requires long-term training/experience, typically between missions.
    *   **Context Examples:** Provides examples for typical CEI ranges for peacetime navies (6-8), corsairs (4-9), and free traders (3-5).

2.  **Deepnight Revelation Campaign:**
    *   **Scope:** Specific to the Deepnight Revelation campaign ship.
    *   **Short-Term Effectiveness:** Uses a **CEI Modifier (CEIM)** applied to the static base **CEI**. The result is called Effective CEI (ECEI). So, ECEI = CEI + CEIM.
    *   **Permanent Change:** Provides specific, detailed rules for permanently altering the base CEI itself (e.g., -1 per 25 casualties, training checks every 2D months).
    *   **Integration:** Deeply integrates CEI/DEI with campaign-specific structures (Divisions, Teams, the "Reach" mechanic, and "Esteem").
    *   **Extreme Measures:** Adds specific casualty rules (1D per point spent) for using Extreme Measures.

3.  **Mercenary Book:**
    *   **Scope:** Rules for mercenary units and military forces (referred to as Company Efficiency Index).
    *   **Short-Term Effectiveness:** Uses **Effective CEI (ECEI)** like the Naval book.
    *   **Campaign & Long-Term Development:** Adds extensive systems not found in the others:
        *   **Post-Mission Analysis:** Uses the Ticket Success Indicator (TSI) to award Experience Points.
        *   **Experience System:** Experience Points can be spent to buy increases (or are lost to buy off decreases) in CEI, DEI, Morale, Reputation, Traits, etc.
        *   **Skill Atrophy:** Experience Points decay when the unit is not on a contract or training.
        *   **Reorganization & Discipline:** Includes rules for unit expansion, cadre duties, promotions, and discipline.

**Contradiction/Divergence to Note:**
*   **Short-Term Tracking Method:** The Naval and Mercenary sources use a variable **ECEI** value that changes up and down. The Deepnight Campaign uses a static base **CEI** with a fluctuating **CEI Modifier**. These are two different mechanical approaches to the same concept. The results are similar (a current effectiveness score), but the tracking method differs.

---
### **Compiled Rules for Crew/Company Effectiveness Index (CEI)**

**1. Core Concept**
The Crew (or Company) Effectiveness Index (CEI) is a score from 0 to 15 that represents the overall training, competence, cohesion, and morale of an organized group, such as a ship's crew or a military unit. It serves as a Die Modifier (DM) for any task the group undertakes as a whole. CEI is dynamic and can change due to the events of a mission or campaign.

**2. CEI Rating Table**
The following table defines the meaning of each CEI score. The "Specialist/General Skills Level" is a quick-reference guide for the typical skill levels of a random group member.

| CEI | Training Level                                                 | Specialist/General Skills Level | Task DM | AOC |
| :-: | -------------------------------------------------------------- | :-----------------------------: | :-----: | :-: |
|  0  | Barely able to do their jobs, with little cohesion             |               0/0               |   -6    | +0  |
|  1  |                                                                |               1/0               |   -5    | +1  |
|  2  | Extremely poorly skilled crew                                  |               1/0               |   -4    | +1  |
|  3  |                                                                |             1+0/1+0             |   -3    | +2  |
|  4  |                                                                |             1+0/1+0             |   -2    | +2  |
|  5  | Low quality or poorly trained crew                             |             1+0/1+0             |   -1    | +2  |
|  6  |                                                                |             2+1/1+1             |   -1    | +3  |
|  7  | Properly trained naval or merchant crew / military force       |             2+1/1+1             |    0    | +3  |
|  8  |                                                                |             2+1/1+1             |    0    | +3  |
|  9  | Highly-trained crew                                            |             3+1/1+1             |   +1    | +4  |
| 10  |                                                                |             3+1/1+1             |   +1    | +4  |
| 11  |                                                                |             3+1/1+1             |   +2    | +4  |
| 12  | Elite or veteran crew                                          |             3+1/2+1             |   +3    | +5  |
| 13  |                                                                |             3+1/2+1             |   +4    | +5  |
| 14  |                                                                |             3+1/2+1             |   +5    | +5  |
| 15  | Legendary crew formed from the cream of veterans and prodigies |             4+2/2+2             |   +6    | +6  |

*   **Skill Notation:** "2+1/1+1" means a typical specialist has a primary skill at level 2, a different secondary specialist skill at level 1, one general skill at level 1, and a second general skill at level 1. This is a template, not an exhaustive skill list.

**3. Determining Current Effectiveness**
There are two presented methods for tracking the group's effectiveness during a mission. **They are mutually exclusive; the Referee must choose one.**

*   **Method A (Naval & Mercenary Source):** The group has a permanent **CEI** and a current **Effective CEI (ECEI)**. ECEI starts equal to CEI at the beginning of a mission/ticket and fluctuates based on events. Task DMs are based on ECEI.
*   **Method B (Deepnight Campaign):** The group has a permanent base **CEI** and a current **CEI Modifier (CEIM)**. **Effective CEI (ECEI)** is calculated as (CEI + CEIM). The CEIM fluctuates based on events. Task DMs are based on ECEI.

**4. Events Affecting Effectiveness (ECEI or CEIM)**
The following events can change the group's current effectiveness (ECEI in Method A, CEIM in Method B).

*   **Crises (Potential Decrease):** Upon a crisis, the commanding officer must make a **Difficult (10+) Leadership check (2D x 5 minutes, INT)**. Failure reduces current effectiveness by 1 and Morale by the negative Effect.
    *   Crises include: 
	    * Taking casualties equal to 5% of the group; 
	    * loss of a major system/key position; 
	    * a Leadership Crisis; 
	    * death/disablement of the commanding officer.
*   **Improvements (Potential Increase):** When a positive event occurs, the commanding officer can make an **Average (8+) Leadership check (2D x 5 minutes, INT)** to try to increase current effectiveness by 1. Morale is modified by the check's Effect.
    *   Events include: 
	    * Generous supply/provisioning; 
	    * a solid, textbook success; 
	    * a period of dedicated training/exercises (at least 2 weeks); 
	    * receiving a draft of high-quality replacement/additional personnel.

**5. Periodic Effectiveness & Morale Check**
Every **2D weeks** of sustained operations, the commanding officer must make a **Difficult (10+) Leadership check (INT)**. Apply the Effect as a DM to the roll on the following table:

| 2D+Effect | Result                                                                                                        |
| --------- | ------------------------------------------------------------------------------------------------------------- |
| 0-        | Morale collapses (-1D+3 Morale) and the group is near mutiny. Reduce current effectiveness (ECEI/CEIM) by -3. |
| 1-2       | -1D Morale. Reduce current effectiveness (ECEI/CEIM) by -2.                                                   |
| 3-4       | -D3 Morale. Reduce current effectiveness (ECEI/CEIM) by -1.                                                   |
| 5-8       | No change.                                                                                                    |
| 9-11      | The group gains confidence. +1 Morale.                                                                        |
| 12+       | Efficiency and morale increase. +1 to current effectiveness (ECEI/CEIM), +D3 Morale.                          |

**6. Heroic / Extreme Measures**
In a dire situation, a leader can inspire a supreme effort.
*   The leader makes a **Difficult (10+) Leadership check**.
*   The positive **Effect** of the check indicates how many points of **CEI/ECEI (Method A) or DEI (Method B)** can be expended.
*   **Each point expended** grants **DM+3 on all task rolls** for the effort **OR** allows the group to attempt to overcome a near-impossible (but physically plausible) obstacle.
*   **Cost:** All expended points are **permanently lost** from the group's CEI/ECEI or the detachment's DEI, representing exhaustion, resource depletion, and sacrifice. *(Deepnight Source adds: This typically inflicts 1D casualties per point spent on the involved detachment.)*

**7. Detachment Efficiency Index (DEI)**
The DEI functions like a CEI for a sub-group (e.g., Engineering department, infantry platoon, science team). It is used for tasks involving only that part of the crew/unit.

**Forming a Detachment:**
1.  **Base Value:** Start with the **DEI** of the parent department or the **CEI/ECEI** of the main group if drawn from multiple sources.
2.  **Leadership Check:** The officer forming the detachment makes a **Difficult (10+) Leadership check**. Note the Effect.
3.  **Size Modifier:** Apply a modifier based on the detachment's size as a percentage of the available personnel it is drawn from:
	*   <1%: +2
	- 1-5%: 0
	- 6-10%: -2
	- 11-20%: -4
	- 21-30%: -6
	- 31-40%: -8
	- 41-50%: -10
4.  **Special Modifiers:**
    *   **"A-Team" (Best Personnel):** DM+3. This weakens the parent group.
    *   **"Anyone Will Do" (Deliberately Weak):** DM-3.
5.  **Final Calculation:** Add the Leadership Effect and all modifiers. Roll 2D and apply the total as a DM on the table below to modify the **Base Value**.

| 2D | DEI Result |
| -- | ---------- |
| 0- | Base - 2D3 |
| 1-2 | Base - 1D |
| 3-4 | Base - D3 |
| 5-6 | Base - 1 |
| 7-8 | Base + 0 |
| 9-10 | Base + 1 |
| 11-12 | Base + D3 |
| 13+ | Base + 1D |

**8. Weakening the Parent Group**
Forming detachments, taking casualties, or absorbing new personnel can weaken a group. When this occurs:
1.  The commanding officer makes an **Average (8+) Leadership or Admin check**.
2.  Apply the following modifiers to the check's **Effect**:
    *   Formed an "A-Team": +2
    *   Formed a weak detachment: -2
    *   Casualties/Detachment Size: 
	    * <1%: 0
	    * 2-4%: -1
	    * 5-9%: -2
	    * Each additional 5%: -1
    * New personnel lack common training: -4
    * Have training but no experience together: -2
    * *(Mercenary Source adds: For each 25% of unit size added in new personnel: -1)*
3.  Roll 2D and add the modified Effect. Consult the table:

| 2D+Mod. Effect | Change to DEI or ECEI |
| :------------: | :-------------------: |
| 0- | -4 |
| 1-3 | -3 |
| 4-6 | -2 |
| 7-9 | -1 |
| 10-12 | 0 |
| 13+ | +1 |

**9. Long-Term Development & Experience (Mercenary Source Only)**
*   After a mission ("ticket"), success is measured by a Ticket Success Indicator (TSI), which grants **Experience Points**.
*   **Experience Points** can be spent to permanently increase (or must be lost to offset decreases in) **CEI, DEI, Morale, Reputation, Influence, and Traits** according to a cost table.
*   **Skill Atrophy:** Experience Points decay at a rate of 1 per week if the unit is not on a mission or engaged in training.

**10. Permanent CEI Changes (Deepnight & Mercenary Sources)**
*   Permanent changes to the base **CEI** require significant time, training, or experience.
*   **Deepnight:** Specific triggers like 25 casualties (-1 CEI), or a training period of 4D days every 2D months with a successful **Difficult (10+) Leadership/Admin check** can increase CEI.
*   **Mercenary:** Permanent CEI changes are purchased with Experience Points between missions.