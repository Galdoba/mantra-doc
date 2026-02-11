---
updated_at: 2026-01-20T11:25:53.004+10:00
tags:
  - objectvalue
---
источник: https://travellermap.com/doc/secondsurvey
# Second Survey Data

The Traveller5 Second Survey format defines a standard for describing star systems in the Official Traveller Universe. It is composed of a series of fields describing various aspects of the Main World (MW) in each system including physical and cultural characteristics, and some details about other bodies in the system of interest to travellers.

The data for Regina (Spinward Marches 1910), home of Duke Norris and an influential world on the edge of the Third Imperium can be summarized in Second Survey format as:

`1910  Regina  A788899-C  Ht Ri Pa Ph An Cp (Amindii)2  { 4 }  (D7E+5)  [9C6D]  BcCeF  NS  -  703  8  ImDd  F7 V BD M3 V`

Different editions of Traveller and different authors and publications may include different fields, present the fields in different orders, or provide different values for fields. The companion document [Sector Data File Formats](https://travellermap.com/doc/fileformats) may be used for reference.

### Contents

- Data Fields
    - [Hex Location](https://travellermap.com/doc/secondsurvey#hex)
    - [Main World Name](https://travellermap.com/doc/secondsurvey#name)
    - [UWP - Universal World Profile](https://travellermap.com/doc/secondsurvey#uwp)
    - [Remarks and Trade Codes](https://travellermap.com/doc/secondsurvey#remarks)
    - [Importance Extension](https://travellermap.com/doc/secondsurvey#ix)
    - [Economic Extension](https://travellermap.com/doc/secondsurvey#ex)
    - [Cultural Extension](https://travellermap.com/doc/secondsurvey#cx)
    - [Nobility](https://travellermap.com/doc/secondsurvey#nobility)
    - [Bases](https://travellermap.com/doc/secondsurvey#bases)
    - [Travel Zone](https://travellermap.com/doc/secondsurvey#zone)
    - [PBG - Population, Belts, Giants](https://travellermap.com/doc/secondsurvey#pbg)
    - [Worlds](https://travellermap.com/doc/secondsurvey#worlds)
    - [Allegiance](https://travellermap.com/doc/secondsurvey#allegiance)
    - [Stellar Data](https://travellermap.com/doc/secondsurvey#stellar)
    - [Resource Units](https://travellermap.com/doc/secondsurvey#ru)
- Appendix
    - [eHex](https://travellermap.com/doc/secondsurvey#ehex)

## Hex Location

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

The location of the world within a sector, given as a four digit number identifying a position on a hexagonal grid. The first two digits are the distance from the spinward edge in parsecs ranging from `01` through `32`. The second two digits are the distance from the coreward edge in parsecs ranging from `01` through `40`.

## Main World Name

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

Name of the primary world in the system. Most characteristics given in the UWP, such as physical and cultural attributes, refer to this world.

## UWP - Universal World Profile

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

The UWP gives a compact description of the physical and cultural aspects of the Main World, as well the starport facilities and available technology. The fields are respectively:

|Code|Description|
|---|---|
|[Starport](https://travellermap.com/doc/secondsurvey#starport)|Type of starport facility on world.|
|[Size](https://travellermap.com/doc/secondsurvey#size)|World diameter (in units of 1,600 kilometers).|
|[Atmosphere](https://travellermap.com/doc/secondsurvey#atmosphere)|World atmosphere type.|
|[Hydrographics](https://travellermap.com/doc/secondsurvey#hydrographics)|World surface covered with water (in tenths).|
|[Population](https://travellermap.com/doc/secondsurvey#population_exponent)|Exponent of intelligent population.|
|[Government](https://travellermap.com/doc/secondsurvey#government)|World government type.|
|[Law Level](https://travellermap.com/doc/secondsurvey#law_level)|Degree of oppression by law.|
|[Technological Level](https://travellermap.com/doc/secondsurvey#tech_level)|Level of technological achievement.|

### Starport

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

Starport indicates the best quality starport in the star system available to travellers.

|Code|Starport Description|
|---|---|
|`A`|Excellent Quality. Refined fuel available. Annual maintenance overhaul available. Shipyard capable of constructing starships and non-starships present. Naval base and/or scout base may be present.|
|`B`|Good Quality. Refined fuel available. Annual maintenance overhaul available. Shipyard capable of constructing non-starships present. Naval base and/or scout base may be present.|
|`C`|Routine Quality. Only unrefined fuel available. Reasonable repair facilities present. Scout base may be present.|
|`D`|Poor Quality. Only unrefined fuel available. No repair facilities present. Scout base may be present.|
|`E`|Frontier Installation. Essentially a marked spot of bedrock with no fuel, facilities, or bases present.|
|`X`|No Starport. No provision is made for any ship landings.|
|`?`|Unknown.|

### Spaceports

Worlds other than the Main World in a system may also be described by UWP codes. The port, if any, is referred to as a spaceport.

|Code|Spaceport Description|
|---|---|
|`F`|Good Quality. Minor damage repairable. Unrefined fuel available.|
|`G`|Poor Quality. Superficial repairs possible. Unrefined fuel available.|
|`H`|Primitive Quality. No repairs or fuel available.|
|`Y`|None.|

### Size

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

Size expresses the diameter of a world in approximately 1,600km units (or 1,000 mile units).

Size code `0` indicates that the main "world" of the system is an asteroid belt.

|   |   |   |   |   |   |
|---|---|---|---|---|---|
|_Size_|_Diameter (km)_|_Mass (Earth = 1)_|_Area (Earth = 1)_|_Gravity (G)_|_Esc. Vel (km/s)_|
|`1`|1,600|0.0019|0.015|0.122|1.35|
|`2`|3,200|0.015|0.063|0.240|2.69|
|`3`|4,800|0.053|0.141|0.377|4.13|
|`4`|6,400|0.125|0.250|0.500|5.49|
|`5`|8,000|0.244|0.391|0.625|6.87|
|`6`|9,600|0.422|0.563|0.840|8.72|
|`7`|11,200|0.670|0.766|0.875|9.62|
|`8`|12,800|1.000|1.000|1.000|11.00|
|`9`|14,400|1.424|1.266|1.120|12.35|
|`A`|16,000|1.953|1.563|1.250|13.73|
|`B`|18,800|2.600|1.891|1.375|15.34|
|`C`|19,200|3.375|2.250|1.500|16.74|
|`D`|20,800|4.291|2.641|1.625|18.13|
|`E`|22,400|5.359|3.063|1.750|19.52|
|`F`|24,000|6.592|3.516|1.875|20.92|
|`?`|Unknown.|   |   |   |   |

The above table assumes a density similar to Earth (5.5 grams per cubic centimeter).

### Atmosphere

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

Atmosphere type shows the general character of the atmosphere for a world on its surface.

|Code|Description|
|---|---|
|`0`|No atmosphere. Requires vacc suit.|
|`1`|Trace. Requires vacc suit.|
|`2`|Very thin. Tainted. Requires combination respirator/filter.|
|`3`|Very thin. Requires respirator.|
|`4`|Thin. Tainted. Requires filter mask.|
|`5`|Thin. Breathable.|
|`6`|Standard. Breathable.|
|`7`|Standard. Tainted. Requires filter mask.|
|`8`|Dense. Breathable.|
|`9`|Dense. Tainted. Requires filter mask.|
|`A`|Exotic. Requires special protective equipment.|
|`B`|Corrosive. Requires protective suit.|
|`C`|Insidious. Requires protective suit.|
|`D`|Dense, high. Breathable above a minimum altitude.|
|`E`|Thin, low. Breathable below certain altitudes.|
|`F`|Unusual.|
|`?`|Unknown.|

### Hydrographics

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

Hydrographics shows the percentage of world surface covered by seas or oceans. If atmosphere type is `A` or above, fluid may be present in place of water.

|Code|Description|
|---|---|
|`0`|No water. Desert World.|
|`1`|10% water.|
|`2`|20% water.|
|`3`|30% water.|
|`4`|40% water.|
|`5`|50% water.|
|`6`|60% water.|
|`7`|70% water. Equivalent to Terra or Vland.|
|`8`|80% water.|
|`9`|90% water.|
|`A`|100% water. Water World.|
|`?`|Unknown.|

### Population Exponent

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

The population exponent gives an order-of-magnitude approximation of a world's population. Together with the population multiplier digit in the [PBG](https://travellermap.com/doc/secondsurvey#pbg) field, the world population can be computed as:

population = multiplier × 10exponent

|Code|Description|
|---|---|
|`0`|Few or no inhabitants.|
|`1`|Tens of inhabitants.|
|`2`|Hundreds of inhabitants.|
|`3`|Thousands of inhabitants.|
|`4`|Tens of thousands of inhabitants.|
|`5`|Hundreds of thousands of inhabitants.|
|`6`|Millions of inhabitants.|
|`7`|Tens of millions of inhabitants.|
|`8`|Hundreds of millions of inhabitants.|
|`9`|Billions of inhabitants.|
|`A`|Tens of billions of inhabitants.|
|`B`|Hundreds of billions of inhabitants.|
|`C`|Trillions of inhabitants.|
|`D`|Tens of trillions of inhabitants.|
|`E`|Hundreds of trillions of inhabitants.|
|`F`|Quadrillions of inhabitants.|
|`?`|Unknown.|

### Government Type

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

Government shows the basic governmental structure for a world.

Codes `G` and above were defined in the Classic Traveller alien modules for sophont-specific govenment types. Other editions of Traveller may define other values for certain codes.

|Code|Description|Allegiance|
|---|---|---|
|`0`|No Government Structure.||
|`1`|Company/Corporation.||
|`2`|Participating Democracy.||
|`3`|Self-Perpetuating Oligarchy.||
|`4`|Representative Democracy.||
|`5`|Feudal Technocracy.||
|`6`|Captive Government / Colony.||
|`7`|Balkanization.||
|`8`|Civil Service Bureaucracy.||
|`9`|Impersonal Bureaucracy.||
|`A`|Charismatic Dictator.||
|`B`|Non-Charismatic Dictator.||
|`C`|Charismatic Oligarchy.||
|`D`|Religious Dictatorship.||
|`E`|Religious Autocracy.||
|`F`|Totalitarian Oligarchy.||
|`G`|Small Station or Facility.|Aslan.|
|`H`|Split Clan Control.|Aslan.|
|`J`|Single On-world Clan Control.|Aslan.|
|`K`|Single Multi-world Clan Control.|Aslan.|
|`L`|Major Clan Control.|Aslan.|
|`M`|Vassal Clan Control.|Aslan.|
|`N`|Major Vassal Clan Control.|Aslan.|
|`P`|Small Station or Facility.|K'kree.|
|`Q`|Krurruna or Krumanak Rule for Off-world Steppelord.|K'kree.|
|`R`|Steppelord On-world Rule|K'kree.|
|`S`|Sept.|Hiver.|
|`T`|Unsupervised Anarchy.|Hiver.|
|`U`|Supervised Anarchy.|Hiver.|
|`V`|||
|`W`|Committee.|Hiver.|
|`X`|Droyne Hierarchy.|Droyne.|
|`Y`|||
|`Z`|||
|`?`|Unknown.|   |

#### Traveller: The New Era

Worlds in the Wilds ([allegiance](https://travellermap.com/doc/secondsurvey#allegiance) code: `Wi`) use a different set of government codes.

|Code|Description|
|---|---|
|`0`|No Government Structure.|
|`1`|Tribal Government.|
|`2`|Participating Democracy.|
|`3`|Representative Democracy.|
|`4`|Charismatic Dictator.|
|`5`|Charismatic Oligarchy.|
|`6`|Technologically Elevated Dictator (TED).|
|`7`|Mystic Dictatorship.|
|`8`|Totalitarian Oligarchy.|
|`9`|Mystic Autocracy.|
|`A`|Civil Service Bureaucracy.|
|`B`|Self-Perpetuating Oligarchy.|
|`C`|Impersonal Bureaucracy.|

### Law Level

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

Law level indicates basic legal status and shows probability of harassment by local enforcers.

|Code|Description|
|---|---|
|`0`|No prohibitions.|
|`1`|Body pistols, explosives, and poison gas prohibited.|
|`2`|Portable energy weapons prohibited.|
|`3`|Machine guns, automatic rifles prohibited.|
|`4`|Light assault weapons prohibited.|
|`5`|Personal concealable weapons prohibited.|
|`6`|All firearms except shotguns prohibited.|
|`7`|Shotguns prohibited.|
|`8`|Long bladed weapons controlled; open possession prohibited.|
|`9`|Possession of weapons outside the home prohibited.|
|`A`|Weapon possession prohibited.|
|`B`|Rigid control of civilian movement.|
|`C`|Unrestricted invasion of privacy.|
|`D`|Paramilitary law enforcement.|
|`E`|Full-fledged police state.|
|`F`|All facets of daily life regularly legislated and controlled.|
|`G`|Severe punishment for petty infractions.|
|`H`|Legalized oppressive practices.|
|`J`|Routinely oppressive and restrictive.|
|`K`|Excessively oppressive and restrictive.|
|`L`|Totally oppressive and restrictive.|
|`S`|Special/Variable situation.|
|`?`|Unknown.|

### Technological Level

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

Technological level shows the degree of technological sophistication to be expected on a world.

|Code|Description|
|---|---|
|`0`|Stone Age. Primitive.|
|`1`|Bronze, Iron. Bronze Age to Middle Ages|
|`2`|Printing Press. circa 1400 to 1700.|
|`3`|Basic Science. circa 1700 to 1860.|
|`4`|External Combustion. circa 1860 to 1900.|
|`5`|Mass Production. circa 1900 to 1939.|
|`6`|Nuclear Power. circa 1940 to 1969.|
|`7`|Miniaturized Electronics. circa 1970 to 1979.|
|`8`|Quality Computers. circa 1980 to 1989.|
|`9`|Anti-Gravity. circa 1990 to 2000.|
|`A`|Interstellar community.|
|`B`|Lower Average Imperial.|
|`C`|Average Imperial.|
|`D`|Above Average Imperial.|
|`E`|Above Average Imperial.|
|`F`|Technical Imperial Maximum.|
|`G`|Robots.|
|`H`|Artificial Intelligence.|
|`J`|Personal Disintegrators.|
|`K`|Plastic Metals.|
|`L`|Comprehensible only as technological magic.|
|`?`|Unknown.|

## Remarks and Trade Codes

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

Remarks and trade classifications indicate obvious or important characteristics for the Main World in the system. They serve to show the potential for a world based on its capacity as a source of trade goods, a market for trade goods, or both.

|Code|Description|
|---|---|
|Planetary|   |
|`As`|Asteroid Belt.Siz 0|
|`De`|Desert.Atm 2-9, Hyd 0|
|`Fl`|Fluid Hydrographics (in place of water).Atm A-C, Hyd 1+|
|`Ga`|Garden World.Siz 6-8, Atm 5,6,8, Hyd 5-7|
|`He`|Hellworld.Siz 3+, Atm 2,4,7,9-C, Hyd 0-2|
|`Ic`|Ice Capped.Atm 0-1, Hyd 1+|
|`Oc`|Ocean World.Siz A+, Hyd A|
|`Va`|Vacuum World.Atm 0|
|`Wa`|Water World.Siz 3-9, Atm 3-9, Hyd A|
|Population|   |
|`Di`|Dieback.PGL 0, TL 1+|
|`Ba`|Barren.PGL 0, TL 0|
|`Lo`|Low Population.Pop 1-3|
|`Ni`|Non-Industrial.Pop 4-6|
|`Ph`|Pre-High Population.Pop 8|
|`Hi`|High Population.Pop 9+|
|Economic|   |
|`Pa`|Pre-Agricultural.Atm 4-9, Hyd 4-8, Pop 4,8|
|`Ag`|Agricultural.Atm 4-9, Hyg 4-8, Pop 5-7|
|`Na`|Non-Agricultural.Atm 3−, Hyd 3−, Pop 6+|
|`Pi`|Pre-Industrial.Atm 0,1,2,4,7,9, Pop 7-8|
|`In`|Industrialized.Atm 0,1,2,4,7,9-C, Pop 9+|
|`Po`|Poor.Atm 2-5, Hyd 3−|
|`Pr`|Pre-Rich.Atm 6,8, Pop 5,9|
|`Ri`|Rich.Atm 6,8, Pop 6-8|
|`Lt`|Low Technology.TL 5−|
|`Ht`|High Technology.TL 12+|
|Climate|   |
|`Fr`|Frozen.Siz 2-9, Hyd 1+, HZ +2 or Outer|
|`Ho`|Hot.HZ −1|
|`Co`|Cold.HZ +1|
|`Lk`|Locked.Close Satellite|
|`Tr`|Tropic.Siz 6-9, Atm 4-9, Hyd 3-7, HZ −1|
|`Tu`|Tundra.Siz 6-9, Atm 4-9, Hyd 3-7, HZ +1|
|`Tz`|Twilight Zone.Orbit 0-1|
|Secondary|   |
|`Fa`|Farming.Atm 4-9, Hyd 4-8, Pop 2-6, Not MW, HZ|
|`Mi`|Mining.Pop 2-6, Not MW, MW=In|
|`Mr`|Military Rule.By regional Allegiance power.|
|`Mr(AAAA)`|Military Rule (by allegiance `AAAA`).|
|`Px`|Prison, Exile Camp.MW|
|`Pe`|Penal Colony.Not MW|
|`Re`|Reserve.|
|Political|   |
|`Cp`|Subsector Capital.|
|`Cs`|Sector Capital.|
|`Cx`|Capital.|
|`Cy`|Colony (see `O:XXYY`).|
|Special|   |
|`Sa`|Satellite (Main World is a moon of a Gas Giant).|
|`Fo`|Forbidden (Red Zone).|
|`Pz`|Puzzle (Amber Zone).Pop 7+|
|`Da`|Danger (Amber Zone).Pop 6-|
|`Ab`|Data Repository.|
|`An`|Ancient Site.|
|`Rs`|Research Station.Imperial.|
|`RsA`|Research Station `A` = Alpha, `B` = Beta, `G` = Gamma, etc.Imperial.|
|Ownership|   |
|`O:XXYY`|Controlled by world in hex `XXYY`.|
|`O:SSSS-XXYY`|Controlled by world in hex `XXYY` in sector `SSSS`.|
|Sophonts|   |
|`[Sophont]`|Homeworld of major race.|
|`(Sophont)`|Homeworld of minor race. `0`-`9` indicates tenths of population if < 100%.|
|`(Sophont)#`|
|`Di(Sophont)`|Homeworld of extinct minor race (Dieback).|
|`Soph0`|Sophont Population. `Soph` is an [abbreviation for the sophont name](https://travellermap.com/doc/secondsurvey#sophont). `0`-`9` indicates tenths of population; `W` is 100%, e.g. `DroyW`.|
|`SophW`|
|Non-Standard / Legacy Codes|   |
|`S0`|Sophont Population. `S` is an [abbreviation for the sophont name](https://travellermap.com/doc/secondsurvey#sophont). `0`-`9` indicates tenths of population; `w` is 100%, e.g. `Dw`|
|`S:0`|
|`Sw`|
|`Nh`|Non-Hiver Population.Hiver.|
|`Nk`|Non-K'kree Population.K'kree.|
|`Tp`|Terra-prime. Adventure 5: Leviathan.|
|`Tn`|Terra-norm. Adventure 5: Leviathan.|
|`Fa`|Fascinating.Hiver.|
|`St`|Steppeworld.K'kree|
|`Ex`|Exile Camp.Imperial.|
|`Pr`|Prison World.Imperial.|
|`Xb`|Xboat Station.Imperial.|
|`Cr`|Reserve Capital.Zhodani.|
|`S#`|Stellar companion orbits (`F` = far), e.g. `S19` indicates companion stars are in orbits 1 and 9; `SF1` indicates far companion is binary (companion in orbit 1)|
|`S##`|
|`{...}`|Other comments or annotations.|

### Sophont Codes

Notable populations of sophonts are present in the [Remarks](https://travellermap.com/doc/secondsurvey#remarks) section. The species of sophont is identified using codes defined below.

### Traveller5

The sophont code is a 4-character code which is selected to uniquely identify a sophont species in Charted Space.

|Code|Sophont|Location|
|---|---|---|
|Adda|Addaxur|Zhodani space|
|Aezo|Aezorgh|[Wind](https://travellermap.com/go/Wind)|
|Akee|Akeed|[Gate](https://travellermap.com/go/Gate)|
|Aqua|Aquans (Daga)/Aquamorphs (Alph)|[Daga](https://travellermap.com/go/Daga) (Aquans)/[Alph](https://travellermap.com/go/Alph) (Aquamorphs)|
|Asla|Aslan|major|
|Bhun|Brunj|[Forn](https://travellermap.com/go/Forn)|
|Brin|Brinn|[Corr](https://travellermap.com/go/Corr)|
|Bruh|Bruhre|[Daib](https://travellermap.com/go/Daib)/[Reav](https://travellermap.com/go/Reav)|
|Buru|Burugdi|[Dagu](https://travellermap.com/go/Dagu)/[Thet](https://travellermap.com/go/Thet)|
|Bwap|Bwaps|Imperial/Vilani space|
|Chir|Chirpers|major|
|Clot|Clotho|[Tien](https://travellermap.com/go/Tien)|
|Darm|Darmine|[Zaru](https://travellermap.com/go/Zaru)|
|Dary|Daryen|[Spin](https://travellermap.com/go/Spin)|
|Dolp|Dolphins|Imperial/Solomani space|
|Droy|Droyne|major|
|Dync|Dynchia|[Leon](https://travellermap.com/go/Leon)|
|Esly|Eslyat|[Beyo](https://travellermap.com/go/Beyo)/[Vang](https://travellermap.com/go/Vang)|
|Flor|Floriani|[Beyo](https://travellermap.com/go/Beyo)/[Troj](https://travellermap.com/go/Troj)|
|Geon|Geonee|[Mass](https://travellermap.com/go/Mass)|
|Gnii|Gniivi|[Hint](https://travellermap.com/go/Hint)|
|Gray|Graytch|[Dagu](https://travellermap.com/go/Dagu)/[Gush](https://travellermap.com/go/Gush)/[Ilel](https://travellermap.com/go/Ilel)|
|Guru|Gurungan|[Solo](https://travellermap.com/go/Solo)|
|Gurv|Gurvin|Hiver space|
|Hama|Hamaran|[Dagu](https://travellermap.com/go/Dagu)|
|Hive|Hiver|Hiver space|
|Huma|Human|Imperial/Solomani space|
|Ithk|Ithklur|Hiver space|
|Jaib|Jaibok|[Thet](https://travellermap.com/go/Thet)|
|Jala|Jala'lak|[Dagu](https://travellermap.com/go/Dagu)|
|Jend|Jenda|[Hint](https://travellermap.com/go/Hint)/[Leon](https://travellermap.com/go/Leon)|
|Jonk|Jonkeereen|[Dene](https://travellermap.com/go/Dene)/[Spin](https://travellermap.com/go/Spin)|
|Kafo|Kafoe|[Cruc](https://travellermap.com/go/Cruc)|
|Kagg|Kaggushus|[Mass](https://travellermap.com/go/Mass)|
|Karh|Karhyri|[Cruc](https://travellermap.com/go/Cruc)|
|Kiak|Kiakh'iee|[Dagu](https://travellermap.com/go/Dagu)|
|K'kr|K'kree|K'[kree](https://travellermap.com/go/kree) space|
|Lamu|Lamura Gav/Teg|[Hint](https://travellermap.com/go/Hint)|
|Lanc|Lancians|[Dagu](https://travellermap.com/go/Dagu)/[Gush](https://travellermap.com/go/Gush)|
|Libe|Liberts|[Daib](https://travellermap.com/go/Daib)/[Dias](https://travellermap.com/go/Dias)|
|Llel|Llellewyloly|[Spin](https://travellermap.com/go/Spin)|
|Luri|Luriani|[Forn](https://travellermap.com/go/Forn)/[Ley](https://travellermap.com/go/Ley)|
|Mal'|Mal'Gnar|[Beyo](https://travellermap.com/go/Beyo)|
|Mask|Maskai|[Glim](https://travellermap.com/go/Glim)|
|Mitz|Mitzene|[Thet](https://travellermap.com/go/Thet)|
|Muri|Murians|[Vang](https://travellermap.com/go/Vang)|
|Orca|Orca|Imperial/Solomani space|
|Ormi|Ormine|[Dark](https://travellermap.com/go/Dark)|
|Ramm|Rammak|[Krus](https://travellermap.com/go/Krus)|
|Scan|Scanians|[Dagu](https://travellermap.com/go/Dagu)|
|Sele|Selenites|[Alph](https://travellermap.com/go/Alph)|
|S'mr|S'mrii|[Dagu](https://travellermap.com/go/Dagu)|
|Sred|Sred*Ni|[Beyo](https://travellermap.com/go/Beyo)|
|Stal|Stalkers|[Hint](https://travellermap.com/go/Hint)|
|Suer|Suerrat|[Ilel](https://travellermap.com/go/Ilel)|
|Sull|Sulliji|[Dene](https://travellermap.com/go/Dene)|
|Swan|Swanfei|[Gate](https://travellermap.com/go/Gate)|
|Sydi|Sydites|[Ley](https://travellermap.com/go/Ley)|
|Syle|Syleans|[Core](https://travellermap.com/go/Core)|
|Tapa|Tapazmal|[Reft](https://travellermap.com/go/Reft)|
|Taur|Taureans|[Alde](https://travellermap.com/go/Alde)|
|Tent|Tentrassi|[Zaru](https://travellermap.com/go/Zaru)|
|Tlye|Tlyetrai|[Reav](https://travellermap.com/go/Reav)|
|UApe|Uplifted Apes|Imperial/Solomani space|
|Ulan|Ulane|[Dark](https://travellermap.com/go/Dark)|
|Ursa|Ursa|[Forn](https://travellermap.com/go/Forn)/[Ley](https://travellermap.com/go/Ley)|
|Urun|Urunishani|[Anta](https://travellermap.com/go/Anta)|
|Varg|Vargr|[Anta](https://travellermap.com/go/Anta)/[Corr](https://travellermap.com/go/Corr)/[Dagu](https://travellermap.com/go/Dagu)/[Dene](https://travellermap.com/go/Dene)/[Empt](https://travellermap.com/go/Empt)/[Ley](https://travellermap.com/go/Ley)/[Lish](https://travellermap.com/go/Lish)/[Spin](https://travellermap.com/go/Spin)/Vargr space|
|Vega|Vegans|[Solo](https://travellermap.com/go/Solo)|
|Yile|Yileans|[Gash](https://travellermap.com/go/Gash)|
|Za't|Za'tachk|[Wren](https://travellermap.com/go/Wren)|
|Zhod|Zhodani|Zhodani space|
|Ziad|Ziadd|[Dagu](https://travellermap.com/go/Dagu)|

### Classic Traveller, MegaTraveller, Traveller: The New Era, Traveller: 4th Edition

Legacy sophont abbreviations include: `A` = Aslan, `C` = Chirper, `D` = Droyne, `F` = Non-Hiver Federation Member, `H` = Hiver, `I` = Ithklur, `M` = Human (e.g. in Vargr space), `V` = Vargr, `X` = Addaxur, `Z` = Zhodani.

## Importance Extension

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

### Traveller5

The Importance Extension is abbreviated Ix and written in braces (`{}`). It is a decimal integer (positive, negative, or zero) ranking the importance of the world within a region.

A world's Importance is calculated by adding or subtracting modifiers based on several factors, such as [Starport](https://travellermap.com/doc/secondsurvey#starport) type, [Tech Level](https://travellermap.com/doc/secondsurvey#tech_level), [Population Exponent](https://travellermap.com/doc/secondsurvey#population_exponent), [Economic codes](https://travellermap.com/doc/secondsurvey#remarks), and [Bases](https://travellermap.com/doc/secondsurvey#bases). Values range from `-3` (very unimportant) to `5` (very important).

## Economic Extension

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

### Traveller5

The Economic Extension is abbreviated Ex and written in parentheses (`()`). It describes the strength of a world's economy. It is given as three [eHex](https://travellermap.com/doc/secondsurvey#ehex) digits representing Resources, Labor and Infrastructure, followed by a decimal integer representing Efficiency written with a leading sign. (Note Efficiency-0 is coded as `+1` as it is functionally equivalent.)

Resources range from `2` (very scarce) to `J` (extremely abundant).

Labor correlates closely with [Population Exponent](https://travellermap.com/doc/secondsurvey#population_exponent).

Infrastructure ranges from `0` (non-existent) to `H` (very comprehensive).

Efficiency ranges from `-5` (extremely poor) to `+5` (very advanced).

## Cultural Extension

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

### Traveller5

The Cultural Extension is abbreviated Cx and written in brackets (`[]`). It gives insight into the social behavior of the world's population. It is given as four [eHex](https://travellermap.com/doc/secondsurvey#ehex) digits representing Heterogeneity, Acceptance, Strangeness, and Symbols. Unpopulated worlds will have `0` for all values.

Heterogeneity ranges from `1` (monolithic) to `G` (fragmented).

Acceptance ranges from `1` (extremely xenophobic) to `F` (extremely xenophilic).

Strangeness ranges from `1` (very typical) to `A` (incomprehensible).

Symbols ranges from `0` (extremely concrete) to `L` (incomprehensibly abstract).

## Nobility

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

### Traveller5

For Imperial worlds, the noble ranks assigned to the world by the Emperor based on importance.

If no noble ranks are assigned to the system, the field may be empty, a blank (space) or `-` (dash).

|Code|Ranking Noble|
|---|---|
|`B`|Knight.|
|`c`|Baronet.|
|`C`|Baron.|
|`D`|Marquis.|
|`e`|Viscount.|
|`E`|Count.|
|`f`|Duke.|
|`F`|Subsector Duke.|
|`G`|Archduke.|
|`H`|Emperor.|

## Bases

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

Base codes show the presence of military bases in a system; special codes deal with the presence of more than one type of base within the same system in order to maintain a single base code letter per system.

If no bases are present in the system, the field may be empty, a blank (space) or `-` (dash).

Base codes have changed significantly between different editions of Traveller. Sector data collections must indicate which set of codes are in use. Sector data collections may also define custom meanings for base codes.

### Traveller5

|Code|Description|Allegiance|
|---|---|---|
|`C`|Corsair Base.|Vargr.|
|`D`|Naval Depot.|Any.|
|`E`|Embassy.|Hiver.|
|`K`|Naval Base.|Any.|
|`M`|Military Base.|Any.|
|`N`|Naval Base.|Imperial.|
|`R`|Clan Base.|Aslan.|
|`S`|Scout Base.|Imperial.|
|`T`|Tlaukhu Base.|Aslan.|
|`V`|Exploration Base.|Any.|
|`W`|Way Station.|Any.|

Multiple codes may be used if multiple bases are present. For example, `NS` indicates that the system contains both an Imperial Naval Base and an Imperial Scout Base. Base codes must appear in alphabetical order, i.e. `MN` rather than `NM`.

### Classic Traveller, MegaTraveller, Traveller: The New Era, Traveller: 4th Edition

|Code|Description|Allegiance|
|---|---|---|
|`A`|Naval Base and Scout Base.|Imperial.|
|`B`|Naval Base and Way Station.|Imperial.|
|`C`|Corsair Base.|Vargr.|
|`D`|Depot.|Imperial.|
|`E`|Embassy Center.|Hiver.|
|`F`|Military and Naval Base.||
|`G`|Naval Base.|Vargr.|
|`H`|Naval Base and Corsair Base.|Vargr.|
|`J`|Naval Base.||
|`K`|Naval Base.|K'kree|
|`L`|Naval Base.|Hiver.|
|`M`|Military Base.||
|`N`|Naval Base.|Imperial.|
|`O`|Naval Outpost.|K'kree|
|`P`|Naval Base.|Droyne.|
|`Q`|Military Garrison.|Droyne.|
|`R`|Clan Base.|Aslan.|
|`S`|Scout Base.|Imperial.|
|`T`|Tlauku Base.|Aslan.|
|`U`|Tlauku and Clan Base.|Aslan.|
|`V`|Scout/Exploration Base.||
|`W`|Way Station.|Imperial.|
|`X`|Relay Station.|Zhodani.|
|`Y`|Depot.|Zhodani.|
|`Z`|Naval/Military Base.|Zhodani.|

Base codes indicate allegiance and general mission or type.

## Travel Zone

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

Codes assigned to the system by government or private institutions about the safety of travel within the system.

Worlds given a Green Zone rating are usually indicated by leaving the data field blank or `-` (dash).

|Code|Description|
|---|---|
|`R`|Red. Interdicted. Dangerous. Prohibited. Imperial.|
|`A`|Amber. Potentially dangerous. Caution advised. Imperial.|
|`G`|Green. Unrestricted. Imperial.|
|`B`|Blue. Balkanized — [code](https://travellermap.com/doc/secondsurvey#government) is dominant government. TNE (circa 1201).|
|`F`|Forbidden. Access prohibited. Zhodani.|
|`U`|Unabsorbed. Access restricted. Zhodani.|

_Imperial travel codes are provided by the Journal of the Travellers' Aid Society, and are used with permission of that publication. Worlds outside the Imperium should be considered Amber Zones by travellers from the Imperium._

## PBG - Population, Belts, Giants

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

### Population Multiplier

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

The first [eHex](https://travellermap.com/doc/secondsurvey#ehex) digit is the population multiplier. Together with the [population exponent](https://travellermap.com/doc/secondsurvey#population_exponent) in the [UWP](https://travellermap.com/doc/secondsurvey#uwp) field, the world population can be computed as:

population = multiplier × 10exponent

Some legacy files may erroneously have `0` for the population multiplier but non-zero for the population exponent. Treat the multiplier as `1` in these cases.

If unknown, `?` can be specified.

### Planetoid Belts

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

The second [eHex](https://travellermap.com/doc/secondsurvey#ehex) digit is the number of planetoid belts in the system. A Main World of size `0` is termed an asteriod belt, and is not counted here.

If unknown, `?` can be specified.

### Gas Giants

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

The third [eHex](https://travellermap.com/doc/secondsurvey#ehex) digit is the number of gas giants in the system, suitable for fuel skimming.

If unknown, `?` can be specified.

## Worlds

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

### Traveller5

The number of "worlds" in the system. This is given as a decimal integer, and will always be at least 1 (the Main World) plus the number of planetoid belts (see [PBG](https://travellermap.com/doc/secondsurvey#pbg)) plus the number of gas giants (see [PBG](https://travellermap.com/doc/secondsurvey#pbg)) but will include other planets orbiting the star(s) in the system.

Note that gas giant satellites are _not_ counted, unless the satellite is the Main World itself - identified with [remark](https://travellermap.com/doc/secondsurvey#remark) `Sa`.

## Allegiance

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

Allegiances indicate the government which dominates a system. A short abbreviation is used.

### Traveller5

The allegiance code is a 4-character code which is selected to uniquely identify a polity in Charted Space. Previous editions used a 2-character code (see below).

|Code|Pre-T5|Description|Location|
|---|---|---|---|
|3EoG|Ga|Third Empire of Gashikan|[Mend](https://travellermap.com/go/Mend)/[Gash](https://travellermap.com/go/Gash)/[Tren](https://travellermap.com/go/Tren)|
|4Wor|Fw|Four Worlds|[Farf](https://travellermap.com/go/Farf)|
|AkUn|Ak|Akeena Union|[Gate](https://travellermap.com/go/Gate)|
|AlCo|Al|Altarean Confederation|[Vang](https://travellermap.com/go/Vang)|
|AnTC|Ac|Anubian Trade Coalition|[Hint](https://travellermap.com/go/Hint)|
|AsIf|As|Iyeaao'fte|[Ustr](https://travellermap.com/go/Ustr)|
|AsMw|As|Aslan Hierate, single multiple-world clan dominates|[Akti](https://travellermap.com/go/Akti)/[Dark](https://travellermap.com/go/Dark)/[Eali](https://travellermap.com/go/Eali)/[Hlak](https://travellermap.com/go/Hlak)/[Iwah](https://travellermap.com/go/Iwah)/[Reav](https://travellermap.com/go/Reav)/[Rift](https://travellermap.com/go/Rift)/[Stai](https://travellermap.com/go/Stai)/[Troj](https://travellermap.com/go/Troj)/[Uist](https://travellermap.com/go/Uist)/[Ustr](https://travellermap.com/go/Ustr)/[Verg](https://travellermap.com/go/Verg)|
|AsOf|As|Oleaiy'fte|[Ustr](https://travellermap.com/go/Ustr)|
|AsSc|As|Aslan Hierate, multiple clans split control|[Akti](https://travellermap.com/go/Akti)/[Dark](https://travellermap.com/go/Dark)/[Eali](https://travellermap.com/go/Eali)/[Hlak](https://travellermap.com/go/Hlak)/[Iwah](https://travellermap.com/go/Iwah)/[Reav](https://travellermap.com/go/Reav)/[Rift](https://travellermap.com/go/Rift)/[Stai](https://travellermap.com/go/Stai)/[Troj](https://travellermap.com/go/Troj)/[Uist](https://travellermap.com/go/Uist)/[Ustr](https://travellermap.com/go/Ustr)|
|AsSF|As|Aslan Hierate, small facility||
|AsT0|A0|Aslan Hierate, Tlaukhu control, Yerlyaruiwo (1), Hrawoao (13), Eisohiyw (14), Ferekhearl (19)|[Akti](https://travellermap.com/go/Akti)/[Dark](https://travellermap.com/go/Dark)/[Eali](https://travellermap.com/go/Eali)/[Hlak](https://travellermap.com/go/Hlak)/[Iwah](https://travellermap.com/go/Iwah)/[Reav](https://travellermap.com/go/Reav)/[Rift](https://travellermap.com/go/Rift)/[Stai](https://travellermap.com/go/Stai)/[Troj](https://travellermap.com/go/Troj)/[Uist](https://travellermap.com/go/Uist)/[Ustr](https://travellermap.com/go/Ustr)|
|AsT1|A1|Aslan Hierate, Tlaukhu control, Khaukheairl (2), Estoieie' (16), Toaseilwi (22)|[Akti](https://travellermap.com/go/Akti)/[Dark](https://travellermap.com/go/Dark)/[Eali](https://travellermap.com/go/Eali)/[Hlak](https://travellermap.com/go/Hlak)/[Iwah](https://travellermap.com/go/Iwah)/[Reav](https://travellermap.com/go/Reav)/[Rift](https://travellermap.com/go/Rift)/[Stai](https://travellermap.com/go/Stai)/[Troj](https://travellermap.com/go/Troj)/[Uist](https://travellermap.com/go/Uist)/[Ustr](https://travellermap.com/go/Ustr)|
|AsT2|A2|Aslan Hierate, Tlaukhu control, Syoisuis (3)|[Akti](https://travellermap.com/go/Akti)/[Dark](https://travellermap.com/go/Dark)/[Eali](https://travellermap.com/go/Eali)/[Hlak](https://travellermap.com/go/Hlak)/[Iwah](https://travellermap.com/go/Iwah)/[Reav](https://travellermap.com/go/Reav)/[Rift](https://travellermap.com/go/Rift)/[Stai](https://travellermap.com/go/Stai)/[Troj](https://travellermap.com/go/Troj)/[Uist](https://travellermap.com/go/Uist)/[Ustr](https://travellermap.com/go/Ustr)|
|AsT3|A3|Aslan Hierate, Tlaukhu control, Tralyeaeawi (4), Yulraleh (12), Aiheilar (25), Riyhalaei (28)|[Akti](https://travellermap.com/go/Akti)/[Dark](https://travellermap.com/go/Dark)/[Eali](https://travellermap.com/go/Eali)/[Hlak](https://travellermap.com/go/Hlak)/[Iwah](https://travellermap.com/go/Iwah)/[Reav](https://travellermap.com/go/Reav)/[Rift](https://travellermap.com/go/Rift)/[Stai](https://travellermap.com/go/Stai)/[Troj](https://travellermap.com/go/Troj)/[Uist](https://travellermap.com/go/Uist)/[Ustr](https://travellermap.com/go/Ustr)|
|AsT4|A4|Aslan Hierate, Tlaukhu control, Eakhtiyho (5), Eteawyolei' (11), Fteweyeakh (23)|[Akti](https://travellermap.com/go/Akti)/[Dark](https://travellermap.com/go/Dark)/[Eali](https://travellermap.com/go/Eali)/[Hlak](https://travellermap.com/go/Hlak)/[Iwah](https://travellermap.com/go/Iwah)/[Reav](https://travellermap.com/go/Reav)/[Rift](https://travellermap.com/go/Rift)/[Stai](https://travellermap.com/go/Stai)/[Troj](https://travellermap.com/go/Troj)/[Uist](https://travellermap.com/go/Uist)/[Ustr](https://travellermap.com/go/Ustr)|
|AsT5|A5|Aslan Hierate, Tlaukhu control, Hlyueawi (6), Isoitiyro (15)|[Akti](https://travellermap.com/go/Akti)/[Dark](https://travellermap.com/go/Dark)/[Eali](https://travellermap.com/go/Eali)/[Hlak](https://travellermap.com/go/Hlak)/[Iwah](https://travellermap.com/go/Iwah)/[Reav](https://travellermap.com/go/Reav)/[Rift](https://travellermap.com/go/Rift)/[Stai](https://travellermap.com/go/Stai)/[Troj](https://travellermap.com/go/Troj)/[Uist](https://travellermap.com/go/Uist)/[Ustr](https://travellermap.com/go/Ustr)|
|AsT6|A6|Aslan Hierate, Tlaukhu control, Uiktawa (7), Iykyasea (17), Faowaou (27)|[Akti](https://travellermap.com/go/Akti)/[Dark](https://travellermap.com/go/Dark)/[Eali](https://travellermap.com/go/Eali)/[Hlak](https://travellermap.com/go/Hlak)/[Iwah](https://travellermap.com/go/Iwah)/[Reav](https://travellermap.com/go/Reav)/[Rift](https://travellermap.com/go/Rift)/[Stai](https://travellermap.com/go/Stai)/[Troj](https://travellermap.com/go/Troj)/[Uist](https://travellermap.com/go/Uist)/[Ustr](https://travellermap.com/go/Ustr)|
|AsT7|A7|Aslan Hierate, Tlaukhu control, Ikhtealyo (8), Tlerlearlyo (20), Yetahikh (24)|[Akti](https://travellermap.com/go/Akti)/[Dark](https://travellermap.com/go/Dark)/[Eali](https://travellermap.com/go/Eali)/[Hlak](https://travellermap.com/go/Hlak)/[Iwah](https://travellermap.com/go/Iwah)/[Reav](https://travellermap.com/go/Reav)/[Rift](https://travellermap.com/go/Rift)/[Stai](https://travellermap.com/go/Stai)/[Troj](https://travellermap.com/go/Troj)/[Uist](https://travellermap.com/go/Uist)/[Ustr](https://travellermap.com/go/Ustr)|
|AsT8|A8|Aslan Hierate, Tlaukhu control, Seieakh (9), Akatoiloh (18), We'okurir (29)|[Akti](https://travellermap.com/go/Akti)/[Dark](https://travellermap.com/go/Dark)/[Eali](https://travellermap.com/go/Eali)/[Hlak](https://travellermap.com/go/Hlak)/[Iwah](https://travellermap.com/go/Iwah)/[Reav](https://travellermap.com/go/Reav)/[Rift](https://travellermap.com/go/Rift)/[Stai](https://travellermap.com/go/Stai)/[Troj](https://travellermap.com/go/Troj)/[Uist](https://travellermap.com/go/Uist)/[Ustr](https://travellermap.com/go/Ustr)|
|AsT9|A9|Aslan Hierate, Tlaukhu control, Aokhalte (10), Sahao' (21), Ouokhoi (26)|[Akti](https://travellermap.com/go/Akti)/[Dark](https://travellermap.com/go/Dark)/[Eali](https://travellermap.com/go/Eali)/[Hlak](https://travellermap.com/go/Hlak)/[Iwah](https://travellermap.com/go/Iwah)/[Reav](https://travellermap.com/go/Reav)/[Rift](https://travellermap.com/go/Rift)/[Stai](https://travellermap.com/go/Stai)/[Troj](https://travellermap.com/go/Troj)/[Uist](https://travellermap.com/go/Uist)/[Ustr](https://travellermap.com/go/Ustr)|
|AsTA|Ta|Tealou Arlaoh|[Uist](https://travellermap.com/go/Uist)/[Ustr](https://travellermap.com/go/Ustr)|
|AsTv|As|Aslan Hierate, Tlaukhu vassal clan dominates|[Akti](https://travellermap.com/go/Akti)/[Dark](https://travellermap.com/go/Dark)/[Eali](https://travellermap.com/go/Eali)/[Hlak](https://travellermap.com/go/Hlak)/[Iwah](https://travellermap.com/go/Iwah)/[Reav](https://travellermap.com/go/Reav)/[Rift](https://travellermap.com/go/Rift)/[Stai](https://travellermap.com/go/Stai)/[Troj](https://travellermap.com/go/Troj)/[Uist](https://travellermap.com/go/Uist)/[Ustr](https://travellermap.com/go/Ustr)|
|AsTz|As|Aslan Hierate, Zodia clan|[Iwah](https://travellermap.com/go/Iwah)|
|AsVc|As|Aslan Hierate, vassal clan dominates|[Akti](https://travellermap.com/go/Akti)/[Dark](https://travellermap.com/go/Dark)/[Eali](https://travellermap.com/go/Eali)/[Hlak](https://travellermap.com/go/Hlak)/[Iwah](https://travellermap.com/go/Iwah)/[Reav](https://travellermap.com/go/Reav)/[Rift](https://travellermap.com/go/Rift)/[Stai](https://travellermap.com/go/Stai)/[Troj](https://travellermap.com/go/Troj)/[Uist](https://travellermap.com/go/Uist)/[Ustr](https://travellermap.com/go/Ustr)|
|AsWc|As|Aslan Hierate, single one-world clan dominates|[Akti](https://travellermap.com/go/Akti)/[Dark](https://travellermap.com/go/Dark)/[Eali](https://travellermap.com/go/Eali)/[Hlak](https://travellermap.com/go/Hlak)/[Iwah](https://travellermap.com/go/Iwah)/[Reav](https://travellermap.com/go/Reav)/[Rift](https://travellermap.com/go/Rift)/[Stai](https://travellermap.com/go/Stai)/[Troj](https://travellermap.com/go/Troj)/[Uist](https://travellermap.com/go/Uist)/[Ustr](https://travellermap.com/go/Ustr)|
|AsXX|As|Aslan Hierate, unknown|[Akti](https://travellermap.com/go/Akti)/[Dark](https://travellermap.com/go/Dark)/[Eali](https://travellermap.com/go/Eali)/[Hlak](https://travellermap.com/go/Hlak)/[Iwah](https://travellermap.com/go/Iwah)/[Reav](https://travellermap.com/go/Reav)/[Rift](https://travellermap.com/go/Rift)/[Stai](https://travellermap.com/go/Stai)/[Troj](https://travellermap.com/go/Troj)/[Uist](https://travellermap.com/go/Uist)/[Ustr](https://travellermap.com/go/Ustr)|
|AvCn|Ac|Avalar Consulate||
|BaCl|Bc|Backman Cluster||
|Bium|Bi|The Biumvirate|[Farf](https://travellermap.com/go/Farf)|
|BlSo|Bs|Belgardian Sojurnate|[Troj](https://travellermap.com/go/Troj)|
|BoWo|Bw|Border Worlds||
|CaAs|Cb|Carrillian Assembly|[Reav](https://travellermap.com/go/Reav)|
|CaPr|Ca|Principality of Caledon|[Reav](https://travellermap.com/go/Reav)|
|CaTe|Ct|Carter Technocracy|[Reav](https://travellermap.com/go/Reav)|
|CoAl|Ca|Corsair Alliance||
|CoBa|Ba|Confederation of Bammesuka|[Mend](https://travellermap.com/go/Mend)|
|CoLg|CL|Corellan League|[Beyo](https://travellermap.com/go/Beyo)/[Vang](https://travellermap.com/go/Vang)|
|CoLp|Lp|Council of Leh Perash|[Hint](https://travellermap.com/go/Hint)|
|CRAk|CA|Anakudnu Cultural Region||
|CRGe|CG|Geonee Cultural Region||
|CRSu|CS|Suerrat Cultural Region||
|CRVi|CV|Vilani Cultural Region||
|CsCa|Ca|Client state, Principality of Caledon|[Reav](https://travellermap.com/go/Reav)|
|CsHv|Hc|Client state, Hive Federation|[Cruc](https://travellermap.com/go/Cruc)/[Spic](https://travellermap.com/go/Spic)|
|CsIm|Cs|Client state, Third Imperium|_various_|
|CsMo|Cm|Client state, Duchy of Mora|[Spin](https://travellermap.com/go/Spin)|
|CsPt|CP|Client state, The Protectorate|[Farf](https://travellermap.com/go/Farf)|
|CsRr|Cr|Client state, Republic of Regina|[Spin](https://travellermap.com/go/Spin)|
|CsTw|KC|Client state, Two Thousand Worlds|_various_|
|CsZh|Cz|Client state, Zhodani Consulate|[Spin](https://travellermap.com/go/Spin)/[Troj](https://travellermap.com/go/Troj)|
|CyUn|Cu|Cytralin Unity|[Hint](https://travellermap.com/go/Hint)|
|DaCf|Da|Darrian Confederation|[Spin](https://travellermap.com/go/Spin)|
|DeHg|Dh|Descarothe Hegemony|[Farf](https://travellermap.com/go/Farf)|
|DeNo|Dn|Demos of Nobles|[Newo](https://travellermap.com/go/Newo)|
|DiGr|Dg|Dienbach Grüpen|[Newo](https://travellermap.com/go/Newo)|
|DoAl|Az|Domain of Alntzar|[Farf](https://travellermap.com/go/Farf)|
|DuCf|Cd|Confederation of Duncinae|[Reav](https://travellermap.com/go/Reav)|
|DuMo|Mo|Duchy of Mora|[Spin](https://travellermap.com/go/Spin)|
|ECRp|EC|Eberhardt Corporate Republic|[Krus](https://travellermap.com/go/Krus)|
|EsMa|Es|Eslyat Magistracy|[Vang](https://travellermap.com/go/Vang)|
|FdAr|Fa|Federation of Arden||
|FdDa|Fd|Federation of Daibei||
|FdIl|Fi|Federation of Ilelish||
|FeAl|Fa|Federation of Alsas|[Farf](https://travellermap.com/go/Farf)|
|FeAm|FA|Federation of Amil|[Cruc](https://travellermap.com/go/Cruc)|
|FeHe|Fh|Federation of Heron|[Glim](https://travellermap.com/go/Glim)|
|FlLe|Fl|Florian League|[Troj](https://travellermap.com/go/Troj)|
|GaFd|Ga|Galian Federation|[Gate](https://travellermap.com/go/Gate)|
|GaRp|Gr|Gamma Republic|[Glim](https://travellermap.com/go/Glim)|
|GdKa|Rm|Grand Duchy of Kalradin|[Cruc](https://travellermap.com/go/Cruc)|
|GdMh|Ma|Grand Duchy of Marlheim|[Reav](https://travellermap.com/go/Reav)|
|GdSt|Gs|Grand Duchy of Stoner|[Glim](https://travellermap.com/go/Glim)|
|GeOr|Go|Gerontocracy of Ormine|[Dark](https://travellermap.com/go/Dark)|
|GlEm|Gl|Glorious Empire|[Troj](https://travellermap.com/go/Troj)|
|GlFe|Gf|Glimmerdrift Federation|[Cruc](https://travellermap.com/go/Cruc)/[Glim](https://travellermap.com/go/Glim)|
|GnCl|Gi|Gniivi Collective|[Hint](https://travellermap.com/go/Hint)|
|HaCo|Hc|Haladon Cooperative|[Farf](https://travellermap.com/go/Farf)|
|HeCo|HC|Hefrin Colony|[Beyo](https://travellermap.com/go/Beyo)|
|HoPA|Ho|Hochiken People's Assembly|[Gate](https://travellermap.com/go/Gate)|
|HvFd|Hv|Hive Federation|[Spic](https://travellermap.com/go/Spic)|
|HyLe|Hy|Hyperion League|[Vang](https://travellermap.com/go/Vang)|
|IHPr|IS|I'Sred*Ni Protectorate|[Beyo](https://travellermap.com/go/Beyo)|
|ImAp|Im|Third Imperium, Amec Protectorate|[Dagu](https://travellermap.com/go/Dagu)|
|ImDa|Im|Third Imperium, Domain of Antares|[Anta](https://travellermap.com/go/Anta)/[Empt](https://travellermap.com/go/Empt)/[Lish](https://travellermap.com/go/Lish)|
|ImDc|Im|Third Imperium, Domain of Sylea|[Core](https://travellermap.com/go/Core)/[Delp](https://travellermap.com/go/Delp)/[Forn](https://travellermap.com/go/Forn)/[Mass](https://travellermap.com/go/Mass)|
|ImDd|Im|Third Imperium, Domain of Deneb|[Dene](https://travellermap.com/go/Dene)/[Reft](https://travellermap.com/go/Reft)/[Spin](https://travellermap.com/go/Spin)/[Troj](https://travellermap.com/go/Troj)|
|ImDg|Im|Third Imperium, Domain of Gateway|[Glim](https://travellermap.com/go/Glim)/[Hint](https://travellermap.com/go/Hint)/[Ley](https://travellermap.com/go/Ley)|
|ImDi|Im|Third Imperium, Domain of Ilelish|[Daib](https://travellermap.com/go/Daib)/[Ilel](https://travellermap.com/go/Ilel)/[Reav](https://travellermap.com/go/Reav)/[Verg](https://travellermap.com/go/Verg)/[Zaru](https://travellermap.com/go/Zaru)|
|ImDs|Im|Third Imperium, Domain of Sol|[Alph](https://travellermap.com/go/Alph)/[Dias](https://travellermap.com/go/Dias)/[Magy](https://travellermap.com/go/Magy)/[Olde](https://travellermap.com/go/Olde)/[Solo](https://travellermap.com/go/Solo)|
|ImDv|Im|Third Imperium, Domain of Vland|[Corr](https://travellermap.com/go/Corr)/[Dagu](https://travellermap.com/go/Dagu)/[Gush](https://travellermap.com/go/Gush)/[Reft](https://travellermap.com/go/Reft)/[Vlan](https://travellermap.com/go/Vlan)|
|ImLa|Im|Third Imperium, League of Antares|[Anta](https://travellermap.com/go/Anta)|
|ImLc|Im|Third Imperium, Lancian Cultural Region|[Corr](https://travellermap.com/go/Corr)/[Dagu](https://travellermap.com/go/Dagu)/[Gush](https://travellermap.com/go/Gush)|
|ImLu|Im|Third Imperium, Luriani Cultural Association|[Ley](https://travellermap.com/go/Ley)/[Forn](https://travellermap.com/go/Forn)|
|ImSy|Im|Third Imperium, Sylean Worlds|[Core](https://travellermap.com/go/Core)|
|ImVd|Ve|Third Imperium, Vegan Autonomous District|[Solo](https://travellermap.com/go/Solo)|
|InRp|Ir|Interstellar Republic|[Krus](https://travellermap.com/go/Krus)|
|IsDo|Id|Islaiat Dominate|[Eali](https://travellermap.com/go/Eali)|
|JAOz|Jo|Julian Protectorate, Alliance of Ozuvon|[Mend](https://travellermap.com/go/Mend)|
|JaPa|Ja|Jarnac Pashalic|[Beyo](https://travellermap.com/go/Beyo)/[Vang](https://travellermap.com/go/Vang)|
|JAsi|Ja|Julian Protectorate, Asimikigir Confederation|[Amdu](https://travellermap.com/go/Amdu)/[Mend](https://travellermap.com/go/Mend)|
|JCoK|Jc|Julian Protectorate, Constitution of Koekhon|[Amdu](https://travellermap.com/go/Amdu)/[Mend](https://travellermap.com/go/Mend)|
|JHhk|Jh|Julian Protectorate, Hhkar Sphere|[Amdu](https://travellermap.com/go/Amdu)/[Mend](https://travellermap.com/go/Mend)|
|JLum|Jd|Julian Protectorate, Lumda Dower|[Mend](https://travellermap.com/go/Mend)|
|JMen|Jm|Julian Protectorate, Commonwealth of Mendan|[Mend](https://travellermap.com/go/Mend)/[Gash](https://travellermap.com/go/Gash)|
|JPSt|Jp|Julian Protectorate, Pirbarish Starlane|[Mend](https://travellermap.com/go/Mend)|
|JRar|Vw|Julian Protectorate, Rar Errall/Wolves Warren|[Mend](https://travellermap.com/go/Mend)|
|JuHl|Hl|Julian Protectorate, Hegemony of Lorean|[Amdu](https://travellermap.com/go/Amdu)/[Empt](https://travellermap.com/go/Empt)/[Mend](https://travellermap.com/go/Mend)|
|JUkh|Ju|Julian Protectorate, Ukhanzi Coordinate|[Mend](https://travellermap.com/go/Mend)|
|JuNa|Jn|Jurisdiction of Nadon|[Cano](https://travellermap.com/go/Cano)|
|JuPr|Jp|Julian Protectorate|[Amdu](https://travellermap.com/go/Amdu)/[Empt](https://travellermap.com/go/Empt)/[Mend](https://travellermap.com/go/Mend)|
|JuRu|Jr|Julian Protectorate, Rukadukaz Republic|[Empt](https://travellermap.com/go/Empt)/[Mend](https://travellermap.com/go/Mend)|
|JVug|Jv|Julian Protectorate, Vugurar Dominion|[Mend](https://travellermap.com/go/Mend)|
|KaCo|KC|Katowice Conquest|[Cruc](https://travellermap.com/go/Cruc)|
|KaEm|KE|Katanga Empire|[Beyo](https://travellermap.com/go/Beyo)|
|KaTr|Kt|Kajaani Triumverate|[Vang](https://travellermap.com/go/Vang)|
|KaWo|KW|Karhyri Worlds|[Cruc](https://travellermap.com/go/Cruc)|
|KhLe|Kl|Khuur League|[Ley](https://travellermap.com/go/Ley)|
|KkTw|Kk|Two Thousand Worlds|_various_|
|KoEm|Ko|Korsumug Empire|[Thet](https://travellermap.com/go/Thet)|
|KoPm|Pm|Percavid Marches|[Thet](https://travellermap.com/go/Thet)|
|KPel|Pe|Kingdom of Peladon|[Thet](https://travellermap.com/go/Thet)|
|KrPr|Kr|Krotan Primacy|[Gzir](https://travellermap.com/go/Gzir)/[Rice](https://travellermap.com/go/Rice)/[KaaG](https://travellermap.com/go/KaaG)|
|LeSu|Ls|League of Suns|[Farf](https://travellermap.com/go/Farf)|
|LnRp|Ln|Loyal Nineworlds Republic|[Glim](https://travellermap.com/go/Glim)|
|LuIm|Li|Lucan's Imperium||
|LyCo|Ly|Lanyard Colonies|[Reav](https://travellermap.com/go/Reav)|
|MaCl|Ma|Mapepire Cluster|[Beyo](https://travellermap.com/go/Beyo)|
|MaEm|Mk|Maskai Empire|[Glim](https://travellermap.com/go/Glim)|
|MaSt|Ma|Maragaret's Domain||
|MaUn|Mu|Malorn Union|[Cano](https://travellermap.com/go/Cano)/[Alde](https://travellermap.com/go/Alde)|
|MeCo|Me|Megusard Corporate|[Gate](https://travellermap.com/go/Gate)|
|MiCo|Mi|Mische Conglomerate|[Cruc](https://travellermap.com/go/Cruc)|
|MnPr|Mn|Mnemosyne Principality|[Farf](https://travellermap.com/go/Farf)|
|MoLo|ML|Monarchy of Lod|[Beyo](https://travellermap.com/go/Beyo)|
|MrCo|MC|Mercantile Concord|[Cruc](https://travellermap.com/go/Cruc)|
|NaAs|As|Non-Aligned, Aslan-dominated|[Akti](https://travellermap.com/go/Akti)/[Dark](https://travellermap.com/go/Dark)/[Eali](https://travellermap.com/go/Eali)/[Rift](https://travellermap.com/go/Rift)/[Uist](https://travellermap.com/go/Uist)/[Ustr](https://travellermap.com/go/Ustr)|
|NaCh|Na|Non-Aligned, TBD|[Sidi](https://travellermap.com/go/Sidi)|
|NaDr|Dr|Non-Aligned, Droyne-dominated|_various_|
|NaHu|Na|Non-Aligned, Human-dominated|_various_|
|NaVa|Va|Non-Aligned, Vargr-dominated|_various_|
|NaXX|Na|Non-Aligned, unclaimed|_various_|
|NkCo|NC|Nakris Confederation|[Beyo](https://travellermap.com/go/Beyo)|
|OcWs|Ow|Outcasts of the Whispering Sky|[Hint](https://travellermap.com/go/Hint)|
|OlWo|Ow|Old Worlds|[Cruc](https://travellermap.com/go/Cruc)|
|PlLe|Pl|Plavian League|[Gate](https://travellermap.com/go/Gate)|
|PrBr|PB|Principality of Bruhkarr|[Beyo](https://travellermap.com/go/Beyo)|
|Prot|Pt|The Protectorate|[Farf](https://travellermap.com/go/Farf)|
|RamW|RW|Rammak Worlds|[Krus](https://travellermap.com/go/Krus)|
|RaRa|Ra|Ral Ranta|[Hint](https://travellermap.com/go/Hint)|
|Reac|Rh|The Reach|[Cruc](https://travellermap.com/go/Cruc)|
|ReUn|Re|Renkard Union|[Gate](https://travellermap.com/go/Gate)|
|Rule|RM|Rule of Man|[Krus](https://travellermap.com/go/Krus)|
|SaCo|Sc|Salinaikin Concordance|[Farf](https://travellermap.com/go/Farf)|
|Sark|Sc|Sarkan Constellation|[Mend](https://travellermap.com/go/Mend)|
|SeFo|Sf|Senlis Foederate|[Troj](https://travellermap.com/go/Troj)|
|SELK|Lk|Sha Elden Lith Kindriu|[Gzir](https://travellermap.com/go/Gzir)/[KaaG](https://travellermap.com/go/KaaG)|
|ShRp|SR|Stormhaven Republic|[Beyo](https://travellermap.com/go/Beyo)|
|SoBF|So|Solomani Confederation, Bootean Federation|[Solo](https://travellermap.com/go/Solo)|
|SoCf|So|Solomani Confederation|[Alph](https://travellermap.com/go/Alph)/[Diab](https://travellermap.com/go/Diab)/[Dark](https://travellermap.com/go/Dark)/[Hint](https://travellermap.com/go/Hint)/[Magy](https://travellermap.com/go/Magy)/[Olde](https://travellermap.com/go/Olde)/[Reav](https://travellermap.com/go/Reav)/[Solo](https://travellermap.com/go/Solo)/[Spic](https://travellermap.com/go/Spic)/[Ustr](https://travellermap.com/go/Ustr)|
|SoCT|So|Solomani Confederation, Consolidation of Turin|[Alph](https://travellermap.com/go/Alph)|
|SoFr|Fr|Solomani Confederation, Third Reformed French Confederate Republic|[Alde](https://travellermap.com/go/Alde)|
|SoHn|Hn|Solomani Confederation, Hanuman Systems|[Lang](https://travellermap.com/go/Lang)|
|SoKE|So|Solomani Confederation, Kruse Enclave|[Krus](https://travellermap.com/go/Krus)|
|SoKv|Kv|Solomani Confederation, Kostov Confederate Republic|[Newo](https://travellermap.com/go/Newo)|
|SoLE|So|Solomani Confederation, Lubbock Enclave|[Lubb](https://travellermap.com/go/Lubb)|
|SoNS|So|Solomani Confederation, New Slavic Solidarity|[Magy](https://travellermap.com/go/Magy)|
|SoQu|Qu|Solomani Confederation, Grand United States of Quesada|[Alde](https://travellermap.com/go/Alde)|
|SoRD|So|Solomani Confederation, Reformed Dootchen Estates|[Magy](https://travellermap.com/go/Magy)|
|SoRz|So|Solomani Confederation, Restricted Zone|[Alde](https://travellermap.com/go/Alde)/[Newo](https://travellermap.com/go/Newo)|
|Sovr|Sv|The Sovereignty|[Krus](https://travellermap.com/go/Krus)|
|SoWu|So|Solomani Confederation, Wuan Technology Association|[Diab](https://travellermap.com/go/Diab)/[Magy](https://travellermap.com/go/Magy)|
|SoXE|So|Solomani Confederation, Xuanzang Enclave|[Xuan](https://travellermap.com/go/Xuan)|
|StCl|Sc|Strend Cluster|[Troj](https://travellermap.com/go/Troj)|
|StIm|St|Strephon's Worlds||
|SwCf|Sw|Sword Worlds Confederation|[Spin](https://travellermap.com/go/Spin)|
|SwFW|Sw|Swanfei Free Worlds|[Gate](https://travellermap.com/go/Gate)|
|SyRe|Sy|Syzlin Republic|[Cruc](https://travellermap.com/go/Cruc)|
|TeCl|Tc|Tellerian Cluster|[Vang](https://travellermap.com/go/Vang)|
|TrBr|Tb|Trita Brotherhood|[Cano](https://travellermap.com/go/Cano)|
|TrCo|Tr|Trindel Confederacy|[Gate](https://travellermap.com/go/Gate)|
|TrDo|Td|Trelyn Domain|[Vang](https://travellermap.com/go/Vang)/[Farf](https://travellermap.com/go/Farf)|
|TroC|Tr|Trooles Confederation|[Thet](https://travellermap.com/go/Thet)|
|UnGa|Ug|Union of Garth|[Farf](https://travellermap.com/go/Farf)|
|UnHa|Uh|Union of Harmony|[Dark](https://travellermap.com/go/Dark)/[Reav](https://travellermap.com/go/Reav)|
|V17D|V7|17th Disjuncture|[Mesh](https://travellermap.com/go/Mesh)/[Wind](https://travellermap.com/go/Wind)|
|V40S|Ve|40th Squadron|[Gvur](https://travellermap.com/go/Gvur)|
|VA16|V6|Assemblage of 1116||
|VAkh|VA|Akhstuti|[Tugl](https://travellermap.com/go/Tugl)|
|VAnP|Vx|Antares Pact|[Mesh](https://travellermap.com/go/Mesh)/[Mend](https://travellermap.com/go/Mend)|
|VARC|Vr|Anti-Rukh Coalition|[Gvur](https://travellermap.com/go/Gvur)|
|VAsP|Vx|Ascendancy Pact|[Knoe](https://travellermap.com/go/Knoe)|
|VAug|Vu|United Followers of Augurgh|[Dene](https://travellermap.com/go/Dene)/[Tugl](https://travellermap.com/go/Tugl)|
|VBkA|Vb|Bakne Alliance|[Tugl](https://travellermap.com/go/Tugl)|
|VCKd|Vk|Commonality of Kedzudh|[Gvur](https://travellermap.com/go/Gvur)|
|VDeG|Vd|Democracy of Greats|[Knoe](https://travellermap.com/go/Knoe)|
|VDrN|VN|Drr'lana Network|[Gash](https://travellermap.com/go/Gash)|
|VDzF|Vf|Dzarrgh Federate|[Dene](https://travellermap.com/go/Dene)/[Prov](https://travellermap.com/go/Prov)/[Tugl](https://travellermap.com/go/Tugl)|
|VFFD|V1|First Fleet of Dzo|[Mesh](https://travellermap.com/go/Mesh)|
|VGoT|Vg|Glory of Taarskoerzn|[Prov](https://travellermap.com/go/Prov)|
|ViCo|Vi|Viyard Concourse|[Gate](https://travellermap.com/go/Gate)|
|VInL|V9|Infinity League|[Knoe](https://travellermap.com/go/Knoe)|
|VIrM|Vh|Irrgh Manifest|[Prov](https://travellermap.com/go/Prov)|
|VJoF|Vj|Jihad of Faarzgaen|[Prov](https://travellermap.com/go/Prov)|
|VKfu|Vk|Kfue|[Tugl](https://travellermap.com/go/Tugl)|
|VLIn|Vi|Llaeghskath Interacterate|[Prov](https://travellermap.com/go/Prov)/[Tugl](https://travellermap.com/go/Tugl)|
|VLPr|Vl|Lair Protectorate|[Prov](https://travellermap.com/go/Prov)|
|VNgC|Vn|Ngath Confederation|[Wind](https://travellermap.com/go/Wind)|
|VNoe|VN|Noefa|[Tugl](https://travellermap.com/go/Tugl)|
|VOpA|Vo|Opposition Alliance|[Knoe](https://travellermap.com/go/Knoe)|
|VOpp|Vo|Opposition Alliance|[Mesh](https://travellermap.com/go/Mesh)|
|VOuz|VO|Ouzvothon|[Tugl](https://travellermap.com/go/Tugl)|
|VPGa|Vg|Pact of Gaerr|[Gvur](https://travellermap.com/go/Gvur)|
|VRo5|V5|Ruler of Five|[Mesh](https://travellermap.com/go/Mesh)|
|VRrS|VW|Rranglloez Stronghold|[Tugl](https://travellermap.com/go/Tugl)|
|VRuk|Vn|Worlds of Leader Rukh|[Gvur](https://travellermap.com/go/Gvur)|
|VSDp|Vs|Saeknouth Dependency|[Gvur](https://travellermap.com/go/Gvur)|
|VSEq|Vd|Society of Equals|[Gvur](https://travellermap.com/go/Gvur)/[Tugl](https://travellermap.com/go/Tugl)|
|VThE|Vt|Thoengling Empire|[Gvur](https://travellermap.com/go/Gvur)/[Tugl](https://travellermap.com/go/Tugl)|
|VTrA|VT|Trae Aggregation|[Tren](https://travellermap.com/go/Tren)|
|VTzE|Vp|Thirz Empire|[Gvur](https://travellermap.com/go/Gvur)/[Ziaf](https://travellermap.com/go/Ziaf)|
|VUru|Vu|Urukhu|[Gvur](https://travellermap.com/go/Gvur)|
|VVar|Ve|Empire of Varroerth|[Prov](https://travellermap.com/go/Prov)/[Tugl](https://travellermap.com/go/Tugl)/[Wind](https://travellermap.com/go/Wind)|
|VVoS|Vv|Voekhaeb Society|[Mesh](https://travellermap.com/go/Mesh)|
|VWan|Vw|People of Wanz|[Tugl](https://travellermap.com/go/Tugl)|
|VWP2|V2|Windhorn Pact of Two|[Tugl](https://travellermap.com/go/Tugl)|
|VYoe|VQ|Union of Yoetyqq|[Gash](https://travellermap.com/go/Gash)|
|WiDe|Wd|Winston Democracy|[Alde](https://travellermap.com/go/Alde)/[Newo](https://travellermap.com/go/Newo)|
|Wild|Wi|Wilds|_various_|
|XXXX|Xx|Unknown|_various_|
|ZePr|Zp|Zelphic Primacy|[Farf](https://travellermap.com/go/Farf)|
|ZhAx|Ax|Zhodani Consulate, Addaxur Reserve|[Tien](https://travellermap.com/go/Tien)|
|ZhCa|Ca|Zhodani Consulate, Colonnade Province|[Vang](https://travellermap.com/go/Vang)/[Farf](https://travellermap.com/go/Farf)|
|ZhCh|Zh|Zhodani Consulate, Chtierabl Province|[Chti](https://travellermap.com/go/Chti)|
|ZhCo|Zh|Zhodani Consulate|_various_|
|ZhIa|Zh|Zhodani Consulate, Iabrensh Province|[Stia](https://travellermap.com/go/Stia)/[Zdie](https://travellermap.com/go/Zdie)|
|ZhIN|Zh|Zhodani Consulate, Iadr Nsobl Province|[Farf](https://travellermap.com/go/Farf)/[Fore](https://travellermap.com/go/Fore)/[Gvur](https://travellermap.com/go/Gvur)/[Spin](https://travellermap.com/go/Spin)/[Yikl](https://travellermap.com/go/Yikl)/[Ziaf](https://travellermap.com/go/Ziaf)|
|ZhJp|Zh|Zhodani Consulate, Jadlapriants Province|[Tien](https://travellermap.com/go/Tien)/[Zhda](https://travellermap.com/go/Zhda)|
|ZhMe|Zh|Zhodani Consulate, Meqlemianz Province|[Eiap](https://travellermap.com/go/Eiap)/[Sidi](https://travellermap.com/go/Sidi)/[Eiap](https://travellermap.com/go/Eiap)|
|ZhOb|Zh|Zhodani Consulate, Obrefripl Province|_various_|
|ZhSh|Zh|Zhodani Consulate, Shtochiadr Province|[Itvi](https://travellermap.com/go/Itvi)/[Tlab](https://travellermap.com/go/Tlab)|
|ZhVQ|Zh|Zhodani Consulate, Vlanchiets Qlom Province|_various_|
|ZiSi|Rv|Restored Vilani Imperium||
|Zuug|Zu|Zuugabish Tripartite|[Mend](https://travellermap.com/go/Mend)|
|ZyCo|Zc|Zydarian Codominium|[Beyo](https://travellermap.com/go/Beyo)|

### Classic Traveller, MegaTraveller, Traveller: The New Era, Traveller: 4th Edition

Older versions of Traveller use a 2-character code. Due to the limited number of useful 2-character codes, these are not unique across Charted Space. In addition to the secondary entries in the above table, the following codes are seen:

|Code|Description|Edition|
|---|---|---|
|`Dr`|Droyne world.||
|`Dd`|Domain of Deneb|MegaTraveller.|
|`Fd`|Federation of Daibei|MegaTraveller.|
|`Fi`|Federation of Ilelish|MegaTraveller.|
|`La`|League of Antares|MegaTraveller.|
|`Li`|Lucan's Imperium|MegaTraveller.|
|`Ma`|Margaret's Stronghold|MegaTraveller.|
|`Rv`|Restored Vilani Empire/Ziru Sirka|MegaTraveller.|
|`St`|Strephon's Imperium|MegaTraveller.|
|`Wi`|Wilds|Traveller: The New Era.|
|`--`|Empty, unclaimed system.||

## Stellar Data

`1910 Regina A788899-C Ht Ri Pa Ph An Cp (Amindii)2 { 4 } (D7E+5) [9C6D] BcCeF NS - 703 8 ImDd F7 V BD M3 V`

Stellar data lists the stars in the system.

Normal stars fuse hydrogen, and occur in a variety of colors (due to temperature) and sizes (sub-dwarfs to super-giants). After a dwarf star exhausts its fuel it will briefly become a cool red giant then leave behind its core as an incredibly dense and bright white dwarf. A giant star will explode and leave behind an even denser neutron star or black hole. An object too small to fuse hydrogen but large enough to fuse deuterium is known as a brown dwarf.

Normal stars are described by their spectral class and luminosity. This uses a subset of the [Morgan–Keenan classification system](https://en.wikipedia.org/wiki/Stellar_classification) (also known as Yerkes or MKK), such as `G2 V` where `G2` is the color/temperature (2⁄10th of the way from yellow `G` to orange `K` on the spectrum) and `V` is the luminosity/size (main sequence).

### Stellar Spectral Class

|Code|Description|Temperature (K)|
|---|---|---|
|`O`|Blue|>33,000|
|`B`|Blue-White|10,000-33,000|
|`A`|Blue-White|7,500-10,000|
|`F`|Yellow-White|6,000-7,500|
|`G`|Yellow|5,200-6,000|
|`K`|Orange|3,700-5,200|
|`M`|Red|2,000-3,700|

Stellar type indicates the spectral classification of a star; colors are the essential perceived colors of the star's visible light.

### Stellar Luminosity

The luminosity gives the magnitude, or brightness of the star. This correlates with the size of the star, although it varies depending on spectral class.

|Code|Description|Diameter (Sol = 1)|
|---|---|---|
|`Ia`|Bright Supergiant.|52 - 3500|
|`Ib`|Weak Supergiant.|30 - 3000|
|`II`|Bright Giant.|14 - 1000|
|`III`|Normal Giant.|4.6 - 360|
|`IV`|Subgiant.|3.3 - 13|
|`V`|Main Sequence Star.|0.2 - 10|
|`VI`|Subdwarf.|0.1 - 1.2|
|`D`|White Dwarf.|0.006 - 0.018|

### Classic Traveller, MegaTraveller, Traveller: The New Era, Traveller: 4th Edition

White dwarfs (the cores of dead stars) are represented by the size code `D`, optionally followed by a spectral code, e.g. `DB`, `DA`, `DF`, `DG`, `DK`, `DM`. _(Note: This differs significantly from the real-world classification of white dwarfs.)_

Some files may use luminosity code `VII` for white dwarfs instead.

The special code `BH` may be used for black holes.

### Traveller5

White dwarfs are represented by the size code `D` with no spectral code.

The special code `BD` may be used for brown dwarf, sub-stellar objects.

The special code `BH` may be used for black holes.

The special code `NS` may be used for neutron stars. The special code `PSR` may be used for pulsars, rapidly spinning neutron stars with active radio beams.

## Resource Units

### Traveller5

Resource Units (RU) provide a summary of the economic activity of a system, and are provided in some data files. RU is calculated by multiplying together each member of the [Economic Extension](https://travellermap.com/doc/secondsurvey#ex), treating any `0` as a `1`.

RU = R × L × I × E

For example, Regina's [Economic Extension](https://travellermap.com/doc/secondsurvey#ex) is `(D7E+5)`, so its RU is 13 × 7 × 14 × +5 = 6370.

## eHex - Extended Hexadecimal

Traveller encodes numbers using "extended hexadecimal", which is really a way of encoding numbers larger than 9 into a single printable character. Numbers in the range 0-9 are represented as `0`-`9` as normal. Like hexadecimal, values in the range 10-15 are represented as `A`-`F` (always uppercase). But the encoding continues for values in the range 16-17 as `G`-`H`, 18-22 as `J`-`N`, and 23-30 as `P`-`W`.

Codes `X`, `Y` and `Z` may be used for 31-33 or may be reserved for exceptional values. Code `X`, in particular, is reserved for _unknown_ values, such as the details of intentionally unexplored systems. Codes `I` and `O` are unused to avoid confusion with `1` and `0`.

In fields such as [UWP](https://travellermap.com/doc/secondsurvey#uwp) and [PBG](https://travellermap.com/doc/secondsurvey#pbg) you can use `?` to represent unknown values.

|eHex|Decimal|
|---|---|
|`0`|0|
|`1`|1|
|`2`|2|
|`3`|3|
|`4`|4|
|`5`|5|
|`6`|6|
|`7`|7|
|`8`|8|
|`9`|9|
|`A`|10|
|`B`|11|
|`C`|12|
|`D`|13|
|`E`|14|
|`F`|15|
|`G`|16|
|`H`|17|
|`J`|18|
|`K`|19|
|`L`|20|
|`M`|21|
|`N`|22|
|`P`|23|
|`Q`|24|
|`R`|25|
|`S`|26|
|`T`|27|
|`U`|28|
|`V`|29|
|`W`|30|
|`X`|31*|
|`Y`|32*|
|`Z`|33*|

The _Traveller_ game in all forms is owned by Mongoose Publishing. Copyright 1977 – 2024 Mongoose Publishing. [Fair Use Policy](https://cdn.shopify.com/s/files/1/0609/6139/0839/files/Traveller_Fair_Use_Policy_2024.pdf?v=1725357857)