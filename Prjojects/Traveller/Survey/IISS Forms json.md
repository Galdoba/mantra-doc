---
updated_at: 2026-03-13T14:23:27.286+10:00
---
```go
package iissforms

// Form0398D0 соответствует IISS Class 0 Sector Survey (FORM 0398D-0)
type Form0398D0 struct {
	SectorGrid      string `json:"Sector | Grid,omitempty"`
	InitialSurvey   string `json:"Initial Survey,omitempty"`
	LastUpdated     string `json:"Last Updated,omitempty"`
	IISSDesignation string `json:"IISS Designation,omitempty"`
	StellarObjects  struct {
		Location string `json:"Location,omitempty"`
		Primary  string `json:"Primary,omitempty"`
		PrimaryP string `json:"Primary+,omitempty"`
		Close    string `json:"Close,omitempty"`
		CloseP   string `json:"Close+,omitempty"`
		Near     string `json:"Near,omitempty"`
		NearP    string `json:"Near+,omitempty"`
		Far      string `json:"Far,omitempty"`
		FarP     string `json:"Far+,omitempty"`
		GG       int    `json:"GG,omitempty"`
		Notes    string `json:"Notes,omitempty"`
	} `json:"Stellar objects,omitempty"`
	Comments string `json:"Comments,omitempty"`
}

// Form0421B0I соответствует IISS Class 0/I Sector Survey (FORM 0421B-0I)
type Form0421B0I struct {
	SectorGrid      string  `json:"Sector | Grid,omitempty"`
	InitialSurvey   string  `json:"Initial Survey,omitempty"`
	LastUpdated     string  `json:"Last Updated,omitempty"`
	IISSDesignation string  `json:"IISS Designation,omitempty"`
	SystemAge       float64 `json:"System Age (Gyr),omitempty"`
	Objects         struct {
		Stellar            int    `json:"Stellar,omitempty"`
		PlanetaryDetections string `json:"Planetary Detections,omitempty"` // может быть "n/a"
		Class1Status       bool   `json:"Class 1 Status,omitempty"`
	} `json:"Objects,omitempty"`
	Stars    []Star0421B0I `json:"Stars,omitempty"`
	Comments string        `json:"Comments,omitempty"`
}

type Star0421B0I struct {
	StarCharacteristics struct {
		StarDesignation string  `json:"Star Designation,omitempty"`
		Component       string  `json:"Component,omitempty"`
		Class           string  `json:"Class,omitempty"`
		Mass            float64 `json:"Mass,omitempty"`
		Temp            int     `json:"Temp,omitempty"`
		Diameter        float64 `json:"Diameter,omitempty"`
		Luminocity      float64 `json:"Luminocity,omitempty"`
		OrbitNum        float64 `json:"Orbit#,omitempty"`
		AU              float64 `json:"AU,omitempty"`
		Ecc             float64 `json:"Ecc,omitempty"`
		Period          string  `json:"Period,omitempty"`
		HZCO            float64 `json:"HZCO,omitempty"`
	} `json:"Star Characteristics,omitempty"`
	Notes string `json:"Notes,omitempty"`
}

// Form0421BIIIII соответствует IISS Class II/III Survey (FORM 0421B-II.III)
type Form0421BIIIII struct {
	SectorGrid      string  `json:"Sector | Grid,omitempty"`
	InitialSurvey   string  `json:"Initial Survey,omitempty"`
	LastUpdated     string  `json:"Last Updated,omitempty"`
	IISSDesignation string  `json:"IISS Designation,omitempty"`
	SystemAge       float64 `json:"System Age (Gyr),omitempty"`
	ObjectCounts    struct {
		Stellar         int  `json:"Stellar,omitempty"`
		GasGiants       int  `json:"Gas Giants,omitempty"`
		PlanetoidBelts  int  `json:"Planetoid Belts,omitempty"`
		Terrestrials    int  `json:"Terrestrials,omitempty"`
		ClassIIIStatus  bool `json:"Class III Status,omitempty"`
	} `json:"ObjectCounts,omitempty"`
	Stars    []Star0421BIIIII `json:"Stars,omitempty"`
	Objects  []PlanetObject   `json:"Objects,omitempty"`
	Comments string           `json:"Comments,omitempty"`
}

type Star0421BIIIII struct {
	StarDesignation string  `json:"Star Designation,omitempty"`
	Component       string  `json:"Component,omitempty"`
	Class           string  `json:"Class,omitempty"`
	Mass            float64 `json:"Mass,omitempty"`
	Temp            int     `json:"Temp,omitempty"`
	Diameter        float64 `json:"Diameter,omitempty"`
	Luminocity      float64 `json:"Luminocity,omitempty"`
	OrbitNum        float64 `json:"Orbit#,omitempty"`
	AU              float64 `json:"AU,omitempty"`
	Ecc             float64 `json:"Ecc,omitempty"`
	Period          string  `json:"Period,omitempty"`
	MAO             string  `json:"MAO,omitempty"`
	HZCO            float64 `json:"HZCO,omitempty"`
}

type PlanetObject struct {
	ObjectName       string  `json:"Object name,omitempty"`
	Primary          string  `json:"Primary,omitempty"`
	ObjectDesignation string  `json:"Object Designation,omitempty"`
	OrbitNum         int     `json:"Orbit#,omitempty"`
	AU               float64 `json:"AU,omitempty"`
	Ecc              float64 `json:"Ecc,omitempty"`
	Period           string  `json:"Period,omitempty"`
	SAHUWP           string  `json:"SAH/UWP,omitempty"`
	Sub              string  `json:"Sub,omitempty"`
	Notes            string  `json:"Notes,omitempty"`
}

// Form0407KIVPartPB соответствует IISS Class IV Survey (FORM 0407K-IV PART P.B)
type Form0407KIVPartPB struct {
	World            string `json:"World,omitempty"`
	SAHUWP           string `json:"SAH/UWP,omitempty"`
	SectorLocation   string `json:"Sector | Location,omitempty"`
	InitialSurvey    string `json:"Initial Survey,omitempty"`
	LastUpdated      string `json:"Last Updated,omitempty"`
	PrimaryObjects   string `json:"Primary Object(s),omitempty"`
	SystemAge        string `json:"System Age,omitempty"`
	TravelZone       string `json:"Travel Zone,omitempty"`
	BeltComposition  struct {
		MTypePercent string `json:"m-type(%),omitempty"`
		STypePercent string `json:"s-type(%),omitempty"`
		CTypePercent string `json:"c-type(%),omitempty"`
		OtherPercent string `json:"other(%),omitempty"`
		Bulk         string `json:"Bulk,omitempty"`
	} `json:"Belt Composition,omitempty"`
	MajorBodiesSize1 []string `json:"Size 1,omitempty"`
	MajorBodiesSizeS []string `json:"Size S,omitempty"`
	Notes            string   `json:"Notes,omitempty"`
	Resources        struct {
		Rating string `json:"Rating,omitempty"`
		Notes  string `json:"Notes,omitempty"`
	} `json:"Resources,omitempty"`
	MajorBodies []MajorBodyDetail `json:"Major Bodies,omitempty"`
	Comments    string            `json:"Comments,omitempty"`
}

type MajorBodyDetail struct {
	Body struct {
		BodyDesignation string  `json:"Body designation,omitempty"`
		SAHUWP          string  `json:"SAH/UWP,omitempty"`
		OrbitNum        int     `json:"Orbit#,omitempty"`
		OrbitAU         float64 `json:"Orbit (AU),omitempty"`
		Ecc             float64 `json:"Ecc,omitempty"`
		Period          string  `json:"Period,omitempty"`
		Type            string  `json:"Type,omitempty"`
		Diameter        float64 `json:"Diameter,omitempty"`
		Density         float64 `json:"Density,omitempty"`
		Mass            float64 `json:"Mass,omitempty"`
	} `json:"Body,omitempty"`
	Notes string `json:"Notes,omitempty"`
}

// Form0407KIVPartP соответствует IISS Class IV Survey (FORM 0407K-IV PART P)
type Form0407KIVPartP struct {
	World            string `json:"World,omitempty"`
	SAHUWP           string `json:"SAH/UWP,omitempty"`
	SectorLocation   string `json:"Sector | Location,omitempty"`
	InitialSurvey    string `json:"Initial Survey,omitempty"`
	LastUpdated      string `json:"Last Updated,omitempty"`
	PrimaryObjects   string `json:"Primary Object(s),omitempty"`
	SystemAge        string `json:"System Age,omitempty"`
	TravelZone       string `json:"Travel Zone,omitempty"`
	Orbit            struct {
		ONum   int     `json:"O#,omitempty"`
		AU     float64 `json:"AU,omitempty"`
		Ecc    float64 `json:"Ecc,omitempty"`
		Period string  `json:"Period,omitempty"`
		Notes  string  `json:"Notes,omitempty"`
	} `json:"Orbit,omitempty"`
	Size struct {
		Diameter    float64 `json:"Diameter,omitempty"`
		Composition string  `json:"Composition,omitempty"`
		Density     float64 `json:"Density,omitempty"`
		Gravity     float64 `json:"Gravity,omitempty"`
		Mass        float64 `json:"Mass,omitempty"`
		EscV        string  `json:"Esc v (kps),omitempty"`
		Notes       string  `json:"Notes,omitempty"`
	} `json:"Size,omitempty"`
	Atmosphere struct {
		Pressure    float64 `json:"Pressure (bar),omitempty"`
		Composition string  `json:"Composition,omitempty"`
		Oxygen      float64 `json:"Oxygen (bar),omitempty"`
		Taints      string  `json:"Taints,omitempty"`
		ScaleHeight string  `json:"Scale Height,omitempty"`
		Notes       string  `json:"Notes,omitempty"`
	} `json:"Atmosphere,omitempty"`
	Hydrographics struct {
		Coverage     string `json:"Coverage (%),omitempty"`
		Composition  string `json:"Composition,omitempty"`
		Distribution string `json:"Distribution,omitempty"`
		MajorBodies  string `json:"Major bodies,omitempty"`
		MinorBodies  string `json:"Minor bodies,omitempty"`
		Other        string `json:"Other,omitempty"`
		Notes        string `json:"Notes,omitempty"`
	} `json:"Hydrographics,omitempty"`
	Rotation struct {
		Sidereal      string  `json:"Sidereal,omitempty"`
		Solar         string  `json:"Solar,omitempty"`
		SolarDaysYear float64 `json:"Solar days/year,omitempty"`
		AxialTilt     string  `json:"Axial Tilt,omitempty"`
		TidalLock     bool    `json:"Tidal Lock,omitempty"`
		Tides         string  `json:"Tides,omitempty"`
		Notes         string  `json:"Notes,omitempty"`
	} `json:"Rotation,omitempty"`
	Temperature struct {
		High            string  `json:"High,omitempty"`
		Mean            string  `json:"Mean,omitempty"`
		Low             string  `json:"Low,omitempty"`
		Luminosity      float64 `json:"Luminosity,omitempty"`
		Albedo          float64 `json:"Albedo,omitempty"`
		Greenhouse      float64 `json:"Greenhouse,omitempty"`
		Notes           string  `json:"Notes,omitempty"`
		SeismicStress   string  `json:"Seismic Stress,omitempty"`
		ResidualStress  string  `json:"Residual Stress,omitempty"`
		TidalStress     string  `json:"Tidal Stress,omitempty"`
		TidalHeating    string  `json:"Tidal Heating,omitempty"`
		MajorTectonicPlates string `json:"Major Tectonic Plates,omitempty"`
	} `json:"Temperature,omitempty"`
	Life struct {
		Biomass        string `json:"Biomass,omitempty"`
		Biocomplexity  string `json:"Biocomplexity,omitempty"`
		Sophonts       string `json:"Sophonts,omitempty"`
		Biodiversity   string `json:"Biodiversity,omitempty"`
		Compatibility  string `json:"Compatibility,omitempty"`
		Notes          string `json:"Notes,omitempty"`
	} `json:"Life,omitempty"`
	Resources struct {
		Rating string `json:"Rating,omitempty"`
		Notes  string `json:"Notes,omitempty"`
	} `json:"Resources,omitempty"`
	Habitability struct {
		Rating string `json:"Rating,omitempty"`
		Notes  string `json:"Notes,omitempty"`
	} `json:"Habitability,omitempty"`
	Subordinated struct {
		Satellites []Satellite `json:"Satelite,omitempty"` // опечатка в оригинале "Satelite"
		Notes      string      `json:"Notes,omitempty"`
	} `json:"Subordinated,omitempty"`
	Comments string `json:"Comments,omitempty"`
}

type Satellite struct {
	Designation string  `json:"Designation,omitempty"`
	SAHUWP      string  `json:"SAH/UWP,omitempty"`
	OrbitPD     string  `json:"Orbit (PD),omitempty"`
	OrbitKm     float64 `json:"Orbit (km),omitempty"`
	Ecc         float64 `json:"Ecc,omitempty"`
	Diameter    float64 `json:"Diameter,omitempty"`
	Density     float64 `json:"Density,omitempty"`
	Mass        float64 `json:"Mass,omitempty"`
	PeriodHours float64 `json:"Period (h),omitempty"`
	SizeDeg     float64 `json:"Size(°),omitempty"`
}

// Form0407FIVPartC соответствует IISS Class IV Survey (FORM 0407F-IV PART C)
type Form0407FIVPartC struct {
	World          string `json:"World,omitempty"`
	UWP            string `json:"UWP,omitempty"`
	SectorLocation string `json:"Sector | Location,omitempty"`
	InitialSurvey  string `json:"Initial Survey,omitempty"`
	LastUpdated    string `json:"Last Updated,omitempty"`
	PrimaryObjects string `json:"Primary Object(s),omitempty"`
	SystemAge      string `json:"System Age,omitempty"`
	TravelZone     string `json:"Travel Zone,omitempty"`
	Population     struct {
		Total          string        `json:"Total,omitempty"`
		Demographics   string        `json:"Demographics,omitempty"`
		PCR            string        `json:"PCR,omitempty"`
		Urbanisation   string        `json:"Urbanisation%,omitempty"`
		MajorCities    int           `json:"Major Cities,omitempty"`
		CapitalPort    string        `json:"Capital/Port,omitempty"`
		MajorCitiesList []CityProfile `json:"Major Cities List,omitempty"`
		Notes          string        `json:"Notes,omitempty"`
	} `json:"Population,omitempty"`
	Government struct {
		Type          string     `json:"Type,omitempty"`
		Centralization string     `json:"Centralization,omitempty"`
		Authority     string     `json:"Authority,omitempty"`
		Profile       string     `json:"Profile,omitempty"`
		Notes         string     `json:"Notes,omitempty"`
		Factions      []FactionInfo `json:"Factions,omitempty"`
	} `json:"Government,omitempty"`
	LawLevel struct {
		Primary               string `json:"Primary,omitempty"`
		Secondary             string `json:"Secondary,omitempty"`
		Uniformity            string `json:"Uniformity,omitempty"`
		PresumptionOfInnocence bool   `json:"Presumption of Innocence,omitempty"`
		DeathPenalty          bool   `json:"Death Penalty,omitempty"`
		Categories            struct {
			Overall       string `json:"Overall,omitempty"`
			Weapons       string `json:"Weapons,omitempty"`
			Economics     string `json:"Economics,omitempty"`
			Criminal      string `json:"Criminal,omitempty"`
			Private       string `json:"Private,omitempty"`
			PersonalRights string `json:"Personal Rights,omitempty"`
		} `json:"Categories,omitempty"`
		Notes string `json:"Notes,omitempty"`
	} `json:"Law Level,omitempty"`
	Technology struct {
		CommonHigh      string `json:"Common High,omitempty"`
		CommonLow       string `json:"Common Low,omitempty"`
		Energy          string `json:"Energy,omitempty"`
		Electronics     string `json:"Electronics,omitempty"`
		Manufacturing   string `json:"Manufacturing,omitempty"`
		Medical         string `json:"Medical,omitempty"`
		Land            string `json:"Land,omitempty"`
		Water           string `json:"Water,omitempty"`
		Air             string `json:"Air,omitempty"`
		Space           string `json:"Space,omitempty"`
		PersonalMilitary string `json:"Personal Military,omitempty"`
		HeavyMilitary   string `json:"Heavy Military,omitempty"`
		Novelty         string `json:"Novelty,omitempty"`
		Environmental   string `json:"Environmental,omitempty"`
		Notes           string `json:"Notes,omitempty"`
	} `json:"Technology,omitempty"`
	Culture struct {
		Diversity      string `json:"Diversity,omitempty"`
		Cohesion       string `json:"Cohesion,omitempty"`
		Xenophilia     string `json:"Xenophilia,omitempty"`
		Progressiveness string `json:"Proggressiveness,omitempty"` // опечатка в оригинале
		Uniqueness     string `json:"Uniquiness,omitempty"`
		Expansionism   string `json:"Expansionism,omitempty"`
		Symbology      string `json:"Symbology,omitempty"`
		Militancy      string `json:"Militancy,omitempty"`
		Notes          string `json:"Notes,omitempty"`
	} `json:"Culture,omitempty"`
	Economics struct {
		TradeCodes      string  `json:"Trade Codes,omitempty"`
		Importance      string  `json:"Importance,omitempty"`
		Resources       string  `json:"Resources,omitempty"`
		Labour          string  `json:"Labour,omitempty"`
		Infrastructure  string  `json:"Infrastructure,omitempty"`
		Efficiency      string  `json:"Efficiency,omitempty"`
		GWPPerCapita    float64 `json:"GWP per capita,omitempty"`
		WTN             float64 `json:"WTN,omitempty"`
		InequalityRating string  `json:"Inequality Rating,omitempty"`
		DevelopmentScore string  `json:"Development Score,omitempty"`
		GWP             float64 `json:"GWP (MCr),omitempty"`
		Tariffs         string  `json:"Tarrifs,omitempty"` // опечатка
		Notes           string  `json:"Notes,omitempty"`
	} `json:"Economics,omitempty"`
	Starport struct {
		Class               string `json:"Class,omitempty"`
		Highport            string `json:"Highport,omitempty"`
		ExpectedWeeklyTraffic string `json:"Expected Weekly Traffic,omitempty"`
		BerthingFees        string `json:"Berthing Fees,omitempty"`
		Docking             string `json:"Docking,omitempty"`
		Shipyard            string `json:"Shipyard,omitempty"`
		AnnualOutput        string `json:"Anual Output,omitempty"`
		NavyBase            string `json:"Navy Base,omitempty"`
		ScoutBase           string `json:"Scout Base,omitempty"`
		MilitaryBase        string `json:"Military Base,omitempty"`
		Other               string `json:"Other,omitempty"`
		Notes               string `json:"Notes,omitempty"`
	} `json:"Starport,omitempty"`
	Military struct {
		EffectiveBudgetPercent string `json:"Effective Budget %,omitempty"`
		Structure              string `json:"Structure,omitempty"`
		Enforcement            string `json:"Enforcement,omitempty"`
		Militia                string `json:"Militia,omitempty"`
		Army                   string `json:"Army,omitempty"`
		WetNavy                string `json:"Wet Navy,omitempty"`
		AirForce               string `json:"Air Force,omitempty"`
		SystemDefence          string `json:"System Defence,omitempty"`
		Navy                   string `json:"Navy,omitempty"`
		Marines                string `json:"Marines,omitempty"`
		Notes                  string `json:"Notes,omitempty"`
	} `json:"Military,omitempty"`
	Comments string `json:"Comments,omitempty"`
}

// CityProfile может быть простой строкой или структурой. Для простоты используем строку.
type CityProfile string

type FactionInfo struct {
	Faction struct {
		Profile     string `json:"Profile,omitempty"`
		Designation string `json:"Designation,omitempty"`
		Description string `json:"Description,omitempty"`
	} `json:"Faction,omitempty"`
	Relationships string `json:"Relationships,omitempty"`
	FactionNotes  string `json:"Faction notes,omitempty"`
}

// Form0407FIVPartCSPF соответствует IISS Class IV Survey (FORM 0407F-IV PART CSPF)
type Form0407FIVPartCSPF struct {
	Subunit        string `json:"Subunit,omitempty"`
	PGL            string `json:"PGL,omitempty"`
	World          string `json:"World,omitempty"`
	UWP            string `json:"UWP,omitempty"`
	Description    string `json:"Description,omitempty"`
	Population     struct {
		Total        string `json:"Total,omitempty"`
		Demographics string `json:"Demographics,omitempty"`
		PCR          string `json:"PCR,omitempty"`
		Urbanisation string `json:"Urbanisation%,omitempty"`
		MajorCities  int    `json:"Major Cities,omitempty"`
		CapitalPort  string `json:"Capital/Port,omitempty"`
		Notes        string `json:"Notes,omitempty"`
	} `json:"Population,omitempty"`
	Government struct {
		Profile  string   `json:"Profile,omitempty"`
		Factions []string `json:"Factions,omitempty"` // предположительно список названий
		Notes    string   `json:"Notes,omitempty"`
	} `json:"Government,omitempty"`
	LawLevel struct {
		Primary string `json:"Primary,omitempty"`
		Profile string `json:"Profile,omitempty"`
		Notes   string `json:"Notes,omitempty"`
	} `json:"Law Level,omitempty"`
	Technology struct {
		Profile string `json:"Profile,omitempty"`
		Notes   string `json:"Notes,omitempty"`
	} `json:"Technology,omitempty"`
	Culture struct {
		Profile string `json:"Profile,omitempty"`
		Notes   string `json:"Notes,omitempty"`
	} `json:"Culture,omitempty"`
	Economics struct {
		GWPPerCapita    float64 `json:"GWP per capita,omitempty"`
		InequalityRating string  `json:"Inequality Rating,omitempty"`
		DevelopmentScore string  `json:"Development Score,omitempty"`
		GWP             float64 `json:"GWP (MCr),omitempty"`
		Notes           string  `json:"Notes,omitempty"`
	} `json:"Economics,omitempty"`
	Military struct {
		EffectiveBudgetPercent string `json:"Effective Budget %,omitempty"`
		Profile                string `json:"Profile,omitempty"`
		Notes                  string `json:"Notes,omitempty"`
	} `json:"Military,omitempty"`
	Comments string `json:"Comments,omitempty"`
}
```